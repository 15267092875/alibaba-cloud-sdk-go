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

// ListEcsInstances invokes the elasticsearch.ListEcsInstances API synchronously
func (client *Client) ListEcsInstances(request *ListEcsInstancesRequest) (response *ListEcsInstancesResponse, err error) {
	response = CreateListEcsInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListEcsInstancesWithChan invokes the elasticsearch.ListEcsInstances API asynchronously
func (client *Client) ListEcsInstancesWithChan(request *ListEcsInstancesRequest) (<-chan *ListEcsInstancesResponse, <-chan error) {
	responseChan := make(chan *ListEcsInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEcsInstances(request)
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

// ListEcsInstancesWithCallback invokes the elasticsearch.ListEcsInstances API asynchronously
func (client *Client) ListEcsInstancesWithCallback(request *ListEcsInstancesRequest, callback func(response *ListEcsInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEcsInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListEcsInstances(request)
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

// ListEcsInstancesRequest is the request struct for api ListEcsInstances
type ListEcsInstancesRequest struct {
	*requests.RoaRequest
	EcsInstanceName string           `position:"Query" name:"ecsInstanceName"`
	EcsInstanceIds  string           `position:"Query" name:"ecsInstanceIds"`
	Size            requests.Integer `position:"Query" name:"size"`
	VpcId           string           `position:"Query" name:"vpcId"`
	Page            requests.Integer `position:"Query" name:"page"`
	Tags            string           `position:"Query" name:"tags"`
}

// ListEcsInstancesResponse is the response struct for api ListEcsInstances
type ListEcsInstancesResponse struct {
	*responses.BaseResponse
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Headers   HeadersInListEcsInstances      `json:"Headers" xml:"Headers"`
	Result    []ResultItemInListEcsInstances `json:"Result" xml:"Result"`
}

// CreateListEcsInstancesRequest creates a request to invoke ListEcsInstances API
func CreateListEcsInstancesRequest() (request *ListEcsInstancesRequest) {
	request = &ListEcsInstancesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListEcsInstances", "/openapi/ecs", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListEcsInstancesResponse creates a response to parse from ListEcsInstances response
func CreateListEcsInstancesResponse() (response *ListEcsInstancesResponse) {
	response = &ListEcsInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
