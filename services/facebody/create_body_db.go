package facebody

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

// CreateBodyDb invokes the facebody.CreateBodyDb API synchronously
func (client *Client) CreateBodyDb(request *CreateBodyDbRequest) (response *CreateBodyDbResponse, err error) {
	response = CreateCreateBodyDbResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBodyDbWithChan invokes the facebody.CreateBodyDb API asynchronously
func (client *Client) CreateBodyDbWithChan(request *CreateBodyDbRequest) (<-chan *CreateBodyDbResponse, <-chan error) {
	responseChan := make(chan *CreateBodyDbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBodyDb(request)
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

// CreateBodyDbWithCallback invokes the facebody.CreateBodyDb API asynchronously
func (client *Client) CreateBodyDbWithCallback(request *CreateBodyDbRequest, callback func(response *CreateBodyDbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBodyDbResponse
		var err error
		defer close(result)
		response, err = client.CreateBodyDb(request)
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

// CreateBodyDbRequest is the request struct for api CreateBodyDb
type CreateBodyDbRequest struct {
	*requests.RpcRequest
	Name string `position:"Body" name:"Name"`
}

// CreateBodyDbResponse is the response struct for api CreateBodyDb
type CreateBodyDbResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateBodyDbRequest creates a request to invoke CreateBodyDb API
func CreateCreateBodyDbRequest() (request *CreateBodyDbRequest) {
	request = &CreateBodyDbRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "CreateBodyDb", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBodyDbResponse creates a response to parse from CreateBodyDb response
func CreateCreateBodyDbResponse() (response *CreateBodyDbResponse) {
	response = &CreateBodyDbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
