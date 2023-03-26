package iplimiter

import (
	"github.com/Aoi-hosizora/ahlib/xreflect"
	"github.com/juju/ratelimit"
	"sync"
	"time"
)

type IPRateLimiter struct {
	limiters    map[string]*ratelimit.Bucket
	limitersMu  sync.RWMutex
	survivors   map[string]int16
	survivorsMu sync.RWMutex

	cleanupInterval time.Duration
	maxSurvived     int16
	bucketCreator   func() *ratelimit.Bucket

	lastCheckTime  time.Time
	cleanupCheckMu sync.Mutex
}

func NewIPRateLimiter(cleanupInterval time.Duration, maxSurvived int16, bucketCreator func() *ratelimit.Bucket) *IPRateLimiter {
	return &IPRateLimiter{
		limiters:  make(map[string]*ratelimit.Bucket),
		survivors: make(map[string]int16),

		cleanupInterval: cleanupInterval,
		maxSurvived:     maxSurvived,
		bucketCreator:   bucketCreator,

		lastCheckTime: time.Now(),
	}
}

func (i *IPRateLimiter) GetOrCreate(clientIP string) (*ratelimit.Bucket, time.Time, func(count int64) int64) {
	defer func() {
		i.cleanupCheckMu.Lock() // exclusive checking
		defer i.cleanupCheckMu.Unlock()
		if time.Now().Sub(i.lastCheckTime) >= i.cleanupInterval {
			i.lastCheckTime = time.Now()
			go i.checkCleanup()
		}
	}()

	i.limitersMu.RLock()
	limiter, ok := i.limiters[clientIP]
	i.limitersMu.RUnlock()
	if !ok {
		limiter = i.bucketCreator()
		i.limitersMu.Lock()
		i.limiters[clientIP] = limiter
		i.limitersMu.Unlock()
	}

	startTime := xreflect.GetUnexportedField(xreflect.FieldValueOf(limiter, "startTime")).Interface().(time.Time)
	return limiter, startTime, func(count int64) int64 { return i.takeAvailable(clientIP, limiter, count) }
}

func (i *IPRateLimiter) takeAvailable(clientIP string, bucket *ratelimit.Bucket, count int64) int64 {
	taken := bucket.TakeAvailable(count)
	i.delSurvivor(clientIP) // must be not filled
	return taken
}

// !!!
func (i *IPRateLimiter) checkCleanup() {
	// wait extra 100ms for finishing taking
	time.Sleep(time.Millisecond * 100)

	// check
	cleanableIPs := make([]string, 0)
	i.limitersMu.RLock()
	for ip, bucket := range i.limiters {
		if bucket.Available() < bucket.Capacity() {
			i.delSurvivor(ip) // not filled
			continue
		}
		survived, ok := i.getSurvivor(ip)
		if !ok {
			i.setSurvivor(ip, 1) // survived firstly
		} else {
			survived++
			if survived <= i.maxSurvived {
				i.setSurvivor(ip, survived) // survived again
			} else {
				i.delSurvivor(ip) // reach max survived, clean
				cleanableIPs = append(cleanableIPs, ip)
			}
		}
	}
	i.limitersMu.RUnlock()

	// update and cleanup
	i.limitersMu.Lock()
	for _, ip := range cleanableIPs {
		delete(i.limiters, ip)
	}
	i.limitersMu.Unlock()
}

func (i *IPRateLimiter) getSurvivor(clientIP string) (int16, bool) {
	i.survivorsMu.RLock()
	survived, ok := i.survivors[clientIP]
	i.survivorsMu.RUnlock()
	return survived, ok
}

func (i *IPRateLimiter) delSurvivor(clientIP string) {
	i.survivorsMu.RLock()
	_, ok := i.survivors[clientIP]
	i.survivorsMu.RUnlock()
	if ok {
		i.survivorsMu.Lock()
		delete(i.survivors, clientIP)
		i.survivorsMu.Unlock()
	}
}

func (i *IPRateLimiter) setSurvivor(clientIP string, newValue int16) {
	i.survivorsMu.Lock()
	i.survivors[clientIP] = newValue
	i.survivorsMu.Unlock()
}
