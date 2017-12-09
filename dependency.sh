#!/bin/bash
echo "Downloading dependency"
go get -u -x github.com/spf13/viper
go get -u -x github.com/fsnotify/fsnotify
go get -u -x github.com/go-sql-driver/mysql
go get -u -x goji.io
echo "Done"