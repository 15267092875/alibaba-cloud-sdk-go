package ims

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListAppSecretIds invokes the ims.ListAppSecretIds API synchronously
func (client *Client) ListAppSecretIds(request *ListAppSecretIdsRequest) (response *ListAppSecretIdsResponse, err error) {
	response = CreateListAppSecretIdsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppSecretIdsWithChan invokes the ims.ListAppSecretIds API asynchronously
func (client *Client) ListAppSecretIdsWithChan(request *ListAppSecretIdsRequest) (<-chan *ListAppSecretIdsResponse, <-chan error) {
	responseChan := make(chan *ListAppSecretIdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppSecretIds(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListAppSecretIdsWithCallback invokes the ims.ListAppSecretIds API asynchronously
func (client *Client) ListAppSecretIdsWithCallback(request *ListAppSecretIdsRequest, callback func(response *ListAppSecretIdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppSecretIdsResponse
		var err error
		defer close(result)
		response, err = client.ListAppSecretIds(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListAppSecretIdsRequest is the request struct for api ListAppSecretIds
type ListAppSecretIdsRequest struct {
	*requests.RpcRequest
	AkProxySuffix string `position:"Query" name:"AkProxySuffix"`
	AppId         string `position:"Query" name:"AppId"`
}

// ListAppSecretIdsResponse is the response struct for api ListAppSecretIds
type ListAppSecretIdsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	AppSecrets AppSecrets `json:"AppSecrets" xml:"AppSecrets"`
}

// CreateListAppSecretIdsRequest creates a request to invoke ListAppSecretIds API
func CreateListAppSecretIdsRequest() (request *ListAppSecretIdsRequest) {
	request = &ListAppSecretIdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "ListAppSecretIds", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAppSecretIdsResponse creates a response to parse from ListAppSecretIds response
func CreateListAppSecretIdsResponse() (response *ListAppSecretIdsResponse) {
	response = &ListAppSecretIdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
