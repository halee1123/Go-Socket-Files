# Go-Socket-Files

### 简介:

#### Go-Socket 文件传输

#

#### 执行之前,先配置好ini文件

#
Server.ini文件配置:

[socket]

ipaddress = 0.0.0.0  (公网IP)  （内网测试:127.0.0.1)

port = 8000  (端口自定义)

#

Client.ini文件配置:

ipaddress = 输入你Server端ip地址 （内网测试:127.0.0.1)

port = 8000  (与你Server端口一致)

#
### server端执行:

##### go run server.go


#
### client端执行:

##### go run client.go

##### 请输入文件的完整路径：/Users/xxxx/Desktop/xxxx.ipa


##### 执行完成之后,client端会自动退出...










