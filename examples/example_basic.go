package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/synadia-io/control-plane-sdk-go/syncp"
)

const (
	systemId = "2bK87TxGCbilvDZPJy04FpIPCEz"
	baseUrl  = "https://cloud.synadia.com"
	// can be generated at baseUrl + /profile/personal-access-tokens
	pat = "uat_app_admin"
)

func main() {
	client := syncp.NewAPIClient(syncp.NewConfiguration())
	ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})
	ctx = context.WithValue(ctx, syncp.ContextAccessToken, pat)
	accountList, _, err := client.SystemAPI.ListAccounts(ctx, systemId).Execute()
	if err != nil {
		handleApiError(err)
	}
	var accountNames []string
	for _, account := range accountList.Items {
		accountNames = append(accountNames, account.Name)
	}
	log.Printf("Accounts in System: %s\n", strings.Join(accountNames, ", "))
	// to make a ptr to nullable (optional nullable):
	// url := syncp.Ptr(syncp.NewNullable("hello"))
	// to make a required nullable:
	// url := syncp.NewNullable("hello")
	accountViewResponse, _, err := client.SystemAPI.CreateAccount(ctx, systemId).
		AccountCreateRequest(
			syncp.AccountCreateRequest{
				Name: "hello",
				JwtSettings: &syncp.AccountJWTSettings{
					Info: syncp.Info{
						Description: syncp.Ptr("hello"),
					},
				},
				// use helper method to get ptr to primitive type
				UserJwtExpiresInSecs: syncp.Ptr(int64(30)),
			},
		).Execute()
	if err != nil {
		handleApiError(err)
	}
	js, _ := json.MarshalIndent(accountViewResponse, "", "\t")
	log.Printf("Newly Created Account:\n%s\n\n", string(js))
}

func handleApiError(err error) {
	// error with body
	apiErr := &syncp.GenericOpenAPIError{}
	if errors.As(err, &apiErr) {
		log.Fatal(apiErr.Error(), string(apiErr.Body()))
	}
	log.Fatal(err)
}
