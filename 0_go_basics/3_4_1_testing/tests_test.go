package tests

import "testing"

func TestAddPositive(t *testing.T) {
	sum, err := Add(1, 2)
	if err != nil {
		t.Error("unexpected error")
	}
	if sum != 3 {
		t.Errorf("sum expected to be 3; got %d", sum)
	}
}

func TestAddNegative(t *testing.T) {
	_, err := Add(-1, 2)
	if err == nil {
		t.Error("first arg negative - expected error not be nil")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error("second arg negative - expected error not be nil")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
	_, err = Add(0, 0)
	if err == nil {
		t.Error("all arg negative - expected error not be nil")
	}
}

func TestEstimateValue(t *testing.T) {
	// args — описывает аргументы тестируемой функции
	type args struct {
		a int
	}
	// описывает структуру тестовых данных и сами тесты
	tests := []struct {
		name    string // название теста
		args    args   // аргументы
		want    string // ожидаемое значение
		wantErr bool   // должна ли функция вернуть ошибку
	}{
		{
			name: "Test Positive",
			args: args{
				a: 9,
			},
			want:    "small",
			wantErr: false,
		},
		{
			name: "Test Negative 1",
			args: args{
				a: 10,
			},
			want:    "medium",
			wantErr: true,
		},
		{
			name: "Test Negative 2",
			args: args{
				a: 100,
			},
			want:    "big",
			wantErr: true,
		},
	}
	// вызываем тестируемую функцию для каждого тестового случая
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EstimateValue(tt.args.a)
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
