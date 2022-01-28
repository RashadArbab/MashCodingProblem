package main

import "fmt"

func main() {

	prTable := map[string]string{
		"br": "br",
		"c":  "c",
		"ca": "ca",
		"h":  "h",
		"he": "he",
		"k":  "k",
		"i":  "i",
		"ki": "ki",
	}

	str := "breakiing badbrc"

	getResults(str, prTable)

}

func getResults(str string, prTable map[string]string) [][]int {
	var arrayOfIndecies [][]int

	for i := 0; i < len(str); i++ {
		char2 := make(chan bool)
		char1a := make(chan bool)

		// for the enf of the array we do not need to check for a 2 char match because then we run
		// into index out of bounds
		if i != len(str)-1 {
			// go find me 2 char solution and 1 char solution
			go findChars(str[i:i+2], prTable, char2)
			go findChars(string(str[i]), prTable, char1a)

			char2bool := <-char2
			char1abool := <-char1a

			// if we get the 2 char we dont care about the 1 char anymore
			if char2bool {
				temp := []int{i, i + 1}
				// fmt.Println("appending char2", temp)
				arrayOfIndecies = append(arrayOfIndecies, temp)
			} else {
				if char1abool {
					temp := []int{i}
					// fmt.Println("appending char1a", temp)
					// if index[i] is not already used as index[i+1] from the previous run we add index[i] as a single
					if arrayOfIndecies[len(arrayOfIndecies)-1][1] != temp[0] {
						arrayOfIndecies = append(arrayOfIndecies, temp)
					}
				}

			}
		} else {
			// this is for the last element in the string if it matches then we can add it as a single
			go findChars(string(str[i]), prTable, char1a)
			char1abool := <-char1a
			if char1abool {
				temp := []int{i}
				// fmt.Println("appending char1a", temp)
				if arrayOfIndecies[len(arrayOfIndecies)-1][0] != temp[0] && arrayOfIndecies[len(arrayOfIndecies)-1][1] != temp[0] {
					arrayOfIndecies = append(arrayOfIndecies, temp)
				}
			}
		}

	}

	fmt.Println(arrayOfIndecies)
	return arrayOfIndecies
}

// runs a simple look up of a piece of the given string with the map, and returns true or false to the chanel
func findChars(str string, prTable map[string]string, ch chan bool) {
	if prTable[str] == str {
		// fmt.Println("returning true ", str)
		ch <- true
	} else {
		ch <- false
	}
}
