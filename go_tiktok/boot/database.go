package boot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"go_tiktok/app/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func MysqlSetup() {
	config := global.Config.Database.Mysql

	db, err := gorm.Open(mysql.Open(config.GetDsn()))
	if err != nil {
		global.Logger.Fatal("initialize mysql failed.", zap.Error(err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxIdleTime(global.Config.Database.Mysql.GetConnMaxIdleTime())
	sqlDB.SetConnMaxLifetime(global.Config.Database.Mysql.GetConnMaxLifeTime())
	sqlDB.SetMaxIdleConns(global.Config.Database.Mysql.MaxIdleconns)
	sqlDB.SetMaxOpenConns(global.Config.Database.Mysql.MaxOpenConns)
	err = sqlDB.Ping()
	if err != nil {
		global.Logger.Fatal("connect to mysql db failed.", zap.Error(err))
	}

	global.MysqlDB = db

	global.Logger.Info("initialize mysql successful")

}

func RedisSetup() {
	config := global.Config.Database.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Addr, config.Port),
		Username: config.Username,
		Password: config.Password,
		DB:       config.Db,
		PoolSize: config.PoolSize,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Fatal("connect to redis instance failed.", zap.Error(err))
	}

	global.Rdb = rdb

	global.Logger.Info("initialize redis client successfully!")
}
