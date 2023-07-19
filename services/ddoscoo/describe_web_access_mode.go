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

// DescribeWebAccessMode invokes the ddoscoo.DescribeWebAccessMode API synchronously
func (client *Client) DescribeWebAccessMode(request *DescribeWebAccessModeRequest) (response *DescribeWebAccessModeResponse, err error) {
	response = CreateDescribeWebAccessModeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebAccessModeWithChan invokes the ddoscoo.DescribeWebAccessMode API asynchronously
func (client *Client) DescribeWebAccessModeWithChan(request *DescribeWebAccessModeRequest) (<-chan *DescribeWebAccessModeResponse, <-chan error) {
	responseChan := make(chan *DescribeWebAccessModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebAccessMode(request)
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

// DescribeWebAccessModeWithCallback invokes the ddoscoo.DescribeWebAccessMode API asynchronously
func (client *Client) DescribeWebAccessModeWithCallback(request *DescribeWebAccessModeRequest, callback func(response *DescribeWebAccessModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebAccessModeResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebAccessMode(request)
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

// DescribeWebAccessModeRequest is the request struct for api DescribeWebAccessMode
type DescribeWebAccessModeRequest struct {
	*requests.RpcRequest
	Domains  *[]string `position:"Query" name:"Domains"  type:"Repeated"`
	SourceIp string    `position:"Query" name:"SourceIp"`
}

// DescribeWebAccessModeResponse is the response struct for api DescribeWebAccessMode
type DescribeWebAccessModeResponse struct {
	*responses.BaseResponse
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	DomainModes []DomainMode `json:"DomainModes" xml:"DomainModes"`
}

// CreateDescribeWebAccessModeRequest creates a request to invoke DescribeWebAccessMode API
func CreateDescribeWebAccessModeRequest() (request *DescribeWebAccessModeRequest) {
	request = &DescribeWebAccessModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebAccessMode", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWebAccessModeResponse creates a response to parse from DescribeWebAccessMode response
func CreateDescribeWebAccessModeResponse() (response *DescribeWebAccessModeResponse) {
	response = &DescribeWebAccessModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
