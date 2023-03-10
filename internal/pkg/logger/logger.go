package logger

import (
	"github.com/Aoi-hosizora/ahlib-more/xlogrus"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"github.com/sirupsen/logrus"
	"time"
)

func Setup() (*logrus.Logger, error) {
	c := xmodule.MustGetByName(sn.SConfig).(*config.Config)

	logger := logrus.New()
	logLevel := logrus.WarnLevel
	if c.Meta.RunMode == "debug" {
		logLevel = logrus.DebugLevel
	}

	logger.SetLevel(logLevel)
	logger.SetReportCaller(false)
	logger.SetFormatter(&xlogrus.CustomFormatter{TimestampFormat: time.RFC3339}) // TODO
	logger.AddHook(xlogrus.NewRotateLogHook(&xlogrus.RotateLogConfig{
		MaxAge:       15 * 24 * time.Hour,
		RotationTime: 24 * time.Hour,
		Filepath:     c.Meta.LogPath,
		Filename:     c.Meta.LogName,
		Level:        logLevel,
		Formatter:    &logrus.JSONFormatter{TimestampFormat: time.RFC3339},
	}))

	return logger, nil
}
