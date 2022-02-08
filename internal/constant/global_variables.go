package constant

import (
	"Go-Blog/internal/adapter/outbound/db"
	"Go-Blog/internal/domain/conveter"
)

var (
	UserConverter = &conveter.UserConverter{}
	tool          = &db.MysqlTool{}
	Connect       = tool.GetInstance().GetDb()
)

const (
	RECORD_NOT_EXIST = "record not exist"
)
