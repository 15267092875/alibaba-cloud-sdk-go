package cloudfw

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

// DescribeTrFirewallsV2Detail invokes the cloudfw.DescribeTrFirewallsV2Detail API synchronously
func (client *Client) DescribeTrFirewallsV2Detail(request *DescribeTrFirewallsV2DetailRequest) (response *DescribeTrFirewallsV2DetailResponse, err error) {
	response = CreateDescribeTrFirewallsV2DetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTrFirewallsV2DetailWithChan invokes the cloudfw.DescribeTrFirewallsV2Detail API asynchronously
func (client *Client) DescribeTrFirewallsV2DetailWithChan(request *DescribeTrFirewallsV2DetailRequest) (<-chan *DescribeTrFirewallsV2DetailResponse, <-chan error) {
	responseChan := make(chan *DescribeTrFirewallsV2DetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTrFirewallsV2Detail(request)
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

// DescribeTrFirewallsV2DetailWithCallback invokes the cloudfw.DescribeTrFirewallsV2Detail API asynchronously
func (client *Client) DescribeTrFirewallsV2DetailWithCallback(request *DescribeTrFirewallsV2DetailRequest, callback func(response *DescribeTrFirewallsV2DetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTrFirewallsV2DetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeTrFirewallsV2Detail(request)
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

// DescribeTrFirewallsV2DetailRequest is the request struct for api DescribeTrFirewallsV2Detail
type DescribeTrFirewallsV2DetailRequest struct {
	*requests.RpcRequest
	FirewallId string `position:"Query" name:"FirewallId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Lang       string `position:"Query" name:"Lang"`
}

// DescribeTrFirewallsV2DetailResponse is the response struct for api DescribeTrFirewallsV2Detail
type DescribeTrFirewallsV2DetailResponse struct {
	*responses.BaseResponse
	FirewallName         string `json:"FirewallName" xml:"FirewallName"`
	RequestId            string `json:"RequestId" xml:"RequestId"`
	FirewallId           string `json:"FirewallId" xml:"FirewallId"`
	FirewallStatus       string `json:"FirewallStatus" xml:"FirewallStatus"`
	CenId                string `json:"CenId" xml:"CenId"`
	TransitRouterId      string `json:"TransitRouterId" xml:"TransitRouterId"`
	RegionNo             string `json:"RegionNo" xml:"RegionNo"`
	FirewallEniId        string `json:"FirewallEniId" xml:"FirewallEniId"`
	FirewallEniVpcId     string `json:"FirewallEniVpcId" xml:"FirewallEniVpcId"`
	FirewallEniVswitchId string `json:"FirewallEniVswitchId" xml:"FirewallEniVswitchId"`
	FirewallDescription  string `json:"FirewallDescription" xml:"FirewallDescription"`
	RouteMode            string `json:"RouteMode" xml:"RouteMode"`
	FirewallSwitchStatus string `json:"FirewallSwitchStatus" xml:"FirewallSwitchStatus"`
}

// CreateDescribeTrFirewallsV2DetailRequest creates a request to invoke DescribeTrFirewallsV2Detail API
func CreateDescribeTrFirewallsV2DetailRequest() (request *DescribeTrFirewallsV2DetailRequest) {
	request = &DescribeTrFirewallsV2DetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeTrFirewallsV2Detail", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTrFirewallsV2DetailResponse creates a response to parse from DescribeTrFirewallsV2Detail response
func CreateDescribeTrFirewallsV2DetailResponse() (response *DescribeTrFirewallsV2DetailResponse) {
	response = &DescribeTrFirewallsV2DetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
