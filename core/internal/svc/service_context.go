package svc

import (
	"liujun/Time_Cloud_Disk/core/internal/config"
	"liujun/Time_Cloud_Disk/core/models"
	"xorm.io/xorm"
	"github.com/go-redis/redis/v8"
)

type ServiceContext struct {
	Config config.Config
	DB *xorm.Engine
	RED *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB: models.InitMysql(c.Mysql.DataSource),
		RED: models.InitRedis(c.Redis.Addr,c.Redis.Password,c.Redis.DB),
	}
}
