package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rd.Intn(8) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, mut *sync.Mutex, ch chan<- Book) {
			if b, ok := queryCache(id, mut); ok {
				ch <- b
			}

			wg.Done()
		}(id, wg, mut, cacheCh)

		go func(id int, wg *sync.WaitGroup, mut *sync.Mutex, ch chan<- Book) {
			if b, ok := queryDatabase(id, mut); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, mut, dbCh)
		time.Sleep(150 * time.Millisecond)

		go func(cacheCh <-chan Book, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from Database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
	}
	wg.Wait()
}

func queryCache(id int, mutex *sync.Mutex) (Book, bool) {
	mutex.Lock()
	b, ok := cache[id]
	mutex.Unlock()
	return b, ok
}

func queryDatabase(id int, mutex *sync.Mutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			mutex.Lock()
			cache[id] = b
			mutex.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
