package jsonTest

import (
	"encoding/json"
	"fmt"
	"log"
)

func Json4map() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(Json_data), &c)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	fmt.Println(c["ip"])

}
