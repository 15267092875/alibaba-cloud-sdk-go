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

// CreateBatchJobs invokes the ccc.CreateBatchJobs API synchronously
func (client *Client) CreateBatchJobs(request *CreateBatchJobsRequest) (response *CreateBatchJobsResponse, err error) {
	response = CreateCreateBatchJobsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBatchJobsWithChan invokes the ccc.CreateBatchJobs API asynchronously
func (client *Client) CreateBatchJobsWithChan(request *CreateBatchJobsRequest) (<-chan *CreateBatchJobsResponse, <-chan error) {
	responseChan := make(chan *CreateBatchJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBatchJobs(request)
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

// CreateBatchJobsWithCallback invokes the ccc.CreateBatchJobs API asynchronously
func (client *Client) CreateBatchJobsWithCallback(request *CreateBatchJobsRequest, callback func(response *CreateBatchJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBatchJobsResponse
		var err error
		defer close(result)
		response, err = client.CreateBatchJobs(request)
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

// CreateBatchJobsRequest is the request struct for api CreateBatchJobs
type CreateBatchJobsRequest struct {
	*requests.RpcRequest
	Description   string           `position:"Query" name:"Description"`
	JobFilePath   string           `position:"Query" name:"JobFilePath"`
	CallingNumber *[]string        `position:"Query" name:"CallingNumber"  type:"Repeated"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Submitted     requests.Boolean `position:"Query" name:"Submitted"`
	StrategyJson  string           `position:"Query" name:"StrategyJson"`
	Name          string           `position:"Query" name:"Name"`
	ScenarioId    string           `position:"Query" name:"ScenarioId"`
}

// CreateBatchJobsResponse is the response struct for api CreateBatchJobs
type CreateBatchJobsResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Code           string   `json:"Code" xml:"Code"`
	Message        string   `json:"Message" xml:"Message"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	JobGroup       JobGroup `json:"JobGroup" xml:"JobGroup"`
}

// CreateCreateBatchJobsRequest creates a request to invoke CreateBatchJobs API
func CreateCreateBatchJobsRequest() (request *CreateBatchJobsRequest) {
	request = &CreateBatchJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "CreateBatchJobs", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBatchJobsResponse creates a response to parse from CreateBatchJobs response
func CreateCreateBatchJobsResponse() (response *CreateBatchJobsResponse) {
	response = &CreateBatchJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
