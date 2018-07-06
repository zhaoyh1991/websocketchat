package main

import (
	"github.com/gorilla/websocket"
	"websocketchat/model"
	"net/http"
	"log"
	"fmt"
)

var connections=make(map[string]*websocket.Conn,10)
var msgchan=make(chan model.Messages,10)

func handConnections(writer http.ResponseWriter, request *http.Request) {
	conn,err:=websocket.Upgrade(writer,request,nil,1024,1024)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
		conn.Close()
	}()
	if e,ok:=err.(websocket.HandshakeError);ok {
		log.Fatal("hand shark error"+e.Error())
		panic(e)
	}
	from:=request.FormValue("from")
	connections[from]=conn
	conn.SetCloseHandler(func(code int, text string) error {
		delete(connections,from)
		return nil
	})
	printmap(connections)
	for  {
		print("get value")
		var msg model.Messages
		e:=conn.ReadJSON(&msg)
		if e!=nil{
			panic(e)
		}else {
			msgchan<-msg

		}
	}


	}
func printmap(m map[string]*websocket.Conn) {
	for k,_:=range m{
		fmt.Printf("one of connection is %s\n",k)
	}
}
func dealMessage() {
	for msg:=range msgchan{
		fmt.Print("deal message")
		to:=msg.To
		c:=connections[to]
		if c!=nil{
			c.WriteJSON(msg)
		}else{
			log.Fatal("connection not exit")
		}
	}
}
func main() {
	go dealMessage()
	http.HandleFunc("/ws",handConnections)
	err:=http.ListenAndServe(":8081",nil)
	if err!=nil {
		fmt.Print("server error")
	}

}

