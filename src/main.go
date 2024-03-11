package main

import (
	"encounters-service/model"
	"fmt"
)

func main() {
	enc := model.Encounter{
		AuthorId: 10,
	}
	fmt.Println(enc)
}
