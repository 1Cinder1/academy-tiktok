package global

import (
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"go_tiktok/app/internal/model/config"
	"gorm.io/gorm"
)

var (
	Config  *model.Config
	Logger  *zap.Logger
	MysqlDB *gorm.DB
	Rdb     *redis.Client
)
