package api

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"os"
)

const (
	SwaggerDocFilename = "./api/spec.json"
	ApibDocFilename    = "./api/spec.apib"
)

func ReadSwaggerDoc() []byte {
	bs, _ := os.ReadFile(SwaggerDocFilename)
	return bs
}

func UpdateApiDoc() {
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Meta
	host := cfg.DocHost
	if host == "" {
		host = fmt.Sprintf("localhost:%d", cfg.Port)
	}
	goapidoc.SetHost(host)
}
