package main
 
import (
	"fmt"
	"time"
)

func sayHello(){
	fmt. Println("Hello from Goroutine!")
}

func main(){
	go sayHello()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Exiting main function")
}