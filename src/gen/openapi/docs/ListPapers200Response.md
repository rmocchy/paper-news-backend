# ListPapers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Papers** | [**[]ListPapers200ResponsePapersInner**](ListPapers200ResponsePapersInner.md) |  | 
**Total** | **int32** | The total number of papers. | 

## Methods

### NewListPapers200Response

`func NewListPapers200Response(papers []ListPapers200ResponsePapersInner, total int32, ) *ListPapers200Response`

NewListPapers200Response instantiates a new ListPapers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPapers200ResponseWithDefaults

`func NewListPapers200ResponseWithDefaults() *ListPapers200Response`

NewListPapers200ResponseWithDefaults instantiates a new ListPapers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPapers

`func (o *ListPapers200Response) GetPapers() []ListPapers200ResponsePapersInner`

GetPapers returns the Papers field if non-nil, zero value otherwise.

### GetPapersOk

`func (o *ListPapers200Response) GetPapersOk() (*[]ListPapers200ResponsePapersInner, bool)`

GetPapersOk returns a tuple with the Papers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPapers

`func (o *ListPapers200Response) SetPapers(v []ListPapers200ResponsePapersInner)`

SetPapers sets Papers field to given value.


### GetTotal

`func (o *ListPapers200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ListPapers200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ListPapers200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


