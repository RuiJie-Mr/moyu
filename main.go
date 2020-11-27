package main

import (
	"fmt"
	"os"
	"strconv"
	"moyu/utils"
)

func main()  {
	if len(os.Args)< 2{
		fmt.Println("请输入文件名")
		return
	}
	path := os.Args[1]
	//获取当前目录
	nowPath,err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//读取日志，获取已读行数
	nowPath = fmt.Sprintf("%s/moyu.log",nowPath)
	//判断当前目录是否有日志文件
	islog ,err := utils.PathExists(nowPath)
	if !islog {
		f,err := os.Create(nowPath)
		defer f.Close()
		if err !=nil {
			fmt.Println(err.Error())
		}else {
			_,err=f.Write([]byte("[]"))
		}
		fmt.Println("日志文件不存在，已重新创建")
	}
	var goLine int
	goLine = utils.ReadLog(nowPath,path)
	if len(os.Args)>2{
		goLine ,_=  strconv.Atoi(os.Args[2])
	}
	utils.Read2(path,goLine,nowPath)
}

