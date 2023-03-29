package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Mutex
	var w sync.WaitGroup

	bisa := []string{"bisa1", "bisa2", "bisa3"}
	coba := []string{"coba1", "coba2", "coba3"}

	ch := make(chan []string)

	go func() {
		for {
			ch <- bisa
			ch <- coba
		}
	}()

	//GOROUTINE keduanya menampilkan secara acak
	for i := 1; i <= 4; i++ {
		w.Add(1)
		go func(id int) {

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			w.Done()
		}(i)
	}

	for i := 1; i <= 4; i++ {
		w.Add(1)
		go func(id int) {

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			w.Done()
		}(i)
	}
	fmt.Println("GOROUTINE keduanya menampilkan secara acak")
	w.Wait()

	//GOROUTINE keduanya menampilkan secara rapih
	//(gunakan mutex golang dengan fungsi lock, dan unlock)
	for i := 1; i <= 4; i++ {
		w.Add(1)
		go func(id int) {
			m.Lock()
			defer m.Unlock()

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			w.Done()
		}(i)
	}

	for i := 1; i <= 4; i++ {
		w.Add(1)
		go func(id int) {
			m.Lock()
			defer m.Unlock()

			arr := <-ch

			fmt.Printf("%v %d\n", arr, id)
			w.Done()
		}(i)
	}
	fmt.Println("\nGOROUTINE keduanya menampilkan secara rapih")
	w.Wait()

}
