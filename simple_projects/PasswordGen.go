package main

// modules 
import (
	"fmt"
	"math/rand"
)

// Generate a password 
func GeneratePassword(length int) {	
	const ( 
	characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits = "1234567890"
	punctuations = "!#$%&'()*+,-./:;<=>?@[\\]^_`{|}~\""
	everything = characters + digits + punctuations
)
	password := make([]byte, length);
	for i := 0; i < length; i++ {
		password[i] = everything[rand.Intn(len(everything))];
	} 
	fmt.Printf("Generated Password: %s", password);
}

// Run everything in main function
func main() {
	fmt.Print("Welcome To a Password Generator in Golang")
	answer := "1";
	var lenght int;
	for answer == "1" {	
		lenght = asker_length();
		GeneratePassword(lenght);
		answer = asker_again();
}
	fmt.Print("Bye!");
}

// This method ask if user want to generate a password
func asker_again() string {
	var answer string;
	fmt.Print("\nDo you want to generate a password?\n1- Yes\n2- No\nUser: "); fmt.Scan(&answer);
	return string(answer);
}

// This method ask length!
func asker_length() int {
	var lenght int;
	fmt.Print("\nChoose your password length\nUser: ");fmt.Scan(&lenght);
	return lenght;
}
