package main

// a quick script for generating questions

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	layouts := []string{time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822,
		time.RFC822Z, time.RFC850, time.RFC1123, time.RFC1123Z, time.RFC3339,
		time.RFC3339Nano, time.Kitchen}
	nowUnix := time.Now().Unix()

	for _, layout := range layouts {
		random := rand.Int63n(nowUnix)
		x := time.Unix(random, 0)

		fmt.Println("{")
		fmt.Printf("\"%s\",\n", x.Format(layout))
		fmt.Println("\"XXX\",")
		fmt.Println(random, ",")
		fmt.Println("false,")
		fmt.Println("},")
	}
}
