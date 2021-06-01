package global

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DBEngine  *gorm.DB
	RedisPool *redis.Pool
	Logger    *logrus.Logger
	Es        *elasticsearch.Client
	Mongo     *mongo.Client
)
