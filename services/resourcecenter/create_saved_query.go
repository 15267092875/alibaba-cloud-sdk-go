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

// CreateSavedQuery invokes the resourcecenter.CreateSavedQuery API synchronously
func (client *Client) CreateSavedQuery(request *CreateSavedQueryRequest) (response *CreateSavedQueryResponse, err error) {
	response = CreateCreateSavedQueryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSavedQueryWithChan invokes the resourcecenter.CreateSavedQuery API asynchronously
func (client *Client) CreateSavedQueryWithChan(request *CreateSavedQueryRequest) (<-chan *CreateSavedQueryResponse, <-chan error) {
	responseChan := make(chan *CreateSavedQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSavedQuery(request)
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

// CreateSavedQueryWithCallback invokes the resourcecenter.CreateSavedQuery API asynchronously
func (client *Client) CreateSavedQueryWithCallback(request *CreateSavedQueryRequest, callback func(response *CreateSavedQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSavedQueryResponse
		var err error
		defer close(result)
		response, err = client.CreateSavedQuery(request)
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

// CreateSavedQueryRequest is the request struct for api CreateSavedQuery
type CreateSavedQueryRequest struct {
	*requests.RpcRequest
	Expression  string `position:"Query" name:"Expression"`
	Description string `position:"Query" name:"Description"`
	Name        string `position:"Query" name:"Name"`
}

// CreateSavedQueryResponse is the response struct for api CreateSavedQuery
type CreateSavedQueryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	QueryId   string `json:"QueryId" xml:"QueryId"`
}

// CreateCreateSavedQueryRequest creates a request to invoke CreateSavedQuery API
func CreateCreateSavedQueryRequest() (request *CreateSavedQueryRequest) {
	request = &CreateSavedQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "CreateSavedQuery", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSavedQueryResponse creates a response to parse from CreateSavedQuery response
func CreateCreateSavedQueryResponse() (response *CreateSavedQueryResponse) {
	response = &CreateSavedQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
