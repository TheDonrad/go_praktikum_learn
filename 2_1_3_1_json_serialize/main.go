package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	Company string `json:"comp,omitempty"`
}

func main() {
	foo := []Data{
		{
			ID:   10,
			Name: "John Doe",
		},
		{
			Name:    "Вася",
			Company: "Яндекс",
		},
	}
	out, err := json.Marshal(foo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))

	var v interface{}
	err = json.Unmarshal([]byte(`[0, 10, 30]`), &v)
	fmt.Printf("%T, %[1]v, %v", v, err)
}
