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

// ListAccountFactoryBaselineItems invokes the governance.ListAccountFactoryBaselineItems API synchronously
func (client *Client) ListAccountFactoryBaselineItems(request *ListAccountFactoryBaselineItemsRequest) (response *ListAccountFactoryBaselineItemsResponse, err error) {
	response = CreateListAccountFactoryBaselineItemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAccountFactoryBaselineItemsWithChan invokes the governance.ListAccountFactoryBaselineItems API asynchronously
func (client *Client) ListAccountFactoryBaselineItemsWithChan(request *ListAccountFactoryBaselineItemsRequest) (<-chan *ListAccountFactoryBaselineItemsResponse, <-chan error) {
	responseChan := make(chan *ListAccountFactoryBaselineItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAccountFactoryBaselineItems(request)
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

// ListAccountFactoryBaselineItemsWithCallback invokes the governance.ListAccountFactoryBaselineItems API asynchronously
func (client *Client) ListAccountFactoryBaselineItemsWithCallback(request *ListAccountFactoryBaselineItemsRequest, callback func(response *ListAccountFactoryBaselineItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAccountFactoryBaselineItemsResponse
		var err error
		defer close(result)
		response, err = client.ListAccountFactoryBaselineItems(request)
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

// ListAccountFactoryBaselineItemsRequest is the request struct for api ListAccountFactoryBaselineItems
type ListAccountFactoryBaselineItemsRequest struct {
	*requests.RpcRequest
	Type       string           `position:"Query" name:"Type"`
	NextToken  string           `position:"Query" name:"NextToken"`
	Names      *[]string        `position:"Query" name:"Names"  type:"Repeated"`
	Versions   *[]string        `position:"Query" name:"Versions"  type:"Repeated"`
	MaxResults requests.Integer `position:"Query" name:"MaxResults"`
}

// ListAccountFactoryBaselineItemsResponse is the response struct for api ListAccountFactoryBaselineItems
type ListAccountFactoryBaselineItemsResponse struct {
	*responses.BaseResponse
	NextToken     string         `json:"NextToken" xml:"NextToken"`
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	BaselineItems []BaselineItem `json:"BaselineItems" xml:"BaselineItems"`
}

// CreateListAccountFactoryBaselineItemsRequest creates a request to invoke ListAccountFactoryBaselineItems API
func CreateListAccountFactoryBaselineItemsRequest() (request *ListAccountFactoryBaselineItemsRequest) {
	request = &ListAccountFactoryBaselineItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("governance", "2021-01-20", "ListAccountFactoryBaselineItems", "governance", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListAccountFactoryBaselineItemsResponse creates a response to parse from ListAccountFactoryBaselineItems response
func CreateListAccountFactoryBaselineItemsResponse() (response *ListAccountFactoryBaselineItemsResponse) {
	response = &ListAccountFactoryBaselineItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
