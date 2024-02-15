package framework

import (
	"github.com/a-aslani/golang_agency_clean_architecture/configs"
	"github.com/a-aslani/golang_agency_clean_architecture/pkg/framework/util"
	"time"
)

type Runner interface {
	Run(cfg *configs.Config) error
}

type ApplicationData struct {
	AppName       string `json:"app_name"`
	AppInstanceID string `json:"app_instance_id"`
	StartTime     string `json:"start_time"`
}

func NewApplicationData(appName string) ApplicationData {
	return ApplicationData{
		AppName:       appName,
		AppInstanceID: util.GenerateID(4),
		StartTime:     time.Now().Format("2006-01-02 15:04:05"),
	}
}
