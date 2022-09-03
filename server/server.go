
package main

import (
	"fmt"
	"net"
	"os"
)


// 接收文件
func recvFile(conn net.Conn, fileName string) {

	// 按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			
		}
	}(f)

	// 从 网络中读数据，写入本地文件
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Printf("Client传过来的:[ %s ] 文件接收完成...\n",fileName)
			return
		}
		// 写入本地文件，读多少，写多少。
		_, err2 := f.Write(buf[:n])
		if err2 != nil {
			return 
		}
	}
}

func main() {
	// 创建用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(" net.Listen err:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			
		}
	}(listener)

	fmt.Println("接收端启动成功，等待发送端发送文件！")

	for {
		// 阻塞监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(" listener.Accept() err:", err)
			return
		}

		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		// 获取文件名，保存
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(" conn.Read err:", err)
			return
		}
		// 获取文件名称..
		fileName := string(buf[:n])

		// 回写 ok 给发送端
		_, err = conn.Write([]byte(fileName + " 文件已接收完毕..."))
		if err != nil {
			return
		}

		// 获取文件内容
		recvFile(conn, fileName)
	}
}












