package address

import (
	"fmt"
	"testing"
)

func TestGenAddr(t *testing.T) {
	privateKey,address := CreateKey()
	fmt.Println(privateKey,address)
}