package main


import (
	"fmt"
)

func main() {
	
	var sum = 0;

	for i := 0; i <= 1000; i++ {
		
		if i%3==0 || i%5==0 {
			sum = sum+i
			fmt.Println(i,"is in the list number can be multiply of 3 or 5")
		}	
	}

	fmt.Println("Total sum multiply of 3 or 5 below 1000 is:",sum)
}