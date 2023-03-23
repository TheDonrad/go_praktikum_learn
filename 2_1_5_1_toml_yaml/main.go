package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

type Data struct {
	ID     int    `toml:"id"`
	Name   string `toml:"name"`
	Values []byte `toml:"values"`
	Item   int
}

const yamlData = `
id: 101
name: John Doe
item1: 1
values:
- 11
- 22
- 33
`

func main() {
	// допишите код
	// 1) десериализуйте yamlData в переменную типа Data
	out := Data{}
	yaml.Unmarshal([]byte(yamlData), &out)
	// 2) преобразуйте полученную переменную в TOML
	rez, _ := toml.Marshal(out)
	// 3) выведите в консоль результат
	fmt.Println(string(rez))

	//var v Data
	//err := yaml.Unmarshal([]byte(yamlData), &v)
	//if err != nil {
	//	panic(err)
	//}
	//out, err := toml.Marshal(v)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(out))
}
