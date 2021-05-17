package ccc

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

// GetJobFileUploadUrl invokes the ccc.GetJobFileUploadUrl API synchronously
func (client *Client) GetJobFileUploadUrl(request *GetJobFileUploadUrlRequest) (response *GetJobFileUploadUrlResponse, err error) {
	response = CreateGetJobFileUploadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobFileUploadUrlWithChan invokes the ccc.GetJobFileUploadUrl API asynchronously
func (client *Client) GetJobFileUploadUrlWithChan(request *GetJobFileUploadUrlRequest) (<-chan *GetJobFileUploadUrlResponse, <-chan error) {
	responseChan := make(chan *GetJobFileUploadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobFileUploadUrl(request)
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

// GetJobFileUploadUrlWithCallback invokes the ccc.GetJobFileUploadUrl API asynchronously
func (client *Client) GetJobFileUploadUrlWithCallback(request *GetJobFileUploadUrlRequest, callback func(response *GetJobFileUploadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobFileUploadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetJobFileUploadUrl(request)
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

// GetJobFileUploadUrlRequest is the request struct for api GetJobFileUploadUrl
type GetJobFileUploadUrlRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	FileName   string `position:"Query" name:"FileName"`
}

// GetJobFileUploadUrlResponse is the response struct for api GetJobFileUploadUrl
type GetJobFileUploadUrlResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	FilePath       string `json:"FilePath" xml:"FilePath"`
	UploadUrl      string `json:"UploadUrl" xml:"UploadUrl"`
}

// CreateGetJobFileUploadUrlRequest creates a request to invoke GetJobFileUploadUrl API
func CreateGetJobFileUploadUrlRequest() (request *GetJobFileUploadUrlRequest) {
	request = &GetJobFileUploadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetJobFileUploadUrl", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetJobFileUploadUrlResponse creates a response to parse from GetJobFileUploadUrl response
func CreateGetJobFileUploadUrlResponse() (response *GetJobFileUploadUrlResponse) {
	response = &GetJobFileUploadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
