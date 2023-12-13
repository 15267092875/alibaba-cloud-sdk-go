package cloud_siem

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

// DescribeStorage invokes the cloud_siem.DescribeStorage API synchronously
func (client *Client) DescribeStorage(request *DescribeStorageRequest) (response *DescribeStorageResponse, err error) {
	response = CreateDescribeStorageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStorageWithChan invokes the cloud_siem.DescribeStorage API asynchronously
func (client *Client) DescribeStorageWithChan(request *DescribeStorageRequest) (<-chan *DescribeStorageResponse, <-chan error) {
	responseChan := make(chan *DescribeStorageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStorage(request)
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

// DescribeStorageWithCallback invokes the cloud_siem.DescribeStorage API asynchronously
func (client *Client) DescribeStorageWithCallback(request *DescribeStorageRequest, callback func(response *DescribeStorageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStorageResponse
		var err error
		defer close(result)
		response, err = client.DescribeStorage(request)
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

// DescribeStorageRequest is the request struct for api DescribeStorage
type DescribeStorageRequest struct {
	*requests.RpcRequest
}

// DescribeStorageResponse is the response struct for api DescribeStorage
type DescribeStorageResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDescribeStorageRequest creates a request to invoke DescribeStorage API
func CreateDescribeStorageRequest() (request *DescribeStorageRequest) {
	request = &DescribeStorageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloud-siem", "2022-06-16", "DescribeStorage", "cloud-siem", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStorageResponse creates a response to parse from DescribeStorage response
func CreateDescribeStorageResponse() (response *DescribeStorageResponse) {
	response = &DescribeStorageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
