package main

import (
	"fmt"
	"log"

	"github.com/z3ntl3/gtranslate/gtranslate"
)

func main() {
	text, err := gtranslate.New().
		Translate("nl", "tr", "ik ga morgen hardlopen en naar de gym")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translation: %s", text)
}
// translation: Yarın koşuya ve spor salonuna gideceğim