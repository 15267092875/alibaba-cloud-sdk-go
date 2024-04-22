package sae

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

// ListApplications invokes the sae.ListApplications API synchronously
func (client *Client) ListApplications(request *ListApplicationsRequest) (response *ListApplicationsResponse, err error) {
	response = CreateListApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListApplicationsWithChan invokes the sae.ListApplications API asynchronously
func (client *Client) ListApplicationsWithChan(request *ListApplicationsRequest) (<-chan *ListApplicationsResponse, <-chan error) {
	responseChan := make(chan *ListApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApplications(request)
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

// ListApplicationsWithCallback invokes the sae.ListApplications API asynchronously
func (client *Client) ListApplicationsWithCallback(request *ListApplicationsRequest, callback func(response *ListApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApplicationsResponse
		var err error
		defer close(result)
		response, err = client.ListApplications(request)
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

// ListApplicationsRequest is the request struct for api ListApplications
type ListApplicationsRequest struct {
	*requests.RoaRequest
	AppSource   string           `position:"Query" name:"AppSource"`
	AppName     string           `position:"Query" name:"AppName"`
	NamespaceId string           `position:"Query" name:"NamespaceId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	OrderBy     string           `position:"Query" name:"OrderBy"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	FieldValue  string           `position:"Query" name:"FieldValue"`
	Reverse     requests.Boolean `position:"Query" name:"Reverse"`
	FieldType   string           `position:"Query" name:"FieldType"`
	Tags        string           `position:"Query" name:"Tags"`
}

// ListApplicationsResponse is the response struct for api ListApplications
type ListApplicationsResponse struct {
	*responses.BaseResponse
	Message     string `json:"Message" xml:"Message"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	ErrorCode   string `json:"ErrorCode" xml:"ErrorCode"`
	Code        string `json:"Code" xml:"Code"`
	Success     bool   `json:"Success" xml:"Success"`
	CurrentPage int    `json:"CurrentPage" xml:"CurrentPage"`
	TotalSize   int    `json:"TotalSize" xml:"TotalSize"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	Data        Data   `json:"Data" xml:"Data"`
}

// CreateListApplicationsRequest creates a request to invoke ListApplications API
func CreateListApplicationsRequest() (request *ListApplicationsRequest) {
	request = &ListApplicationsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("sae", "2019-05-06", "ListApplications", "/pop/v1/sam/app/listApplications", "serverless", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListApplicationsResponse creates a response to parse from ListApplications response
func CreateListApplicationsResponse() (response *ListApplicationsResponse) {
	response = &ListApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
