package libs

import (
	"encoding/json"
	"strconv"
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

func SaveChannelStatus(channel *models.ChannelEntity) (string, error) {
	url := "/channel"
	param := "id=" + channel.ID.Hex()

	if channel.Group != "" {
		param += "&group=" + channel.Group
	}
	if channel.Single != "" {
		param += "&single=" + channel.Single
	}
	if channel.Vod != "" {
		param += "&vod=" + channel.Vod
	}
	if channel.TSoc != "" {
		param += "&tSoc=" + channel.TSoc
	}

	return doRequestGet(url, param)

}

func SaveChannelEntity(channel *models.ChannelEntity, param string) (string, error) {
	url := "/manage"

	param += "&name=" + channel.Name
	param += "&src=" + channel.Src
	param += "&netcardin=" + channel.NetCardin
	param += "&program=" + strconv.Itoa(channel.Program)
	param += "&group=" + channel.Group
	param += "&single=" + channel.Single
	param += "&vod=" + channel.Vod
	param += "&tSoc=" + channel.TSoc

	return doRequestGet(url, param)
}

func ManageAll() interface{} {
	status, dat := doRequest("/manage", "query=all")

	if status == 200 {
		return dat.(map[string]interface{})["all"]
	}
	utils.ConsoleLogs.Info("status{},{}", status, dat)
	return ""
}

func doRequestGet(url string, param string) (string, error) {
	requestUrl := "http://" + models.IptvUrl + url + "?" + param
	return utils.HttpGet(requestUrl)
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
