package global

import (
	"Go-details/pkg/logger"
	"Go-details/pkg/setting"
)

/**
* @Author: super
* @Date: 2020-09-18 08:32
* @Description: 全局配置包括：服务，数据库，Email，JWT和日志
**/

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	Logger          *logger.Logger
)
