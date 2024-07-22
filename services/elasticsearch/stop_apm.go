package elasticsearch

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

// StopApm invokes the elasticsearch.StopApm API synchronously
func (client *Client) StopApm(request *StopApmRequest) (response *StopApmResponse, err error) {
	response = CreateStopApmResponse()
	err = client.DoAction(request, response)
	return
}

// StopApmWithChan invokes the elasticsearch.StopApm API asynchronously
func (client *Client) StopApmWithChan(request *StopApmRequest) (<-chan *StopApmResponse, <-chan error) {
	responseChan := make(chan *StopApmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopApm(request)
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

// StopApmWithCallback invokes the elasticsearch.StopApm API asynchronously
func (client *Client) StopApmWithCallback(request *StopApmRequest, callback func(response *StopApmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopApmResponse
		var err error
		defer close(result)
		response, err = client.StopApm(request)
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

// StopApmRequest is the request struct for api StopApm
type StopApmRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
}

// StopApmResponse is the response struct for api StopApm
type StopApmResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateStopApmRequest creates a request to invoke StopApm API
func CreateStopApmRequest() (request *StopApmRequest) {
	request = &StopApmRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "StopApm", "/openapi/apm/[instanceId]/actions/stop", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopApmResponse creates a response to parse from StopApm response
func CreateStopApmResponse() (response *StopApmResponse) {
	response = &StopApmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
