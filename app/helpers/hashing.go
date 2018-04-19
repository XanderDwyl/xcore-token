package helpers

import (
	"crypto/sha512"
	"fmt"
	"log"
)

// Sha512 ...
func Sha512(input string) string {
	s512 := sha512.New()
	s512.Write([]byte(input))
	ret := fmt.Sprintf("sha512:\t%x", s512.Sum(nil))
	log.Printf("ret:%s", ret)
	return ret
}
