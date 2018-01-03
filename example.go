package main

import (
	"fmt"
	"os"

	"github.com/akshaykarle/mongodb-atlas-go/mongodb"
)

func main() {
	username := os.Args[0]
	password := os.Args[1]
	gid := os.Args[2]
	client := mongodb.New(username, password)
	clusters, resp, err := client.Cluster.List(gid)
	fmt.Println(clusters)
	fmt.Println(resp)
	fmt.Println(err)
}
