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

// GetJobList invokes the ccc.GetJobList API synchronously
func (client *Client) GetJobList(request *GetJobListRequest) (response *GetJobListResponse, err error) {
	response = CreateGetJobListResponse()
	err = client.DoAction(request, response)
	return
}

// GetJobListWithChan invokes the ccc.GetJobList API asynchronously
func (client *Client) GetJobListWithChan(request *GetJobListRequest) (<-chan *GetJobListResponse, <-chan error) {
	responseChan := make(chan *GetJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetJobList(request)
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

// GetJobListWithCallback invokes the ccc.GetJobList API asynchronously
func (client *Client) GetJobListWithCallback(request *GetJobListRequest, callback func(response *GetJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetJobListResponse
		var err error
		defer close(result)
		response, err = client.GetJobList(request)
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

// GetJobListRequest is the request struct for api GetJobList
type GetJobListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	QueryAll   requests.Boolean `position:"Query" name:"QueryAll"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	JobGroupId string           `position:"Query" name:"JobGroupId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Status     requests.Integer `position:"Query" name:"Status"`
}

// GetJobListResponse is the response struct for api GetJobList
type GetJobListResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Jobs           Jobs   `json:"Jobs" xml:"Jobs"`
}

// CreateGetJobListRequest creates a request to invoke GetJobList API
func CreateGetJobListRequest() (request *GetJobListRequest) {
	request = &GetJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "GetJobList", "", "")
	request.Method = requests.POST
	return
}

// CreateGetJobListResponse creates a response to parse from GetJobList response
func CreateGetJobListResponse() (response *GetJobListResponse) {
	response = &GetJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}