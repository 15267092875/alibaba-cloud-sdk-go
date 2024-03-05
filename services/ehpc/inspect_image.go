package ehpc

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

// InspectImage invokes the ehpc.InspectImage API synchronously
func (client *Client) InspectImage(request *InspectImageRequest) (response *InspectImageResponse, err error) {
	response = CreateInspectImageResponse()
	err = client.DoAction(request, response)
	return
}

// InspectImageWithChan invokes the ehpc.InspectImage API asynchronously
func (client *Client) InspectImageWithChan(request *InspectImageRequest) (<-chan *InspectImageResponse, <-chan error) {
	responseChan := make(chan *InspectImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InspectImage(request)
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

// InspectImageWithCallback invokes the ehpc.InspectImage API asynchronously
func (client *Client) InspectImageWithCallback(request *InspectImageRequest, callback func(response *InspectImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InspectImageResponse
		var err error
		defer close(result)
		response, err = client.InspectImage(request)
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

// InspectImageRequest is the request struct for api InspectImage
type InspectImageRequest struct {
	*requests.RpcRequest
	ClusterId     string `position:"Query" name:"ClusterId"`
	ContainerType string `position:"Query" name:"ContainerType"`
	ImageName     string `position:"Query" name:"ImageName"`
}

// InspectImageResponse is the response struct for api InspectImage
type InspectImageResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	ImageStatus ImageStatus `json:"ImageStatus" xml:"ImageStatus"`
}

// CreateInspectImageRequest creates a request to invoke InspectImage API
func CreateInspectImageRequest() (request *InspectImageRequest) {
	request = &InspectImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "InspectImage", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateInspectImageResponse creates a response to parse from InspectImage response
func CreateInspectImageResponse() (response *InspectImageResponse) {
	response = &InspectImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
