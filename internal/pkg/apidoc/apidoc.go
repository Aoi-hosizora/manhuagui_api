package apidoc

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-mx/xgin"
	"github.com/Aoi-hosizora/ahlib/xmodule"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/config"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module/sn"
	"os"
)

var (
	ParamPage  = goapidoc.NewQueryParam("page", "integer#int32", false, "current page")
	ParamLimit = goapidoc.NewQueryParam("limit", "integer#int32", false, "page size")
	ParamOrder = goapidoc.NewQueryParam("order", "string", false, "order string")
	ParamToken = goapidoc.NewHeaderParam("Authorization", "string", true, "access token")
)

const (
	swaggerDocFilename = "./api/doc.json"
	apibDocFilename    = "./api/doc.apib"
)

func ReadSwaggerDoc() []byte {
	bs, _ := os.ReadFile(swaggerDocFilename)
	return bs
}

func SwaggerOptions() []xgin.SwaggerOption {
	return []xgin.SwaggerOption{
		xgin.WithSwaggerDefaultModelExpandDepth(999),
		xgin.WithSwaggerDisplayRequestDuration(true),
		xgin.WithSwaggerShowExtensions(true),
		xgin.WithSwaggerShowCommonExtensions(true),
	}
}

func UpdateAndSave() error {
	// update host
	cfg := xmodule.MustGetByName(sn.SConfig).(*config.Config).Meta
	host := cfg.DocHost
	if host == "" {
		host = fmt.Sprintf("localhost:%d", cfg.Port)
	}
	goapidoc.SetHost(host)

	// update param
	allowCache := goapidoc.NewQueryParam("allow_cache", "boolean", false, "flag to allow using cache").Default(false) // for backward-compatibility
	for _, op := range goapidoc.GetOperations() {
		if op.GetMethod() == "GET" {
			op.AddParams(allowCache)
		}
	}

	// save
	_, err := goapidoc.SaveSwaggerJson(swaggerDocFilename)
	if err != nil {
		return err
	}
	_, err = goapidoc.SaveApib(apibDocFilename)
	if err != nil {
		return err
	}
	return nil
}
