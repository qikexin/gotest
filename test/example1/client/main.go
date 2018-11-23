package main

import (
	"net"
	"fmt"
	"os"
	"strings"
	"bufio"
)

inport (
"bufio"
"fmt"
"net"
"os"
"strings"
)
func main(){
	conn,err := net.Dial("tcp","localhost:50000")
	if err != nil {
		fmt.Println("error dialing",err.Error())
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input,_  := inputReader.ReadString('\n')
		trimmerInput := strings.Trim(input,"\r\n")
		if trimmerInput == "Q"{
			return
		}
		_,err := conn.Write([]byte(trimmerInput))
		if err != nil {
			return
		}
	}
}