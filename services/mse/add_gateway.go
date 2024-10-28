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

// AddGateway invokes the mse.AddGateway API synchronously
func (client *Client) AddGateway(request *AddGatewayRequest) (response *AddGatewayResponse, err error) {
	response = CreateAddGatewayResponse()
	err = client.DoAction(request, response)
	return
}

// AddGatewayWithChan invokes the mse.AddGateway API asynchronously
func (client *Client) AddGatewayWithChan(request *AddGatewayRequest) (<-chan *AddGatewayResponse, <-chan error) {
	responseChan := make(chan *AddGatewayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddGateway(request)
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

// AddGatewayWithCallback invokes the mse.AddGateway API asynchronously
func (client *Client) AddGatewayWithCallback(request *AddGatewayRequest, callback func(response *AddGatewayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddGatewayResponse
		var err error
		defer close(result)
		response, err = client.AddGateway(request)
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

// AddGatewayRequest is the request struct for api AddGateway
type AddGatewayRequest struct {
	*requests.RpcRequest
	MseSessionId               string                `position:"Query" name:"MseSessionId"`
	InternetSlbSpec            string                `position:"Query" name:"InternetSlbSpec"`
	EnableXtrace               requests.Boolean      `position:"Query" name:"EnableXtrace"`
	Replica                    requests.Integer      `position:"Query" name:"Replica"`
	EnableHardwareAcceleration requests.Boolean      `position:"Query" name:"EnableHardwareAcceleration"`
	EnableSls                  requests.Boolean      `position:"Query" name:"EnableSls"`
	Spec                       string                `position:"Query" name:"Spec"`
	ResourceGroupId            string                `position:"Query" name:"ResourceGroupId"`
	RequestPars                string                `position:"Query" name:"RequestPars"`
	EnterpriseSecurityGroup    requests.Boolean      `position:"Query" name:"EnterpriseSecurityGroup"`
	Tag                        *[]AddGatewayTag      `position:"Query" name:"Tag"  type:"Repeated"`
	VSwitchId                  string                `position:"Query" name:"VSwitchId"`
	SlbSpec                    string                `position:"Query" name:"SlbSpec"`
	Name                       string                `position:"Query" name:"Name"`
	Region                     string                `position:"Query" name:"Region"`
	MserVersion                string                `position:"Query" name:"MserVersion"`
	ZoneInfo                   *[]AddGatewayZoneInfo `position:"Query" name:"ZoneInfo"  type:"Json"`
	XtraceRatio                string                `position:"Query" name:"XtraceRatio"`
	VSwitchId2                 string                `position:"Query" name:"VSwitchId2"`
	ClbNetworkType             string                `position:"Query" name:"ClbNetworkType"`
	Vpc                        string                `position:"Query" name:"Vpc"`
	NlbNetworkType             string                `position:"Query" name:"NlbNetworkType"`
	AcceptLanguage             string                `position:"Query" name:"AcceptLanguage"`
	ChargeType                 string                `position:"Query" name:"ChargeType"`
}

// AddGatewayZoneInfo is a repeated param struct in AddGatewayRequest
type AddGatewayZoneInfo struct {
	VSwitchId string `name:"VSwitchId"`
	ZoneId    string `name:"ZoneId"`
}

// AddGatewayTag is a repeated param struct in AddGatewayRequest
type AddGatewayTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// AddGatewayResponse is the response struct for api AddGateway
type AddGatewayResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	Code           int    `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateAddGatewayRequest creates a request to invoke AddGateway API
func CreateAddGatewayRequest() (request *AddGatewayRequest) {
	request = &AddGatewayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "AddGateway", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddGatewayResponse creates a response to parse from AddGateway response
func CreateAddGatewayResponse() (response *AddGatewayResponse) {
	response = &AddGatewayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
