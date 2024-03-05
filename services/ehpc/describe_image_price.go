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

// DescribeImagePrice invokes the ehpc.DescribeImagePrice API synchronously
func (client *Client) DescribeImagePrice(request *DescribeImagePriceRequest) (response *DescribeImagePriceResponse, err error) {
	response = CreateDescribeImagePriceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeImagePriceWithChan invokes the ehpc.DescribeImagePrice API asynchronously
func (client *Client) DescribeImagePriceWithChan(request *DescribeImagePriceRequest) (<-chan *DescribeImagePriceResponse, <-chan error) {
	responseChan := make(chan *DescribeImagePriceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeImagePrice(request)
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

// DescribeImagePriceWithCallback invokes the ehpc.DescribeImagePrice API asynchronously
func (client *Client) DescribeImagePriceWithCallback(request *DescribeImagePriceRequest, callback func(response *DescribeImagePriceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeImagePriceResponse
		var err error
		defer close(result)
		response, err = client.DescribeImagePrice(request)
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

// DescribeImagePriceRequest is the request struct for api DescribeImagePrice
type DescribeImagePriceRequest struct {
	*requests.RpcRequest
	Period    requests.Integer `position:"Query" name:"Period"`
	Amount    requests.Integer `position:"Query" name:"Amount"`
	ImageId   string           `position:"Query" name:"ImageId"`
	SkuCode   string           `position:"Query" name:"SkuCode"`
	PriceUnit string           `position:"Query" name:"PriceUnit"`
	OrderType string           `position:"Query" name:"OrderType"`
}

// DescribeImagePriceResponse is the response struct for api DescribeImagePrice
type DescribeImagePriceResponse struct {
	*responses.BaseResponse
	Amount        int     `json:"Amount" xml:"Amount"`
	RequestId     string  `json:"RequestId" xml:"RequestId"`
	DiscountPrice float64 `json:"DiscountPrice" xml:"DiscountPrice"`
	TradePrice    float64 `json:"TradePrice" xml:"TradePrice"`
	OriginalPrice float64 `json:"OriginalPrice" xml:"OriginalPrice"`
	ImageId       string  `json:"ImageId" xml:"ImageId"`
}

// CreateDescribeImagePriceRequest creates a request to invoke DescribeImagePrice API
func CreateDescribeImagePriceRequest() (request *DescribeImagePriceRequest) {
	request = &DescribeImagePriceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DescribeImagePrice", "ehs", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeImagePriceResponse creates a response to parse from DescribeImagePrice response
func CreateDescribeImagePriceResponse() (response *DescribeImagePriceResponse) {
	response = &DescribeImagePriceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
