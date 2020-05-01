package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type code struct {
	Code  string `json:"Code"`
	Value int    `json:"Value"`
}


func main() {
	var code []code
	rcvd := `[{"Code":"StatusContinue","Value":100},{"Code":"StatusSwitchingProtocols","Value":101},{"Code":"StatusProcessing","Value":102},{"Code":"StatusEarlyHints","Value":103},{"Code":"StatusOK","Value":200}]`
	err := json.Unmarshal([]byte(rcvd), &code)
	if err != nil {
		log.Fatalln(err)
	}

	for _, value := range code {
		fmt.Println(value.Code, " - ", value.Value)
	}
}