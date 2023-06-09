package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	// второе возвращаемое значение — это дочерний контекст, который будет отменён
	// при первой ошибке, возвращённой обработчиками, или при завершении g.Wait()
	//
	// так вы можете получать сигнал об остановке в своём коде,
	// но в данном примере контекст не используется

	hostsToCheck := []string{
		"https://yandex.ru",
		"https://eda.yandex.ru",
		"https://lavka.yandex.ru",
	}
	for _, hostToCheck := range hostsToCheck {
		log.Println("checking", hostToCheck)

		// тело функции healthCheck вставлено анонимной функцией для компактности
		hostToCheck := hostToCheck
		g.Go(func() error {
			resp, err := http.Get(hostToCheck)
			if err != nil {
				return fmt.Errorf("healthcheck failed: %w", err)
			}
			if resp.StatusCode != http.StatusOK {
				return errors.New("healthcheck failed: status not ok")
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Println(err)
		return
	}

	log.Println("successful healthcheck")
}
