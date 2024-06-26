package global

import (
	"github.com/ksks2012/ubiquitous-sniffle/pkg/logger"
	"github.com/ksks2012/ubiquitous-sniffle/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	Logger          *logger.Logger
)
