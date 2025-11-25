package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
)

type ChatServer struct {
	Messages []string   // stores the chat history
	mu       sync.Mutex // ensures safe access when multiple clients send messages
}

type Message struct {
	Sender  string
	Content string
}

type Reply struct {
	History []string
}

func (cs *ChatServer) SendMessage(msg Message, reply *Reply) error { //appends a new message and returns the full chat history.
	cs.mu.Lock()
	defer cs.mu.Unlock() // to prevent deadlocks

	formatted := fmt.Sprintf("%s: %s", msg.Sender, msg.Content)
	cs.Messages = append(cs.Messages, formatted)

	// Copy the chat history to avoid sharing internal slice reference (more secure)
	reply.History = append([]string(nil), cs.Messages...)
	return nil
}

func main() {
	chatServer := new(ChatServer)

	// Register the ChatServer allows the servers methods to be called by clients
	err := rpc.Register(chatServer)
	if err != nil {
		log.Fatal("Error registering ChatServer:", err)
	}

	// Start listening for TCP connections on port 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listener error:", err)
	}
	defer listener.Close()

	fmt.Println("Server running on port 1234...")

	// Accept client connections forever
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		// (go) handles each client in a separate goroutine to ensure multiclients
		go rpc.ServeConn(conn)
	}
}
