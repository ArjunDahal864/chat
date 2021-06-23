package main

import (
	"fmt"

	"github.com/arjun/chat/backend/pkg/db"
)

func main() {
	conn, err := db.ConnectDatabase()
	if err != nil {
		panic(err)
	}
	fmt.Println(*conn)

}
