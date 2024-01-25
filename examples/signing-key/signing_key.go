package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/synadia-io/control-plane-sdk-go/syncp"
)

const (
	accountId = "2bKAVoSBsfwNDAEj5MdMaiwqj8e"
	baseUrl   = "https://cloud.synadia.com"
	// can be generated at baseUrl + /profile/personal-access-tokens
	pat = "uat_app_admin"
)

func main() {
	client := syncp.NewAPIClient(syncp.NewConfiguration())
	ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})
	ctx = context.WithValue(ctx, syncp.ContextAccessToken, pat)

	// create signing keys
	skResponse, _, err := client.AccountAPI.CreateAccountSkGroup(ctx, accountId).
		SigningKeyGroupCreateRequest(syncp.SigningKeyGroupCreateRequest{
			Name:         "sk group 1",
			Programmatic: false,
		}).Execute()
	if err != nil {
		handleApiError(err)
	}
	js, _ := json.MarshalIndent(skResponse, "", "\t")
	log.Printf("On Demand SK Group 1:\n%s\n\n", string(js))

	skResponse, _, err = client.AccountAPI.CreateAccountSkGroup(ctx, accountId).
		SigningKeyGroupCreateRequest(syncp.SigningKeyGroupCreateRequest{
			Name:         "sk group 2",
			Programmatic: false,
			Scope:        &syncp.UserPermissionLimits{BearerToken: syncp.Ptr(false)},
		}).Execute()
	if err != nil {
		handleApiError(err)
	}
	js, _ = json.MarshalIndent(skResponse, "", "\t")
	log.Printf("On Demand SK Group 2:\n%s\n\n", string(js))

	// check the account JWT now:
	account, _, err := client.AccountAPI.GetAccount(ctx, accountId).Execute()
	js, _ = json.MarshalIndent(account.Claims, "", "\t")
	log.Printf("Updated claims:\n%s\n\n", string(js))
}

func handleApiError(err error) {
	// error with body
	apiErr := &syncp.GenericOpenAPIError{}
	if errors.As(err, &apiErr) {
		log.Fatal(apiErr.Error(), string(apiErr.Body()))
	}
	log.Fatal(err)
}
