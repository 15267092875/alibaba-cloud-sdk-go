package domain

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

// SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId invokes the domain.SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId API synchronously
func (client *Client) SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId(request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) (response *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, err error) {
	response = CreateSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithChan invokes the domain.SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId API asynchronously
func (client *Client) SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithChan(request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) (<-chan *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, <-chan error) {
	responseChan := make(chan *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId(request)
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

// SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithCallback invokes the domain.SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId API asynchronously
func (client *Client) SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdWithCallback(request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest, callback func(response *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse
		var err error
		defer close(result)
		response, err = client.SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId(request)
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

// SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest is the request struct for api SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId
type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest struct {
	*requests.RpcRequest
	ContactType           string           `position:"Query" name:"ContactType"`
	RegistrantProfileId   requests.Integer `position:"Query" name:"RegistrantProfileId"`
	DomainName            *[]string        `position:"Query" name:"DomainName"  type:"Repeated"`
	TransferOutProhibited requests.Boolean `position:"Query" name:"TransferOutProhibited"`
	UserClientIp          string           `position:"Query" name:"UserClientIp"`
	Lang                  string           `position:"Query" name:"Lang"`
}

// SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse is the response struct for api SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId
type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest creates a request to invoke SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId API
func CreateSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest() (request *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) {
	request = &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse creates a response to parse from SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId response
func CreateSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse() (response *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse) {
	response = &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
