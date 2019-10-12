package main

import (
	"fmt"
	"strconv"
	"log"
)

func parseBools(values []string) ([]bool, error) {
	var bools[] bool
	for i, value := range values {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf("invalid value %s at %d", value, i)
		}
		bools = append(bools, parsed)
	}
	return bools, nil
}

func main() {
	booleanStrigs := []string{"true", "false", "foobar"}
	x, err := parseBools(booleanStrigs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(x)
	
}