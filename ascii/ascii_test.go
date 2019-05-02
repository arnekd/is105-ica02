package ascii

import (
	"testing"
	"unicode"
)

const want = "Hello :-)"

//testen TestGreetingASCII er en hjelpetest som tester om verdien vi får fra GreetingASCII() er lik "Hello :-)"
func TestGreetingASCII(t *testing.T) {
	got := GreetingASCII()
	if want != got {
		t.Errorf("greetingASCII() returned %s, expected %v", got, want)
	}
}

//tester om karakterene i GreetingASCII() er i ASCII-tegnsettet
func TestGreetingASCIIContainsASCII(t *testing.T) {
	got := GreetingASCII()

	for i := range got {
		if got[i] > unicode.MaxASCII {
			t.Errorf("Value contains: %q, which is not an ASCII value", got[i])
		}
	}
}
