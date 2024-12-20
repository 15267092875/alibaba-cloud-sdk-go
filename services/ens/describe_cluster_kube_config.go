package ens

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

// DescribeClusterKubeConfig invokes the ens.DescribeClusterKubeConfig API synchronously
func (client *Client) DescribeClusterKubeConfig(request *DescribeClusterKubeConfigRequest) (response *DescribeClusterKubeConfigResponse, err error) {
	response = CreateDescribeClusterKubeConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeClusterKubeConfigWithChan invokes the ens.DescribeClusterKubeConfig API asynchronously
func (client *Client) DescribeClusterKubeConfigWithChan(request *DescribeClusterKubeConfigRequest) (<-chan *DescribeClusterKubeConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeClusterKubeConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeClusterKubeConfig(request)
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

// DescribeClusterKubeConfigWithCallback invokes the ens.DescribeClusterKubeConfig API asynchronously
func (client *Client) DescribeClusterKubeConfigWithCallback(request *DescribeClusterKubeConfigRequest, callback func(response *DescribeClusterKubeConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeClusterKubeConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeClusterKubeConfig(request)
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

// DescribeClusterKubeConfigRequest is the request struct for api DescribeClusterKubeConfig
type DescribeClusterKubeConfigRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// DescribeClusterKubeConfigResponse is the response struct for api DescribeClusterKubeConfig
type DescribeClusterKubeConfigResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ClusterId  string `json:"ClusterId" xml:"ClusterId"`
	Kubeconfig string `json:"Kubeconfig" xml:"Kubeconfig"`
}

// CreateDescribeClusterKubeConfigRequest creates a request to invoke DescribeClusterKubeConfig API
func CreateDescribeClusterKubeConfigRequest() (request *DescribeClusterKubeConfigRequest) {
	request = &DescribeClusterKubeConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeClusterKubeConfig", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeClusterKubeConfigResponse creates a response to parse from DescribeClusterKubeConfig response
func CreateDescribeClusterKubeConfigResponse() (response *DescribeClusterKubeConfigResponse) {
	response = &DescribeClusterKubeConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
