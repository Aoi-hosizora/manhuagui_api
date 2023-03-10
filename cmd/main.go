package main

import (
	"fmt"
	"github.com/Aoi-hosizora/ahlib-more/xpflag"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/api"
	"github.com/Aoi-hosizora/manhuagui-api/internal/pkg/module"
	"github.com/Aoi-hosizora/manhuagui-api/internal/server"
	"log"
)

var (
	fConfig = xpflag.Cmd().StringP("config", "c", "./config.yaml", "config file path")
	fHelp   = xpflag.Cmd().BoolP("help", "h", false, "show help")
)

func main() {
	// flag
	if xpflag.MustParse(); *fHelp {
		xpflag.PrintUsage()
		return
	}

	// module
	err := module.Provide(*fConfig) // may call fatal
	if err != nil {
		log.Fatalln("Failed to provide all modules:", err)
	}

	// document
	api.UpdateApiDoc()
	_, err = goapidoc.SaveSwaggerJson(api.SwaggerDocFilename)
	if err != nil {
		log.Fatalln("Failed to generate swagger:", err)
	}
	_, err = goapidoc.SaveApib(api.ApibDocFilename)
	if err != nil {
		log.Fatalln("Failed to generate apib:", err)
	}

	// server
	s, err := server.NewServer()
	if err != nil {
		log.Fatalln("Failed to create server:", err)
	}

	// start
	fmt.Println()
	s.Serve()
}
