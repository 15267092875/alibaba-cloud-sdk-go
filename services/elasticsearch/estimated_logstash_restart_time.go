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

// EstimatedLogstashRestartTime invokes the elasticsearch.EstimatedLogstashRestartTime API synchronously
func (client *Client) EstimatedLogstashRestartTime(request *EstimatedLogstashRestartTimeRequest) (response *EstimatedLogstashRestartTimeResponse, err error) {
	response = CreateEstimatedLogstashRestartTimeResponse()
	err = client.DoAction(request, response)
	return
}

// EstimatedLogstashRestartTimeWithChan invokes the elasticsearch.EstimatedLogstashRestartTime API asynchronously
func (client *Client) EstimatedLogstashRestartTimeWithChan(request *EstimatedLogstashRestartTimeRequest) (<-chan *EstimatedLogstashRestartTimeResponse, <-chan error) {
	responseChan := make(chan *EstimatedLogstashRestartTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EstimatedLogstashRestartTime(request)
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

// EstimatedLogstashRestartTimeWithCallback invokes the elasticsearch.EstimatedLogstashRestartTime API asynchronously
func (client *Client) EstimatedLogstashRestartTimeWithCallback(request *EstimatedLogstashRestartTimeRequest, callback func(response *EstimatedLogstashRestartTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EstimatedLogstashRestartTimeResponse
		var err error
		defer close(result)
		response, err = client.EstimatedLogstashRestartTime(request)
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

// EstimatedLogstashRestartTimeRequest is the request struct for api EstimatedLogstashRestartTime
type EstimatedLogstashRestartTimeRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Force      requests.Boolean `position:"Query" name:"force"`
	Body       string           `position:"Body" name:"body"`
}

// EstimatedLogstashRestartTimeResponse is the response struct for api EstimatedLogstashRestartTime
type EstimatedLogstashRestartTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateEstimatedLogstashRestartTimeRequest creates a request to invoke EstimatedLogstashRestartTime API
func CreateEstimatedLogstashRestartTimeRequest() (request *EstimatedLogstashRestartTimeRequest) {
	request = &EstimatedLogstashRestartTimeRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "EstimatedLogstashRestartTime", "/openapi/logstashes/[InstanceId]/estimated-time/restart-time", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateEstimatedLogstashRestartTimeResponse creates a response to parse from EstimatedLogstashRestartTime response
func CreateEstimatedLogstashRestartTimeResponse() (response *EstimatedLogstashRestartTimeResponse) {
	response = &EstimatedLogstashRestartTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
