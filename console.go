package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Clear() {
	fmt.Print("\033[2J\033[H")
}

func Reset() {
	fmt.Print("\033[00m")
}

func SetColor(color int) {
	fmt.Printf("\033[%dm", color)
}
func SetBackgroundColor(color int) {
	fmt.Printf("\033[%dm", color+10)
}

func SetCursorPosition(x, y int) {
	fmt.Printf("\033[%d;%dH", y+1, x+1)
}

func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Println(err)
		return ""
	}
	return str
}
