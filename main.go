package main

import (
	"flag"
	"fmt"
	"github.com/a-aslani/golang_agency_clean_architecture/application"
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework"
	"os"
)

var Version = "0.0.1"

func main() {

	configFile := os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = "config.local.yml"
	}

	cfg, err := configs.InitConfig(configFile)
	if err != nil {
		fmt.Printf("config file error: %s", err.Error())
		return
	}

	appMap := map[string]framework.Runner{
		"swagger":       application.NewSwagger(),
		"newsletter":    application.NewNewsletter(),
		"file_uploader": application.NewFileUploader(),
		"support":       application.NewSupport(),
		"project":       application.NewProject(),
	}

	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if !exist {
		fmt.Printf("You may try :\n\n")
		for appName := range appMap {
			fmt.Printf("    go run main.go %s\n", appName)
		}
		fmt.Printf("\n")
		return
	}

	fmt.Printf("Config: %s - Version: %s\n", configFile, Version)

	err = app.Run(cfg)
	if err != nil {
		fmt.Printf("run error: %s", err.Error())
		return
	}
}
