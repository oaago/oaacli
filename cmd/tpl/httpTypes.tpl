package {{.PackageName}}

{{range $s := .StructMap}}
type {{$s}} struct {

}
{{end}}
