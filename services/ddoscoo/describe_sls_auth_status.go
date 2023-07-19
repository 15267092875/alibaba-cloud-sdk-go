package ddoscoo

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

// DescribeSlsAuthStatus invokes the ddoscoo.DescribeSlsAuthStatus API synchronously
func (client *Client) DescribeSlsAuthStatus(request *DescribeSlsAuthStatusRequest) (response *DescribeSlsAuthStatusResponse, err error) {
	response = CreateDescribeSlsAuthStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSlsAuthStatusWithChan invokes the ddoscoo.DescribeSlsAuthStatus API asynchronously
func (client *Client) DescribeSlsAuthStatusWithChan(request *DescribeSlsAuthStatusRequest) (<-chan *DescribeSlsAuthStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeSlsAuthStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSlsAuthStatus(request)
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

// DescribeSlsAuthStatusWithCallback invokes the ddoscoo.DescribeSlsAuthStatus API asynchronously
func (client *Client) DescribeSlsAuthStatusWithCallback(request *DescribeSlsAuthStatusRequest, callback func(response *DescribeSlsAuthStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSlsAuthStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeSlsAuthStatus(request)
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

// DescribeSlsAuthStatusRequest is the request struct for api DescribeSlsAuthStatus
type DescribeSlsAuthStatusRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeSlsAuthStatusResponse is the response struct for api DescribeSlsAuthStatus
type DescribeSlsAuthStatusResponse struct {
	*responses.BaseResponse
	SlsAuthStatus bool   `json:"SlsAuthStatus" xml:"SlsAuthStatus"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeSlsAuthStatusRequest creates a request to invoke DescribeSlsAuthStatus API
func CreateDescribeSlsAuthStatusRequest() (request *DescribeSlsAuthStatusRequest) {
	request = &DescribeSlsAuthStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeSlsAuthStatus", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSlsAuthStatusResponse creates a response to parse from DescribeSlsAuthStatus response
func CreateDescribeSlsAuthStatusResponse() (response *DescribeSlsAuthStatusResponse) {
	response = &DescribeSlsAuthStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
