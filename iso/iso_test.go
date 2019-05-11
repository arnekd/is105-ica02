package iso

import (
	"testing"
)

//Tester om karakterene er en extended-ASCII karakter, fra 0x80 til 0xFF
func TestGreetingExtendedASCII(t *testing.T) {
	got := GreetingExtendedASCII()

	for _, i := range got {
		if !(i >= 0x80 && i <= 0xFF) {
			t.Errorf("Found value: %q - %X, which is not part of ASCII-extended", i, i)
		}
	}
}
