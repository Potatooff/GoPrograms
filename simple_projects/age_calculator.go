package main;
import (
	"fmt"
	"time"
)
func find_age(time uint16, birthyear uint16) uint16 {
		result := time - birthyear
		return result
}
func main() {
	var birthyear uint16
	time := time.Now().Year()
	fmt.Print("What is your birth year?\nYou: ")
	fmt.Scan(&birthyear)
	var lol uint8 = find_age(time, birthyear)
	fmt.Printf("You are %d years old!", lol)
}
