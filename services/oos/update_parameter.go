package oos

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

// UpdateParameter invokes the oos.UpdateParameter API synchronously
// api document: https://help.aliyun.com/api/oos/updateparameter.html
func (client *Client) UpdateParameter(request *UpdateParameterRequest) (response *UpdateParameterResponse, err error) {
	response = CreateUpdateParameterResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateParameterWithChan invokes the oos.UpdateParameter API asynchronously
// api document: https://help.aliyun.com/api/oos/updateparameter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateParameterWithChan(request *UpdateParameterRequest) (<-chan *UpdateParameterResponse, <-chan error) {
	responseChan := make(chan *UpdateParameterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateParameter(request)
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

// UpdateParameterWithCallback invokes the oos.UpdateParameter API asynchronously
// api document: https://help.aliyun.com/api/oos/updateparameter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateParameterWithCallback(request *UpdateParameterRequest, callback func(response *UpdateParameterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateParameterResponse
		var err error
		defer close(result)
		response, err = client.UpdateParameter(request)
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

// UpdateParameterRequest is the request struct for api UpdateParameter
type UpdateParameterRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	Name        string `position:"Query" name:"Name"`
	Value       string `position:"Query" name:"Value"`
}

// UpdateParameterResponse is the response struct for api UpdateParameter
type UpdateParameterResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Parameter Parameter `json:"Parameter" xml:"Parameter"`
}

// CreateUpdateParameterRequest creates a request to invoke UpdateParameter API
func CreateUpdateParameterRequest() (request *UpdateParameterRequest) {
	request = &UpdateParameterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("oos", "2019-06-01", "UpdateParameter", "oos", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateParameterResponse creates a response to parse from UpdateParameter response
func CreateUpdateParameterResponse() (response *UpdateParameterResponse) {
	response = &UpdateParameterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}