package cloudauth

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

// DescribeWhitelist invokes the cloudauth.DescribeWhitelist API synchronously
func (client *Client) DescribeWhitelist(request *DescribeWhitelistRequest) (response *DescribeWhitelistResponse, err error) {
	response = CreateDescribeWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWhitelistWithChan invokes the cloudauth.DescribeWhitelist API asynchronously
func (client *Client) DescribeWhitelistWithChan(request *DescribeWhitelistRequest) (<-chan *DescribeWhitelistResponse, <-chan error) {
	responseChan := make(chan *DescribeWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWhitelist(request)
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

// DescribeWhitelistWithCallback invokes the cloudauth.DescribeWhitelist API asynchronously
func (client *Client) DescribeWhitelistWithCallback(request *DescribeWhitelistRequest, callback func(response *DescribeWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWhitelistResponse
		var err error
		defer close(result)
		response, err = client.DescribeWhitelist(request)
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

// DescribeWhitelistRequest is the request struct for api DescribeWhitelist
type DescribeWhitelistRequest struct {
	*requests.RpcRequest
	ValidEndDate   string           `position:"Query" name:"ValidEndDate"`
	Valid          string           `position:"Query" name:"Valid"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Lang           string           `position:"Query" name:"Lang"`
	CurrentPage    requests.Integer `position:"Query" name:"CurrentPage"`
	BizType        string           `position:"Query" name:"BizType"`
	IdCardNum      string           `position:"Query" name:"IdCardNum"`
	BizId          string           `position:"Query" name:"BizId"`
	ValidStartDate string           `position:"Query" name:"ValidStartDate"`
}

// DescribeWhitelistResponse is the response struct for api DescribeWhitelist
type DescribeWhitelistResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	TotalCount  int    `json:"TotalCount" xml:"TotalCount"`
	CurrentPage int    `json:"CurrentPage" xml:"CurrentPage"`
	PageSize    int    `json:"PageSize" xml:"PageSize"`
	Items       []Item `json:"Items" xml:"Items"`
}

// CreateDescribeWhitelistRequest creates a request to invoke DescribeWhitelist API
func CreateDescribeWhitelistRequest() (request *DescribeWhitelistRequest) {
	request = &DescribeWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "DescribeWhitelist", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeWhitelistResponse creates a response to parse from DescribeWhitelist response
func CreateDescribeWhitelistResponse() (response *DescribeWhitelistResponse) {
	response = &DescribeWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
