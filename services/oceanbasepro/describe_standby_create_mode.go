package oceanbasepro

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

// DescribeStandbyCreateMode invokes the oceanbasepro.DescribeStandbyCreateMode API synchronously
func (client *Client) DescribeStandbyCreateMode(request *DescribeStandbyCreateModeRequest) (response *DescribeStandbyCreateModeResponse, err error) {
	response = CreateDescribeStandbyCreateModeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStandbyCreateModeWithChan invokes the oceanbasepro.DescribeStandbyCreateMode API asynchronously
func (client *Client) DescribeStandbyCreateModeWithChan(request *DescribeStandbyCreateModeRequest) (<-chan *DescribeStandbyCreateModeResponse, <-chan error) {
	responseChan := make(chan *DescribeStandbyCreateModeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStandbyCreateMode(request)
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

// DescribeStandbyCreateModeWithCallback invokes the oceanbasepro.DescribeStandbyCreateMode API asynchronously
func (client *Client) DescribeStandbyCreateModeWithCallback(request *DescribeStandbyCreateModeRequest, callback func(response *DescribeStandbyCreateModeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStandbyCreateModeResponse
		var err error
		defer close(result)
		response, err = client.DescribeStandbyCreateMode(request)
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

// DescribeStandbyCreateModeRequest is the request struct for api DescribeStandbyCreateMode
type DescribeStandbyCreateModeRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Body" name:"InstanceId"`
	TenantId   string `position:"Body" name:"TenantId"`
}

// DescribeStandbyCreateModeResponse is the response struct for api DescribeStandbyCreateMode
type DescribeStandbyCreateModeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeStandbyCreateModeRequest creates a request to invoke DescribeStandbyCreateMode API
func CreateDescribeStandbyCreateModeRequest() (request *DescribeStandbyCreateModeRequest) {
	request = &DescribeStandbyCreateModeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeStandbyCreateMode", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStandbyCreateModeResponse creates a response to parse from DescribeStandbyCreateMode response
func CreateDescribeStandbyCreateModeResponse() (response *DescribeStandbyCreateModeResponse) {
	response = &DescribeStandbyCreateModeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
