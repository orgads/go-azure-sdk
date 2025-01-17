package deployments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentExtended
}

type ListAtManagementGroupScopeCompleteResult struct {
	Items []DeploymentExtended
}

type ListAtManagementGroupScopeOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultListAtManagementGroupScopeOperationOptions() ListAtManagementGroupScopeOperationOptions {
	return ListAtManagementGroupScopeOperationOptions{}
}

func (o ListAtManagementGroupScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAtManagementGroupScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListAtManagementGroupScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListAtManagementGroupScope ...
func (c DeploymentsClient) ListAtManagementGroupScope(ctx context.Context, id commonids.ManagementGroupId, options ListAtManagementGroupScopeOperationOptions) (result ListAtManagementGroupScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Resources/deployments", id.ID()),
		OptionsObject: options,
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
		Values *[]DeploymentExtended `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAtManagementGroupScopeComplete retrieves all the results into a single object
func (c DeploymentsClient) ListAtManagementGroupScopeComplete(ctx context.Context, id commonids.ManagementGroupId, options ListAtManagementGroupScopeOperationOptions) (ListAtManagementGroupScopeCompleteResult, error) {
	return c.ListAtManagementGroupScopeCompleteMatchingPredicate(ctx, id, options, DeploymentExtendedOperationPredicate{})
}

// ListAtManagementGroupScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentsClient) ListAtManagementGroupScopeCompleteMatchingPredicate(ctx context.Context, id commonids.ManagementGroupId, options ListAtManagementGroupScopeOperationOptions, predicate DeploymentExtendedOperationPredicate) (result ListAtManagementGroupScopeCompleteResult, err error) {
	items := make([]DeploymentExtended, 0)

	resp, err := c.ListAtManagementGroupScope(ctx, id, options)
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

	result = ListAtManagementGroupScopeCompleteResult{
		Items: items,
	}
	return
}
