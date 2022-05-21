package common

import (
	"bufio"
	"fmt"
	"net"
	"sync"
	"time"
)

//声明服务器启用状态,预留状态变量
var SERVER_STATUS int8 = STATUS_UN_START

//状态变量枚举
const (
	STATUS_UN_START            int8 = 0
	STATUS_START_START_ING     int8 = 1
	STATUS_START_START_SUCCESS int8 = 2
	STATUS_CLOSE               int8 = 3
)

//启动server
func StartTcpServer(listenPort string, sw *sync.WaitGroup) {
	go func() {
		//将状态更新为正在启动
		SERVER_STATUS = STATUS_START_START_ING
		//监听tcp协议
		listen, err := net.Listen("tcp", "0.0.0.0:"+listenPort)
		if err != nil {
			fmt.Println("listen failed,err:", err)
		}
		if listen != nil {
			fmt.Println("tcp port " + listenPort + " listen success.")

		} else {
			return
		}
		//将状态改为启动成功
		SERVER_STATUS = STATUS_START_START_SUCCESS
		defer func() {
			//如果发生了异常,则直接关闭服务
			SERVER_STATUS = STATUS_CLOSE
			if listen != nil {
				_ = listen.Close()
			}
		}()
		for {
			//关闭监听
			//如果全局变量变为已关闭,则退出循环
			if SERVER_STATUS == STATUS_CLOSE {
				break
			}
			//接收客户端连接

			conn, err := listen.Accept()
			if err != nil {
				continue
			}
			//新开协程处理客户端数据
			go tcpProcess(conn)
		}
		sw.Done()
	}()

}

func tcpProcess(conn net.Conn) {
	t := time.Now().Format("2006-01-02 15:04:05")
	LogSuccess("[+]" + t + "   " + conn.RemoteAddr().String() + " connect " + conn.LocalAddr().String())

	//如果发生异常或者结束栈,则关闭此连接
	defer func() {
		conn.Close()
	}()

	for true {
		//关闭监听
		if SERVER_STATUS == STATUS_CLOSE {
			break
		}
		//读取数据
		reader := bufio.NewReader(conn)
		//声明一个二进制数组
		var buf [128]byte
		//读取
		n, err := reader.Read(buf[:])
		if err != nil {
			conn.Close()
			return
		}
		//将二进制数据转为string
		recvStr := string(buf[:n])
		t2 := time.Now().Format("2006-01-02 15:04:05")
		LogSuccess("[*]" + t2 + "   " + conn.RemoteAddr().String() + " send to " + conn.LocalAddr().String() + "\r\n" + recvStr)

		//发送数据到客户端
		conn.Write([]byte(""))
	}
}
