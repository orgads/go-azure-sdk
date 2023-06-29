package feature

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureClient struct {
	Client  autorest.Client
	baseUri string
}

func NewFeatureClientWithBaseURI(endpoint string) FeatureClient {
	return FeatureClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
