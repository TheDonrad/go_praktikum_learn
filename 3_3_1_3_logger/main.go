package main

import (
	"bytes"
	"log"
)

func main() {
	var buf bytes.Buffer
	// допишите код
	// 1) создайте переменную типа *log.Logger
	logger := log.New(&buf, "mylog: ", 0)
	logger.Print("Hello, world!")
	logger.Print("Goodbye")

	// должна вывести
	// mylog: Hello, world!
	// mylog: Goodbye
}
