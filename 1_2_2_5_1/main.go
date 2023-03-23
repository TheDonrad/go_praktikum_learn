package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	v := Abs(3)
	fmt.Println(v)
	users()
	familys()
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
// Покрыть тестами нужно эту функцию.
func Abs(value float64) float64 {
	return math.Abs(value)
}

// User в системе.
type User struct {
	FirstName string
	LastName  string
}

// FullName возвращает фамилию и имя человека.
func (u User) FullName() string {
	return u.FirstName + " " + u.LastName
}

func users() {
	u := User{
		FirstName: "Misha",
		LastName:  "Popov",
	}

	fmt.Println(u.FullName())
}

// Relationship определяет положение в семье.
type Relationship string

// Возможные роли в семье.
const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

// Family описывает семью.
type Family struct {
	Members map[Relationship]Person
}

// Person описывает конкретного человека в семье.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

var (
	// ErrRelationshipAlreadyExists возвращает ошибку, если роль уже занята.
	// Подробнее об ошибках поговорим в девятой теме: «Errors, log».
	ErrRelationshipAlreadyExists = errors.New("relationship already exists")
)

// AddNew добавляет нового члена семьи.
// Если в семье ещё нет людей, создаётся пустой map.
// Если роль уже занята, метод выдаёт ошибку.
func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; ok {
		return ErrRelationshipAlreadyExists
	}
	f.Members[r] = p
	return nil
}

func familys() {
	f := Family{}
	err := f.AddNew(Father, Person{
		FirstName: "Misha",
		LastName:  "Popov",
		Age:       56,
	})
	fmt.Println(f, err)

	err = f.AddNew(Father, Person{
		FirstName: "Drug",
		LastName:  "Mishi",
		Age:       57,
	})
	fmt.Println(f, err)
}
