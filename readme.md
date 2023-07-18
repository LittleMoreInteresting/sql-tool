# Sql To Struct Tool

## Install

> go install github.com/LittleMoreInteresting/sql-tool@latest
> 

## Usage

> sql-tool sql -u root -p root --db bss_sys 

```shell
Usage:
   sql [flags]

Flags:
  -c, --charset string    请输入数据库编码 (default "utf8mb4")
  -d, --db string         请输入数据库
  -h, --help              help for struct
      --host string       请输入数据库的HOST (default "127.0.0.1:3306")
  -p, --password string   请输入数据库的密码
  -t, --table string      请输入表名(不输入将全库导出)
  -m, --tmpl string       模版文件 (default "./template/model.tmpl")
      --type string       请输入数据库的类型 (default "mysql")
  -u, --username string   请输入数据库的账号

```