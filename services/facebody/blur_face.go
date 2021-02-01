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

// BlurFace invokes the facebody.BlurFace API synchronously
func (client *Client) BlurFace(request *BlurFaceRequest) (response *BlurFaceResponse, err error) {
	response = CreateBlurFaceResponse()
	err = client.DoAction(request, response)
	return
}

// BlurFaceWithChan invokes the facebody.BlurFace API asynchronously
func (client *Client) BlurFaceWithChan(request *BlurFaceRequest) (<-chan *BlurFaceResponse, <-chan error) {
	responseChan := make(chan *BlurFaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BlurFace(request)
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

// BlurFaceWithCallback invokes the facebody.BlurFace API asynchronously
func (client *Client) BlurFaceWithCallback(request *BlurFaceRequest, callback func(response *BlurFaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BlurFaceResponse
		var err error
		defer close(result)
		response, err = client.BlurFace(request)
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

// BlurFaceRequest is the request struct for api BlurFace
type BlurFaceRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Body" name:"ImageURL"`
}

// BlurFaceResponse is the response struct for api BlurFace
type BlurFaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateBlurFaceRequest creates a request to invoke BlurFace API
func CreateBlurFaceRequest() (request *BlurFaceRequest) {
	request = &BlurFaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("facebody", "2019-12-30", "BlurFace", "facebody", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBlurFaceResponse creates a response to parse from BlurFace response
func CreateBlurFaceResponse() (response *BlurFaceResponse) {
	response = &BlurFaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
