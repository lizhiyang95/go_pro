package client

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	Conn       net.Conn
	flag       int
}

func NewClient(serverIp string, serverPort int) *Client {
	return &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       9999,
	}
}

func (this *Client) Connect() *Client {
	addr := fmt.Sprintf("%s:%d", this.ServerIp, this.ServerPort)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return nil
	}
	this.Conn = conn
	return this
}

// 读取服务器返回的消息
func (this *Client) Read() {

	buf := make([]byte, 4096)
	for {
		n, err := this.Conn.Read(buf)
		if n == 0 {
			fmt.Println("连接已断开")
			return
		}
		//错误不为空
		if err != nil {
			//等于结束符表示连接关闭
			if err == io.EOF {
				fmt.Println("Connection closed by remote peer")
			} else {
				fmt.Println("Other error occurred")
			}
			return
		}
		//n是读取的字节数，截取全部，再去掉'\r\n'
		//msg := strings.TrimRight(string(buf[:n]), "\r\n")
		msg := string(buf[:n])
		fmt.Printf("服务器返回消息：%#v \n", msg)
	}
}

// 向服务器写入消息
func (this *Client) Write() {
	for this.flag != 0 {
		flag := this.menu()
		if flag {
			msg := this.doMenu(this.flag)
			// Send a message to the server
			_, err := this.Conn.Write([]byte(msg))
			if err != nil {
				fmt.Println("Error sending message:", err.Error())
			}
		}
	}
}

// 处理菜单选项
func (this *Client) doMenu(menu int) string {
	var cmd string
	var message string
	switch menu {
	case -1:
		fmt.Println(this.Name + "已退出...")
		os.Exit(0)
	case 1:
		cmd = ""
		fmt.Println("请输入广播消息")
	case 2:
		cmd = "to|"
		fmt.Println("请输入对方用户名")
		var user string
		fmt.Scanln(&user)
		cmd = cmd + user + "|"
		fmt.Println("请输入想发送的内容")
	case 3:
		cmd = "rename|"
		fmt.Println("请输入用户名")
	}
	fmt.Scanln(&message)
	cmd = cmd + message
	return cmd
}

// 展示可选菜单项
func (this *Client) menu() bool {
	var flag int
	menuMap := make(map[int]string)
	menuMap[-1] = "退出"
	menuMap[1] = "发送广播消息"
	menuMap[2] = "发送私聊消息"
	menuMap[3] = "更新名称"

	fmt.Println("请选择模式")
	for index, menu := range menuMap {
		fmt.Printf("%v: %v\n", index, menu)
	}
	var input string
	fmt.Scanln(&input)
	flag, _ = strconv.Atoi(input)
	_, ok := menuMap[flag]
	if ok {
		this.flag = flag
		fmt.Println("已选择：" + menuMap[this.flag] + "模式")
		return true
	} else {
		fmt.Println("该指令不存在")
		return false
	}

}
