package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

// String должен уметь сериализовать переменную типа в строку.
func (u *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", u.Host, u.Port)
}

// Set связывает переменную типа со значением флага
// и устанавливает правила парсинга для пользовательского типа.
func (u *NetAddress) Set(flagValue string) (err error) {
	addr := strings.Split(flagValue, ":")
	if len(addr) != 2 {
		return errors.New("wrong")
	}
	u.Host = addr[0]
	u.Port, err = strconv.Atoi(addr[1])
	return err
}

// допишите код реализации методов интерфейса

func main() {
	addr := new(NetAddress)
	// если интерфейс не реализован,
	// здесь будет ошибка компиляции
	_ = flag.Value(addr)
	// проверка реализации
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
