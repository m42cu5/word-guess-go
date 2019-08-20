package main

import "fmt"
import "math/rand"
import "time"
import "strconv"
import "strings"

func main() {
  words := [3](string){"apple", "banana", "orange"}
  rand.Seed(time.Now().UnixNano())
  word := words[rand.Intn(3)]
  fmt.Println(word)
  wordLength := len(word)
  clue := make([]string, wordLength)
	for i := 0; i < wordLength; i++ {
		clue[i] = "_" 
	}
  var guess string
  var nbrOfGuesses int
  for guess != word {
    nbrOfGuesses++
    fmt.Println("Guess the word \"" + strings.Join(clue, "") + "\":")
    fmt.Scanln(&guess)
    if guess != word {
      if len(guess) == 1 {
        index := -1
        indexRelative := 0
        for indexRelative != -1 {
          index++
          if index > wordLength-1 {
            break
          } else {
            indexRelative = strings.Index(word[index:], guess)
            if indexRelative != -1 {
              index = index + indexRelative
              clue[index] = guess
            }
          }
        }
        guess = strings.Join(clue, "")
      }
    }
  }
  fmt.Println("Correct! The word was \"" + word + "\".\nYou guessed the word in " + strconv.Itoa(nbrOfGuesses) +  " guess(es)") 
}
