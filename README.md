<!-- header -->

<div align="center">   
    <div>
        <img src="image.png" width=60><br><br>
        <div>
                <img alt="GitHub License" src="https://img.shields.io/github/license/z3ntl3/ProxyBeast" >
                <img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/z3ntl3/ProxyBeast">
                <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/z3ntl3/ProxyBeast">
        </div>
    </div>
</div>
<br>
<hr>

<div style="text-align: center">

**>> Enhance your Go programs with translations <<**

</div>


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
		Translate("nl", "tr", "ik ga morgen hardlopen en naar de gym")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translation: %s", text)
}
// translation: Yarın koşuya ve spor salonuna gideceğim
```