package ddoscoo

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

// ConfigLayer4RulePolicy invokes the ddoscoo.ConfigLayer4RulePolicy API synchronously
func (client *Client) ConfigLayer4RulePolicy(request *ConfigLayer4RulePolicyRequest) (response *ConfigLayer4RulePolicyResponse, err error) {
	response = CreateConfigLayer4RulePolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigLayer4RulePolicyWithChan invokes the ddoscoo.ConfigLayer4RulePolicy API asynchronously
func (client *Client) ConfigLayer4RulePolicyWithChan(request *ConfigLayer4RulePolicyRequest) (<-chan *ConfigLayer4RulePolicyResponse, <-chan error) {
	responseChan := make(chan *ConfigLayer4RulePolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigLayer4RulePolicy(request)
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

// ConfigLayer4RulePolicyWithCallback invokes the ddoscoo.ConfigLayer4RulePolicy API asynchronously
func (client *Client) ConfigLayer4RulePolicyWithCallback(request *ConfigLayer4RulePolicyRequest, callback func(response *ConfigLayer4RulePolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigLayer4RulePolicyResponse
		var err error
		defer close(result)
		response, err = client.ConfigLayer4RulePolicy(request)
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

// ConfigLayer4RulePolicyRequest is the request struct for api ConfigLayer4RulePolicy
type ConfigLayer4RulePolicyRequest struct {
	*requests.RpcRequest
	Listeners string `position:"Query" name:"Listeners"`
	SourceIp  string `position:"Query" name:"SourceIp"`
}

// ConfigLayer4RulePolicyResponse is the response struct for api ConfigLayer4RulePolicy
type ConfigLayer4RulePolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateConfigLayer4RulePolicyRequest creates a request to invoke ConfigLayer4RulePolicy API
func CreateConfigLayer4RulePolicyRequest() (request *ConfigLayer4RulePolicyRequest) {
	request = &ConfigLayer4RulePolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "ConfigLayer4RulePolicy", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigLayer4RulePolicyResponse creates a response to parse from ConfigLayer4RulePolicy response
func CreateConfigLayer4RulePolicyResponse() (response *ConfigLayer4RulePolicyResponse) {
	response = &ConfigLayer4RulePolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
