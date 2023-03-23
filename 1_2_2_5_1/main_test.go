package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name   string
		values float64
		want   float64
	}{
		{
			name:   "zero", // описывается каждый тест
			values: 0,      // значения, которые будет принимать функция
			want:   0,      // ожидаемое значение
		},
		{
			name:   "one",
			values: -3,
			want:   3,
		},
		{
			name:   "with negative values",
			values: 3,
			want:   3,
		},
		{
			name:   "with negative zero",
			values: -2.000001,
			want:   2.000001,
		},
		{
			name:   "a lot of values",
			values: -0.000000003,
			want:   0.000000003,
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			if sum := Abs(tt.values); sum != tt.want {
				t.Errorf("Add() = %v, want %v", sum, tt.want)
			}
		})
	}

}

func TestFullName(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name   string
		values User
		want   string
	}{
		{
			name: "zero", // описывается каждый тест
			values: User{
				FirstName: "Misha",
				LastName:  "Popov",
			}, // значения, которые будет принимать функция
			want: "Misha Popov", // ожидаемое значение
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			if sum := tt.values.FullName(); sum != tt.want {
				t.Errorf("Add() = %v, want %v", sum, tt.want)
			}
		})
	}
}

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Family{
				Members: tt.existedMembers,
			}
			if err := f.AddNew(tt.newPerson.r, tt.newPerson.p); (err != nil) != tt.wantErr {
				t.Errorf("AddNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAbs_assert(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name   string
		values float64
		want   float64
	}{
		{
			name:   "zero", // описывается каждый тест
			values: 0,      // значения, которые будет принимать функция
			want:   0,      // ожидаемое значение
		},
		{
			name:   "one",
			values: -3,
			want:   3,
		},
		{
			name:   "with negative values",
			values: 3,
			want:   3,
		},
		{
			name:   "with negative zero",
			values: -2.000001,
			want:   2.000001,
		},
		{
			name:   "a lot of values",
			values: -0.000000003,
			want:   0.000000003,
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			v := Abs(tt.values)
			assert.Equal(t, tt.want, v)
		})
	}
}

func TestFullName_assert(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name   string
		values User
		want   string
	}{
		{
			name: "zero", // описывается каждый тест
			values: User{
				FirstName: "Misha",
				LastName:  "Popov",
			}, // значения, которые будет принимать функция
			want: "Misha Popov", // ожидаемое значение
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			name := tt.values.FullName()
			assert.Equal(t, name, tt.want)
		})
	}
}

func TestFamily_AddNew_assert(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Family{
				Members: tt.existedMembers,
			}
			err := f.AddNew(tt.newPerson.r, tt.newPerson.p)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Contains(t, f.Members, tt.newPerson.r)
			}
		})
	}
}
