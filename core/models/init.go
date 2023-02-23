package models

import (
	"xorm.io/xorm"
	"log"
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-redis/redis/v8"
)

func InitMysql(dataSource string)*xorm.Engine{
	engine,err := xorm.NewEngine("mysql",dataSource)
	if err != nil{
		log.Println("mysql init failed......err:",err)
		return nil
	}
	log.Println("mysql init success......")
	return engine
}

func InitRedis(addr,password string,db int)*redis.Client{
	client:= redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, 
		DB:       db, 
	})
	ping,err := client.Ping(context.Background()).Result()
	if err != nil{
		log.Println("redis init failed......err:",err)
		return nil
	}
	log.Println("redis init success......",ping)
	return client
}