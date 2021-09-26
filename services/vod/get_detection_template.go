package vod

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

// GetDetectionTemplate invokes the vod.GetDetectionTemplate API synchronously
func (client *Client) GetDetectionTemplate(request *GetDetectionTemplateRequest) (response *GetDetectionTemplateResponse, err error) {
	response = CreateGetDetectionTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetDetectionTemplateWithChan invokes the vod.GetDetectionTemplate API asynchronously
func (client *Client) GetDetectionTemplateWithChan(request *GetDetectionTemplateRequest) (<-chan *GetDetectionTemplateResponse, <-chan error) {
	responseChan := make(chan *GetDetectionTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDetectionTemplate(request)
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

// GetDetectionTemplateWithCallback invokes the vod.GetDetectionTemplate API asynchronously
func (client *Client) GetDetectionTemplateWithCallback(request *GetDetectionTemplateRequest, callback func(response *GetDetectionTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDetectionTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetDetectionTemplate(request)
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

// GetDetectionTemplateRequest is the request struct for api GetDetectionTemplate
type GetDetectionTemplateRequest struct {
	*requests.RpcRequest
	TemplateId string `position:"Query" name:"TemplateId"`
}

// GetDetectionTemplateResponse is the response struct for api GetDetectionTemplate
type GetDetectionTemplateResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	DetectionTemplate DetectionTemplate `json:"DetectionTemplate" xml:"DetectionTemplate"`
}

// CreateGetDetectionTemplateRequest creates a request to invoke GetDetectionTemplate API
func CreateGetDetectionTemplateRequest() (request *GetDetectionTemplateRequest) {
	request = &GetDetectionTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetDetectionTemplate", "vod", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetDetectionTemplateResponse creates a response to parse from GetDetectionTemplate response
func CreateGetDetectionTemplateResponse() (response *GetDetectionTemplateResponse) {
	response = &GetDetectionTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}