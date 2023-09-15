package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/synadia-io/control-plane-sdk-go/syncp"
)

type mockSystemApi struct {
	*syncp.SystemAPIService
	accountList *syncp.AccountListResponse
	resp        *http.Response
	err         error
}

func (a *mockSystemApi) ListAccountsExecute(r syncp.ApiListAccountsRequest) (*syncp.AccountListResponse, *http.Response, error) {
	return a.accountList, a.resp, a.err
}

func (a *mockSystemApi) ListAccounts(ctx context.Context, systemId string) syncp.ApiListAccountsRequest {
	return syncp.ApiListAccountsRequest{
		ApiService: a,
	}
}

func TestExample(t *testing.T) {
	client := syncp.NewAPIClient(syncp.NewConfiguration())
	ctx := context.WithValue(context.Background(), syncp.ContextServerVariables, map[string]string{
		"baseUrl": "http://test.url",
	})
	ctx = context.WithValue(ctx, syncp.ContextAccessToken, "test_bearer_token")
	client.SystemAPI = &mockSystemApi{
		accountList: syncp.NewAccountListResponse([]syncp.AccountViewResponse{
			{
				Id:   "testid1",
				Name: "test account",
			},
			{
				Id:   "testid2",
				Name: "test account 2",
			},
		}),
	}
	accountList, _, err := client.SystemAPI.ListAccounts(ctx, "2Ui3Xj44Ly8BG13qQT0nQWzZ0ss").Execute()
	require.NoError(t, err)
	require.Len(t, accountList.Items, 2)
	require.Equal(t, accountList.Items[0].Name, "test account")
}
