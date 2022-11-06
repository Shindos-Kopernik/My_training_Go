package main

import (
	"My_training_Go/My_training_Go/The Art of Development/Lessons_Golang2_Advanced_Playlist2/Lesson1_Rest_API/internal/user"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println("create router")
	router := httprouter.New()
	handler := user.NewHandler()
	handler.Register(router)

	start(router)

}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second, // Цифры беруться империческим путем
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))

}
