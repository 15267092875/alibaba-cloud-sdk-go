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

// GetUserMFAInfo invokes the ims.GetUserMFAInfo API synchronously
func (client *Client) GetUserMFAInfo(request *GetUserMFAInfoRequest) (response *GetUserMFAInfoResponse, err error) {
	response = CreateGetUserMFAInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserMFAInfoWithChan invokes the ims.GetUserMFAInfo API asynchronously
func (client *Client) GetUserMFAInfoWithChan(request *GetUserMFAInfoRequest) (<-chan *GetUserMFAInfoResponse, <-chan error) {
	responseChan := make(chan *GetUserMFAInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserMFAInfo(request)
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

// GetUserMFAInfoWithCallback invokes the ims.GetUserMFAInfo API asynchronously
func (client *Client) GetUserMFAInfoWithCallback(request *GetUserMFAInfoRequest, callback func(response *GetUserMFAInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserMFAInfoResponse
		var err error
		defer close(result)
		response, err = client.GetUserMFAInfo(request)
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

// GetUserMFAInfoRequest is the request struct for api GetUserMFAInfo
type GetUserMFAInfoRequest struct {
	*requests.RpcRequest
	AkProxySuffix     string `position:"Query" name:"AkProxySuffix"`
	UserPrincipalName string `position:"Query" name:"UserPrincipalName"`
}

// GetUserMFAInfoResponse is the response struct for api GetUserMFAInfo
type GetUserMFAInfoResponse struct {
	*responses.BaseResponse
	RequestId   string    `json:"RequestId" xml:"RequestId"`
	IsMFAEnable bool      `json:"IsMFAEnable" xml:"IsMFAEnable"`
	MFADevice   MFADevice `json:"MFADevice" xml:"MFADevice"`
}

// CreateGetUserMFAInfoRequest creates a request to invoke GetUserMFAInfo API
func CreateGetUserMFAInfoRequest() (request *GetUserMFAInfoRequest) {
	request = &GetUserMFAInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ims", "2019-08-15", "GetUserMFAInfo", "ims", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetUserMFAInfoResponse creates a response to parse from GetUserMFAInfo response
func CreateGetUserMFAInfoResponse() (response *GetUserMFAInfoResponse) {
	response = &GetUserMFAInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
