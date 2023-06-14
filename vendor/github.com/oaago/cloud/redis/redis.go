package redis

import (
	"github.com/go-redis/redis"
	"github.com/oaago/cloud/config"
	"github.com/oaago/cloud/logx"
	"os"
	"strconv"
	"time"
)

var Client *redis.Client

type ClientType *redis.Client
type Cli struct {
	Client *redis.Client
	Name   string
}

type RedisType struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
	Name     string `yaml:"name"`
}

var RedisOptions = &RedisType{}

var RedisClient = &Cli{}

func init() {
	enable := config.Op.GetBool("redis.enable")
	arg := os.Args
	if !enable || arg[0] == "oaago" {
		return
	}
	redisStr := config.Op.GetStringMapString("redis")
	RedisOptions.DB, _ = strconv.Atoi(redisStr["db"])
	RedisOptions.Addr = redisStr["addr"]
	RedisOptions.Password = redisStr["password"]
	RedisClient = RedisOptions.NewRedis()
	Client = RedisClient.Client
}

func (op *RedisType) NewRedis() *Cli {
	redisCli := &Cli{}
	redisCli.Client = redis.NewClient(&redis.Options{
		Addr:        op.Addr,
		Password:    op.Password, // no password set
		DB:          op.DB,       // use default DB
		ReadTimeout: time.Second * time.Duration(10),
	})
	redisCli.Name = op.Name
	_, err := redisCli.Client.Ping().Result()
	if err != nil {
		logx.Logger.Info("redis连接异常", err)
		return nil
	} else {
		logx.Logger.Info("redis连接成功")
	}
	return redisCli
}
