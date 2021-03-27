package myutils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// GetString : returns a string from console
func GetString(prompt string) string {
	fmt.Print(prompt + " > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// GetInteger : returns an integer from the console
func GetInteger(prompt string) int {
	for {
		s := GetString(prompt)
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Print("wrong input ... ")
		} else {
			return int(i)
		}
	}
}

var _randSource rand.Source
var _randRand *rand.Rand

// GetRandom : returns a random integer between 1 and 2_000_000_000
func GetRandom() int {
	return GetRandom2(0, 0)
}

// GetRandom2 : return a random integer between min and max
func GetRandom2(min, max int) int {
	if _randRand == nil {
		_randSource = rand.NewSource(time.Now().UnixNano())
		_randRand = rand.New(_randSource)
	}
	if max <= min {
		min = 1
		max = 2000000000
	}
	a := max - min + 1
	return _randRand.Intn(a) + min
}
