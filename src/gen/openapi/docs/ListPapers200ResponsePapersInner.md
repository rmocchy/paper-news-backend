# ListPapers200ResponsePapersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | The title of the paper. | 
**Authors** | **[]string** | The authors of the paper. | 
**Abstract** | **string** | The abstract of the paper. | 
**Tags** | Pointer to **[]string** | The tags of the paper. | [optional] 
**Url** | **string** | The URL of the paper. | 
**PublishedAt** | **time.Time** | The date the paper was published. | 

## Methods

### NewListPapers200ResponsePapersInner

`func NewListPapers200ResponsePapersInner(title string, authors []string, abstract string, url string, publishedAt time.Time, ) *ListPapers200ResponsePapersInner`

NewListPapers200ResponsePapersInner instantiates a new ListPapers200ResponsePapersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPapers200ResponsePapersInnerWithDefaults

`func NewListPapers200ResponsePapersInnerWithDefaults() *ListPapers200ResponsePapersInner`

NewListPapers200ResponsePapersInnerWithDefaults instantiates a new ListPapers200ResponsePapersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ListPapers200ResponsePapersInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListPapers200ResponsePapersInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListPapers200ResponsePapersInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetAuthors

`func (o *ListPapers200ResponsePapersInner) GetAuthors() []string`

GetAuthors returns the Authors field if non-nil, zero value otherwise.

### GetAuthorsOk

`func (o *ListPapers200ResponsePapersInner) GetAuthorsOk() (*[]string, bool)`

GetAuthorsOk returns a tuple with the Authors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthors

`func (o *ListPapers200ResponsePapersInner) SetAuthors(v []string)`

SetAuthors sets Authors field to given value.


### GetAbstract

`func (o *ListPapers200ResponsePapersInner) GetAbstract() string`

GetAbstract returns the Abstract field if non-nil, zero value otherwise.

### GetAbstractOk

`func (o *ListPapers200ResponsePapersInner) GetAbstractOk() (*string, bool)`

GetAbstractOk returns a tuple with the Abstract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstract

`func (o *ListPapers200ResponsePapersInner) SetAbstract(v string)`

SetAbstract sets Abstract field to given value.


### GetTags

`func (o *ListPapers200ResponsePapersInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListPapers200ResponsePapersInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListPapers200ResponsePapersInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListPapers200ResponsePapersInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUrl

`func (o *ListPapers200ResponsePapersInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListPapers200ResponsePapersInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListPapers200ResponsePapersInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPublishedAt

`func (o *ListPapers200ResponsePapersInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ListPapers200ResponsePapersInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ListPapers200ResponsePapersInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


