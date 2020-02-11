package cvt2json

import (
	"fmt"
	"regexp"
	"testing"
)

func TestBug(t *testing.T) {

	r, _ := regexp.Compile(`p\[ \[([a-z]+)ch`)
	fmt.Println(r.MatchString("p[ [each"))

}
