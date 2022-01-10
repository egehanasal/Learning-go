//? Just wanted to see how to create a json with Go

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	encodeJson()
	decodeJson()
}

type weatherData struct {
	LocationName string `json:"şehir"`
	Weather      string `json:"hava durumu"`
	Temperature  int    `json:"-"` //json'da bu görünmeyecek
}

func encodeJson() {

	weathers := []weatherData{
		{LocationName: "Istanbul",
			Weather:     "sunny",
			Temperature: 34},
		{LocationName: "Antalya",
			Weather:     "sunny",
			Temperature: 38},
		{LocationName: "Ankara",
			Weather:     "rainy",
			Temperature: 27},
	}
	//* Package this data as JSON data
	finalJson, err := json.MarshalIndent(weathers, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", finalJson)
}

func decodeJson() {
	jsonDataFramWeb := []byte(`
	{
		"şehir": "Istanbul",
		"hava durumu": "sunny"
	}
	`)

	var myVar weatherData
	checkValid := json.Valid(jsonDataFramWeb)

	if checkValid {
		fmt.Println("Valid")
		json.Unmarshal(jsonDataFramWeb, &myVar)
		fmt.Printf("%#v\n", myVar)
	} else {
		fmt.Println("\nJson was not valid")
	}

	//* some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFramWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("\nKey is %v and value is %v and type is %T", key, value, value)
	}
}
