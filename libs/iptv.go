package libs

import (
	"encoding/json"
	"streamConsole/models"
	"streamConsole/utils"
)

type NetCard struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
}

func GetNetCards() []NetCard {
	status, dat := doRequest("/netcards", "")
	if status != 200 {
		return nil
	}
	netcards := dat.(map[string]interface{})["netcards"]

	b, err := json.Marshal(netcards)
	if err != nil {
		return nil
	}

	result := []NetCard{}
	if err = json.Unmarshal(b, &result); err != nil {
		return nil
	}
	return result
}

func ManageAll() interface{} {
	status, dat := doRequest("/manage", "query=all")

	if status == 200 {
		return dat.(map[string]interface{})["all"]
	}
	utils.ConsoleLogs.Info("status{},{}", status, dat)
	return ""
}

func doRequest(url string, param string) (int, interface{}) {
	requestUrl := "http://" + models.IptvUrl + url + "?" + param
	html, err := utils.HttpGet(requestUrl)
	if err != nil {
		return 500, err.Error()
	}
	return formatToJson(html)
}

func formatToJson(jsonString string) (int, interface{}) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &dat); err != nil {
		return 500, "返回的内容不是json格式，无法识别"
	}
	status := dat["status"].(float64)
	return int(status), dat
}
