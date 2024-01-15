// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
    ans := make(chan string)

    //A go-routine that receives all questions,
    // and for each incoming question, creates a separate go-routine that answers that question
	go func(){
		for question := range questions {
			go queAnswers(question, ans)
			}
		}()

    // A go-routine that generates predictions
	go func(){
		for{
            time.Sleep(10 * time.Second)
            go prophecy("",ans)
		}
	}()

    // A go-routine that receives all answers and predictions, and prints then to stdout

	go func(){
		for answer := range ans{
			characters := strings.Split(answer, "")
			for _, c := range characters {
				fmt.Printf(c)
				time.Sleep(time.Millisecond*time.Duration((60*rand.Intn(5))))
			}
		}

	}()

	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	//time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"The moon is dark.",
		"The sun is bright.",
		"NFT is the Future",
		"We will pay with BTC in 5 years",
		"Summer in Sweden is Amazing!",
	}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))] + "\n"
}
func queAnswers(pred string, answer chan<- string){


	//Declaring and initializing an array with the different answers
	possibleAnswers := []string{
		"Sorry, I don't understand",
		"Yes.",
		"What do you mean?",
		"Wouldn't you want to know",
		"No",
		"Interesting...",
		"I see",
	}

	answer <- possibleAnswers[rand.Intn(len(possibleAnswers))] + "\n"+prompt

}
func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}