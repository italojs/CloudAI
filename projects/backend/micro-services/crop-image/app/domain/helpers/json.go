package helper

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
)

func ToMap(body io.ReadCloser) (jsonMap map[string]interface{}, err error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalln(err)
	}
	bodyString := string(bodyBytes)
	jsonMap = make(map[string]interface{})
	err = json.Unmarshal([]byte(bodyString), &jsonMap)
	return
}
