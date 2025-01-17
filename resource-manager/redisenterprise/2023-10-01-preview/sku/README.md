
## `github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/sku` Documentation

The `sku` SDK allows for interaction with the Azure Resource Manager Service `redisenterprise` (API Version `2023-10-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/redisenterprise/2023-10-01-preview/sku"
```


### Client Initialization

```go
client := sku.NewSKUClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `SKUClient.List`

```go
ctx := context.TODO()
id := sku.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
