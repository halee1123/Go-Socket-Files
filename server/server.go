
package main

// Server
// go get github.com/gookit/ini/v2

import (
	"fmt"
	"github.com/gookit/ini/v2"
	"net"
	"os"
)

// init函数
func init() {
	// 获取当前路径
	str, _ := os.Getwd()
	// 在当前路径下创建cLIent.ini文件
	var filePath = str + "/Server.ini"

	// ini文件路径
	_, err := os.Stat(filePath)

	if err == nil {
		return
		//fmt.Printf(" 当前路径:%s/%s 文件存在\n", str, err)
	}
	if os.IsNotExist(err) {

		_, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			return
		}
	}
}


// ReadServeriniFile // 读取ini文件
func ReadServeriniFile(Text string) string {
	// 获取当前路径

	err := ini.LoadExists("./Server.ini")
	if err != nil {
		panic(err)
	}
	value := ini.String(Text)
	//fmt.Println(value)
	return value
}



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
	// 获取ini文件数据
	ipaddress := ReadServeriniFile("socket.ipaddress")
	port := ReadServeriniFile("socket.port")
	ipAndPort := ipaddress + ":" + port

	if ipAndPort != "" && port != "" {

		// 创建用于监听的socket
		//listener, err := net.Listen("tcp", "127.0.0.1:8000")
		listener, err := net.Listen("tcp", ipAndPort)
		if err != nil {
			fmt.Println(" net.Listen err:", err)
			return
		}
		defer func(listener net.Listener) {
			err := listener.Close()
			if err != nil {

			}
		}(listener)

		fmt.Printf("Server地址与端口:[ %s:%s ] ---> Server端启动成功，等待服务端发送文件！\n",ipaddress,port)

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
	}else{
		fmt.Printf("ip地址与端口不能为空,请查看Server.ini文件是否填写参数...")
	}

}












