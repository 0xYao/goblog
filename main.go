package main

import (
	"0AlexZhong0/goblog/config"
	"0AlexZhong0/goblog/internal/articles"
	"0AlexZhong0/goblog/internal/users"
	"sync"
)

func main() {
	// start all servers
	config.LoadConfig()
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		articles.Run()
		wg.Done()
	}()

	go func() {
		users.Run()
		wg.Done()
	}()

	wg.Wait()
}
