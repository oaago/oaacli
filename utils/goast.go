package utils

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Tags struct {
	Name string
	Type string
	Tags map[string]string
}

func GetAllStruct(path string) (string, map[string][]Tags) {
	file, err := GetAllFile(path)
	if err != nil {
		return "", nil
	}
	var structList = make(map[string][]Tags)
	var packageName string
	for _, s := range file {
		packageName, structList = MapStruct(s)
		fmt.Println(s, packageName, structList)
	}
	return packageName, structList
}

func GetAllFile(pathname string) ([]string, error) {
	var files []string
	err := filepath.Walk(pathname, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "types.go") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func MapStruct(path string) (string, map[string][]Tags) {
	set := token.NewFileSet()
	byteF, _ := os.ReadFile(path)
	f, err := parser.ParseFile(set, "", string(byteF), parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	// Collect the struct types in this slice.
	//var structTypes = make(map[string][]string)
	var structStr = map[string][]Tags{}
	var structKey = map[string]bool{}
	// Use the Inspect function to walk AST looking for struct
	// type nodes.
	ast.Inspect(f, func(n ast.Node) bool {
		if _, ok := n.(*ast.StructType); ok {
			//fmt.Println(f.Name.Name, f.Scope.Objects, n)
			for _, object := range f.Scope.Objects {
				//fmt.Println(s, object.Decl.(*ast.TypeSpec))
				if nn, oo := object.Decl.(*ast.TypeSpec); oo {
					if nnn, ooo := nn.Type.(*ast.StructType); ooo {
						//fmt.Println(object.Name, nnn.Fields.List)
						for _, fields := range nnn.Fields.List {
							if nnnn, oooo := fields.Type.(*ast.Ident); oooo {
								if nnnn.Obj != nil {
									panic("在" + object.Name + "里面的" + nnnn.Obj.Name + "不可用")
								}
								key := object.Name + "-" + fields.Names[0].Name
								if !structKey[key] {
									var Tag = make(map[string]string)
									//sf := fields.Names[0].Name + "," + nnnn.Name + "," + fields.Tag.Value
									fmt.Println(fields.Tag.Value)
									tagStr := strings.Split(fields.Tag.Value, " ")
									for i, s := range tagStr {
										fmt.Println(i, s)
										ss := strings.Replace(s, "`", "", 1)
										nss := strings.Split(ss, ":")
										fmt.Println(nss)
										Tag[nss[0]] = strings.Replace(nss[1], "\"", "", -1)
									}
									sf := Tags{
										Name: fields.Names[0].Name,
										Type: nnnn.Name,
										Tags: Tag,
									}
									structStr[object.Name] = append(structStr[object.Name], sf)
									structKey[key] = true
								}
							}
						}
					}
				}
			}
		}
		return true
	})
	return f.Name.Name, structStr
}
