package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	times()

	sleeps()
}

func sleeps() {

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("hi from AfterFunc")
	})
	fmt.Println("hi")
	time.Sleep(2 * time.Second)
	fmt.Println("goodbye")

	start := time.Now()
	timer := time.NewTimer(2 * time.Second) // создаём таймер
	t := <-timer.C                          // ожидаем срабатывания таймера
	fmt.Println(t.Sub(start).Seconds())     // выводим разницу во времени

	start = time.Now()
	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 10; i++ {
		t := <-ticker.C
		fmt.Println(int(t.Sub(start).Seconds()))
	}
}

func times() {
	// time format
	now := time.Now()
	timeStr := now.Format("1.2.06 3:4:5 -07 MST")
	fmt.Println(timeStr)
	timeStr = now.Format("Mon 02 Jan 2006 15:01:05 MST")
	fmt.Println(timeStr)
	timeStr = now.Format(time.RFC1123)
	fmt.Println(timeStr)
	currentTimeStr := "2021-09-19T15:59:41+03:00"

	// time parse
	layout := time.RFC3339
	currentTime, err := time.Parse(layout, currentTimeStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(currentTime)

	// time compare
	fmt.Println("Is", now, "before", currentTime, "? Answer:", now.Before(currentTime))
	fmt.Println("Is", now, "after", currentTime, "? Answer:", now.After(currentTime))
	fmt.Println("Is", now, "equal", currentTime, "? Answer:", now.Equal(currentTime))

	// time parts
	fmt.Printf("Год: %d. Месяц: %d. Число: %d. День недели: %d\n",
		now.Year(), now.Month(), now.Day(), now.Weekday())
	hour, min, sec := now.Clock()
	fmt.Printf("Время: %d, %d, %d\n", hour, min, sec)
	fmt.Println("Часовой пояс:", now.Location())
	fmt.Println("timestamp в секундах:", now.Unix())
	fmt.Println("timestamp в наносекундах:", now.UnixNano())

	// time manipulation
	fmt.Println(now.Add(time.Second))          // добавить интервал к текущему времени
	fmt.Println(now.Round(time.Hour))          // округлить время до часа
	fmt.Println(now.Truncate(3 * time.Minute)) // округлить время в меньшую сторону
	now = time.Now()
	fmt.Println(now.Truncate(24 * time.Hour)) // округлить до начала дня

	// time zones
	loc, _ := time.LoadLocation("Europe/Berlin")
	fmt.Println(now.In(loc))

	birthday := time.Date(1993, time.November, 26, 0, 0, 0, 0, time.Local)
	hundred := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	days := int(hundred.Sub(birthday).Hours() / 24)
	daysNow := int(time.Until(hundred).Hours() / 24)
	fmt.Printf("Time from birthday: %d. Time from now: %d\n", days, daysNow)
}
