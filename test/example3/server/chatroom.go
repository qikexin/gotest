package main

import (
	"fmt"
	"net"
	"strings"
)
//发送的数据格式事先规定如：来源ip#消息内容
//127.0.0.1:14133#你好
var onlneConns = make(map[string]net.Conn)
var messageQuere = make(chan string,1000)
var quitChan = make(chan bool)

func CheckError(err error){
	if err != nil {
		fmt.Println("error: %s", err.Error())
		panic(err)
		//os.Exit(1)
	}
}
func ProcessInfo(conn net.Conn){
	buf := make([]byte,1024)
	defer conn.Close()

	for {
		numofBytes,err := conn.Read(buf)
		//CheckError(err)
		if err != nil {
			break
		}
		if numofBytes !=0 {
			remoteAddr := conn.RemoteAddr()
			//fmt.Println(remoteAddr)
			fmt.Printf("%s ,has received this message : %s\n",remoteAddr,string(buf[:numofBytes]))
			message := string(buf[:numofBytes])
			messageQuere <- message
		}
	}
}
func ConsumeMessage(){
	for {
		select {
		case message <- messageQuere:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

func doProcessMessage(message string){
	contents := strings.Split(message,"#")
	if len(contents) > 1{
		addr := contents[0]
		sendMessage := contents[1]

		addr = strings.Trim(addr," ")
		if conn,ok := onlneConns[addr];ok {
			_ ,err := conn.Write([]byte(sendMessage)){
				if err != nil {
					fmt.Println("online conns send failure!")
				}
			}
		}
	}
}
func main(){
	listen_socket, err := net.Listen("tcp","127.0.0.1:8080")
	CheckError(err)
	defer listen_socket.Close()
	fmt.Println("server begin !")
	 go ConsumeMessage()
	for {
		conn, err := listen_socket.Accept()
		CheckError(err)
		//存储连接对象到map中（onlineConns中）
		addr := fmt.Sprintf("%s",conn.RemoteAddr())
		onlneConns[addr] = conn
		go ProcessInfo(conn)
	}
}