package sn

import "github.com/Aoi-hosizora/ahlib/xdi"

const (
	// common
	SConfig xdi.ServiceName = "config" // *config.Config
	SLogger xdi.ServiceName = "logger" // *logrus.Logger

	// service
	SMangaService xdi.ServiceName = "manga-service" // *service.MangaService
)
