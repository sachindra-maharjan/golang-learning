package main

//Singleton
/*func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func(wg *sync.WaitGroup) {
			GetInstance()
			wg.Done()
		}(wg)

		go func(wg *sync.WaitGroup) {
			GetInstance()
			wg.Done()
		}(wg)
	}
	time.Sleep(100 * time.Millisecond)
	wg.Wait()
}*/

// Builder
func main() {

}
