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

	kv, _, err := client.AccountAPI.CreateKvBucket(ctx, accountId).JSKVBucketConfig(syncp.JSKVBucketConfig{
		Bucket:      "test-kv-1",
		Description: syncp.Ptr("example kv"),
		Storage:     syncp.STORAGETYPE_FILE,
		MaxBytes:    syncp.Ptr(int64(100)),
	}).Execute()
	if err != nil {
		handleApiError(err)
	}

	js, _ := json.MarshalIndent(kv, "", "\t")
	log.Printf("kv:\n%s\n\n", string(js))
}

func handleApiError(err error) {
	// error with body
	apiErr := &syncp.GenericOpenAPIError{}
	if errors.As(err, &apiErr) {
		log.Fatal(apiErr.Error(), string(apiErr.Body()))
	}
	log.Fatal(err)
}
