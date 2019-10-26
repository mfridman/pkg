package pp

import (
	"encoding/json"
	"fmt"
)

// PP is a convenience function to pretty print structs.
func PP(i interface{}) {
	b, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(string(b))
}
