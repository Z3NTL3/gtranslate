<!-- header -->

<div align="center">   
    <div>
        <img src="image.png" width=60><br><br>
        <div>
                <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/z3ntl3/gtranslate">
        </div>
        <p> <strong> >> Enhance your Go programs with translations << </strong> </p>
    </div>
</div>
<br>
<hr>



# GTranslate
GTranslate is a powerful Go library made especially to enhance your programs with translations over a very ease and comfortable API.

#### Installation
> ``go get github.com/z3ntl3/gtranslate/gtranslate``

### Features
- Translate freely, without limits
- No API key required
- Supports timeout
- Supports proxy tunneling (``http``, ``sock4`` & ``https``)
- Modular interface, you can customize anything the way you wish


### Minimal example

```go
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
```