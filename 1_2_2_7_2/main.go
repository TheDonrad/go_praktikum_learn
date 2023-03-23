package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var cars = map[string]string{
	"id1": "Renault",
	"id2": "BMW",
	"id3": "VW",
	"id4": "Audi",
}

func main() {
	r := chi.NewRouter()
	// зададим встроенные middleware, чтобы улучшить стабильность приложения
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/cars", func(rw http.ResponseWriter, r *http.Request) {
		carsList := carsListFunc()
		_, err := io.WriteString(rw, strings.Join(carsList, ","))
		if err != nil {
			panic(err)
		}
	})

	r.Get("/car/{carID}", func(rw http.ResponseWriter, r *http.Request) {
		carID := chi.URLParam(r, "carID")
		if carID == "" {
			http.Error(rw, "carID param is missed", http.StatusBadRequest)
			return
		}
		rw.Write([]byte(carFunc(carID)))
	})

	r.Route("/{brand}", func(r chi.Router) {
		r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
			// код вывода всех автомобилей бренда
			brand := chi.URLParam(r, "brand")
			rw.Write([]byte(brand))
		})
		r.Get("/{model}", func(rw http.ResponseWriter, r *http.Request) {
			// код вывода определённой модели
			// параметр brand будет доступен
			brand := chi.URLParam(r, "brand")
			model := chi.URLParam(r, "model")
			rw.Write([]byte(brand + " " + model))
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return ""
}

func TimerTrace(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// перед началом выполнения функции сохраняем текущее время
		tStart := time.Now()
		// выполняем основную функцию
		next(w, r)
		// после завершения замеряем время выполнения функции
		tEnd := time.Since(tStart)
		// выводим результат в формате 7h23m0.5s
		fmt.Printf("duration for a request %s\r\n", tEnd)
	}
}
