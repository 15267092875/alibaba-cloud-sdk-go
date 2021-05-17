package ccc

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

// ListCallEventDetailByContactId invokes the ccc.ListCallEventDetailByContactId API synchronously
func (client *Client) ListCallEventDetailByContactId(request *ListCallEventDetailByContactIdRequest) (response *ListCallEventDetailByContactIdResponse, err error) {
	response = CreateListCallEventDetailByContactIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListCallEventDetailByContactIdWithChan invokes the ccc.ListCallEventDetailByContactId API asynchronously
func (client *Client) ListCallEventDetailByContactIdWithChan(request *ListCallEventDetailByContactIdRequest) (<-chan *ListCallEventDetailByContactIdResponse, <-chan error) {
	responseChan := make(chan *ListCallEventDetailByContactIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCallEventDetailByContactId(request)
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

// ListCallEventDetailByContactIdWithCallback invokes the ccc.ListCallEventDetailByContactId API asynchronously
func (client *Client) ListCallEventDetailByContactIdWithCallback(request *ListCallEventDetailByContactIdRequest, callback func(response *ListCallEventDetailByContactIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCallEventDetailByContactIdResponse
		var err error
		defer close(result)
		response, err = client.ListCallEventDetailByContactId(request)
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

// ListCallEventDetailByContactIdRequest is the request struct for api ListCallEventDetailByContactId
type ListCallEventDetailByContactIdRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	ContactId  string `position:"Query" name:"ContactId"`
}

// ListCallEventDetailByContactIdResponse is the response struct for api ListCallEventDetailByContactId
type ListCallEventDetailByContactIdResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListCallEventDetailByContactIdRequest creates a request to invoke ListCallEventDetailByContactId API
func CreateListCallEventDetailByContactIdRequest() (request *ListCallEventDetailByContactIdRequest) {
	request = &ListCallEventDetailByContactIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ListCallEventDetailByContactId", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCallEventDetailByContactIdResponse creates a response to parse from ListCallEventDetailByContactId response
func CreateListCallEventDetailByContactIdResponse() (response *ListCallEventDetailByContactIdResponse) {
	response = &ListCallEventDetailByContactIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
