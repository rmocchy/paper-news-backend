# \PapersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPapers**](PapersAPI.md#ListPapers) | **Get** /papers | List Papers With Pagination



## ListPapers

> ListPapers200Response ListPapers(ctx).Page(page).PageSize(pageSize).Title(title).Authors(authors).Tags(tags).Execute()

List Papers With Pagination

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	page := int32(56) // int32 | The page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of papers per page. (optional) (default to 10)
	title := "title_example" // string | The title of the paper. (optional)
	authors := []string{"Inner_example"} // []string | The authors of the paper. (optional)
	tags := []string{"Inner_example"} // []string | The tags of the paper (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PapersAPI.ListPapers(context.Background()).Page(page).PageSize(pageSize).Title(title).Authors(authors).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PapersAPI.ListPapers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPapers`: ListPapers200Response
	fmt.Fprintf(os.Stdout, "Response from `PapersAPI.ListPapers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPapersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number. | [default to 1]
 **pageSize** | **int32** | The number of papers per page. | [default to 10]
 **title** | **string** | The title of the paper. | 
 **authors** | **[]string** | The authors of the paper. | 
 **tags** | **[]string** | The tags of the paper | 

### Return type

[**ListPapers200Response**](ListPapers200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

