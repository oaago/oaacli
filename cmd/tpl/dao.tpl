package dao_{{.Package}}

import (
	"github.com/oaago/cloud/mysql"
	"github.com/oaago/cloud/redis"
	"gorm.io/gorm"
	"{{.Module}}/internal/model/{{.Package}}_model"
	"{{.Module}}/internal/model/{{.Package}}_query"
)

type {{.UpMethod}}Dao struct {
	DB      *gorm.DB
	Query   *{{.Package}}_query.Query
	Table   string
	Columns {{.Package}}_model.{{.UpMethod}}
	RedisClient redis.Cli
}

var (
	{{.UpMethod}}Model = {{.Package}}_model.{{.UpMethod}}{}
	{{.UpMethod}} = {{.UpMethod}}Dao{
		DB:      mysql.GetDBByName("{{.Package}}"),
		Query:   {{.Package}}_query.Use(mysql.GetDBByName("{{.Package}}")),
		Table:   {{.UpMethod}}Model.TableName(),
		Columns: {{.Package}}_model.{{.UpMethod}}{},
		RedisClient: *redis.RedisClient,
	}
)