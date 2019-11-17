package libs

import (
	"encoding/json"
	"fmt"
	"streamConsole/models"
	"streamConsole/utils"
	"strings"
	"time"
)

func CheckSerial(serial *models.Serial) bool {

	hardwareCode, _ := GetSerialCode()
	if hardwareCode == "" {
		serial.Remark = "获取硬件代码失败"
		return false
	} else {
		code, dat := FormatToJson(hardwareCode)
		if code == 200 || code == 404 {
			netcards := dat.(map[string]interface{})["serial"]
			serial.HardwareCode = netcards.(string)
		}
	}
	fmt.Printf("serial：%s\n", serial.SerialCode)

	if serial.SerialCode != "" {

		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("捕获到的错误：%s\n", r)
				serial.Remark = "解析注册码失败"
			}
		}()
		//解密序列号
		jsonData := utils.AesDecrypt(serial.SerialCode)
		if jsonData == "" {
			serial.Remark = "解析注册码"
			return false
		}

		fmt.Printf("jsonData：%s\n", jsonData)

		//解析JSON数据
		var serialData models.SerialData
		if err := json.Unmarshal([]byte(jsonData), &serialData); err != nil {
			serial.Remark = "注册信息不正确"
			return false
		}

		if !strings.EqualFold(serialData.HardwareCode, serial.HardwareCode) {

			serial.Remark = "注册信息与机器码不一致"
			return false
		}

		//有效期检查
		dateFormat := "2006-01-02 15:04:05"
		currentTime := time.Now()
		//转化注册码时间
		validTimeString := serialData.ValidTime + " 23:59:59"

		validTime, _ := time.ParseInLocation(dateFormat, validTimeString, time.Local)

		//fmt.Printf("validTime：%v\n",validTime)
		//fmt.Printf("validTimeString：%v\n",validTimeString)

		if currentTime.After(validTime) {

			//fmt.Printf("validTime：%v\n",validTime)
			//fmt.Printf("currentTime：%v\n",currentTime)

			serial.Remark = "当前序列号已过期，请更换新的序列号。"
			return false
		}

		//fmt.Printf("validTime：%v\n",validTime)
		//fmt.Printf("validTimeString：%v\n",validTimeString)

		//更新注册信息
		serial.ValidTime = validTime.Unix()
		serial.HardwareCode = serial.HardwareCode
		serial.Remark = "注册成功"

		return true
	} else {
		serial.Remark = "序列号不能为空1"
		return false
	}
}

func AutoCheckSerial() {
	serial, _ := models.SerialGetById(1)

	if serial != nil {
		valid := CheckSerial(serial)

		if valid {
			models.SerialValid = true
		}

		serial.Update()
	}
}

func ManualCheckSerial(serialCode string) *models.Serial {
	serial, _ := models.SerialGetById(1)
	serial.SerialCode = serialCode

	valid := CheckSerial(serial)
	if valid {
		models.SerialValid = true
	}
	serial.Update()

	return serial
}
