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

// QueryDomainSpecialBizInfoByDomain invokes the domain.QueryDomainSpecialBizInfoByDomain API synchronously
func (client *Client) QueryDomainSpecialBizInfoByDomain(request *QueryDomainSpecialBizInfoByDomainRequest) (response *QueryDomainSpecialBizInfoByDomainResponse, err error) {
	response = CreateQueryDomainSpecialBizInfoByDomainResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDomainSpecialBizInfoByDomainWithChan invokes the domain.QueryDomainSpecialBizInfoByDomain API asynchronously
func (client *Client) QueryDomainSpecialBizInfoByDomainWithChan(request *QueryDomainSpecialBizInfoByDomainRequest) (<-chan *QueryDomainSpecialBizInfoByDomainResponse, <-chan error) {
	responseChan := make(chan *QueryDomainSpecialBizInfoByDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDomainSpecialBizInfoByDomain(request)
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

// QueryDomainSpecialBizInfoByDomainWithCallback invokes the domain.QueryDomainSpecialBizInfoByDomain API asynchronously
func (client *Client) QueryDomainSpecialBizInfoByDomainWithCallback(request *QueryDomainSpecialBizInfoByDomainRequest, callback func(response *QueryDomainSpecialBizInfoByDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDomainSpecialBizInfoByDomainResponse
		var err error
		defer close(result)
		response, err = client.QueryDomainSpecialBizInfoByDomain(request)
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

// QueryDomainSpecialBizInfoByDomainRequest is the request struct for api QueryDomainSpecialBizInfoByDomain
type QueryDomainSpecialBizInfoByDomainRequest struct {
	*requests.RpcRequest
	DomainName   string `position:"Body" name:"DomainName"`
	BizType      string `position:"Body" name:"BizType"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
}

// QueryDomainSpecialBizInfoByDomainResponse is the response struct for api QueryDomainSpecialBizInfoByDomain
type QueryDomainSpecialBizInfoByDomainResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode    string   `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string   `json:"DynamicMessage" xml:"DynamicMessage"`
	Synchro        bool     `json:"Synchro" xml:"Synchro"`
	ErrorMsg       string   `json:"ErrorMsg" xml:"ErrorMsg"`
	ErrorCode      string   `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool     `json:"Success" xml:"Success"`
	AllowRetry     bool     `json:"AllowRetry" xml:"AllowRetry"`
	AppName        string   `json:"AppName" xml:"AppName"`
	ErrorArgs      []string `json:"ErrorArgs" xml:"ErrorArgs"`
	Module         Module   `json:"Module" xml:"Module"`
}

// CreateQueryDomainSpecialBizInfoByDomainRequest creates a request to invoke QueryDomainSpecialBizInfoByDomain API
func CreateQueryDomainSpecialBizInfoByDomainRequest() (request *QueryDomainSpecialBizInfoByDomainRequest) {
	request = &QueryDomainSpecialBizInfoByDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "QueryDomainSpecialBizInfoByDomain", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDomainSpecialBizInfoByDomainResponse creates a response to parse from QueryDomainSpecialBizInfoByDomain response
func CreateQueryDomainSpecialBizInfoByDomainResponse() (response *QueryDomainSpecialBizInfoByDomainResponse) {
	response = &QueryDomainSpecialBizInfoByDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
