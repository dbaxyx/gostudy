package main

import (
	"log"
	"os"
)

func main() {
	//log.Println("这是一条很普通的日志。")
	//v := "很普通的"
	//log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会出发fatal的日志。")
	//log.Panicln("这是一条会出发panic的日志。")

	//log.Println("-----------------------------")
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.Println("这是一条很普通的日志")
	//
	//
	//log.Println("--------配置日志输出位置--------")
	//logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil {
	//	fmt.Println("open log file faied, err:", err)
	//	return
	//}
	//log.SetOutput(logFile)
	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.SetPrefix("[小王子]")
	//log.Println("这是一条很普通的日志")

	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志")

}
