package main

import (
	"fmt"
	"net/http"
	"time"

	tata "github.com/jameszhujj/tata-go"
)

func main() {
	headerUsername := "some-credential"
	headerPassword := "some-credential"
	bodyUsername := "some-credential"
	bodyPassword := "some-credential"
	toNumber := "some-number"
	fromNumber := "Uber"
	text := "Message from tata-go"
	callback := "some-callback-url"

	client := tata.Client{
		HeaderAuthentication: tata.Credentials{
			Username: headerUsername,
			Password: headerPassword,
		},
		BodyAuthentication: tata.Credentials{
			Username: bodyUsername,
			Password: bodyPassword,
		},
		HTTPClient: http.DefaultClient, //optional parameter
	}

	sendSms(client, fromNumber, toNumber, text)

	sendSmsWithCallback(client, fromNumber, toNumber, text, callback)

	time.Sleep(time.Minute * 5)
}

func sendSms(client tata.Client, from string, to string, text string) string {
	res, err := client.SendSMSMessage(from, to, text, "")

	if err != nil {
		fmt.Println("SMS sending error:")
		fmt.Println(err)
		return ""
	}

	fmt.Println("SMS sending success:")
	fmt.Printf("%+v\n", res)
	return res
}

func sendSmsWithCallback(client tata.Client, from string, to string, text string, callbackUrl string) string {
	res, err := client.SendSMSMessage(from, to, text+" with call back", callbackUrl)

	if err != nil {
		fmt.Println("SMS sending with callback error:")
		fmt.Println(err)
		return ""
	}

	fmt.Println("SMS sending with callback success:")
	fmt.Printf("%+v\n", res)
	return res
}
