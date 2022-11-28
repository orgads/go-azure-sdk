package objectdatatypes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectDataTypesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewObjectDataTypesClientWithBaseURI(endpoint string) ObjectDataTypesClient {
	return ObjectDataTypesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}