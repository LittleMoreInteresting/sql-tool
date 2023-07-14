package sqlToStruct

import (
	"fmt"
	"text/template"

	"github.com/sql-tool/pkg/file"
	"github.com/sql-tool/pkg/word"
)

type StructTemplate struct {
	structTpl string
	Dir       string
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

func NewStructTemplate(dir, tmpl string) *StructTemplate {
	return &StructTemplate{structTpl: tmpl, Dir: "dist/" + dir}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, len(tbColumns))
	for i, v := range tbColumns {
		tag := fmt.Sprintf("`json:\"%s\"`", v.ColumnName)
		tplColumns[i] = &StructColumn{
			Name:    v.ColumnName,
			Type:    DBTypeToStructType[v.DataType],
			Tag:     tag,
			Comment: v.ColumnComment,
		}
	}

	return tplColumns
}

func (t *StructTemplate) Generate(tbName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sqlToSturct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tbName,
		Columns:   tplColumns,
	}
	out, err := file.CreateWriter(t.Dir + "/" + tbName + ".go")
	if err != nil {
		return err
	}
	return tpl.Execute(out, tplDB)
}

func (t *StructTemplate) CheckDir() string {
	if file.CheckSavePath(t.Dir) {
		file.CreateSavePath(t.Dir)
	}
	return t.Dir
}
