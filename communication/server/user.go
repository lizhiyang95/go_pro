package server

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	Conn   net.Conn
	Server *Server
	isLive chan bool
}

// 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		Conn:   conn,
		Server: server,
		isLive: make(chan bool),
	}
	//创建用户时启动协程监听
	go user.ListenMessage()

	return user
}

// 监听当前user 的 channel ，有消息写入，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		//chan关闭返回对应类型的零值，比如字符串返回空，
		if msg != "" {
			_, err := this.Conn.Write([]byte(msg))
			if err != nil {
				fmt.Printf("err: %v\n", err)
			}
		}
	}
}

func (this *User) SendMessage(msg string) {
	this.C <- msg
}

func (this *User) SendMessageDirect(msg string) {
	this.Conn.Write([]byte(msg))
}

// 监听当前客户端发送的消息
func (this *User) ReceiveMessages() {
	buf := make([]byte, 4096)
	for {
		n, err := this.Conn.Read(buf)
		if n == 0 {
			this.Server.UserOffline(this)
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
		msg := strings.TrimRight(string(buf[:n]), "\r\n")
		fmt.Println("接收到消息:", msg)
		this.DoMessage(msg)
		this.isLive <- true
	}
}

func (this *User) GetUserIdPrefix() string {
	return "[" + this.Addr + "]" + this.Name + "："
}

// 处理接收到的消息规则
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		sendMsg := "总共" + strconv.Itoa(len(this.Server.OnlineUserMap)) + "在线----"
		for _, onLineUser := range this.Server.OnlineUserMap {
			sendMsg = sendMsg + onLineUser.GetUserIdPrefix() + "在线中"
		}
		this.SendMessage(sendMsg)
		return
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]
		_, ok := this.Server.OnlineUserMap[newName]
		if ok {
			this.SendMessage("该用户名已存在")
		} else {
			this.Server.OnlineUserMapLock.Lock()
			delete(this.Server.OnlineUserMap, this.Name)
			this.Name = newName
			this.Server.OnlineUserMap[this.Name] = this
			this.Server.OnlineUserMapLock.Unlock()
			this.SendMessage("修改成功，新用户名：" + this.Name)
		}
		return
	} else if len(msg) > 3 && msg[:3] == "to|" {
		msg := strings.Split(msg, "|")
		remoteName := msg[1]
		user, ok := this.Server.OnlineUserMap[remoteName]
		if ok {
			index := 2
			if len(msg) <= index {
				msg = append(msg, "")
			}
			user.SendMessage(this.Name + "对你说：" + msg[index])
		} else {
			this.SendMessage("远程用户不存在")
		}
		return
	}
	this.Server.BroadCast(this, msg)
}
