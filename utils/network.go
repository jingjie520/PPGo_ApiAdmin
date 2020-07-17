package utils

import (
	"io/ioutil"
	"net/http"
)

/**
* 发起HTTP请求
 */
func HttpGet(url string) (string, error) {
	ConsoleLogs.Debug("发起请求：%s", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ConsoleLogs.Debug("Resp：%s", body)
	return string(body), nil
}
