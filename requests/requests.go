package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var ELK = "10.128.190.12:9200"



func GetData(url string, body string) string {
	//start_time :=time.Now()
	req, err := http.NewRequest("GET", "http://"  + ELK + url, strings.NewReader(body))
	if err != nil {
		// handle err
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	bodyString := ""
	if err != nil {
		fmt.Println(err)
	} else {
		if res.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString = string(bodyBytes)
		}
		res.Body.Close()
	}
	//log.Print(time.Since(start_time))
	return bodyString
}
