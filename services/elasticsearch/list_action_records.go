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

// ListActionRecords invokes the elasticsearch.ListActionRecords API synchronously
func (client *Client) ListActionRecords(request *ListActionRecordsRequest) (response *ListActionRecordsResponse, err error) {
	response = CreateListActionRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// ListActionRecordsWithChan invokes the elasticsearch.ListActionRecords API asynchronously
func (client *Client) ListActionRecordsWithChan(request *ListActionRecordsRequest) (<-chan *ListActionRecordsResponse, <-chan error) {
	responseChan := make(chan *ListActionRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListActionRecords(request)
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

// ListActionRecordsWithCallback invokes the elasticsearch.ListActionRecords API asynchronously
func (client *Client) ListActionRecordsWithCallback(request *ListActionRecordsRequest, callback func(response *ListActionRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListActionRecordsResponse
		var err error
		defer close(result)
		response, err = client.ListActionRecords(request)
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

// ListActionRecordsRequest is the request struct for api ListActionRecords
type ListActionRecordsRequest struct {
	*requests.RoaRequest
	Filter      string           `position:"Query" name:"filter"`
	ActionNames string           `position:"Query" name:"actionNames"`
	InstanceId  string           `position:"Path" name:"InstanceId"`
	Size        requests.Integer `position:"Query" name:"size"`
	RequestId   string           `position:"Query" name:"requestId"`
	EndTime     requests.Integer `position:"Query" name:"endTime"`
	Page        requests.Integer `position:"Query" name:"page"`
	StartTime   requests.Integer `position:"Query" name:"startTime"`
	UserId      string           `position:"Query" name:"userId"`
}

// ListActionRecordsResponse is the response struct for api ListActionRecords
type ListActionRecordsResponse struct {
	*responses.BaseResponse
	RequestId string         `json:"RequestId" xml:"RequestId"`
	Result    []ActionRecord `json:"Result" xml:"Result"`
}

// CreateListActionRecordsRequest creates a request to invoke ListActionRecords API
func CreateListActionRecordsRequest() (request *ListActionRecordsRequest) {
	request = &ListActionRecordsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListActionRecords", "/openapi/instances/[InstanceId]/action-records", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListActionRecordsResponse creates a response to parse from ListActionRecords response
func CreateListActionRecordsResponse() (response *ListActionRecordsResponse) {
	response = &ListActionRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
