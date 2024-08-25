# JsonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **map[string]interface{}** |  | [optional] 
**MessagePrototype** | Pointer to [**MessagePrototype**](MessagePrototype.md) |  | [optional] 

## Methods

### NewJsonResponse

`func NewJsonResponse() *JsonResponse`

NewJsonResponse instantiates a new JsonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonResponseWithDefaults

`func NewJsonResponseWithDefaults() *JsonResponse`

NewJsonResponseWithDefaults instantiates a new JsonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *JsonResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *JsonResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *JsonResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *JsonResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *JsonResponse) GetMessage() map[string]interface{}`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *JsonResponse) GetMessageOk() (*map[string]interface{}, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *JsonResponse) SetMessage(v map[string]interface{})`

SetMessage sets Message field to given value.

### HasMessage

`func (o *JsonResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessagePrototype

`func (o *JsonResponse) GetMessagePrototype() MessagePrototype`

GetMessagePrototype returns the MessagePrototype field if non-nil, zero value otherwise.

### GetMessagePrototypeOk

`func (o *JsonResponse) GetMessagePrototypeOk() (*MessagePrototype, bool)`

GetMessagePrototypeOk returns a tuple with the MessagePrototype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePrototype

`func (o *JsonResponse) SetMessagePrototype(v MessagePrototype)`

SetMessagePrototype sets MessagePrototype field to given value.

### HasMessagePrototype

`func (o *JsonResponse) HasMessagePrototype() bool`

HasMessagePrototype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


