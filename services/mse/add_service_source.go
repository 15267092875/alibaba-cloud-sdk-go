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

// AddServiceSource invokes the mse.AddServiceSource API synchronously
func (client *Client) AddServiceSource(request *AddServiceSourceRequest) (response *AddServiceSourceResponse, err error) {
	response = CreateAddServiceSourceResponse()
	err = client.DoAction(request, response)
	return
}

// AddServiceSourceWithChan invokes the mse.AddServiceSource API asynchronously
func (client *Client) AddServiceSourceWithChan(request *AddServiceSourceRequest) (<-chan *AddServiceSourceResponse, <-chan error) {
	responseChan := make(chan *AddServiceSourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddServiceSource(request)
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

// AddServiceSourceWithCallback invokes the mse.AddServiceSource API asynchronously
func (client *Client) AddServiceSourceWithCallback(request *AddServiceSourceRequest, callback func(response *AddServiceSourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddServiceSourceResponse
		var err error
		defer close(result)
		response, err = client.AddServiceSource(request)
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

// AddServiceSourceRequest is the request struct for api AddServiceSource
type AddServiceSourceRequest struct {
	*requests.RpcRequest
	IngressOptionsRequest     AddServiceSourceIngressOptionsRequest        `position:"Query" name:"IngressOptionsRequest"  type:"Struct"`
	MseSessionId              string                                       `position:"Query" name:"MseSessionId"`
	GatewayUniqueId           string                                       `position:"Query" name:"GatewayUniqueId"`
	Source                    string                                       `position:"Query" name:"Source"`
	Type                      string                                       `position:"Query" name:"Type"`
	PathList                  *[]string                                    `position:"Query" name:"PathList"  type:"Json"`
	Address                   string                                       `position:"Query" name:"Address"`
	ToAuthorizeSecurityGroups *[]AddServiceSourceToAuthorizeSecurityGroups `position:"Query" name:"ToAuthorizeSecurityGroups"  type:"Json"`
	Name                      string                                       `position:"Query" name:"Name"`
	AcceptLanguage            string                                       `position:"Query" name:"AcceptLanguage"`
	GroupList                 *[]string                                    `position:"Query" name:"GroupList"  type:"Json"`
}

// AddServiceSourceToAuthorizeSecurityGroups is a repeated param struct in AddServiceSourceRequest
type AddServiceSourceToAuthorizeSecurityGroups struct {
	PortRange       string `name:"PortRange"`
	SecurityGroupId string `name:"SecurityGroupId"`
	Description     string `name:"Description"`
}

// AddServiceSourceIngressOptionsRequest is a repeated param struct in AddServiceSourceRequest
type AddServiceSourceIngressOptionsRequest struct {
	EnableStatus   string `name:"EnableStatus"`
	EnableIngress  string `name:"EnableIngress"`
	WatchNamespace string `name:"WatchNamespace"`
	IngressClass   string `name:"IngressClass"`
}

// AddServiceSourceResponse is the response struct for api AddServiceSource
type AddServiceSourceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           int64  `json:"Data" xml:"Data"`
}

// CreateAddServiceSourceRequest creates a request to invoke AddServiceSource API
func CreateAddServiceSourceRequest() (request *AddServiceSourceRequest) {
	request = &AddServiceSourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "AddServiceSource", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddServiceSourceResponse creates a response to parse from AddServiceSource response
func CreateAddServiceSourceResponse() (response *AddServiceSourceResponse) {
	response = &AddServiceSourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
