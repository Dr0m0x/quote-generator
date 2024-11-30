package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var quotes = []string{
	"Trying to define yourself is like trying to bite your own teeth. - Alan Watts",
	"You are the universe experiencing itself. - Alan Watts",
	"This is the real secret of life—to be completely engaged with what you are doing in the here and now. And instead of calling it work, realize it is play. - Alan Watts",
	"The meaning of life is just to be alive. It is so plain and so obvious and so simple. And yet, everybody rushes around in a great panic as if it were necessary to achieve something beyond themselves. - Alan Watts",
	"You don’t look out there for God, something in the sky, you look in you. - Alan Watts",
	"The only way to make sense out of change is to plunge into it, move with it, and join the dance. - Alan Watts",
	"Muddy water is best cleared by leaving it alone. - Alan Watts",
	"Man suffers only because he takes seriously what the gods made for fun. - Alan Watts",
	"We seldom realize, for example, that our most private thoughts and emotions are not actually our own. For we think in terms of languages and images which we did not invent, but which were given to us by our society. - Alan Watts",
	"To have faith is to trust yourself to the water. When you swim you don’t grab hold of the water, because if you do you will sink and drown. Instead, you relax and float. - Alan Watts",
	"Believe you can and you're halfway there. - Theodore Roosevelt",
	"The only way to do great work is to love what you do. - Steve Jobs",
	"Do what you can with all you have, wherever you are. - Theodore Roosevelt",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nQuote Generator")
		fmt.Println("1. Get a Random Quote")
		fmt.Println("2. Add a New Quote")
		fmt.Println("3. View All Quotes")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			fmt.Println("\nRandom Quote:")
			fmt.Println(quotes[rand.Intn(len(quotes))])
		case "2":
			addQuote(reader)
		case "3":
			viewQuotes()
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addQuote(reader *bufio.Reader) {
	fmt.Print("\nEnter a new quote: ")
	newQuote, _ := reader.ReadString('\n')
	newQuote = strings.TrimSpace(newQuote)
	quotes = append(quotes, newQuote)
	fmt.Println("Quote added!")
}

func viewQuotes() {
	fmt.Println("\nAll Quotes:")
	for i, quote := range quotes {
		fmt.Printf("%d. %s\n", i+1, quote)
	}
}
