package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

type LogPass struct {
	logging  string
	password string
}

func main() {
	var fileName string
	var count int
	flag.StringVar(&fileName, "file", "default", "File name")
	flag.IntVar(&count, "passCount", 0, "The pass count param")
	flag.Parse()
	dat, err := os.ReadFile(fileName)
	check(err)
	loggins := strings.Split(string(dat), "\n")
	createPairLogPass(loggins, count)
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func check(e error) {
	if e != nil {
		fmt.Print(e)
		panic(e)
	}
}

func createPairLogPass(logging []string, countPass int) {
	var itemsArray []LogPass
	rand.Seed(time.Now().Unix())
	for j := 0; j < len(logging); j++ {
		for i := 0; i < countPass; i++ {
			minSpecialChar := 1
			minNum := 1
			minUpperCase := 1
			passwordLength := rand.Intn(19) + 1
			password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
			if checkUniquePasswd(password, itemsArray) {
				i--
				continue
			}

			fmt.Printf("%s", logging[j]+":"+password+"\n")
			newItem := LogPass{logging: logging[j], password: password}
			itemsArray = append(itemsArray, newItem)
		}
	}

	f, err := os.Create("logPas.txt")
	check(err)
	defer f.Close()
	for i := 0; i < len(itemsArray); i++ {
		_, err := f.WriteString(itemsArray[i].logging + ":" + itemsArray[i].password + "\n")
		check(err)
	}
}

func checkUniquePasswd(password string, passArray []LogPass) bool {
	for i := 0; i < len(passArray); i++ {
		switch password {
		case
			passArray[i].password:
			return true
		}
		return false
	}
	return false
}
