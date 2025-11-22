package utils

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data any) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s \n", p)
}
