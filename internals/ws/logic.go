package ws

import (
	"encoding/json"
	"log"

	wsmodels "github.com/suhas-developer07/GuessVibe-Server/internals/models/Websocket_model"
)

func HandleIncomingMessage(c *Client, msg []byte){
	var input wsmodels.ClientMessage

	if err := json.Unmarshal(msg,&input);err!=nil{
		log.Println("invalid ws msg",err)
		return 
	}

	switch input.Type{
	case "init":
	//TODO: here i should handle init message from client when he connects first time to socket 
	// send this request to redis and grpc server

	case "answer":
		//user answered a questino
		//TODO:send this request to grpc server and get the next question
	}

}