package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Let's start to 7-Up!")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
	//log.SetLevel(log.DebugLevel)
	maxNumber := 69
	checkDigit := 7
	for curNumber := 1; curNumber < maxNumber; curNumber++ {
		if multipleOf(checkDigit, curNumber) || containsDigit(checkDigit, curNumber) {
			fmt.Println("Up")
		} else {
			fmt.Println(curNumber)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func containsDigit(digit int, num int) bool {
	log.Debug("Checking if ", num, " contains ", digit)
	d := strconv.Itoa(digit)
	n := strconv.Itoa(num)
	log.Debug("d is ", d)
	log.Debug("n is ", n)
	return strings.Contains(n, d)
}

func multipleOf(digit int, num int) bool {
	log.Debug("Checking if ", num, " is a multiple of ", digit)
	if num%digit == 0 {
		return true
	}
	return false
}
