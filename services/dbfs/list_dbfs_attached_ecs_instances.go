package dbfs

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

// ListDbfsAttachedEcsInstances invokes the dbfs.ListDbfsAttachedEcsInstances API synchronously
func (client *Client) ListDbfsAttachedEcsInstances(request *ListDbfsAttachedEcsInstancesRequest) (response *ListDbfsAttachedEcsInstancesResponse, err error) {
	response = CreateListDbfsAttachedEcsInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDbfsAttachedEcsInstancesWithChan invokes the dbfs.ListDbfsAttachedEcsInstances API asynchronously
func (client *Client) ListDbfsAttachedEcsInstancesWithChan(request *ListDbfsAttachedEcsInstancesRequest) (<-chan *ListDbfsAttachedEcsInstancesResponse, <-chan error) {
	responseChan := make(chan *ListDbfsAttachedEcsInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDbfsAttachedEcsInstances(request)
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

// ListDbfsAttachedEcsInstancesWithCallback invokes the dbfs.ListDbfsAttachedEcsInstances API asynchronously
func (client *Client) ListDbfsAttachedEcsInstancesWithCallback(request *ListDbfsAttachedEcsInstancesRequest, callback func(response *ListDbfsAttachedEcsInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDbfsAttachedEcsInstancesResponse
		var err error
		defer close(result)
		response, err = client.ListDbfsAttachedEcsInstances(request)
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

// ListDbfsAttachedEcsInstancesRequest is the request struct for api ListDbfsAttachedEcsInstances
type ListDbfsAttachedEcsInstancesRequest struct {
	*requests.RpcRequest
	FsId string `position:"Query" name:"FsId"`
}

// ListDbfsAttachedEcsInstancesResponse is the response struct for api ListDbfsAttachedEcsInstances
type ListDbfsAttachedEcsInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string      `json:"RequestId" xml:"RequestId"`
	EcsLabelInfo []LabelInfo `json:"EcsLabelInfo" xml:"EcsLabelInfo"`
}

// CreateListDbfsAttachedEcsInstancesRequest creates a request to invoke ListDbfsAttachedEcsInstances API
func CreateListDbfsAttachedEcsInstancesRequest() (request *ListDbfsAttachedEcsInstancesRequest) {
	request = &ListDbfsAttachedEcsInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "ListDbfsAttachedEcsInstances", "", "")
	request.Method = requests.POST
	return
}

// CreateListDbfsAttachedEcsInstancesResponse creates a response to parse from ListDbfsAttachedEcsInstances response
func CreateListDbfsAttachedEcsInstancesResponse() (response *ListDbfsAttachedEcsInstancesResponse) {
	response = &ListDbfsAttachedEcsInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
