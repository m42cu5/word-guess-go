package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	words := [3](string){"apple", "banana", "orange"}
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn(3)]
	wordLength := len(word)
	// Store the clue in an array for easier manipulation.
	clue := make([]string, wordLength)
	for i := 0; i < wordLength; i++ {
		clue[i] = "_"
	}
	var guess string
	var nbrOfGuesses int
	for guess != word {
		nbrOfGuesses++
		// Print the clue as a string instead of an array.
		fmt.Println("Guess the word \"" + strings.Join(clue, "") + "\":")
		fmt.Scanln(&guess)
		if guess != word {
			if len(guess) == 1 {
				index := -1
				indexOfGuess := 0
				for indexOfGuess != -1 {
					index++
					if index > wordLength-1 {
						break
					} else {
						// Replace the first occurrence of the guess in the clue, starting from the current index.
						indexOfGuess = strings.Index(word[index:], guess)
						if indexOfGuess != -1 {
							index = index + indexOfGuess
							clue[index] = guess
						}
					}
				}
				guess = strings.Join(clue, "")
			}
		}
	}
	fmt.Println("Correct! The word was \"" + word + "\".\nYou guessed the word in " + strconv.Itoa(nbrOfGuesses) + " guess(es)")
}
