package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

//IServer的接口实现，定义一个sERVER的服务器模块

type Server struct{
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPversion string
	//服务器监听的IP
	IP string
	//服务器监听的port
	Port int
}

func (s *Server)Start(){
	fmt.Println("[Start] Server Listener at IP %s,Port %d",s.IP,s.Port)
	go func(){
		//1.获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPversion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error",err)
			return
		}

		//2.监听服务器的地址
		listener, err := net.ListenTCP(s.IPversion, addr)
		if err != nil {
			fmt.Println("listern ",s.IPversion," err ",err)
			return
		}
		fmt.Println("start zinx server succ,",s.Name," succ, Listenning..")
		//3.阻塞的等待客户端链接，处理客户端链接业务
		for {
			conn,err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err",err)
				continue
			}
			//已经与客户端建立键列，做一些业务，做一些基本最大512字节长度的回显业务
			go func(){
				for {
					buf := make([]byte,512)
					cnt,err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err",err)
						continue
					}
					//回显功能
					if _,err := conn.Write(buf[:cnt]);err != nil {
						fmt.Println("write back buf err",err)
						continue
					}
				}
			}()
		}
	}()
	fmt.Println("hello world" )
}


func(s *Server)Stop(){

}

func(s *Server)Serve(){
	//启动服务功能
	s.Start()
	//todo 做一些启动服务器之后的额外业务
	//阻塞
	select{}
}

//初始化server的模块
func NewServer(name string)ziface.IServer{
	s := &Server{
		Name:      name,
		IPversion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
