package live

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

// ModifyLiveMessageUserInfo invokes the live.ModifyLiveMessageUserInfo API synchronously
func (client *Client) ModifyLiveMessageUserInfo(request *ModifyLiveMessageUserInfoRequest) (response *ModifyLiveMessageUserInfoResponse, err error) {
	response = CreateModifyLiveMessageUserInfoResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLiveMessageUserInfoWithChan invokes the live.ModifyLiveMessageUserInfo API asynchronously
func (client *Client) ModifyLiveMessageUserInfoWithChan(request *ModifyLiveMessageUserInfoRequest) (<-chan *ModifyLiveMessageUserInfoResponse, <-chan error) {
	responseChan := make(chan *ModifyLiveMessageUserInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLiveMessageUserInfo(request)
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

// ModifyLiveMessageUserInfoWithCallback invokes the live.ModifyLiveMessageUserInfo API asynchronously
func (client *Client) ModifyLiveMessageUserInfoWithCallback(request *ModifyLiveMessageUserInfoRequest, callback func(response *ModifyLiveMessageUserInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLiveMessageUserInfoResponse
		var err error
		defer close(result)
		response, err = client.ModifyLiveMessageUserInfo(request)
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

// ModifyLiveMessageUserInfoRequest is the request struct for api ModifyLiveMessageUserInfo
type ModifyLiveMessageUserInfoRequest struct {
	*requests.RpcRequest
	DataCenter   string `position:"Query" name:"DataCenter"`
	UserId       string `position:"Query" name:"UserId"`
	UserMetaInfo string `position:"Query" name:"UserMetaInfo"`
	AppId        string `position:"Query" name:"AppId"`
}

// ModifyLiveMessageUserInfoResponse is the response struct for api ModifyLiveMessageUserInfo
type ModifyLiveMessageUserInfoResponse struct {
	*responses.BaseResponse
	RequestId   string          `json:"RequestId" xml:"RequestId"`
	SuccessList []SuccessGroups `json:"SuccessList" xml:"SuccessList"`
	FailList    []FailGroups    `json:"FailList" xml:"FailList"`
}

// CreateModifyLiveMessageUserInfoRequest creates a request to invoke ModifyLiveMessageUserInfo API
func CreateModifyLiveMessageUserInfoRequest() (request *ModifyLiveMessageUserInfoRequest) {
	request = &ModifyLiveMessageUserInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyLiveMessageUserInfo", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLiveMessageUserInfoResponse creates a response to parse from ModifyLiveMessageUserInfo response
func CreateModifyLiveMessageUserInfoResponse() (response *ModifyLiveMessageUserInfoResponse) {
	response = &ModifyLiveMessageUserInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
