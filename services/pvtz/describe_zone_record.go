package pvtz

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

// DescribeZoneRecord invokes the pvtz.DescribeZoneRecord API synchronously
func (client *Client) DescribeZoneRecord(request *DescribeZoneRecordRequest) (response *DescribeZoneRecordResponse, err error) {
	response = CreateDescribeZoneRecordResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeZoneRecordWithChan invokes the pvtz.DescribeZoneRecord API asynchronously
func (client *Client) DescribeZoneRecordWithChan(request *DescribeZoneRecordRequest) (<-chan *DescribeZoneRecordResponse, <-chan error) {
	responseChan := make(chan *DescribeZoneRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeZoneRecord(request)
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

// DescribeZoneRecordWithCallback invokes the pvtz.DescribeZoneRecord API asynchronously
func (client *Client) DescribeZoneRecordWithCallback(request *DescribeZoneRecordRequest, callback func(response *DescribeZoneRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeZoneRecordResponse
		var err error
		defer close(result)
		response, err = client.DescribeZoneRecord(request)
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

// DescribeZoneRecordRequest is the request struct for api DescribeZoneRecord
type DescribeZoneRecordRequest struct {
	*requests.RpcRequest
	RecordId requests.Integer `position:"Query" name:"RecordId"`
}

// DescribeZoneRecordResponse is the response struct for api DescribeZoneRecord
type DescribeZoneRecordResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	Status          string `json:"Status" xml:"Status"`
	Type            string `json:"Type" xml:"Type"`
	Value           string `json:"Value" xml:"Value"`
	Ttl             int    `json:"Ttl" xml:"Ttl"`
	Remark          string `json:"Remark" xml:"Remark"`
	RecordId        int64  `json:"RecordId" xml:"RecordId"`
	Rr              string `json:"Rr" xml:"Rr"`
	Priority        int    `json:"Priority" xml:"Priority"`
	RegionId        string `json:"RegionId" xml:"RegionId"`
	Line            string `json:"Line" xml:"Line"`
	Weight          int    `json:"Weight" xml:"Weight"`
	CreateTime      string `json:"CreateTime" xml:"CreateTime"`
	UpdateTime      string `json:"UpdateTime" xml:"UpdateTime"`
	CreateTimestamp int64  `json:"CreateTimestamp" xml:"CreateTimestamp"`
	UpdateTimestamp int64  `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	ZoneId          string `json:"ZoneId" xml:"ZoneId"`
}

// CreateDescribeZoneRecordRequest creates a request to invoke DescribeZoneRecord API
func CreateDescribeZoneRecordRequest() (request *DescribeZoneRecordRequest) {
	request = &DescribeZoneRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("pvtz", "2018-01-01", "DescribeZoneRecord", "pvtz", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeZoneRecordResponse creates a response to parse from DescribeZoneRecord response
func CreateDescribeZoneRecordResponse() (response *DescribeZoneRecordResponse) {
	response = &DescribeZoneRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
