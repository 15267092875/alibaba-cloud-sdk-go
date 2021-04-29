package unimkt

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

// QueryAppPromotion invokes the unimkt.QueryAppPromotion API synchronously
func (client *Client) QueryAppPromotion(request *QueryAppPromotionRequest) (response *QueryAppPromotionResponse, err error) {
	response = CreateQueryAppPromotionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAppPromotionWithChan invokes the unimkt.QueryAppPromotion API asynchronously
func (client *Client) QueryAppPromotionWithChan(request *QueryAppPromotionRequest) (<-chan *QueryAppPromotionResponse, <-chan error) {
	responseChan := make(chan *QueryAppPromotionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAppPromotion(request)
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

// QueryAppPromotionWithCallback invokes the unimkt.QueryAppPromotion API asynchronously
func (client *Client) QueryAppPromotionWithCallback(request *QueryAppPromotionRequest, callback func(response *QueryAppPromotionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAppPromotionResponse
		var err error
		defer close(result)
		response, err = client.QueryAppPromotion(request)
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

// QueryAppPromotionRequest is the request struct for api QueryAppPromotion
type QueryAppPromotionRequest struct {
	*requests.RpcRequest
	Extra     string `position:"Query" name:"Extra"`
	UserType  string `position:"Query" name:"UserType"`
	UserId    string `position:"Query" name:"UserId"`
	ChannelId string `position:"Query" name:"ChannelId"`
}

// QueryAppPromotionResponse is the response struct for api QueryAppPromotion
type QueryAppPromotionResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Result       string `json:"Result" xml:"Result"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateQueryAppPromotionRequest creates a request to invoke QueryAppPromotion API
func CreateQueryAppPromotionRequest() (request *QueryAppPromotionRequest) {
	request = &QueryAppPromotionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "QueryAppPromotion", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryAppPromotionResponse creates a response to parse from QueryAppPromotion response
func CreateQueryAppPromotionResponse() (response *QueryAppPromotionResponse) {
	response = &QueryAppPromotionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}