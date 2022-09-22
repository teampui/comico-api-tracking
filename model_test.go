package tracking

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	hmacNew := hmac.New(sha256.New, []byte("a"))
	hmacNew.Write([]byte("b"))
	sign := fmt.Sprintf("%x", hmacNew.Sum(nil))

	if sign != Signature() {
		t.Errorf("Signature() = %v, want %v", sign, Signature())
	}

	fmt.Println(sign)
}

func TestCheckSignature(t *testing.T) {
	if !CheckSignature("08de329931e295683776aa9a43529bd0b275286df3160300c49ba4e841833013") {
		t.Errorf("CheckSignature() = %v, want %v", false, true)
	}
}
