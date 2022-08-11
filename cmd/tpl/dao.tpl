package dao_{{.Package}}

import (
	"github.com/oaago/cloud/mysql"
	"github.com/oaago/cloud/redis"
	"gorm.io/gorm"
	"{{.Module}}/internal/model/{{.Package}}_model"
	"{{.Module}}/internal/model/{{.Package}}_query"
)

type {{.UpMethod}}ModelType {{.Package}}_model.{{.UpMethod}}
type {{.UpMethod}}DaoType struct {
	DB      *gorm.DB
	Query   *{{.Package}}_query.Query
	TableName   string
	Model {{.UpMethod}}ModelType
	RedisClient redis.Cli
}

var (
    {{.UpMethod}}Model = {{.UpMethod}}ModelType{}
	{{.UpMethod}}Dao = {{.UpMethod}}DaoType{
		DB:      mysql.GetDBByName("{{.Package}}"),
		Query:   {{.Package}}_query.Use(mysql.GetDBByName("{{.Package}}")),
		TableName: {{.Package}}_model.TableName{{.UpMethod}},
		Model: {{.UpMethod}}Model,
		RedisClient: *redis.RedisClient,
	}
)