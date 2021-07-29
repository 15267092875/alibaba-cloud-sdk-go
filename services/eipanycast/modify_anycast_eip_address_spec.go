package eipanycast

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

// ModifyAnycastEipAddressSpec invokes the eipanycast.ModifyAnycastEipAddressSpec API synchronously
func (client *Client) ModifyAnycastEipAddressSpec(request *ModifyAnycastEipAddressSpecRequest) (response *ModifyAnycastEipAddressSpecResponse, err error) {
	response = CreateModifyAnycastEipAddressSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyAnycastEipAddressSpecWithChan invokes the eipanycast.ModifyAnycastEipAddressSpec API asynchronously
func (client *Client) ModifyAnycastEipAddressSpecWithChan(request *ModifyAnycastEipAddressSpecRequest) (<-chan *ModifyAnycastEipAddressSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyAnycastEipAddressSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyAnycastEipAddressSpec(request)
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

// ModifyAnycastEipAddressSpecWithCallback invokes the eipanycast.ModifyAnycastEipAddressSpec API asynchronously
func (client *Client) ModifyAnycastEipAddressSpecWithCallback(request *ModifyAnycastEipAddressSpecRequest, callback func(response *ModifyAnycastEipAddressSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyAnycastEipAddressSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyAnycastEipAddressSpec(request)
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

// ModifyAnycastEipAddressSpecRequest is the request struct for api ModifyAnycastEipAddressSpec
type ModifyAnycastEipAddressSpecRequest struct {
	*requests.RpcRequest
	Bandwidth string `position:"Query" name:"Bandwidth"`
	AnycastId string `position:"Query" name:"AnycastId"`
}

// ModifyAnycastEipAddressSpecResponse is the response struct for api ModifyAnycastEipAddressSpec
type ModifyAnycastEipAddressSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyAnycastEipAddressSpecRequest creates a request to invoke ModifyAnycastEipAddressSpec API
func CreateModifyAnycastEipAddressSpecRequest() (request *ModifyAnycastEipAddressSpecRequest) {
	request = &ModifyAnycastEipAddressSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Eipanycast", "2020-03-09", "ModifyAnycastEipAddressSpec", "eipanycast", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyAnycastEipAddressSpecResponse creates a response to parse from ModifyAnycastEipAddressSpec response
func CreateModifyAnycastEipAddressSpecResponse() (response *ModifyAnycastEipAddressSpecResponse) {
	response = &ModifyAnycastEipAddressSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}