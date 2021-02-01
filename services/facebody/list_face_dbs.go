package facebody

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

// ListFaceDbs invokes the facebody.ListFaceDbs API synchronously
func (client *Client) ListFaceDbs(request *ListFaceDbsRequest) (response *ListFaceDbsResponse, err error) {
	response = CreateListFaceDbsResponse()
	err = client.DoAction(request, response)
	return
}

// ListFaceDbsWithChan invokes the facebody.ListFaceDbs API asynchronously
func (client *Client) ListFaceDbsWithChan(request *ListFaceDbsRequest) (<-chan *ListFaceDbsResponse, <-chan error) {
	responseChan := make(chan *ListFaceDbsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFaceDbs(request)
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

// ListFaceDbsWithCallback invokes the facebody.ListFaceDbs API asynchronously
func (client *Client) ListFaceDbsWithCallback(request *ListFaceDbsRequest, callback func(response *ListFaceDbsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFaceDbsResponse
		var err error
		defer close(result)
		response, err = client.ListFaceDbs(request)
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

// ListFaceDbsRequest is the request struct for api ListFaceDbs
type ListFaceDbsRequest struct {
	*requests.RpcRequest
}

// ListFaceDbsResponse is the response struct for api ListFaceDbs
type ListFaceDbsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateListFaceDbsRequest creates a request to invoke ListFaceDbs API
func CreateListFaceDbsRequest() (request *ListFaceDbsRequest) {
	request = &ListFaceDbsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "ListFaceDbs", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFaceDbsResponse creates a response to parse from ListFaceDbs response
func CreateListFaceDbsResponse() (response *ListFaceDbsResponse) {
	response = &ListFaceDbsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
