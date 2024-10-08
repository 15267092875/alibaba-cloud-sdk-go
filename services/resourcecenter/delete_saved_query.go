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

// DeleteSavedQuery invokes the resourcecenter.DeleteSavedQuery API synchronously
func (client *Client) DeleteSavedQuery(request *DeleteSavedQueryRequest) (response *DeleteSavedQueryResponse, err error) {
	response = CreateDeleteSavedQueryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSavedQueryWithChan invokes the resourcecenter.DeleteSavedQuery API asynchronously
func (client *Client) DeleteSavedQueryWithChan(request *DeleteSavedQueryRequest) (<-chan *DeleteSavedQueryResponse, <-chan error) {
	responseChan := make(chan *DeleteSavedQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSavedQuery(request)
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

// DeleteSavedQueryWithCallback invokes the resourcecenter.DeleteSavedQuery API asynchronously
func (client *Client) DeleteSavedQueryWithCallback(request *DeleteSavedQueryRequest, callback func(response *DeleteSavedQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSavedQueryResponse
		var err error
		defer close(result)
		response, err = client.DeleteSavedQuery(request)
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

// DeleteSavedQueryRequest is the request struct for api DeleteSavedQuery
type DeleteSavedQueryRequest struct {
	*requests.RpcRequest
	QueryId string `position:"Query" name:"QueryId"`
}

// DeleteSavedQueryResponse is the response struct for api DeleteSavedQuery
type DeleteSavedQueryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSavedQueryRequest creates a request to invoke DeleteSavedQuery API
func CreateDeleteSavedQueryRequest() (request *DeleteSavedQueryRequest) {
	request = &DeleteSavedQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "DeleteSavedQuery", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSavedQueryResponse creates a response to parse from DeleteSavedQuery response
func CreateDeleteSavedQueryResponse() (response *DeleteSavedQueryResponse) {
	response = &DeleteSavedQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
