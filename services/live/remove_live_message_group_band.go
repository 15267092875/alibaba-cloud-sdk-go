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

// RemoveLiveMessageGroupBand invokes the live.RemoveLiveMessageGroupBand API synchronously
func (client *Client) RemoveLiveMessageGroupBand(request *RemoveLiveMessageGroupBandRequest) (response *RemoveLiveMessageGroupBandResponse, err error) {
	response = CreateRemoveLiveMessageGroupBandResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveLiveMessageGroupBandWithChan invokes the live.RemoveLiveMessageGroupBand API asynchronously
func (client *Client) RemoveLiveMessageGroupBandWithChan(request *RemoveLiveMessageGroupBandRequest) (<-chan *RemoveLiveMessageGroupBandResponse, <-chan error) {
	responseChan := make(chan *RemoveLiveMessageGroupBandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveLiveMessageGroupBand(request)
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

// RemoveLiveMessageGroupBandWithCallback invokes the live.RemoveLiveMessageGroupBand API asynchronously
func (client *Client) RemoveLiveMessageGroupBandWithCallback(request *RemoveLiveMessageGroupBandRequest, callback func(response *RemoveLiveMessageGroupBandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveLiveMessageGroupBandResponse
		var err error
		defer close(result)
		response, err = client.RemoveLiveMessageGroupBand(request)
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

// RemoveLiveMessageGroupBandRequest is the request struct for api RemoveLiveMessageGroupBand
type RemoveLiveMessageGroupBandRequest struct {
	*requests.RpcRequest
	GroupId       string    `position:"Query" name:"GroupId"`
	DataCenter    string    `position:"Query" name:"DataCenter"`
	UnbannedUsers *[]string `position:"Query" name:"UnbannedUsers"  type:"Repeated"`
	AppId         string    `position:"Query" name:"AppId"`
}

// RemoveLiveMessageGroupBandResponse is the response struct for api RemoveLiveMessageGroupBand
type RemoveLiveMessageGroupBandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRemoveLiveMessageGroupBandRequest creates a request to invoke RemoveLiveMessageGroupBand API
func CreateRemoveLiveMessageGroupBandRequest() (request *RemoveLiveMessageGroupBandRequest) {
	request = &RemoveLiveMessageGroupBandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "RemoveLiveMessageGroupBand", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRemoveLiveMessageGroupBandResponse creates a response to parse from RemoveLiveMessageGroupBand response
func CreateRemoveLiveMessageGroupBandResponse() (response *RemoveLiveMessageGroupBandResponse) {
	response = &RemoveLiveMessageGroupBandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
