package main

import (
	"flag"
	"fmt"
	"go_pro/communication/client/client"
)

var serverIp string
var serverPort int

// 从终端接收参数 格式是-ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器的ip地址（默认值是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器的ip地址（默认值是8888）")
	flag.Parse()
}

func main() {
	client := client.NewClient(serverIp, serverPort)
	client = client.Connect()
	if client == nil {
		fmt.Println(">>>>>>>>>连接服务器失败")
		return
	}
	fmt.Println(">>>>>>>>>服务器连接成功")
	defer client.Conn.Close()
	go client.Read()
	go client.Write()
	select {}
}
