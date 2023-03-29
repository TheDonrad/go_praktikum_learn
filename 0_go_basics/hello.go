package main

import (
	"encoding/json"
	"fmt"
	"helloWorld/randbyte"
	"log"
	"math"
	"reflect"
	"rsc.io/quote"
	"strings"
	"time"
	"unicode/utf8"
)

func main() {
	//defer metricTime(time.Now())
	//variables()
	//mainIfs()
	//generations(time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC))
	//fizzbuzz()
	//links()
	//readLine.ReadLine()
	//collections()
	//sliceFromString()
	//maps()
	//pares := find([]int{1, 5, 2, 4, 3, 3, 0}, 6)
	//fmt.Println(pares)
	//removeDuplicates()
	//toJson()
	//lo()
	//LogTest()
	//callRandByte()
	// создаём strings.Builder
	//singBatman()
	//limReader()
	//Mul("str", 2)
	//reflection()
	//reflections.Reflections()
	//reflections.ReflectHW()
	//panics.Panic()
	y := 1

	factors := map[int]int{
		1: 2,
		2: 3,
		3: 1,
	}

	for _, val := range factors {
		y *= factors[val]
	}
	fmt.Println(y)
}

func reflection() {
	var varBool *bool
	fmt.Println(reflect.ValueOf(varBool).Kind()) // ptr — указатель
	fmt.Println(reflect.ValueOf(varBool).Type()) // *bool — указатель на bool

	var varFloat float32
	fmt.Println(reflect.ValueOf(varFloat).Kind()) // float32
	fmt.Println(reflect.ValueOf(varFloat).Type()) // float32
	fmt.Println(reflect.ValueOf(varBool).IsNil()) // true

	trueVal := true
	reflect.ValueOf(&varBool).Elem().Set(reflect.ValueOf(&trueVal))

	fmt.Println(reflect.ValueOf(varBool).IsNil())       // false
	fmt.Println(reflect.ValueOf(varBool).Elem().Bool()) // true — получить значение через рефлексию
	fmt.Println(*varBool)

	var varMap map[string]int
	fmt.Println(reflect.ValueOf(varMap).Kind()) // map
	fmt.Println(reflect.ValueOf(varMap).Type()) // map[string]int

	varStruct := struct{ Value int }{}
	fmt.Println(reflect.ValueOf(varStruct).Kind()) // struct
	fmt.Println(reflect.ValueOf(varStruct).Type()) // struct { Value int }

	varInt := 100
	varIntValue := reflect.ValueOf(varInt)
	fmt.Println(varIntValue.IsZero()) // false
	fmt.Println(varIntValue.Int())    // 100

	var varPtr *int

	varPtrValue := reflect.ValueOf(varPtr)
	fmt.Println(varPtrValue.IsNil())  // true
	fmt.Println(varPtrValue.IsZero()) // true

}

func Mul(a any, b int) interface{} {
	switch c := a.(type) {
	case int:
		return c * b
	case string:
		return strings.Repeat(c, b)
	case fmt.Stringer:
		return strings.Repeat(c.String(), b)
	default:
		return nil
	}
}

func singBatman() {
	w := strings.Builder{}

	for i := 0; i < 50; i++ {
		// функция fmt.Fprintf принимает аргументом io.Writer
		// благодаря этому можно записывать форматированный вывод
		fmt.Fprintf(&w, "%v", math.NaN())
	}

	w.Write([]byte("... BATMAN!"))
	// выводим собранную строку
	fmt.Printf("%s\n", &w)
}

func callRandByte() {
	// создаём генератор случайных чисел
	generator := randbyte.New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.

	buf := make([]byte, 24)

	for i := 0; i < 5; i++ {
		n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
		fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
	}
}

func links() {
	var a int8 = 5
	p := &a

	fmt.Println(a, p)
}

func variables() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	var a, b int
	a, b = 5, 10 // a == 5, b == 10
	a, b = b, a  // swap: a == 10, b == 5
	fmt.Println(a, b)

	var c bool
	c = true
	c = false
	fmt.Println(c)

	var d string
	d = "abc"
	println(len(d), d[0]) // 3 97

	var rus string
	rus = "абц"
	println(len(rus), rus[5])            // 6 134
	println(utf8.RuneCountInString(rus)) // 3
	println(string(d[2]))

	var stringFormattedVar string
	// следующие выражения равнозначны
	stringFormattedVar = "Hello,\nworld!\n\n\t\t\"quote!\""
	println(stringFormattedVar)
	stringFormattedVar = `Hello,
world!

        "quote!"`
	println(stringFormattedVar)

	ver := "v0.0.1"
	var id int
	pi := float32(3.1415)

	fmt.Println("ver =", ver, "id =", id, "pi =", pi)

	const (
		one = iota*2 + 1
		three
		five
		seven
		nine
		eleven
	)
	fmt.Println(one, three, five, seven, nine, eleven)
}

func mainIfs() {
	a := 1

	if a == 1 {
		fmt.Println("if1")
	} else if a == 2 {
		fmt.Println("if2")
	} else {
		fmt.Println("if3")
	}
	a = 5
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("Default case")
	}

}

func generations(date time.Time) {
	switch y := date.Year(); {
	case y >= 1946 && y < 1965:
		fmt.Println("Привет, бумер!")
	case y >= 1965 && y < 1981:
		fmt.Println("Привет, представитель X!")
	case y >= 1981 && y < 1997:
		fmt.Println("Привет, миллениал!")
	case y >= 1997 && y < 2013:
		fmt.Println("Привет, зумер!")
	case y >= 2013:
		fmt.Println("Привет, альфа!")
	default:
		fmt.Println("Здравствуйте!")
	}
}

func fizzbuzz() {
	timeStart := time.Now()
	start := timeStart.UnixMilli()
	limit := 1000

	for i := 0; i <= limit; i++ {
		if i%5 == 0 {
			if i%3 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Buzz")
			}
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
	timeFinish := time.Now()
	finish := timeFinish.UnixMilli()
	println(finish - start)

}

func collections() {
	//arrays()

	//weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	//workDaysSlice := weekTempArr[:5]
	//weekendSlice := weekTempArr[5:]
	//fromTuesdayToThursDaySlice := weekTempArr[1:4]
	//weekTempSlice := weekTempArr[:]
	//
	//weekTempArr[1] = 20
	//weekTempArr[6] = 18
	//fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))
	//fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))
	//fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice))
	//fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))

	var _ []int // лайс [], базовый массив []
	mySlice := make([]int, 5, 10)
	fmt.Println(mySlice)
	sl := []int{1, 2, 3} // [1 2 3]
	sl = sl[:len(sl)-1]
	sl = append(sl, 10)
	fmt.Println(sl)

	dim := 100
	s := make([]int, 0, dim)
	// заполняем слайс числами
	for i := 0; i < dim; i++ {
		s = append(s, i+1)
	}
	// оставляем первые и последние 10 элементов
	s = append(s[:10], s[dim-10:]...)
	dim = len(s)
	// разворачиваем слайс
	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}
	fmt.Println(s)

}

func sliceFromString() {
	input := "The quick brown 狐 jumped over the lazy 犬"
	// Get Unicode code points.
	n := 0
	// создаём слайс рун
	runes := make([]rune, len(input))
	// добавляем руны в слайс
	for _, r := range input {
		runes[n] = r
		n++
	}
	// убираем лишние нулевые руны
	runes = runes[0:n]
	// разворачиваем
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	// преобразуем в строку UTF-8.
	output := string(runes)
	fmt.Println(output)
}

func arrays() {
	thisWeekTemp := [7]int{6: 11, 2: 3}

	fmt.Println(thisWeekTemp)

	sumTemp := 0

	for id, temp := range &thisWeekTemp {
		sumTemp += temp // не изменяет значения
		println(id)
	}

	average := sumTemp / len(thisWeekTemp)
	println(average)

	for i := range &thisWeekTemp {
		thisWeekTemp[i] = 1 // изменяет значения
	}
}

func maps() {
	m := make(map[string]string)
	m["foo"] = "bar"
	m["ping"] = "pong"

	for k, v := range m {
		m[k] = "key:" + k + ";value:" + v
	}
	fmt.Println(m)

	price := make(map[string]int)
	price["хлеб"] = 50
	price["молоко"] = 100
	price["масло"] = 200
	price["колбаса"] = 500
	price["соль"] = 20
	price["огурцы"] = 200
	price["сыр"] = 600
	price["ветчина"] = 700
	price["буженина"] = 900
	price["помидоры"] = 250
	price["рыба"] = 300
	price["хамон"] = 1500

	for k, v := range price {
		if v > 500 {
			fmt.Println(k)
		}
	}

	order := [4]string{"хлеб", "буженина", "сыр", "огурцы"}

	sum := 0
	for i := range &order {
		sum += price[order[i]]
	}
	fmt.Println(sum)
}

func find(arr []int, k int) []int {
	// Создадим пустую map
	m := make(map[int]int)
	// будем складывать в неё индексы массива, а в качестве ключей использовать само значение
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // если значение k-a уже есть в массиве, значит, arr[j] + arr[i] = k и мы нашли, то что нужно
			return []int{i, j}
		}
		// если искомого значения нет, то добавляем текущий индекс и значение в map
		m[a] = i
	}
	// не нашли пары подходящих чисел
	return nil
}

func removeDuplicates() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	output := make([]string, 0)
	inputSet := make(map[string]struct{}, len(input))
	for _, v := range input {
		if _, ok := inputSet[v]; !ok {
			output = append(output, v)

		}
		inputSet[v] = struct{}{}
	}
	fmt.Println(output)
}

type Person struct {
	Name        string
	Email       string
	DateOfBirth time.Time
}

func toJson() {
	man := Person{Name: "Alex", Email: "alex@yandex.ru"}
	man.DateOfBirth = time.Date(200, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	st, err := json.Marshal(man)
	if err == nil {
		fmt.Printf("Man %v", string(st))
	} else {
		fmt.Println(err)
		log.Fatalln("unable marshal to json")
	}
}

type Response struct {
	header struct {
		code    int
		message string
	}
	data []struct {
		_type string
		id    int
	}
}

func ReadResponse(rawResp string) (Response, error) {
	stand := Response{}
	byteResp := []byte(rawResp)
	err := json.Unmarshal(byteResp, stand)
	return stand, err
}

const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}
`

type (
	Response1 struct {
		Header ResponseHeader `json:"header"`
		Data   ResponseData   `json:"data,omitempty"`
	}

	ResponseHeader struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	}

	ResponseData []ResponseDataItem

	ResponseDataItem struct {
		Type       string                `json:"type"`
		Id         int                   `json:"id"`
		Attributes ResponseDataItemAttrs `json:"attributes"`
	}

	ResponseDataItemAttrs struct {
		Email      string `json:"email"`
		ArticleIds []int  `json:"article_ids"`
	}
)

func ReadResponse1(rawResp string) (Response1, error) {
	resp := Response1{}
	if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
		return Response1{}, fmt.Errorf("JSON unmarshal: %w", err)
	}

	return resp, nil
}

func ReadResponseMain() {
	resp, err := ReadResponse(rawResp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", resp)
}

func metricTime(start time.Time) {
	fmt.Println(time.Now().Sub(start))
}
