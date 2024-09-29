package main

import (
	"fmt"
	"log"

	"github.com/z3ntl3/gtranslate/gtranslate"
)

func main() {
	text, err := gtranslate.New().
	Translate(gtranslate.TranslationOptions{
		SourceLang: "nl",
		TargetLang: "tr",
		Query: "ik ga morgen hardlopen en naar de gym",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translation: %s", text)
}
// translation: Yarın koşuya ve spor salonuna gideceğim