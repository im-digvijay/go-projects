package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ApiConfigData struct {
	Token string `json:"Token"`
}

func ReadConfig() (ApiConfigData, error) {
	fmt.Println("Reading config file...")

	bytes, err := os.ReadFile("./.apiConfig")
	if err != nil {
		fmt.Println(err.Error())
		return ApiConfigData{}, err
	}

	fmt.Println(string(bytes))

	var apiConfig ApiConfigData
	err = json.Unmarshal(bytes, &apiConfig)
	if err != nil {
		fmt.Println(err.Error())
		return ApiConfigData{}, err
	}

	return apiConfig, nil
}
