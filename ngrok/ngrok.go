package ngrok

import (
	"encoding/json"
	"fmt"
	"log"
)

func Parse(str string) string {

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	var text string
	var data map[string]interface{}
	err := json.Unmarshal([]byte(str), &data)
	if err != nil {
		panic("解析失败")
	}
	for _, tunnel := range data["tunnels"].([]interface{}) {
		t := tunnel.(map[string]interface{})
		line := fmt.Sprintf("name:%s, url:%s, local:%s<br/>\n", t["name"], t["public_url"], (t["config"].(map[string]interface{}))["addr"])
		text += line
	}
	return text
}
