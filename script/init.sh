#!/usr/bin/env bash


go get github.com/astaxie/beego
go get github.com/mattn/go-sqlite3
go get github.com/satori/go.uuid
# 执行数据初始化
if [ "$DEBUG" == "True" ]; then
    echo "use debug mode"
else
    echo "use product mode"

fi



exec "$@"