package main;
import (
	"fmt"
	"time"
)
func find_age(time int, birthyear int) int {
		result := time - birthyear
		return result
}
func main() {
	var birthyear int
	time := time.Now().Year()
	fmt.Print("What is your birth year?\nYou: ")
	fmt.Scan(&birthyear)
	var lol int = find_age(time, birthyear)
	fmt.Printf("You are %d years old!", lol)
	if lol >= 18 {
		fmt.Println("\nYou're an adult!")
	} else if lol < 18 { 
		fmt.Println("\nYou're young!")
	} else {
		fmt.Println("impossible!")
	}
}
