package logger

import (
	"github.com/Aoi-hosizora/ahlib-more/xlogrus"
	"github.com/Aoi-hosizora/ahlib-more/xrotation"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/ahlib/xtime"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/sirupsen/logrus"
	"time"
)

func Setup() (*logrus.Logger, error) {
	logger := logrus.New()
	level := logrus.WarnLevel
	if config.IsDebugMode() {
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)
	logger.SetReportCaller(false)
	logger.SetFormatter(xlogrus.NewSimpleFormatter(
		xlogrus.WithTimestampFormat(time.RFC3339),
		xlogrus.WithTimeLocation(time.Local),
	))

	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Meta
	rotation, err := xrotation.New(
		cfg.LogName+".%Y%m%d.log",
		xrotation.WithSymlinkFilename(cfg.LogName+".current.log"),
		xrotation.WithRotationTime(24*time.Hour),
		xrotation.WithRotationMaxAge(15*24*time.Hour),
		xrotation.WithClock(xtime.Local),
	)
	if err != nil {
		return nil, err
	}
	logger.AddHook(xlogrus.NewRotationHook(
		rotation,
		xlogrus.WithRotateLevel(level),
		xlogrus.WithRotateFormatter(xlogrus.RFC3339JsonFormatter()),
	))

	return logger, nil
}
