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

// GetEffectiveDays invokes the outboundbot.GetEffectiveDays API synchronously
func (client *Client) GetEffectiveDays(request *GetEffectiveDaysRequest) (response *GetEffectiveDaysResponse, err error) {
	response = CreateGetEffectiveDaysResponse()
	err = client.DoAction(request, response)
	return
}

// GetEffectiveDaysWithChan invokes the outboundbot.GetEffectiveDays API asynchronously
func (client *Client) GetEffectiveDaysWithChan(request *GetEffectiveDaysRequest) (<-chan *GetEffectiveDaysResponse, <-chan error) {
	responseChan := make(chan *GetEffectiveDaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetEffectiveDays(request)
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

// GetEffectiveDaysWithCallback invokes the outboundbot.GetEffectiveDays API asynchronously
func (client *Client) GetEffectiveDaysWithCallback(request *GetEffectiveDaysRequest, callback func(response *GetEffectiveDaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetEffectiveDaysResponse
		var err error
		defer close(result)
		response, err = client.GetEffectiveDays(request)
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

// GetEffectiveDaysRequest is the request struct for api GetEffectiveDays
type GetEffectiveDaysRequest struct {
	*requests.RpcRequest
	StrategyLevel requests.Integer `position:"Query" name:"StrategyLevel"`
	EntryId       string           `position:"Query" name:"EntryId"`
}

// GetEffectiveDaysResponse is the response struct for api GetEffectiveDays
type GetEffectiveDaysResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	EffectiveDays  int    `json:"EffectiveDays" xml:"EffectiveDays"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
}

// CreateGetEffectiveDaysRequest creates a request to invoke GetEffectiveDays API
func CreateGetEffectiveDaysRequest() (request *GetEffectiveDaysRequest) {
	request = &GetEffectiveDaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OutboundBot", "2019-12-26", "GetEffectiveDays", "outboundbot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetEffectiveDaysResponse creates a response to parse from GetEffectiveDays response
func CreateGetEffectiveDaysResponse() (response *GetEffectiveDaysResponse) {
	response = &GetEffectiveDaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
