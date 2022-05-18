package common

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"zeropiece/conf"
)

var (
	CONFIG conf.Config
	VP     *viper.Viper
	LOG    *zap.Logger
	DB     *gorm.DB
	CACHE  *redis.Client
)
