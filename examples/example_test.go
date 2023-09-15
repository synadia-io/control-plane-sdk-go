package main

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	scp "github.com/synadia-io/control-plane-sdk-go/syncp"
)

type mockSystemApi struct {
	*scp.SystemAPIService
	accountList *scp.AccountListResponse
	resp        *http.Response
	err         error
}

func (a *mockSystemApi) ListAccountsExecute(r scp.ApiListAccountsRequest) (*scp.AccountListResponse, *http.Response, error) {
	return a.accountList, a.resp, a.err
}

func (a *mockSystemApi) ListAccounts(ctx context.Context, systemId string) scp.ApiListAccountsRequest {
	return scp.ApiListAccountsRequest{
		ApiService: a,
	}
}

func TestExample(t *testing.T) {
	client := scp.NewAPIClient(scp.NewConfiguration())
	ctx := context.WithValue(context.Background(), scp.ContextServerVariables, map[string]string{
		"baseUrl": "http://test.url",
	})
	ctx = context.WithValue(ctx, scp.ContextAccessToken, "test_bearer_token")
	client.SystemAPI = &mockSystemApi{
		accountList: scp.NewAccountListResponse([]scp.AccountViewResponse{
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
