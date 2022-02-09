package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	ID           string `json:"id"`
	Nome         string `json:"nome"`
	Clima        string `json:"clima"`
	Terreno      string `json:"terreno"`
	QtdAparicoes int    `json:"qtdaparicoes"`
}

func main() {
	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://swapi.co/", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObj Response
	json.Unmarshal(bodyBytes, &responseObj)
	fmt.Printf("Api Response as struct %+v\n", responseObj)
}
