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

// ModifyEmptyNumberNoMoreCallsInfo invokes the outboundbot.ModifyEmptyNumberNoMoreCallsInfo API synchronously
func (client *Client) ModifyEmptyNumberNoMoreCallsInfo(request *ModifyEmptyNumberNoMoreCallsInfoRequest) (response *ModifyEmptyNumberNoMoreCallsInfoResponse, err error) {
	response = CreateModifyEmptyNumberNoMoreCallsInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyEmptyNumberNoMoreCallsInfoWithChan invokes the outboundbot.ModifyEmptyNumberNoMoreCallsInfo API asynchronously
func (client *Client) ModifyEmptyNumberNoMoreCallsInfoWithChan(request *ModifyEmptyNumberNoMoreCallsInfoRequest) (<-chan *ModifyEmptyNumberNoMoreCallsInfoResponse, <-chan error) {
	responseChan := make(chan *ModifyEmptyNumberNoMoreCallsInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyEmptyNumberNoMoreCallsInfo(request)
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

// ModifyEmptyNumberNoMoreCallsInfoWithCallback invokes the outboundbot.ModifyEmptyNumberNoMoreCallsInfo API asynchronously
func (client *Client) ModifyEmptyNumberNoMoreCallsInfoWithCallback(request *ModifyEmptyNumberNoMoreCallsInfoRequest, callback func(response *ModifyEmptyNumberNoMoreCallsInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyEmptyNumberNoMoreCallsInfoResponse
		var err error
		defer close(result)
		response, err = client.ModifyEmptyNumberNoMoreCallsInfo(request)
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

// ModifyEmptyNumberNoMoreCallsInfoRequest is the request struct for api ModifyEmptyNumberNoMoreCallsInfo
type ModifyEmptyNumberNoMoreCallsInfoRequest struct {
	*requests.RpcRequest
	StrategyLevel          requests.Integer `position:"Query" name:"StrategyLevel"`
	EmptyNumberNoMoreCalls requests.Boolean `position:"Query" name:"EmptyNumberNoMoreCalls"`
	EntryId                string           `position:"Query" name:"EntryId"`
}

// ModifyEmptyNumberNoMoreCallsInfoResponse is the response struct for api ModifyEmptyNumberNoMoreCallsInfo
type ModifyEmptyNumberNoMoreCallsInfoResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateModifyEmptyNumberNoMoreCallsInfoRequest creates a request to invoke ModifyEmptyNumberNoMoreCallsInfo API
func CreateModifyEmptyNumberNoMoreCallsInfoRequest() (request *ModifyEmptyNumberNoMoreCallsInfoRequest) {
	request = &ModifyEmptyNumberNoMoreCallsInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "ModifyEmptyNumberNoMoreCallsInfo", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyEmptyNumberNoMoreCallsInfoResponse creates a response to parse from ModifyEmptyNumberNoMoreCallsInfo response
func CreateModifyEmptyNumberNoMoreCallsInfoResponse() (response *ModifyEmptyNumberNoMoreCallsInfoResponse) {
	response = &ModifyEmptyNumberNoMoreCallsInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
