package myutils

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Version : which tools used
func Version() {
	fmt.Println("myutils Version: 1.0.0 (from github)")
}

// H1 : prints a heading like html <h1>
func H1(msg string) {
	str := strings.Repeat("=", len(msg))
	fmt.Println()
	fmt.Println("/===" + str + "===\\")
	fmt.Println("|   " + msg + "   |")
	fmt.Println("\\===" + str + "===/")
}

// H2 : prints a heading like html <h2>
func H2(msg string) {
	str := strings.Repeat("-", len(msg))
	fmt.Println()
	fmt.Println("+---" + str + "---+")
	fmt.Println("+   " + msg + "   !")
	fmt.Println("+---" + str + "---+")
}

// Hr : prints a horizontal Line like html <hr>
func Hr() {
	fmt.Println("------------------------")
}

// Br : prints a new Line like html <br>
func Br() {
	NewLines(1)
}

// NewLines : prints n new Lines
func NewLines(n int) {
	// n must be between 1 an 10
	if n < 1 || n > 10 {
		n = 1
	}
	for n > 0 {
		fmt.Println()
		n--
	}
}

// Wait : prints *** and waits for ENTER
// this technique is inspired from the IBM Host TSO/SPF
func Wait() {
	fmt.Print("***")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	_ = scanner.Text()
}

// Comment : prints a comment
func Comment(msg string) {
	Br()
	fmt.Println("===== " + msg + " =====")
}

// Description : prints description (one or more lines)
func Description(msg ...string) {
	if len(msg) > 0 {
		Br()
		fmt.Println("===")
		for _, m := range msg {
			fmt.Println("      " + m)
		}
		fmt.Println("===")
	}
}

// ShowObjects : for objects show kind, type and value type (one or many)
func ShowObjects(arr ...interface{}) {
	for _, a := range arr { // index not used
		k := reflect.ValueOf(a).Kind().String()
		t := reflect.TypeOf(a)
		fmt.Printf("Kind: %-8v Type: %-10v Value: %v \n",
			k, t, a,
		)
	}
}

// ShowSlice : for a slice show kind, type, capacity, length and value
func ShowSlice(a interface{}, cap, len int) {
	k := reflect.ValueOf(a).Kind().String()
	t := reflect.TypeOf(a)
	fmt.Printf("Kind: %-8v Type: %-10v Cap: %-3v Len: %-3v Value: %v \n",
		k, t, cap, len, a,
	)
}
