package main

import (
"fmt"
"io/ioutil"
"net/http"
"net/url"
"bytes"
)

func main() {
	send()
}

func send() {
	// 预约成功提交车主信息 (POST http://118.194.246.76/HKCIBY/hkcOrderInfo/checkUser2HkcOrderInfo.action)

	params := url.Values{}
	params.Set("car_plate_number", "京N590Z1")
	params.Set("reverseTimeStr", "2016-04-30")
	params.Set("vehicle_info_id", "2121507")
	params.Set("vehicle_owner_id", "1496749")
	params.Set("reverseDealerId", "765")
	params.Set("brandName", "嘉实多")
	params.Set("reverseMileage", "10000")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "http://118.194.246.76/HKCIBY/hkcOrderInfo/checkUser2HkcOrderInfo.action", body)

	// Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}


