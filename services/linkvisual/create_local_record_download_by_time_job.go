package linkvisual

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

// CreateLocalRecordDownloadByTimeJob invokes the linkvisual.CreateLocalRecordDownloadByTimeJob API synchronously
func (client *Client) CreateLocalRecordDownloadByTimeJob(request *CreateLocalRecordDownloadByTimeJobRequest) (response *CreateLocalRecordDownloadByTimeJobResponse, err error) {
	response = CreateCreateLocalRecordDownloadByTimeJobResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLocalRecordDownloadByTimeJobWithChan invokes the linkvisual.CreateLocalRecordDownloadByTimeJob API asynchronously
func (client *Client) CreateLocalRecordDownloadByTimeJobWithChan(request *CreateLocalRecordDownloadByTimeJobRequest) (<-chan *CreateLocalRecordDownloadByTimeJobResponse, <-chan error) {
	responseChan := make(chan *CreateLocalRecordDownloadByTimeJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLocalRecordDownloadByTimeJob(request)
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

// CreateLocalRecordDownloadByTimeJobWithCallback invokes the linkvisual.CreateLocalRecordDownloadByTimeJob API asynchronously
func (client *Client) CreateLocalRecordDownloadByTimeJobWithCallback(request *CreateLocalRecordDownloadByTimeJobRequest, callback func(response *CreateLocalRecordDownloadByTimeJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLocalRecordDownloadByTimeJobResponse
		var err error
		defer close(result)
		response, err = client.CreateLocalRecordDownloadByTimeJob(request)
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

// CreateLocalRecordDownloadByTimeJobRequest is the request struct for api CreateLocalRecordDownloadByTimeJob
type CreateLocalRecordDownloadByTimeJobRequest struct {
	*requests.RpcRequest
	Speed         requests.Float   `position:"Query" name:"Speed"`
	IotId         string           `position:"Query" name:"IotId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	StreamType    requests.Integer `position:"Query" name:"StreamType"`
	EndTime       requests.Integer `position:"Query" name:"EndTime"`
	BeginTime     requests.Integer `position:"Query" name:"BeginTime"`
	ProductKey    string           `position:"Query" name:"ProductKey"`
	ApiProduct    string           `position:"Body" name:"ApiProduct"`
	ApiRevision   string           `position:"Body" name:"ApiRevision"`
	DeviceName    string           `position:"Query" name:"DeviceName"`
}

// CreateLocalRecordDownloadByTimeJobResponse is the response struct for api CreateLocalRecordDownloadByTimeJob
type CreateLocalRecordDownloadByTimeJobResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateCreateLocalRecordDownloadByTimeJobRequest creates a request to invoke CreateLocalRecordDownloadByTimeJob API
func CreateCreateLocalRecordDownloadByTimeJobRequest() (request *CreateLocalRecordDownloadByTimeJobRequest) {
	request = &CreateLocalRecordDownloadByTimeJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "CreateLocalRecordDownloadByTimeJob", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLocalRecordDownloadByTimeJobResponse creates a response to parse from CreateLocalRecordDownloadByTimeJob response
func CreateCreateLocalRecordDownloadByTimeJobResponse() (response *CreateLocalRecordDownloadByTimeJobResponse) {
	response = &CreateLocalRecordDownloadByTimeJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
