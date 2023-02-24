package main

import (
	"fmt"
)
func worker(id int, jobs <- chan int, results chan <- int){
	for j := range jobs {
		fmt.Println("Worker", id, "Processing Job", j)
		results <- j*2
	}
}

func main(){
	jobs := make(chan int, 50)
	results := make(chan int, 50)

	for w := 1; w <=4; w++ {
		go worker(w,jobs,results)
	}
	for j:=1; j<=10; j++{
		jobs <-j
	}
	close(jobs)

	for a :=1; a<=10; a++{
		r := <-results
		fmt.Println("Result", a, "is", r)
	}

}