package jsonTest

import (
	"encoding/json"
	"fmt"
	"log"
)

func Map2Json() {
	c := make(map[string]interface{})
	c["ip"] = "116.62.202.230"
	c["city"] = "betjing"
	c["region"] = "Beijing"
	c["country"] = "CN"
	c["log"] = "39.9075,116.3972"
	c["org"] = "AS23724 IDC, China Telecommunications Corporation"
	c["timezone"] = "Asia/Shanghai"
	c["readme"] = "https://ipinfo.io/missingauth"
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	fmt.Println(string(data))
}
