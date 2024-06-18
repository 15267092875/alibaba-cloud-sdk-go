package computenest

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

// StopServiceInstance invokes the computenest.StopServiceInstance API synchronously
func (client *Client) StopServiceInstance(request *StopServiceInstanceRequest) (response *StopServiceInstanceResponse, err error) {
	response = CreateStopServiceInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// StopServiceInstanceWithChan invokes the computenest.StopServiceInstance API asynchronously
func (client *Client) StopServiceInstanceWithChan(request *StopServiceInstanceRequest) (<-chan *StopServiceInstanceResponse, <-chan error) {
	responseChan := make(chan *StopServiceInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopServiceInstance(request)
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

// StopServiceInstanceWithCallback invokes the computenest.StopServiceInstance API asynchronously
func (client *Client) StopServiceInstanceWithCallback(request *StopServiceInstanceRequest, callback func(response *StopServiceInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopServiceInstanceResponse
		var err error
		defer close(result)
		response, err = client.StopServiceInstance(request)
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

// StopServiceInstanceRequest is the request struct for api StopServiceInstance
type StopServiceInstanceRequest struct {
	*requests.RpcRequest
	ClientToken       string `position:"Query" name:"ClientToken"`
	ServiceInstanceId string `position:"Query" name:"ServiceInstanceId"`
}

// StopServiceInstanceResponse is the response struct for api StopServiceInstance
type StopServiceInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopServiceInstanceRequest creates a request to invoke StopServiceInstance API
func CreateStopServiceInstanceRequest() (request *StopServiceInstanceRequest) {
	request = &StopServiceInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNest", "2021-06-01", "StopServiceInstance", "computenest", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopServiceInstanceResponse creates a response to parse from StopServiceInstance response
func CreateStopServiceInstanceResponse() (response *StopServiceInstanceResponse) {
	response = &StopServiceInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
