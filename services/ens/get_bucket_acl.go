package ens

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

// GetBucketAcl invokes the ens.GetBucketAcl API synchronously
func (client *Client) GetBucketAcl(request *GetBucketAclRequest) (response *GetBucketAclResponse, err error) {
	response = CreateGetBucketAclResponse()
	err = client.DoAction(request, response)
	return
}

// GetBucketAclWithChan invokes the ens.GetBucketAcl API asynchronously
func (client *Client) GetBucketAclWithChan(request *GetBucketAclRequest) (<-chan *GetBucketAclResponse, <-chan error) {
	responseChan := make(chan *GetBucketAclResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetBucketAcl(request)
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

// GetBucketAclWithCallback invokes the ens.GetBucketAcl API asynchronously
func (client *Client) GetBucketAclWithCallback(request *GetBucketAclRequest, callback func(response *GetBucketAclResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetBucketAclResponse
		var err error
		defer close(result)
		response, err = client.GetBucketAcl(request)
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

// GetBucketAclRequest is the request struct for api GetBucketAcl
type GetBucketAclRequest struct {
	*requests.RpcRequest
	BucketName string `position:"Query" name:"BucketName"`
}

// GetBucketAclResponse is the response struct for api GetBucketAcl
type GetBucketAclResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	BucketAcl string `json:"BucketAcl" xml:"BucketAcl"`
}

// CreateGetBucketAclRequest creates a request to invoke GetBucketAcl API
func CreateGetBucketAclRequest() (request *GetBucketAclRequest) {
	request = &GetBucketAclRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "GetBucketAcl", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetBucketAclResponse creates a response to parse from GetBucketAcl response
func CreateGetBucketAclResponse() (response *GetBucketAclResponse) {
	response = &GetBucketAclResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
