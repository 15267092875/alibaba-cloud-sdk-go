package kms

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

// SetSecretPolicy invokes the kms.SetSecretPolicy API synchronously
func (client *Client) SetSecretPolicy(request *SetSecretPolicyRequest) (response *SetSecretPolicyResponse, err error) {
	response = CreateSetSecretPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// SetSecretPolicyWithChan invokes the kms.SetSecretPolicy API asynchronously
func (client *Client) SetSecretPolicyWithChan(request *SetSecretPolicyRequest) (<-chan *SetSecretPolicyResponse, <-chan error) {
	responseChan := make(chan *SetSecretPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetSecretPolicy(request)
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

// SetSecretPolicyWithCallback invokes the kms.SetSecretPolicy API asynchronously
func (client *Client) SetSecretPolicyWithCallback(request *SetSecretPolicyRequest, callback func(response *SetSecretPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetSecretPolicyResponse
		var err error
		defer close(result)
		response, err = client.SetSecretPolicy(request)
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

// SetSecretPolicyRequest is the request struct for api SetSecretPolicy
type SetSecretPolicyRequest struct {
	*requests.RpcRequest
	PolicyName string `position:"Query" name:"PolicyName"`
	SecretName string `position:"Query" name:"SecretName"`
	Policy     string `position:"Query" name:"Policy"`
}

// SetSecretPolicyResponse is the response struct for api SetSecretPolicy
type SetSecretPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetSecretPolicyRequest creates a request to invoke SetSecretPolicy API
func CreateSetSecretPolicyRequest() (request *SetSecretPolicyRequest) {
	request = &SetSecretPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "SetSecretPolicy", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetSecretPolicyResponse creates a response to parse from SetSecretPolicy response
func CreateSetSecretPolicyResponse() (response *SetSecretPolicyResponse) {
	response = &SetSecretPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
