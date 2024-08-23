package governance

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

// ListEvaluationResults invokes the governance.ListEvaluationResults API synchronously
func (client *Client) ListEvaluationResults(request *ListEvaluationResultsRequest) (response *ListEvaluationResultsResponse, err error) {
	response = CreateListEvaluationResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEvaluationResultsWithChan invokes the governance.ListEvaluationResults API asynchronously
func (client *Client) ListEvaluationResultsWithChan(request *ListEvaluationResultsRequest) (<-chan *ListEvaluationResultsResponse, <-chan error) {
	responseChan := make(chan *ListEvaluationResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEvaluationResults(request)
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

// ListEvaluationResultsWithCallback invokes the governance.ListEvaluationResults API asynchronously
func (client *Client) ListEvaluationResultsWithCallback(request *ListEvaluationResultsRequest, callback func(response *ListEvaluationResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEvaluationResultsResponse
		var err error
		defer close(result)
		response, err = client.ListEvaluationResults(request)
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

// ListEvaluationResultsRequest is the request struct for api ListEvaluationResults
type ListEvaluationResultsRequest struct {
	*requests.RpcRequest
	AccountId   requests.Integer `position:"Query" name:"AccountId"`
	PartnerCode string           `position:"Query" name:"PartnerCode"`
}

// ListEvaluationResultsResponse is the response struct for api ListEvaluationResults
type ListEvaluationResultsResponse struct {
	*responses.BaseResponse
	AccountId int64   `json:"AccountId" xml:"AccountId"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Results   Results `json:"Results" xml:"Results"`
}

// CreateListEvaluationResultsRequest creates a request to invoke ListEvaluationResults API
func CreateListEvaluationResultsRequest() (request *ListEvaluationResultsRequest) {
	request = &ListEvaluationResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("governance", "2021-01-20", "ListEvaluationResults", "governance", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEvaluationResultsResponse creates a response to parse from ListEvaluationResults response
func CreateListEvaluationResultsResponse() (response *ListEvaluationResultsResponse) {
	response = &ListEvaluationResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
