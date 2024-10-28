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

// ListFlowRules invokes the mse.ListFlowRules API synchronously
func (client *Client) ListFlowRules(request *ListFlowRulesRequest) (response *ListFlowRulesResponse, err error) {
	response = CreateListFlowRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowRulesWithChan invokes the mse.ListFlowRules API asynchronously
func (client *Client) ListFlowRulesWithChan(request *ListFlowRulesRequest) (<-chan *ListFlowRulesResponse, <-chan error) {
	responseChan := make(chan *ListFlowRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowRules(request)
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

// ListFlowRulesWithCallback invokes the mse.ListFlowRules API asynchronously
func (client *Client) ListFlowRulesWithCallback(request *ListFlowRulesRequest, callback func(response *ListFlowRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowRulesResponse
		var err error
		defer close(result)
		response, err = client.ListFlowRules(request)
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

// ListFlowRulesRequest is the request struct for api ListFlowRules
type ListFlowRulesRequest struct {
	*requests.RpcRequest
	MseSessionId      string           `position:"Query" name:"MseSessionId"`
	AppName           string           `position:"Query" name:"AppName"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	PageIndex         requests.Integer `position:"Query" name:"PageIndex"`
	Resource          string           `position:"Query" name:"Resource"`
	AppId             string           `position:"Query" name:"AppId"`
	Namespace         string           `position:"Query" name:"Namespace"`
	AcceptLanguage    string           `position:"Query" name:"AcceptLanguage"`
	ResourceSearchKey string           `position:"Query" name:"ResourceSearchKey"`
}

// ListFlowRulesResponse is the response struct for api ListFlowRules
type ListFlowRulesResponse struct {
	*responses.BaseResponse
	Code           int    `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListFlowRulesRequest creates a request to invoke ListFlowRules API
func CreateListFlowRulesRequest() (request *ListFlowRulesRequest) {
	request = &ListFlowRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListFlowRules", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFlowRulesResponse creates a response to parse from ListFlowRules response
func CreateListFlowRulesResponse() (response *ListFlowRulesResponse) {
	response = &ListFlowRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
