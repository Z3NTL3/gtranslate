package tests

import (
	"testing"

	"github.com/z3ntl3/gtranslate/gtranslate"
)

// go test -timeout 30s -run ^TestTranslation$ github.com/z3ntl3/gtranslate/tests -v
func TestTranslation(t *testing.T) 	{
	text, err := gtranslate.New().
	Translate("nl", "tr", "ik ga morgen hardlopen en naar de gym")

	if err != nil {
		t.Fatalf("error: %s", err)
	}

	t.Logf("translation: %s", text)
}
/*
~\Documents\gtranslate via 🐹 v1.22.5 
❯ go test -timeout 30s -run ^TestTranslation$ github.com/z3ntl3/gtranslate/tests -v
Selam Dünya
=== RUN   TestTranslation
    gtranslate_test.go:17: translation: Yarın koşuya ve spor salonuna gideceğim
--- PASS: TestTranslation (0.10s)
PASS
ok      github.com/z3ntl3/gtranslate/tests      1.036s
*/