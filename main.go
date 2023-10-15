package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// generate random 12-digit ID
	id := ""
	for i := 0; i < 12; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}

	// RID simply abbreviated as "Random ID"
	result := "RID" + id
	fmt.Println("Random 12-digit ID:", result)
}
