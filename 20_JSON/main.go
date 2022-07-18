package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Nome  string `json:"nome"`
	Idade uint   `json:"idade"`
	Raca  string `json:"raca"`
}

func main() {

	c := dog{"rex", 3, "vira_lata"}

	fmt.Println(c)

	cJSON, err := json.Marshal(c)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cJSON)
	fmt.Println(bytes.NewBuffer(cJSON))

	jsonDog := `{"nome":"rex","idade":3,"raca":"vira_lata"}`
	var c1 dog

	fmt.Println(c1)

	if er := json.Unmarshal([]byte(jsonDog), &c1); er != nil {
		log.Fatal(er)
	}

	fmt.Println(c1)

}
