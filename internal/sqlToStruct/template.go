package sqlToStruct

import (
	"fmt"
	"os"
	"text/template"
)

const structTpl = `type {{. TableName | ToCamelCase} struct {
{{range .Columns}} {{ length := len .Comment}} {{ if gt $length 0 }}
	//{{.Comment}} {{else}}// {{.Name}}{{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }} {{.Name | ToCamelCase}} {{.Type}} {{.Tag}} {{ else }} {{.Name}} {{ end }}
{{end}}}
func (model {{.TableName | ToCamelCase}}} TableName() string {
	return "{{. TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, len(tbColumns))
	for _, v := range tbColumns {
		tag := fmt.Sprintf("`json:\"%s\"`", v.ColumnName)
		tplColumns = append(tplColumns,&StructColumn{
			Name:v.ColumnName,
			Type: DBTypeToStructType[v.ColumnType],
			Tag: tag,
			Comment: v.ColumnComment,
		})
	}

	return tplColumns
}

func (t *StructTemplate) Generate (tbName string,tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sqlToSturct").Funcs(template.FuncMap{
		"ToCamelCase" :nil,
	}).Parse(t.structTpl))
	
	tplDB := StructTemplateDB{
		TableName:tbName,
		Columns: tplColumns,
	}
	
	return tpl.Execute(os.Stdout,tplDB)
}