package cmd

import (
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/LittleMoreInteresting/sql-tool/internal/sqlToStruct"
	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string
var tmpl string

var sqlToStructCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql to struct",
	Long:  "sql to struct",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sqlToStruct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sqlToStruct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect error :%v", err)
		}
		template, _ := ioutil.ReadFile(tmpl)
		if len(template) == 0 {
			template = []byte(defaultTemp)
		}
		var dir string
		if len(tableName) > 0 {
			columns, err := dbModel.GetCloumns(dbName, tableName)

			if err != nil {
				log.Fatalf("dbModel.GetCloumns error :%v", err)
			}
			tpl := sqlToStruct.NewStructTemplate(dbName, string(template), dbName)
			dir = tpl.CheckDir()
			tplColumns := tpl.AssemblyColumns(columns)
			err = tpl.Generate(tableName, tplColumns)
			if err != nil {
				log.Fatalf("tpl.Generate error :%v", err)
			}
			exec.Command("gofmt", "-w", dir).Run()
			return
		}

		tables, err := dbModel.GetTableNames(dbName)
		if err != nil {
			log.Fatalf("dbModel.GetTableNames error :%v", err)
		}
		for _, t := range tables {
			columns, err := dbModel.GetCloumns(dbName, t)

			if err != nil {
				log.Fatalf("dbModel.GetCloumns error :%v", err)
			}
			tpl := sqlToStruct.NewStructTemplate(dbName, string(template), dbName)
			dir = tpl.CheckDir()
			tplColumns := tpl.AssemblyColumns(columns)
			err = tpl.Generate(t, tplColumns)
			if err != nil {
				log.Fatalf("tpl.Generate error :%v", err)
			}
		}
		err = exec.Command("gofmt", "-w", dir).Run()
	},
}

func init() {
	sqlToStructCmd.Flags().StringVarP(&username, "username", "u", "", "请输入数据库的账号")
	sqlToStructCmd.Flags().StringVarP(&password, "password", "p", "", "请输入数据库的密码")
	sqlToStructCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sqlToStructCmd.Flags().StringVarP(&charset, "charset", "c", "utf8mb4", "请输入数据库编码")
	sqlToStructCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库的类型")
	sqlToStructCmd.Flags().StringVarP(&dbName, "db", "d", "", "请输入数据库")
	sqlToStructCmd.Flags().StringVarP(&tableName, "table", "t", "", "请输入表名(不输入将全库导出)")
	sqlToStructCmd.Flags().StringVarP(&tmpl, "tmpl", "m", "./template/model.tmpl", "模版文件")
}

var defaultTemp = "package {{ .Database }}\n\nimport (\n     \"context\"\n     \"gorm.io/gorm\"\n)\n\ntype {{ .TableName | ToCamelCase }} struct {\n\t{{range $index, $element := .Columns}}{{ if gt $index 0 }}\n\t{{ end }}{{ $typeLen := len .Type }}{{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}\t{{.Type}}\t{{.Tag}}{{ else }}{{.Name}}{{ end }}{{ $length := len .Comment}}{{ if gt $length 0 }}// {{ .Comment }}{{else}}// {{.Name}}{{ end }}{{end}}\n}\n\nfunc (model {{ .TableName | ToCamelCase}}) TableName() string {\n    return \"{{ .TableName }}\"\n}\n\ntype {{ .TableName | ToCamelCase }}Interface interface {\n    Create{{ .TableName | ToCamelCase }}(ctx context.Context,t *{{ .TableName | ToCamelCase }}) (*{{ .TableName | ToCamelCase }},error)\n    Get{{ .TableName | ToCamelCase }}(ctx context.Context,id int64) (*{{ .TableName | ToCamelCase }},error)\n    Update{{ .TableName | ToCamelCase }}(ctx context.Context,t *{{ .TableName | ToCamelCase }}) error\n    Delete{{ .TableName | ToCamelCase }}(ctx context.Context,id int64) error\n}\nvar _ {{ .TableName | ToCamelCase }}Interface = (*{{ .TableName | ToCamelCase }}Repo)(nil)\n\ntype {{ .TableName | ToCamelCase }}Repo struct {\n    db  *gorm.DB\n}\n\nfunc New{{ .TableName | ToCamelCase }}Repo(db *gorm.DB) *{{ .TableName | ToCamelCase }}Repo {\n    return &{{ .TableName | ToCamelCase }}Repo{\n        db:db,\n    }\n}\n\n// Get DB client\nfunc(m *{{ .TableName | ToCamelCase }}Repo)DB(ctx context.Context) *gorm.DB {\n    return m.db\n}\n\n//Create One record\nfunc (m *{{ .TableName | ToCamelCase }}Repo) Create{{ .TableName | ToCamelCase }}(ctx context.Context,t *{{ .TableName | ToCamelCase }}) (*{{ .TableName | ToCamelCase }},error){\n    err := m.DB(ctx).Create(t).Error\n    if err != nil {\n        return nil,err\n    }\n    return t,err\n}\n\n//Get One record\nfunc (m *{{ .TableName | ToCamelCase }}Repo) Get{{ .TableName | ToCamelCase }}(ctx context.Context,id int64) (*{{ .TableName | ToCamelCase }},error){\n    record := &{{ .TableName | ToCamelCase }}{}\n    err:= m.DB(ctx).First(record,id).Error\n    return record,err\n}\n\n//Update One record\nfunc (m *{{ .TableName | ToCamelCase }}Repo) Update{{ .TableName | ToCamelCase }}(ctx context.Context,t *{{ .TableName | ToCamelCase }}) error {\n    return m.DB(ctx).Updates(t).Error\n}\n\n//Delete One record\nfunc (m *{{ .TableName | ToCamelCase }}Repo) Delete{{ .TableName | ToCamelCase }}(ctx context.Context,id int64) error {\n    return m.DB(ctx).Delete(&{{ .TableName | ToCamelCase }}{},id).Error\n}"
