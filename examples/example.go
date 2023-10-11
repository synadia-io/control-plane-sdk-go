package main

import (
	"context"
	"log"

	scp "github.com/synadia-io/control-plane-sdk-go/syncp"
)

func main() {
	client := scp.NewAPIClient(scp.NewConfiguration())
	ctx := context.WithValue(context.Background(), scp.ContextServerVariables, map[string]string{
		"baseUrl": "http://localhost:3000",
	})
	ctx = context.WithValue(ctx, scp.ContextAccessToken, "uat_app_admin")
	accountList, _, err := client.SystemAPI.ListAccounts(ctx, "2Ui3Xj44Ly8BG13qQT0nQWzZ0ss").Execute()
	if err != nil {
		log.Fatal(err)
	}
	for i, account := range accountList.Items {
		log.Printf("%d: %s\n", i+1, account.Name)
	}
}
