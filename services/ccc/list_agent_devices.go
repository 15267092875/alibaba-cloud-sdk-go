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

// ListAgentDevices invokes the ccc.ListAgentDevices API synchronously
func (client *Client) ListAgentDevices(request *ListAgentDevicesRequest) (response *ListAgentDevicesResponse, err error) {
	response = CreateListAgentDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// ListAgentDevicesWithChan invokes the ccc.ListAgentDevices API asynchronously
func (client *Client) ListAgentDevicesWithChan(request *ListAgentDevicesRequest) (<-chan *ListAgentDevicesResponse, <-chan error) {
	responseChan := make(chan *ListAgentDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAgentDevices(request)
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

// ListAgentDevicesWithCallback invokes the ccc.ListAgentDevices API asynchronously
func (client *Client) ListAgentDevicesWithCallback(request *ListAgentDevicesRequest, callback func(response *ListAgentDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAgentDevicesResponse
		var err error
		defer close(result)
		response, err = client.ListAgentDevices(request)
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

// ListAgentDevicesRequest is the request struct for api ListAgentDevices
type ListAgentDevicesRequest struct {
	*requests.RpcRequest
	RamIds     string           `position:"Query" name:"RamIds"`
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	StopTime   requests.Integer `position:"Query" name:"StopTime"`
	InstanceId string           `position:"Query" name:"InstanceId"`
}

// ListAgentDevicesResponse is the response struct for api ListAgentDevices
type ListAgentDevicesResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	Success         bool            `json:"Success" xml:"Success"`
	Code            string          `json:"Code" xml:"Code"`
	Message         string          `json:"Message" xml:"Message"`
	HttpStatusCode  int             `json:"HttpStatusCode" xml:"HttpStatusCode"`
	AgentDeviceList AgentDeviceList `json:"AgentDeviceList" xml:"AgentDeviceList"`
}

// CreateListAgentDevicesRequest creates a request to invoke ListAgentDevices API
func CreateListAgentDevicesRequest() (request *ListAgentDevicesRequest) {
	request = &ListAgentDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListAgentDevices", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAgentDevicesResponse creates a response to parse from ListAgentDevices response
func CreateListAgentDevicesResponse() (response *ListAgentDevicesResponse) {
	response = &ListAgentDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
