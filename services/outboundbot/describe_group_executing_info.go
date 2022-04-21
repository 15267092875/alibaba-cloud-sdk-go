package outboundbot

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

// DescribeGroupExecutingInfo invokes the outboundbot.DescribeGroupExecutingInfo API synchronously
func (client *Client) DescribeGroupExecutingInfo(request *DescribeGroupExecutingInfoRequest) (response *DescribeGroupExecutingInfoResponse, err error) {
	response = CreateDescribeGroupExecutingInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGroupExecutingInfoWithChan invokes the outboundbot.DescribeGroupExecutingInfo API asynchronously
func (client *Client) DescribeGroupExecutingInfoWithChan(request *DescribeGroupExecutingInfoRequest) (<-chan *DescribeGroupExecutingInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeGroupExecutingInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGroupExecutingInfo(request)
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

// DescribeGroupExecutingInfoWithCallback invokes the outboundbot.DescribeGroupExecutingInfo API asynchronously
func (client *Client) DescribeGroupExecutingInfoWithCallback(request *DescribeGroupExecutingInfoRequest, callback func(response *DescribeGroupExecutingInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGroupExecutingInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeGroupExecutingInfo(request)
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

// DescribeGroupExecutingInfoRequest is the request struct for api DescribeGroupExecutingInfo
type DescribeGroupExecutingInfoRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	JobGroupId string `position:"Query" name:"JobGroupId"`
}

// DescribeGroupExecutingInfoResponse is the response struct for api DescribeGroupExecutingInfo
type DescribeGroupExecutingInfoResponse struct {
	*responses.BaseResponse
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	GroupId        string        `json:"GroupId" xml:"GroupId"`
	Success        bool          `json:"Success" xml:"Success"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	InstanceId     string        `json:"InstanceId" xml:"InstanceId"`
	ExecutingInfo  ExecutingInfo `json:"ExecutingInfo" xml:"ExecutingInfo"`
}

// CreateDescribeGroupExecutingInfoRequest creates a request to invoke DescribeGroupExecutingInfo API
func CreateDescribeGroupExecutingInfoRequest() (request *DescribeGroupExecutingInfoRequest) {
	request = &DescribeGroupExecutingInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "DescribeGroupExecutingInfo", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGroupExecutingInfoResponse creates a response to parse from DescribeGroupExecutingInfo response
func CreateDescribeGroupExecutingInfoResponse() (response *DescribeGroupExecutingInfoResponse) {
	response = &DescribeGroupExecutingInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
