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
