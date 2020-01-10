//Jonathan Burdette
//Text Cloud Implementation

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

//maps for the exclude and input file
var excludeMap = make(map[string]string)
var inputMap = make(map[string]int)

//struct to hold word and corresponding count
type wordAndCount struct {
	word  string
	count int
}

//slice to hold top fifty words
var finalSlice = make([]wordAndCount, 50)

//slice to hold all keys
var keys = make([]string, 0)

//reads in the exclude file
func readExclude(excludeFileName string) {
	file, err := os.Open(excludeFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//stores words from exclude file in hash map (blank values)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		excludeMap[scanner.Text()] = ""
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//reads the input file and store for further use
func readInput(inputFileName string) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	//stores words and counts where key is string and value is an integer used to hold the count
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		r := regexp.MustCompile("[^a-z'A-Z]+")
		stringSlice := r.Split(line, -1) //stringSlice is a slice containing all words on a line
		for i := 0; i < len(stringSlice); i++ {
			word := strings.ToLower(stringSlice[i])
			if _, ok := excludeMap[word]; !ok { //doesn't allow excluded words
				re := regexp.MustCompile("[a-zA-Z]+")
				if re.MatchString(word) == true && len(word) > 1 { // doesn't allow punctuation as a word
					count := inputMap[word]
					if count == 0 {
						inputMap[word] = 1
					} else {
						inputMap[word] = count + 1
					}
				}
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

//finds the 50 most common words based on counts
func findCommonWords() {

	//creates a slice filled with words from input file and sorts it based on corresponding word counts in the map
	keys = make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return inputMap[keys[i]] > inputMap[keys[j]]
	})

	//assigns top fifty words in slice of words and corresponding counts from input map into the array
	i := 0
	for _, key := range keys {
		if i < 50 {
			finalSlice[i].word = key
			finalSlice[i].count = inputMap[key]
		}
		i++
	}

	//sort based on words
	sort.Slice(finalSlice, func(i, j int) bool {
		return finalSlice[i].word < finalSlice[j].word
	})
}

//writes html for text cloud
func writeHTMLFile(HTMLFileName string) {

	//info for html file
	rangeValue := inputMap[keys[0]] - inputMap[keys[49]]
	sizeFactor := 1000.0 / float64(rangeValue)
	colors := []string{"#623224", "#FABE8B", "#EC1D23", "#0160B0"}

	file, err := os.Create(HTMLFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	colorCount := 0
	for i := 0; i < 50; i++ {
		if colorCount == 4 { //colors restart after 4 iterations
			colorCount = 0
		}
		fontSize := float64(finalSlice[i].count) * sizeFactor
		file.WriteString("<span style=\"font-size:")
		file.WriteString(fmt.Sprintf("%f", fontSize))
		file.WriteString("%; color:")
		file.WriteString(colors[colorCount])
		file.WriteString(";\">")
		file.WriteString(finalSlice[i].word)
		file.WriteString("</span> &nbsp; &nbsp;\n")
		colorCount++
	}
}

func main() {

	//check command line syntax
	if len(os.Args) < 3 {
		fmt.Println("Invalid syntax. To run program, enter: \"go run TextCloud <input file name> <exclude file name> <output file name>\"")
	} else {
		readExclude(os.Args[2])
		readInput(os.Args[1])
		findCommonWords()
		writeHTMLFile(os.Args[3])
	}
}
