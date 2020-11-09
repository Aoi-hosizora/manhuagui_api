package main

import (
	"flag"
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/provide"
	"github.com/Aoi-hosizora/manhuagui-backend/src/server"
	"log"
)

var (
	fConfig = flag.String("config", "./config.yaml", "change config path")
	fHelp   = flag.Bool("h", false, "show help")
)

func main() {
	flag.Parse()
	if *fHelp {
		flag.Usage()
	} else {
		run()
	}
}

func run() {
	_, err := goapidoc.GenerateSwaggerJson("./docs/doc.json")
	if err != nil {
		log.Fatalln("Failed to generate swagger:", err)
	}
	_, err = goapidoc.GenerateApib("./docs/doc.apib")
	if err != nil {
		log.Fatalln("Failed to generate apib:", err)
	}

	err = provide.Provide(*fConfig)
	if err != nil {
		log.Fatalln("Failed to load some service:", err)
	}

	s := server.NewServer()
	s.Serve()
}
