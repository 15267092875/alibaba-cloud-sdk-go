package resourcecenter

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

// ListSavedQueries invokes the resourcecenter.ListSavedQueries API synchronously
func (client *Client) ListSavedQueries(request *ListSavedQueriesRequest) (response *ListSavedQueriesResponse, err error) {
	response = CreateListSavedQueriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListSavedQueriesWithChan invokes the resourcecenter.ListSavedQueries API asynchronously
func (client *Client) ListSavedQueriesWithChan(request *ListSavedQueriesRequest) (<-chan *ListSavedQueriesResponse, <-chan error) {
	responseChan := make(chan *ListSavedQueriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSavedQueries(request)
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

// ListSavedQueriesWithCallback invokes the resourcecenter.ListSavedQueries API asynchronously
func (client *Client) ListSavedQueriesWithCallback(request *ListSavedQueriesRequest, callback func(response *ListSavedQueriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSavedQueriesResponse
		var err error
		defer close(result)
		response, err = client.ListSavedQueries(request)
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

// ListSavedQueriesRequest is the request struct for api ListSavedQueries
type ListSavedQueriesRequest struct {
	*requests.RpcRequest
	NextToken  string `position:"Query" name:"NextToken"`
	MaxResults string `position:"Query" name:"MaxResults"`
}

// ListSavedQueriesResponse is the response struct for api ListSavedQueries
type ListSavedQueriesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	MaxResults   string       `json:"MaxResults" xml:"MaxResults"`
	NextToken    string       `json:"NextToken" xml:"NextToken"`
	SavedQueries []SavedQuery `json:"SavedQueries" xml:"SavedQueries"`
}

// CreateListSavedQueriesRequest creates a request to invoke ListSavedQueries API
func CreateListSavedQueriesRequest() (request *ListSavedQueriesRequest) {
	request = &ListSavedQueriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "ListSavedQueries", "", "")
	request.Method = requests.POST
	return
}

// CreateListSavedQueriesResponse creates a response to parse from ListSavedQueries response
func CreateListSavedQueriesResponse() (response *ListSavedQueriesResponse) {
	response = &ListSavedQueriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
