package dyplsapi

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

// LockSecretNo invokes the dyplsapi.LockSecretNo API synchronously
func (client *Client) LockSecretNo(request *LockSecretNoRequest) (response *LockSecretNoResponse, err error) {
	response = CreateLockSecretNoResponse()
	err = client.DoAction(request, response)
	return
}

// LockSecretNoWithChan invokes the dyplsapi.LockSecretNo API asynchronously
func (client *Client) LockSecretNoWithChan(request *LockSecretNoRequest) (<-chan *LockSecretNoResponse, <-chan error) {
	responseChan := make(chan *LockSecretNoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.LockSecretNo(request)
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

// LockSecretNoWithCallback invokes the dyplsapi.LockSecretNo API asynchronously
func (client *Client) LockSecretNoWithCallback(request *LockSecretNoRequest, callback func(response *LockSecretNoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *LockSecretNoResponse
		var err error
		defer close(result)
		response, err = client.LockSecretNo(request)
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

// LockSecretNoRequest is the request struct for api LockSecretNo
type LockSecretNoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecretNo             string           `position:"Query" name:"SecretNo"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PoolKey              string           `position:"Query" name:"PoolKey"`
}

// LockSecretNoResponse is the response struct for api LockSecretNo
type LockSecretNoResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateLockSecretNoRequest creates a request to invoke LockSecretNo API
func CreateLockSecretNoRequest() (request *LockSecretNoRequest) {
	request = &LockSecretNoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "LockSecretNo", "", "")
	request.Method = requests.POST
	return
}

// CreateLockSecretNoResponse creates a response to parse from LockSecretNo response
func CreateLockSecretNoResponse() (response *LockSecretNoResponse) {
	response = &LockSecretNoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
