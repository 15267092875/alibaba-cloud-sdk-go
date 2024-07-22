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

// AddSnapshotRepo invokes the elasticsearch.AddSnapshotRepo API synchronously
func (client *Client) AddSnapshotRepo(request *AddSnapshotRepoRequest) (response *AddSnapshotRepoResponse, err error) {
	response = CreateAddSnapshotRepoResponse()
	err = client.DoAction(request, response)
	return
}

// AddSnapshotRepoWithChan invokes the elasticsearch.AddSnapshotRepo API asynchronously
func (client *Client) AddSnapshotRepoWithChan(request *AddSnapshotRepoRequest) (<-chan *AddSnapshotRepoResponse, <-chan error) {
	responseChan := make(chan *AddSnapshotRepoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddSnapshotRepo(request)
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

// AddSnapshotRepoWithCallback invokes the elasticsearch.AddSnapshotRepo API asynchronously
func (client *Client) AddSnapshotRepoWithCallback(request *AddSnapshotRepoRequest, callback func(response *AddSnapshotRepoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddSnapshotRepoResponse
		var err error
		defer close(result)
		response, err = client.AddSnapshotRepo(request)
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

// AddSnapshotRepoRequest is the request struct for api AddSnapshotRepo
type AddSnapshotRepoRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"InstanceId"`
	Body       string `position:"Body" name:"body"`
}

// AddSnapshotRepoResponse is the response struct for api AddSnapshotRepo
type AddSnapshotRepoResponse struct {
	*responses.BaseResponse
	Result    bool   `json:"Result" xml:"Result"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddSnapshotRepoRequest creates a request to invoke AddSnapshotRepo API
func CreateAddSnapshotRepoRequest() (request *AddSnapshotRepoRequest) {
	request = &AddSnapshotRepoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "AddSnapshotRepo", "/openapi/instances/[InstanceId]/snapshot-repos", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddSnapshotRepoResponse creates a response to parse from AddSnapshotRepo response
func CreateAddSnapshotRepoResponse() (response *AddSnapshotRepoResponse) {
	response = &AddSnapshotRepoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
