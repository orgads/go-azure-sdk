package reports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByTimeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReportRecordContract
}

type ListByTimeCompleteResult struct {
	Items []ReportRecordContract
}

type ListByTimeOperationOptions struct {
	Filter   *string
	Interval *string
	Orderby  *string
	Skip     *int64
	Top      *int64
}

func DefaultListByTimeOperationOptions() ListByTimeOperationOptions {
	return ListByTimeOperationOptions{}
}

func (o ListByTimeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByTimeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByTimeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Interval != nil {
		out.Append("interval", fmt.Sprintf("%v", *o.Interval))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListByTime ...
func (c ReportsClient) ListByTime(ctx context.Context, id ServiceId, options ListByTimeOperationOptions) (result ListByTimeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/reports/byTime", id.ID()),
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
		Values *[]ReportRecordContract `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByTimeComplete retrieves all the results into a single object
func (c ReportsClient) ListByTimeComplete(ctx context.Context, id ServiceId, options ListByTimeOperationOptions) (ListByTimeCompleteResult, error) {
	return c.ListByTimeCompleteMatchingPredicate(ctx, id, options, ReportRecordContractOperationPredicate{})
}

// ListByTimeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ReportsClient) ListByTimeCompleteMatchingPredicate(ctx context.Context, id ServiceId, options ListByTimeOperationOptions, predicate ReportRecordContractOperationPredicate) (result ListByTimeCompleteResult, err error) {
	items := make([]ReportRecordContract, 0)

	resp, err := c.ListByTime(ctx, id, options)
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

	result = ListByTimeCompleteResult{
		Items: items,
	}
	return
}
