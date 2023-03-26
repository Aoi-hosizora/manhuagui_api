package gincache

import (
	"bufio"
	"errors"
	"github.com/Aoi-hosizora/ahlib/xstring"
	"github.com/bluele/gcache"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

type ResponseData struct {
	status int
	header http.Header
	data   []byte

	expiredAt time.Time // used by ResponseDataCacher
}

func (r *ResponseData) Status() int {
	return r.status
}

func (r *ResponseData) Header() http.Header {
	return r.header
}

func (r *ResponseData) Data() []byte {
	return r.data
}

func Is2XXStatus(code int) bool {
	return code >= 200 && code < 300
}

type ResponseDataCacher interface {
	Get(key string) (*ResponseData, error)
	Set(key string, data *ResponseData) error
	GetWithExpiration(key string) (*ResponseData, time.Time, error)
	SetWithExpiration(key string, data *ResponseData, expiration time.Duration) error
}

type CacheStorage struct {
	cache      gcache.Cache
	expiration time.Duration
}

var _ ResponseDataCacher = (*CacheStorage)(nil)

func NewCacheStorage(size int16, expiration time.Duration) *CacheStorage {
	return &CacheStorage{
		cache:      gcache.New(int(size)).LRU().Expiration(expiration).Build(),
		expiration: expiration,
	}
}

func (c *CacheStorage) Get(key string) (*ResponseData, error) {
	data, err := c.cache.Get(key)
	if err != nil {
		return nil, err
	}
	rdata, ok := data.(*ResponseData)
	if !ok {
		return nil, errors.New("data is not of ResponseData type")
	}
	return rdata, nil
}

func (c *CacheStorage) GetWithExpiration(key string) (*ResponseData, time.Time, error) {
	data, err := c.Get(key)
	if err != nil {
		return nil, time.Time{}, err
	}
	return data, data.expiredAt, nil
}

func (c *CacheStorage) Set(key string, data *ResponseData) error {
	data.expiredAt = time.Now().Add(c.expiration)
	return c.cache.Set(key, data)
}

func (c *CacheStorage) SetWithExpiration(key string, data *ResponseData, expiration time.Duration) error {
	data.expiredAt = time.Now().Add(expiration)
	return c.cache.SetWithExpire(key, data, expiration)
}

type CachedWriter struct {
	origin   gin.ResponseWriter
	storage  ResponseDataCacher
	cacheKey string

	data []byte
}

var _ gin.ResponseWriter = (*CachedWriter)(nil)

func NewCachedWriter(writer gin.ResponseWriter, storage ResponseDataCacher, cacheKey string) *CachedWriter {
	return &CachedWriter{
		origin:   writer,
		storage:  storage,
		cacheKey: cacheKey,
	}
}

func (r *CachedWriter) Origin() gin.ResponseWriter {
	return r.origin
}

func (r *CachedWriter) WriteCache() error {
	data := &ResponseData{
		status: r.Status(),
		header: r.Header(),
		data:   r.Data(),
	}
	return r.storage.Set(r.cacheKey, data)
}

func (r *CachedWriter) Write(bytes []byte) (int, error) {
	n, err := r.origin.Write(bytes)
	if err == nil {
		r.data = append(r.data, bytes...)
	}
	return n, err
}
func (r *CachedWriter) WriteString(s string) (int, error) {
	n, err := r.origin.WriteString(s)
	if err == nil {
		r.data = append(r.data, xstring.FastStob(s)...)
	}
	return n, err
}

func (r *CachedWriter) WriteHeader(statusCode int) {
	r.origin.WriteHeader(statusCode)
}

func (r *CachedWriter) Status() int {
	return r.origin.Status()
}

func (r *CachedWriter) Header() http.Header {
	return r.origin.Header()
}

func (r *CachedWriter) Data() []byte {
	return r.data
}

func (r *CachedWriter) Size() int {
	return r.origin.Size()
}

func (r *CachedWriter) Written() bool {
	return r.origin.Written()
}

func (r *CachedWriter) WriteHeaderNow() {
	r.origin.WriteHeaderNow()
}

func (r *CachedWriter) Flush() {
	r.origin.Flush()
}

func (r *CachedWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return r.origin.Hijack()
}

func (r *CachedWriter) CloseNotify() <-chan bool {
	return r.origin.CloseNotify()
}

func (r *CachedWriter) Pusher() http.Pusher {
	return r.origin.Pusher()
}
