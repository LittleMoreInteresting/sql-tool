package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/sql-tool/internal/sqlToStruct"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sqlToStructCmd = &cobra.Command{
	Use: "struct",
	Short: "sql to struct",
	Long : "sql to struct",
	Run: func(cmd *cobra.Command,args []string){
		dbInfo := &sqlToStruct.DBInfo{
			DBType:dbType,
			Host:host,
			UserName:username,
			Password:password,
			Charset:charset,
		}
		dbModel := sqlToStruct.NewDBModel(dbInfo)
		err := dbModel.Connect();
		if err != nil {
			log.Fatalf("dbModel.Connect error :%v",err)
		}
		if len(tableName) >= 0{
			columns, err := dbModel.GetCloumns(dbName, tableName)
		
			if err != nil {
				log.Fatalf("dbModel.GetCloumns error :%v",err)
			}
			tpl := sqlToStruct.NewStructTemplate()
			tplColumns := tpl.AssemblyColumns(columns)
			err = tpl.Generate(tableName, tplColumns)
			if err != nil {
				log.Fatalf("tpl.Generate error :%v",err)
			}
		}
		
		tables,err := dbModel.GetTableNames(dbName)
		if err != nil {
			log.Fatalf("dbModel.GetTableNames error :%v",err)
		}
		for _, t := range tables {
			columns, err := dbModel.GetCloumns(dbName, t)
		
			if err != nil {
				log.Fatalf("dbModel.GetCloumns error :%v",err)
			}
			tpl := sqlToStruct.NewStructTemplate()
			tplColumns := tpl.AssemblyColumns(columns)
			err = tpl.Generate(t, tplColumns)
			if err != nil {
				log.Fatalf("tpl.Generate error :%v",err)
			}
		}
		
	},
}

func init() {
	sqlCmd.AddCommand(sqlToStructCmd)
	sqlToStructCmd.Flags().StringVarP(&username, "username", "U", "","请输入数据库的账号")
	sqlToStructCmd.Flags().StringVarP(&password, "password", "P", "","请输入数据库的密码")
	sqlToStructCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1:3306","请输入数据库的HOST")
	sqlToStructCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4","请输入数据库编码")
	sqlToStructCmd.Flags().StringVarP(&dbType, "type", "", "mysql","请输入数据库的类型")
	sqlToStructCmd.Flags().StringVarP(&dbName, "db", "", "","请输入数据库")
	sqlToStructCmd.Flags().StringVarP(&tableName, "table", "", "","请输入表名")
}