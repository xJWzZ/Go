package main

import "os"
import "fmt"
import "math/rand"
import "time"

func main() {
	m := make(map[string]string)

	var a = "Great food, no atmosphere."
	m["Did you hear about the restaurant on the moon?"] = a
	m["What do you call a fake noodle?"] = "An Impasta."
	m["How many apples grow on a tree?"] = "All of them."
	m["Want to hear a joke about paper?"] = "Nevermind it's tearable."
	m["I just watched a program about beavers."] = "It was the best dam program I've ever seen."

	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithoutProg)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var i = r1.Intn(len(m))
	var j = 0

	for k := range m {
		if j == i {
			fmt.Println(k)
		}
		j++
	}

}
