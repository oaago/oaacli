package cli

import (
	"fmt"
	tpl2 "github.com/oaago/oaago/cmd/tpl"
	_const2 "github.com/oaago/oaago/const"
	"github.com/oaago/oaago/utils"
	"html/template"
	"os"
	"strings"
)

func genTypes(CurrentDBName, dirName, fileName string, hasTable bool) {
	// 先选择数据库
	if len(_const2.TableMap) > 1 {
		fmt.Println("请输入要关联的数据库")
		fmt.Scanln(&CurrentDBName)
	}
	typesDir := utils.Camel2Case(_const2.ServicePath) + utils.Camel2Case(dirName)
	tableName := utils.Camel2Case(dirName + "_" + fileName)
	hasDbName := ""
	// 搜索数据库是否存在表
	for _, table := range _const2.TableMap[CurrentDBName] {
		if table == tableName {
			hasTable = true
			hasDbName = CurrentDBName
		}
	}
	// 如果存在
	if hasTable {
		utils.TableStruct(CurrentDBName, utils.Camel2Case(dirName+"_"+fileName), typesDir+"/"+utils.Camel2Case(fileName))
		fmt.Println("表名称："+utils.Camel2Case(dirName+"_"+fileName), "结构：", hasDbName)
		return
	} else {
		// 如果表不存在证明是单纯增加一个接口
		genTypesFiles(dirName, fileName)
	}
}

func genTypesFiles(dirName, fileName string) {
	var structMap = make([]string, 0)
	for _, funcName := range _const2.SemanticMap {
		structName := strings.Replace(funcName.FunctionName, "$", utils.Ucfirst(dirName)+utils.Case2Camel(utils.Ucfirst(fileName)), 1)
		structMap = append(structMap, structName+"Req", structName+"Res")
	}
	type DefinedType struct {
		PackageName string
		StructMap   []string
	}
	var typeData = DefinedType{
		PackageName: utils.Camel2Case(dirName + "_" + fileName),
		StructMap:   structMap,
	}

	//创建模板
	defined := "http-type"
	tmpl := template.New(defined)
	//解析模板
	text := tpl2.HttpTypes
	tpl, err := tmpl.Parse(text)
	if err != nil {
		panic(err)
	}
	typesDir := utils.Camel2Case(_const2.ServicePath) + utils.Camel2Case(dirName)
	hasDir, _ := utils.PathExists(typesDir)
	if !hasDir {
		err := os.Mkdir(typesDir, os.ModePerm)
		err = os.Mkdir(typesDir+"/"+utils.Camel2Case(fileName), os.ModePerm)
		if err != nil {
			panic("目录初始化失败" + err.Error())
		}
	}
	filePath := typesDir + "/" + utils.Camel2Case(fileName) + "/types.go"
	//hasFile, _ := utils.PathExists(filePath)
	//if hasFile {
	//	fmt.Println(filePath + "文件已存在, 不会继续生成")
	//	return
	//}
	//渲染输出
	fs, _ := os.Create(filePath)
	err = tpl.ExecuteTemplate(fs, defined, typeData)
	if err != nil {
		panic(err)
	}
	fs.Close()
	fmt.Println("写入types模版成功 " + typesDir)
}
