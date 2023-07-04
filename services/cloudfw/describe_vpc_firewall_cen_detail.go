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

// DescribeVpcFirewallCenDetail invokes the cloudfw.DescribeVpcFirewallCenDetail API synchronously
func (client *Client) DescribeVpcFirewallCenDetail(request *DescribeVpcFirewallCenDetailRequest) (response *DescribeVpcFirewallCenDetailResponse, err error) {
	response = CreateDescribeVpcFirewallCenDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpcFirewallCenDetailWithChan invokes the cloudfw.DescribeVpcFirewallCenDetail API asynchronously
func (client *Client) DescribeVpcFirewallCenDetailWithChan(request *DescribeVpcFirewallCenDetailRequest) (<-chan *DescribeVpcFirewallCenDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcFirewallCenDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcFirewallCenDetail(request)
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

// DescribeVpcFirewallCenDetailWithCallback invokes the cloudfw.DescribeVpcFirewallCenDetail API asynchronously
func (client *Client) DescribeVpcFirewallCenDetailWithCallback(request *DescribeVpcFirewallCenDetailRequest, callback func(response *DescribeVpcFirewallCenDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcFirewallCenDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcFirewallCenDetail(request)
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

// DescribeVpcFirewallCenDetailRequest is the request struct for api DescribeVpcFirewallCenDetail
type DescribeVpcFirewallCenDetailRequest struct {
	*requests.RpcRequest
	NetworkInstanceId string `position:"Query" name:"NetworkInstanceId"`
	SourceIp          string `position:"Query" name:"SourceIp"`
	Lang              string `position:"Query" name:"Lang"`
	VpcFirewallId     string `position:"Query" name:"VpcFirewallId"`
}

// DescribeVpcFirewallCenDetailResponse is the response struct for api DescribeVpcFirewallCenDetail
type DescribeVpcFirewallCenDetailResponse struct {
	*responses.BaseResponse
	ConnectType          string      `json:"ConnectType" xml:"ConnectType"`
	VpcFirewallName      string      `json:"VpcFirewallName" xml:"VpcFirewallName"`
	VpcFirewallId        string      `json:"VpcFirewallId" xml:"VpcFirewallId"`
	FirewallSwitchStatus string      `json:"FirewallSwitchStatus" xml:"FirewallSwitchStatus"`
	RequestId            string      `json:"RequestId" xml:"RequestId"`
	LocalVpc             LocalVpc    `json:"LocalVpc" xml:"LocalVpc"`
	FirewallVpc          FirewallVpc `json:"FirewallVpc" xml:"FirewallVpc"`
}

// CreateDescribeVpcFirewallCenDetailRequest creates a request to invoke DescribeVpcFirewallCenDetail API
func CreateDescribeVpcFirewallCenDetailRequest() (request *DescribeVpcFirewallCenDetailRequest) {
	request = &DescribeVpcFirewallCenDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudfw", "2017-12-07", "DescribeVpcFirewallCenDetail", "cloudfirewall", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpcFirewallCenDetailResponse creates a response to parse from DescribeVpcFirewallCenDetail response
func CreateDescribeVpcFirewallCenDetailResponse() (response *DescribeVpcFirewallCenDetailResponse) {
	response = &DescribeVpcFirewallCenDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
