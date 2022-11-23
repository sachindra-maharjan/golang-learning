package main

import "fmt"

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
	x := GetBuilder("X")
	y := GetBuilder("Y")

	d := NewDirector(x)
	var house = d.BuildHouse()

	fmt.Printf("X Builder has built %+v \n", house)

	d = NewDirector(y)
	house = d.BuildHouse()
	fmt.Printf("Y Builder has built %+v \n", house)

}
