package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/synadia-io/control-plane-sdk-go/syncp"
)

const (
	SystemId = "2XoUBoGZoQzDpq3kFwh4YQKOqlQ"
)

func main() {
	client := syncp.NewAPIClient(syncp.NewConfiguration())
	ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
		"baseUrl": "http://localhost:8080",
	})
	ctx = context.WithValue(ctx, syncp.ContextAccessToken, "uat_app_admin")
	accountList, _, err := client.SystemAPI.ListAccounts(ctx, SystemId).Execute()
	if err != nil {
		log.Fatal(err)
	}
	for i, account := range accountList.Items {
		log.Printf("%d: %s\n", i+1, account.Name)
	}
	// to make a ptr to nullable (optional nullable):
	// urls := syncp.Ptr(syncp.NewNullable([]string{"server.url.1", "server.url.2"}))
	// to make a required nullable:
	// urls := syncp.NewNullable("hello")
	accountViewResponse, _, err := client.SystemAPI.CreateAccount(ctx, SystemId).
		AccountCreateRequest(
			syncp.AccountCreateRequest{
				Name: "hello",
				JwtSettings: &syncp.AccountJWTSettings{
					Info: syncp.Info{
						Description: syncp.Ptr("hello"),
					},
				},
				// use helper method to get ptr to primitive type
				UserJwtExpiresInSecs: syncp.Ptr(int64(0)),
			},
		).Execute()
	js, _ := json.Marshal(accountViewResponse)
	log.Printf("%s", string(js))
}
