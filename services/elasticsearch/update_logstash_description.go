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

// UpdateLogstashDescription invokes the elasticsearch.UpdateLogstashDescription API synchronously
func (client *Client) UpdateLogstashDescription(request *UpdateLogstashDescriptionRequest) (response *UpdateLogstashDescriptionResponse, err error) {
	response = CreateUpdateLogstashDescriptionResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLogstashDescriptionWithChan invokes the elasticsearch.UpdateLogstashDescription API asynchronously
func (client *Client) UpdateLogstashDescriptionWithChan(request *UpdateLogstashDescriptionRequest) (<-chan *UpdateLogstashDescriptionResponse, <-chan error) {
	responseChan := make(chan *UpdateLogstashDescriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLogstashDescription(request)
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

// UpdateLogstashDescriptionWithCallback invokes the elasticsearch.UpdateLogstashDescription API asynchronously
func (client *Client) UpdateLogstashDescriptionWithCallback(request *UpdateLogstashDescriptionRequest, callback func(response *UpdateLogstashDescriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLogstashDescriptionResponse
		var err error
		defer close(result)
		response, err = client.UpdateLogstashDescription(request)
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

// UpdateLogstashDescriptionRequest is the request struct for api UpdateLogstashDescription
type UpdateLogstashDescriptionRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	Body        string `position:"Body" name:"body"`
}

// UpdateLogstashDescriptionResponse is the response struct for api UpdateLogstashDescription
type UpdateLogstashDescriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateUpdateLogstashDescriptionRequest creates a request to invoke UpdateLogstashDescription API
func CreateUpdateLogstashDescriptionRequest() (request *UpdateLogstashDescriptionRequest) {
	request = &UpdateLogstashDescriptionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateLogstashDescription", "/openapi/logstashes/[InstanceId]/description", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLogstashDescriptionResponse creates a response to parse from UpdateLogstashDescription response
func CreateUpdateLogstashDescriptionResponse() (response *UpdateLogstashDescriptionResponse) {
	response = &UpdateLogstashDescriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
