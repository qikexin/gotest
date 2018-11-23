package client

import (
	"fmt"
	"net"
)

func CheckError(err error){
	if err != nil {
		fmt.Println("error: %s", err.Error())
		panic(err)
		//os.Exit(1)
	}
}

func main(){
	conn, err := net.Dial("tcp","127.0.0.1:8080")
	CheckError(err)
	defer conn.Close()

	conn.Write([]byte("hello"))
	fmt.Println("has sent the messessge " )
}