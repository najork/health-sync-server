// This file was generated by Conjure and should not be manually edited.

package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/palantir/conjure-go-runtime/v2/conjure-go-client/httpclient"
)

type ApolloServiceClient interface {
	// Triggers metrics collection from the given provider.
	Collect(ctx context.Context, providerArg Provider, requestArg ProviderRequest) error
}

type apolloServiceClient struct {
	client httpclient.Client
}

func NewApolloServiceClient(client httpclient.Client) ApolloServiceClient {
	return &apolloServiceClient{client: client}
}

func (c *apolloServiceClient) Collect(ctx context.Context, providerArg Provider, requestArg ProviderRequest) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Collect"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("POST"))
	requestParams = append(requestParams, httpclient.WithPathf("/api/collect/%s", url.PathEscape(fmt.Sprint(providerArg))))
	requestParams = append(requestParams, httpclient.WithJSONRequest(requestArg))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}
