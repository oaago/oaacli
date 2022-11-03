package utils

import (
	"bytes"
	"fmt"
	"github.com/oaago/cloud/mysql"
	"github.com/oaago/cloud/op"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/pterm/pterm"
)

func RunCmd(cmd string, shell bool) []byte {
	if shell {
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			log.Fatal(err)
			panic("some error found")
		}
		return out
	}
	out, err := exec.Command(cmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

// PathExists 判断目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(GetCurrentPath() + "/" + path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Camel2Case 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// Camel2Comm 中划线转下划线
func Camel2Comm(name string) string {
	str := strings.ReplaceAll(name, "-", "_")
	return Case2Camel(str)
}

// Case2Camel 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Ucfirst 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Lcfirst 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// Buffer 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

// GetCurrentPath 获取当前程序路径
func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	path := strings.Replace(dir, "\\", "/", -1) + "/"
	return path
}

func CLIScreen() {
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Oaago CLI"))
	introSpinner, _ := pterm.DefaultSpinner.WithShowTimer(false).WithRemoveWhenDone(true).Start("Waiting for ...")
	pterm.DefaultSection.Println("安装完成，请使用 oaago 命令进行操作")
	introSpinner.Stop() //nolint:errcheck
}

func LoadAllTables() map[string][]string {
	tables := make(map[string][]string)
	for s, v := range op.ConfigData.Mysql {
		fmt.Println(s, v)
		if s != "enable" {
			tables[s] = GetTables(s)
		}
	}
	return tables
}

func GetTables(dbName string) []string {
	db, _ := mysql.NewConnect(dbName)
	rows, _ := db.Raw("show tables").Rows()
	defer rows.Close()
	var tables = make([]string, 0)
	for rows.Next() {
		var name string
		rows.Scan(&name) //nolint:errcheck
		tables = append(tables, name)
	}
	return tables
}

func TableStruct(dbName, tableName, path string) map[string]map[string]string {
	tables := make(map[string]map[string]string)
	fmt.Println("生成types " +
		"dbName:" + dbName + "" +
		"tableName:" + tableName + "" +
		"path:" + path)
	t2t := NewTable2Struct()
	// 个性化配置
	t2t.Config(&T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
	})
	// 开始迁移转换
	tables, err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(tableName).
		// 表前缀
		Prefix("").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName(tableName).
		// tag字段的key值,默认是orm
		TagKey("").
		// 是否添加结构体方法获取表名
		RealNameMethod("").
		// 生成的结构体保存路径
		SavePath(path + "/types.go").
		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn(op.ConfigData.Mysql[dbName] + "?charset=utf8").
		// 执行
		Run()
	fmt.Println(op.ConfigData.Mysql[dbName])
	if err != nil {
		fmt.Println(err, "op.ConfigData.Mysql[dbName]")
		panic(err)
	}
	return tables
}
