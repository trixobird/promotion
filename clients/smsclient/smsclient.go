package smsclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	u "promotion/utils"
	"strings"
)

func SendSms(phone string, msg string) (map[string]interface{}, bool) {
	apiUrl := "https://rest.nexmo.com/sms/json"
	data := url.Values{}
	data.Set("from", "inreach")
	data.Set("text", msg)
	data.Set("to", phone)
	data.Set("api_key", "94b64155")
	data.Set("api_secret", "ab63s0WvgdvETJVk")

	client := &http.Client{}
	r, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return u.Message(false, "Unexpected Error during building the request"), false
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return u.Message(false, "The sms server is offline, please try again later"), false
	}
	fmt.Print(string(bodyBytes))
	return u.Message(true, "Message sent successfully"), true
}
