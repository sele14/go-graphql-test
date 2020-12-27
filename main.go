package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Source; https://www.youtube.com/watch?v=1OwMtqxJDnA&list=PLII40EOBF0-voAGYkZHVipwlIVfgVxSQT&ab_channel=ThePolyglotDeveloper

func main() {
	// declare our api query we want to send
	jsonData := map[string]string{
		"query": `
		{
			countries {
				code,
				name,
				currency
			}
		}
		`,
	}
	// convert our map to a readable format, e.g. byte
	jsonValue, err := json.Marshal(jsonData)

	if err != nil {
		panic(err)
	}

	// create our post request to the API of choice.
	// also catch the error.
	// Converting bytes to buffer so the API can read it.
	request, err := http.NewRequest("POST", "https://countries.trevorblades.com/", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	// add a header to specify the content type as json
	request.Header.Add("content-type", "application/json")

	// create http client
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
