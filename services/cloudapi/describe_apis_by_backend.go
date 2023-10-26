package cloudapi

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

// DescribeApisByBackend invokes the cloudapi.DescribeApisByBackend API synchronously
func (client *Client) DescribeApisByBackend(request *DescribeApisByBackendRequest) (response *DescribeApisByBackendResponse, err error) {
	response = CreateDescribeApisByBackendResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeApisByBackendWithChan invokes the cloudapi.DescribeApisByBackend API asynchronously
func (client *Client) DescribeApisByBackendWithChan(request *DescribeApisByBackendRequest) (<-chan *DescribeApisByBackendResponse, <-chan error) {
	responseChan := make(chan *DescribeApisByBackendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeApisByBackend(request)
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

// DescribeApisByBackendWithCallback invokes the cloudapi.DescribeApisByBackend API asynchronously
func (client *Client) DescribeApisByBackendWithCallback(request *DescribeApisByBackendRequest, callback func(response *DescribeApisByBackendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeApisByBackendResponse
		var err error
		defer close(result)
		response, err = client.DescribeApisByBackend(request)
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

// DescribeApisByBackendRequest is the request struct for api DescribeApisByBackend
type DescribeApisByBackendRequest struct {
	*requests.RpcRequest
	StageName     string           `position:"Query" name:"StageName"`
	BackendId     string           `position:"Query" name:"BackendId"`
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeApisByBackendResponse is the response struct for api DescribeApisByBackend
type DescribeApisByBackendResponse struct {
	*responses.BaseResponse
	PageNumber  int                                `json:"PageNumber" xml:"PageNumber"`
	RequestId   string                             `json:"RequestId" xml:"RequestId"`
	PageSize    int                                `json:"PageSize" xml:"PageSize"`
	TotalCount  int                                `json:"TotalCount" xml:"TotalCount"`
	ApiInfoList ApiInfoListInDescribeApisByBackend `json:"ApiInfoList" xml:"ApiInfoList"`
}

// CreateDescribeApisByBackendRequest creates a request to invoke DescribeApisByBackend API
func CreateDescribeApisByBackendRequest() (request *DescribeApisByBackendRequest) {
	request = &DescribeApisByBackendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeApisByBackend", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeApisByBackendResponse creates a response to parse from DescribeApisByBackend response
func CreateDescribeApisByBackendResponse() (response *DescribeApisByBackendResponse) {
	response = &DescribeApisByBackendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
