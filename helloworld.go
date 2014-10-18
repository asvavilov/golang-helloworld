package main

import (
	"fmt"
	"strings"
	"time"
)

func say(lbl, msg string, c chan string) {
	c <- lbl + ":" + msg
	fmt.Println(lbl, msg)
}

func saySync(lbl, msg string, c chan string) (b bool, i int) {
	b = true
	i = 555
	c <- lbl + ":" + msg
	fmt.Println(lbl, msg)
	return b, i
}

func main() {
	c := make(chan string, 2)
	const greating = "hello"
	fmt.Printf("%s, World!\n", strings.Title(greating))
	go say("mylbl1", "mymsg1", c)
	b, i := saySync("mylbl2", "mymsg2", c)
	fmt.Println(b, i)
	fromChan1, fromChan2 := <-c, <-c
	fmt.Println("fromChan1, fromChan2 = ", fromChan1, fromChan2)
	time.Sleep(1000)
}
