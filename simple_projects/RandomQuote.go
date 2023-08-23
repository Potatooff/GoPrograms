package main;
import (
  "fmt";
  "math/rand";
)

func main() {
	// Run answer is also main function just lazy to clean it!
	answer();
}

// Choose a random quote from "quotes"
func randomize()string {
	var answer string;
	quotes := [] string {
		"I love you",
		"I hate you",
		"I like you",
		"I love chickens",
		"I eat chickens",
	}
	RandomQuote := rand.Intn(len(quotes));
	fmt.Printf("\nRandom Quote of the day: %v", quotes[RandomQuote]);
	fmt.Print("\nGenerate another one?\n1- yes\n2- No\nUser:");fmt.Scan(&answer);
	return answer;
}

// Check if answer is 1 or 2 before running randomizer again!
func answer() {
	var answer string;
	fmt.Println("Welcome to A Random Quote randomizer.");
	answer = randomize()
	for answer == "1" {
		answer = randomize();
	}
}
