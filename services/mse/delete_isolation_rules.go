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

// DeleteIsolationRules invokes the mse.DeleteIsolationRules API synchronously
func (client *Client) DeleteIsolationRules(request *DeleteIsolationRulesRequest) (response *DeleteIsolationRulesResponse, err error) {
	response = CreateDeleteIsolationRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIsolationRulesWithChan invokes the mse.DeleteIsolationRules API asynchronously
func (client *Client) DeleteIsolationRulesWithChan(request *DeleteIsolationRulesRequest) (<-chan *DeleteIsolationRulesResponse, <-chan error) {
	responseChan := make(chan *DeleteIsolationRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIsolationRules(request)
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

// DeleteIsolationRulesWithCallback invokes the mse.DeleteIsolationRules API asynchronously
func (client *Client) DeleteIsolationRulesWithCallback(request *DeleteIsolationRulesRequest, callback func(response *DeleteIsolationRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIsolationRulesResponse
		var err error
		defer close(result)
		response, err = client.DeleteIsolationRules(request)
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

// DeleteIsolationRulesRequest is the request struct for api DeleteIsolationRules
type DeleteIsolationRulesRequest struct {
	*requests.RpcRequest
	MseSessionId   string    `position:"Query" name:"MseSessionId"`
	AppName        string    `position:"Query" name:"AppName"`
	Namespace      string    `position:"Query" name:"Namespace"`
	AcceptLanguage string    `position:"Query" name:"AcceptLanguage"`
	Ids            *[]string `position:"Query" name:"Ids"  type:"Json"`
}

// DeleteIsolationRulesResponse is the response struct for api DeleteIsolationRules
type DeleteIsolationRulesResponse struct {
	*responses.BaseResponse
	Code           int     `json:"Code" xml:"Code"`
	Message        string  `json:"Message" xml:"Message"`
	RequestId      string  `json:"RequestId" xml:"RequestId"`
	Success        bool    `json:"Success" xml:"Success"`
	HttpStatusCode int     `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           []int64 `json:"Data" xml:"Data"`
}

// CreateDeleteIsolationRulesRequest creates a request to invoke DeleteIsolationRules API
func CreateDeleteIsolationRulesRequest() (request *DeleteIsolationRulesRequest) {
	request = &DeleteIsolationRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "DeleteIsolationRules", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteIsolationRulesResponse creates a response to parse from DeleteIsolationRules response
func CreateDeleteIsolationRulesResponse() (response *DeleteIsolationRulesResponse) {
	response = &DeleteIsolationRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
