package cvt2json

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestTmp(t *testing.T) {
	m := map[string]interface{}{
		"USER_ID":   "JD",
		"USER_NAME": "John Doe",
	}
	sub := map[string]interface{}{
		"Root": m,
	}
	b, err := json.Marshal(sub)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
