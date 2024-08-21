package ess

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

// CreateNotificationConfiguration invokes the ess.CreateNotificationConfiguration API synchronously
func (client *Client) CreateNotificationConfiguration(request *CreateNotificationConfigurationRequest) (response *CreateNotificationConfigurationResponse, err error) {
	response = CreateCreateNotificationConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNotificationConfigurationWithChan invokes the ess.CreateNotificationConfiguration API asynchronously
func (client *Client) CreateNotificationConfigurationWithChan(request *CreateNotificationConfigurationRequest) (<-chan *CreateNotificationConfigurationResponse, <-chan error) {
	responseChan := make(chan *CreateNotificationConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNotificationConfiguration(request)
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

// CreateNotificationConfigurationWithCallback invokes the ess.CreateNotificationConfiguration API asynchronously
func (client *Client) CreateNotificationConfigurationWithCallback(request *CreateNotificationConfigurationRequest, callback func(response *CreateNotificationConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNotificationConfigurationResponse
		var err error
		defer close(result)
		response, err = client.CreateNotificationConfiguration(request)
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

// CreateNotificationConfigurationRequest is the request struct for api CreateNotificationConfiguration
type CreateNotificationConfigurationRequest struct {
	*requests.RpcRequest
	ScalingGroupId       string           `position:"Query" name:"ScalingGroupId"`
	TimeZone             string           `position:"Query" name:"TimeZone"`
	NotificationArn      string           `position:"Query" name:"NotificationArn"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NotificationType     *[]string        `position:"Query" name:"NotificationType"  type:"Repeated"`
}

// CreateNotificationConfigurationResponse is the response struct for api CreateNotificationConfiguration
type CreateNotificationConfigurationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateNotificationConfigurationRequest creates a request to invoke CreateNotificationConfiguration API
func CreateCreateNotificationConfigurationRequest() (request *CreateNotificationConfigurationRequest) {
	request = &CreateNotificationConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ess", "2014-08-28", "CreateNotificationConfiguration", "ess", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNotificationConfigurationResponse creates a response to parse from CreateNotificationConfiguration response
func CreateCreateNotificationConfigurationResponse() (response *CreateNotificationConfigurationResponse) {
	response = &CreateNotificationConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
