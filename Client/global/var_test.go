package global

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	fmt.Println(Init())
	fmt.Println(Cfg.Path)
}
