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

// ListDiagnosisItems invokes the elasticsearch.ListDiagnosisItems API synchronously
func (client *Client) ListDiagnosisItems(request *ListDiagnosisItemsRequest) (response *ListDiagnosisItemsResponse, err error) {
	response = CreateListDiagnosisItemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDiagnosisItemsWithChan invokes the elasticsearch.ListDiagnosisItems API asynchronously
func (client *Client) ListDiagnosisItemsWithChan(request *ListDiagnosisItemsRequest) (<-chan *ListDiagnosisItemsResponse, <-chan error) {
	responseChan := make(chan *ListDiagnosisItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDiagnosisItems(request)
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

// ListDiagnosisItemsWithCallback invokes the elasticsearch.ListDiagnosisItems API asynchronously
func (client *Client) ListDiagnosisItemsWithCallback(request *ListDiagnosisItemsRequest, callback func(response *ListDiagnosisItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDiagnosisItemsResponse
		var err error
		defer close(result)
		response, err = client.ListDiagnosisItems(request)
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

// ListDiagnosisItemsRequest is the request struct for api ListDiagnosisItems
type ListDiagnosisItemsRequest struct {
	*requests.RoaRequest
	Lang string `position:"Query" name:"lang"`
}

// ListDiagnosisItemsResponse is the response struct for api ListDiagnosisItems
type ListDiagnosisItemsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListDiagnosisItemsRequest creates a request to invoke ListDiagnosisItems API
func CreateListDiagnosisItemsRequest() (request *ListDiagnosisItemsRequest) {
	request = &ListDiagnosisItemsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListDiagnosisItems", "/openapi/diagnosis/items", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDiagnosisItemsResponse creates a response to parse from ListDiagnosisItems response
func CreateListDiagnosisItemsResponse() (response *ListDiagnosisItemsResponse) {
	response = &ListDiagnosisItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
