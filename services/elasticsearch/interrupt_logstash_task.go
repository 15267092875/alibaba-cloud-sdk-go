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

// InterruptLogstashTask invokes the elasticsearch.InterruptLogstashTask API synchronously
func (client *Client) InterruptLogstashTask(request *InterruptLogstashTaskRequest) (response *InterruptLogstashTaskResponse, err error) {
	response = CreateInterruptLogstashTaskResponse()
	err = client.DoAction(request, response)
	return
}

// InterruptLogstashTaskWithChan invokes the elasticsearch.InterruptLogstashTask API asynchronously
func (client *Client) InterruptLogstashTaskWithChan(request *InterruptLogstashTaskRequest) (<-chan *InterruptLogstashTaskResponse, <-chan error) {
	responseChan := make(chan *InterruptLogstashTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InterruptLogstashTask(request)
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

// InterruptLogstashTaskWithCallback invokes the elasticsearch.InterruptLogstashTask API asynchronously
func (client *Client) InterruptLogstashTaskWithCallback(request *InterruptLogstashTaskRequest, callback func(response *InterruptLogstashTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InterruptLogstashTaskResponse
		var err error
		defer close(result)
		response, err = client.InterruptLogstashTask(request)
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

// InterruptLogstashTaskRequest is the request struct for api InterruptLogstashTask
type InterruptLogstashTaskRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
}

// InterruptLogstashTaskResponse is the response struct for api InterruptLogstashTask
type InterruptLogstashTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateInterruptLogstashTaskRequest creates a request to invoke InterruptLogstashTask API
func CreateInterruptLogstashTaskRequest() (request *InterruptLogstashTaskRequest) {
	request = &InterruptLogstashTaskRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "InterruptLogstashTask", "/openapi/logstashes/[InstanceId]/actions/interrupt", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInterruptLogstashTaskResponse creates a response to parse from InterruptLogstashTask response
func CreateInterruptLogstashTaskResponse() (response *InterruptLogstashTaskResponse) {
	response = &InterruptLogstashTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
