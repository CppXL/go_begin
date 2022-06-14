package jsonTest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	IpInfo struct {
		Ip       string `json:"ip"`
		City     string `json:"city"`
		Region   string `json:"region"`
		Country  string `json:"country"`
		Loc      string `json:"loc"`
		Org      string `json:"org"`
		Timezone string `json:"timezone"`
		Readme   string `json:"readme"`
	}
)
type duration int64

var Json_data = `{
	"ip": "220.181.41.32",
	"city": "Beijing",
	"region": "Beijing",
	"country": "CN",
	"loc": "39.9075,116.3972",
	"org": "AS23724 IDC, China Telecommunications Corporation",
	"timezone": "Asia/Shanghai",
	"readme": "https://ipinfo.io/missingauth"
}`

func TJson1() {
	url := "http://116.62.202.230/myip.php"
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	var ipInfo IpInfo
	err = json.NewDecoder(resp.Body).Decode(&ipInfo)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", ipInfo)

	var c IpInfo
	err = json.Unmarshal([]byte(Json_data), &c)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println("---------------------")
	fmt.Println(c)
}
