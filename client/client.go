package main

// Client
// go get github.com/gookit/ini/v2

import (
	"fmt"
	"github.com/gookit/ini/v2"
	"io"
	"net"
	"os"
)

// init函数
func init() {

	// 获取当前路径
	str, _ := os.Getwd()

	// 在当前路径下创建cLIent.ini文件
	var filePath = str + "/Client.ini"

	// 当前ini文件路径
	_, err := os.Stat("Client.ini")

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

	err := ini.LoadExists("./Client.ini")
	if err != nil {
		panic(err)
	}
	value := ini.String(Text)
	return value
}

// SendFile 发送文件到服务端
func SendFile(filePath string, fileSize int64, conn net.Conn,revData string) {

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	var count int64

	for {
		buf := make([]byte, 2048)
		//读取文件内容
		n, err := f.Read(buf)
		if err != nil && io.EOF == err {
			fmt.Printf("Server端发来的消息:[ %s ]",revData)

			//告诉服务端结束文件接收
			_, err := conn.Write([]byte("finish"))
			if err != nil {
				return 
			}
			return
		}
		//发送给服务端
		_, err = conn.Write(buf[:n])
		if err != nil {
			return
		}

		count += int64(n)
		sendPercent := float64(count) / float64(fileSize) * 100
		value := fmt.Sprintf("%.2f", sendPercent)

		//打印上传进度
		fmt.Println("文件上传：" + value + "%")
	}
}

//main
func main() {

	// 获取ini文件数据参数
	ipaddress := ReadServeriniFile("socket.ipaddress")
	port := ReadServeriniFile("socket.port")
	ipAndPort := ipaddress + ":" + port

	// 判断ip和端口是否为空
	if ipaddress != "" && port != "" {
		fmt.Print("请输入文件的完整路径：")
		//创建切片，用于存储输入的路径
		var str string
		_, err := fmt.Scan(&str)
		if err != nil {
			return
		}
		//获取文件信息
		fileInfo, err := os.Stat(str)
		if err != nil {
			fmt.Println(err)
			return
		}
		//创建客户端连接
		conn, err := net.Dial("tcp", ipAndPort)
		//conn, err := net.Dial("tcp", "127.0.0.1:8000")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		//文件名称
		fileName := fileInfo.Name()

		//文件大小
		fileSize := fileInfo.Size()

		//发送文件名称到服务端
		_, err = conn.Write([]byte(fileName))
		if err != nil {
			return
		}
		buf := make([]byte, 2048)

		//读取服务端内容
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		revData := string(buf[:n])
		if revData == fileName + " 文件已接收完毕..." {

			//发送文件数据
			SendFile(str, fileSize, conn,revData)
		}

	}else{
		fmt.Printf("ip地址与端口不能为空,请查看Client.ini文件是否填写参数...")
	}


}