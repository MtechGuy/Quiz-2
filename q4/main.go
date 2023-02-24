package main

import (
	"fmt"
	"sync"
)
func worker (id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d starting \n", id)

	for i := 0; i < 5; i++{
		fmt.Printf("Worker %d working...%d\n",id,i)
	}
	fmt.Printf("Worker %d done \n", id)
}

func main(){
	var wg sync.WaitGroup

	for i :=1; i <= 4; i++{
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All workers finished")
}
