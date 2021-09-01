package mts

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

// QueryIProductionJob invokes the mts.QueryIProductionJob API synchronously
func (client *Client) QueryIProductionJob(request *QueryIProductionJobRequest) (response *QueryIProductionJobResponse, err error) {
	response = CreateQueryIProductionJobResponse()
	err = client.DoAction(request, response)
	return
}

// QueryIProductionJobWithChan invokes the mts.QueryIProductionJob API asynchronously
func (client *Client) QueryIProductionJobWithChan(request *QueryIProductionJobRequest) (<-chan *QueryIProductionJobResponse, <-chan error) {
	responseChan := make(chan *QueryIProductionJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryIProductionJob(request)
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

// QueryIProductionJobWithCallback invokes the mts.QueryIProductionJob API asynchronously
func (client *Client) QueryIProductionJobWithCallback(request *QueryIProductionJobRequest, callback func(response *QueryIProductionJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryIProductionJobResponse
		var err error
		defer close(result)
		response, err = client.QueryIProductionJob(request)
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

// QueryIProductionJobRequest is the request struct for api QueryIProductionJob
type QueryIProductionJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
}

// QueryIProductionJobResponse is the response struct for api QueryIProductionJob
type QueryIProductionJobResponse struct {
	*responses.BaseResponse
	FunctionName string `json:"FunctionName" xml:"FunctionName"`
	Input        string `json:"Input" xml:"Input"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	UserData     string `json:"UserData" xml:"UserData"`
	State        string `json:"State" xml:"State"`
	Output       string `json:"Output" xml:"Output"`
	PipelineId   string `json:"PipelineId" xml:"PipelineId"`
	JobParams    string `json:"JobParams" xml:"JobParams"`
	JobId        string `json:"JobId" xml:"JobId"`
	Result       string `json:"Result" xml:"Result"`
}

// CreateQueryIProductionJobRequest creates a request to invoke QueryIProductionJob API
func CreateQueryIProductionJobRequest() (request *QueryIProductionJobRequest) {
	request = &QueryIProductionJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryIProductionJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryIProductionJobResponse creates a response to parse from QueryIProductionJob response
func CreateQueryIProductionJobResponse() (response *QueryIProductionJobResponse) {
	response = &QueryIProductionJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
