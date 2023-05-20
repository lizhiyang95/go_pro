package server

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	//在线用户列表
	OnlineUserMap map[string]*User
	//读写锁，可同时拥有读锁，但只有一个写锁
	OnlineUserMapLock sync.RWMutex
	//全部消息广播
	Message chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:            ip,
		Port:          port,
		OnlineUserMap: make(map[string]*User),
		Message:       make(chan string),
	}
}

// 用户上线
func (this *Server) UserOnline(user *User) {
	this.OnlineUserMapLock.Lock()
	this.OnlineUserMap[user.Name] = user
	this.OnlineUserMapLock.Unlock()
	fmt.Printf("this.OnlineUserMap: %v\n", this.OnlineUserMap)
	this.BroadCast(user, "已上线")
}

// 用户下线
func (this *Server) UserOffline(user *User) {
	this.OnlineUserMapLock.Lock()
	delete(this.OnlineUserMap, user.Name)
	this.OnlineUserMapLock.Unlock()
	fmt.Printf("this.OnlineUserMap: %v\n", this.OnlineUserMap)
	this.BroadCast(user, "已下线")
}

// 监听Message消息，循环发送给全部在线用户
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message
		this.OnlineUserMapLock.Lock()
		for _, user := range this.OnlineUserMap {
			user.SendMessage(msg)
		}
		this.OnlineUserMapLock.Unlock()
	}
}

// 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := user.GetUserIdPrefix() + msg
	this.Message <- sendMsg
}

// 判断当前客户端是否活跃
func (this *Server) IsUserLive(user *User) {
	for {
		select {
		case <-user.isLive:

		case <-time.After(time.Second * 3000):
			user.SendMessageDirect("你已被踢出")
			defer close(user.C)
			defer close(user.isLive)
			defer user.Conn.Close()
			return
		}

	}
}

// 当用户连接处理的业务
func (this *Server) Handler(conn net.Conn) {
	fmt.Printf("链接成功conn: %v\n", conn)
	//用户上线添加OnlineUser中
	user := NewUser(conn, this)
	this.UserOnline(user)
	//监听当前用户客户端发送消息
	go user.ReceiveMessages()
	go this.IsUserLive(user)
}

func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		panic(err)
	}
	fmt.Println("服务器已启动。。。")

	//close listener socket
	defer listener.Close()

	//监听message
	go this.ListenMessage()

	for {
		//accept
		conn, err2 := listener.Accept()
		if err2 != nil {
			fmt.Printf("err2: %v\n", err2)
			continue
		}
		//do handle
		go this.Handler(conn)
	}

}
