package main

import (
	"context"
	"fmt"

	binance_connector "github.com/AlekseiSolonkov/binance-connector-go"
)

func main() {
	EnableFastWithdrawSwitchService()
}

func EnableFastWithdrawSwitchService() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance_connector.NewClient(apiKey, secretKey, baseURL)

	// EnableFastWithdrawSwitchService - /sapi/v1/account/enableFastWithdrawSwitch
	res, err := client.NewEnableFastWithdrawSwitchService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance_connector.PrettyPrint(res))
}
