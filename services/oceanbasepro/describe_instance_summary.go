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

// DescribeInstanceSummary invokes the oceanbasepro.DescribeInstanceSummary API synchronously
func (client *Client) DescribeInstanceSummary(request *DescribeInstanceSummaryRequest) (response *DescribeInstanceSummaryResponse, err error) {
	response = CreateDescribeInstanceSummaryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceSummaryWithChan invokes the oceanbasepro.DescribeInstanceSummary API asynchronously
func (client *Client) DescribeInstanceSummaryWithChan(request *DescribeInstanceSummaryRequest) (<-chan *DescribeInstanceSummaryResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceSummaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceSummary(request)
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

// DescribeInstanceSummaryWithCallback invokes the oceanbasepro.DescribeInstanceSummary API asynchronously
func (client *Client) DescribeInstanceSummaryWithCallback(request *DescribeInstanceSummaryRequest, callback func(response *DescribeInstanceSummaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceSummaryResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceSummary(request)
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

// DescribeInstanceSummaryRequest is the request struct for api DescribeInstanceSummary
type DescribeInstanceSummaryRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Body" name:"PageNumber"`
	PageSize   requests.Integer `position:"Body" name:"PageSize"`
}

// DescribeInstanceSummaryResponse is the response struct for api DescribeInstanceSummary
type DescribeInstanceSummaryResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	InstanceSummary InstanceSummary `json:"InstanceSummary" xml:"InstanceSummary"`
}

// CreateDescribeInstanceSummaryRequest creates a request to invoke DescribeInstanceSummary API
func CreateDescribeInstanceSummaryRequest() (request *DescribeInstanceSummaryRequest) {
	request = &DescribeInstanceSummaryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeInstanceSummary", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceSummaryResponse creates a response to parse from DescribeInstanceSummary response
func CreateDescribeInstanceSummaryResponse() (response *DescribeInstanceSummaryResponse) {
	response = &DescribeInstanceSummaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
