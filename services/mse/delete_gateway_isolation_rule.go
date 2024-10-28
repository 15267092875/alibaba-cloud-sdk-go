package mse

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

// DeleteGatewayIsolationRule invokes the mse.DeleteGatewayIsolationRule API synchronously
func (client *Client) DeleteGatewayIsolationRule(request *DeleteGatewayIsolationRuleRequest) (response *DeleteGatewayIsolationRuleResponse, err error) {
	response = CreateDeleteGatewayIsolationRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteGatewayIsolationRuleWithChan invokes the mse.DeleteGatewayIsolationRule API asynchronously
func (client *Client) DeleteGatewayIsolationRuleWithChan(request *DeleteGatewayIsolationRuleRequest) (<-chan *DeleteGatewayIsolationRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteGatewayIsolationRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteGatewayIsolationRule(request)
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

// DeleteGatewayIsolationRuleWithCallback invokes the mse.DeleteGatewayIsolationRule API asynchronously
func (client *Client) DeleteGatewayIsolationRuleWithCallback(request *DeleteGatewayIsolationRuleRequest, callback func(response *DeleteGatewayIsolationRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteGatewayIsolationRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteGatewayIsolationRule(request)
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

// DeleteGatewayIsolationRuleRequest is the request struct for api DeleteGatewayIsolationRule
type DeleteGatewayIsolationRuleRequest struct {
	*requests.RpcRequest
	MseSessionId    string           `position:"Query" name:"MseSessionId"`
	GatewayUniqueId string           `position:"Query" name:"GatewayUniqueId"`
	RouteId         requests.Integer `position:"Query" name:"RouteId"`
	AcceptLanguage  string           `position:"Query" name:"AcceptLanguage"`
	RuleId          requests.Integer `position:"Query" name:"RuleId"`
}

// DeleteGatewayIsolationRuleResponse is the response struct for api DeleteGatewayIsolationRule
type DeleteGatewayIsolationRuleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteGatewayIsolationRuleRequest creates a request to invoke DeleteGatewayIsolationRule API
func CreateDeleteGatewayIsolationRuleRequest() (request *DeleteGatewayIsolationRuleRequest) {
	request = &DeleteGatewayIsolationRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteGatewayIsolationRule", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteGatewayIsolationRuleResponse creates a response to parse from DeleteGatewayIsolationRule response
func CreateDeleteGatewayIsolationRuleResponse() (response *DeleteGatewayIsolationRuleResponse) {
	response = &DeleteGatewayIsolationRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
