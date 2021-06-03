package hitsdb

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

// SwitchHiTSDBInstancePublicNet invokes the hitsdb.SwitchHiTSDBInstancePublicNet API synchronously
func (client *Client) SwitchHiTSDBInstancePublicNet(request *SwitchHiTSDBInstancePublicNetRequest) (response *SwitchHiTSDBInstancePublicNetResponse, err error) {
	response = CreateSwitchHiTSDBInstancePublicNetResponse()
	err = client.DoAction(request, response)
	return
}

// SwitchHiTSDBInstancePublicNetWithChan invokes the hitsdb.SwitchHiTSDBInstancePublicNet API asynchronously
func (client *Client) SwitchHiTSDBInstancePublicNetWithChan(request *SwitchHiTSDBInstancePublicNetRequest) (<-chan *SwitchHiTSDBInstancePublicNetResponse, <-chan error) {
	responseChan := make(chan *SwitchHiTSDBInstancePublicNetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SwitchHiTSDBInstancePublicNet(request)
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

// SwitchHiTSDBInstancePublicNetWithCallback invokes the hitsdb.SwitchHiTSDBInstancePublicNet API asynchronously
func (client *Client) SwitchHiTSDBInstancePublicNetWithCallback(request *SwitchHiTSDBInstancePublicNetRequest, callback func(response *SwitchHiTSDBInstancePublicNetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SwitchHiTSDBInstancePublicNetResponse
		var err error
		defer close(result)
		response, err = client.SwitchHiTSDBInstancePublicNet(request)
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

// SwitchHiTSDBInstancePublicNetRequest is the request struct for api SwitchHiTSDBInstancePublicNet
type SwitchHiTSDBInstancePublicNetRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	SwitchAction         requests.Integer `position:"Query" name:"SwitchAction"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// SwitchHiTSDBInstancePublicNetResponse is the response struct for api SwitchHiTSDBInstancePublicNet
type SwitchHiTSDBInstancePublicNetResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSwitchHiTSDBInstancePublicNetRequest creates a request to invoke SwitchHiTSDBInstancePublicNet API
func CreateSwitchHiTSDBInstancePublicNetRequest() (request *SwitchHiTSDBInstancePublicNetRequest) {
	request = &SwitchHiTSDBInstancePublicNetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("hitsdb", "2017-06-01", "SwitchHiTSDBInstancePublicNet", "hitsdb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSwitchHiTSDBInstancePublicNetResponse creates a response to parse from SwitchHiTSDBInstancePublicNet response
func CreateSwitchHiTSDBInstancePublicNetResponse() (response *SwitchHiTSDBInstancePublicNetResponse) {
	response = &SwitchHiTSDBInstancePublicNetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}