package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricsConfigurationsListByClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterMetricsConfiguration
}

type MetricsConfigurationsListByClusterCompleteResult struct {
	Items []ClusterMetricsConfiguration
}

// MetricsConfigurationsListByCluster ...
func (c NetworkcloudsClient) MetricsConfigurationsListByCluster(ctx context.Context, id ClusterId) (result MetricsConfigurationsListByClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/metricsConfigurations", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ClusterMetricsConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MetricsConfigurationsListByClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) MetricsConfigurationsListByClusterComplete(ctx context.Context, id ClusterId) (MetricsConfigurationsListByClusterCompleteResult, error) {
	return c.MetricsConfigurationsListByClusterCompleteMatchingPredicate(ctx, id, ClusterMetricsConfigurationOperationPredicate{})
}

// MetricsConfigurationsListByClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) MetricsConfigurationsListByClusterCompleteMatchingPredicate(ctx context.Context, id ClusterId, predicate ClusterMetricsConfigurationOperationPredicate) (result MetricsConfigurationsListByClusterCompleteResult, err error) {
	items := make([]ClusterMetricsConfiguration, 0)

	resp, err := c.MetricsConfigurationsListByCluster(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = MetricsConfigurationsListByClusterCompleteResult{
		Items: items,
	}
	return
}
