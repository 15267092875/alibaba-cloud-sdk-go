package config

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

// UpdateAggregateConfigDeliveryChannel invokes the config.UpdateAggregateConfigDeliveryChannel API synchronously
func (client *Client) UpdateAggregateConfigDeliveryChannel(request *UpdateAggregateConfigDeliveryChannelRequest) (response *UpdateAggregateConfigDeliveryChannelResponse, err error) {
	response = CreateUpdateAggregateConfigDeliveryChannelResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAggregateConfigDeliveryChannelWithChan invokes the config.UpdateAggregateConfigDeliveryChannel API asynchronously
func (client *Client) UpdateAggregateConfigDeliveryChannelWithChan(request *UpdateAggregateConfigDeliveryChannelRequest) (<-chan *UpdateAggregateConfigDeliveryChannelResponse, <-chan error) {
	responseChan := make(chan *UpdateAggregateConfigDeliveryChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAggregateConfigDeliveryChannel(request)
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

// UpdateAggregateConfigDeliveryChannelWithCallback invokes the config.UpdateAggregateConfigDeliveryChannel API asynchronously
func (client *Client) UpdateAggregateConfigDeliveryChannelWithCallback(request *UpdateAggregateConfigDeliveryChannelRequest, callback func(response *UpdateAggregateConfigDeliveryChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAggregateConfigDeliveryChannelResponse
		var err error
		defer close(result)
		response, err = client.UpdateAggregateConfigDeliveryChannel(request)
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

// UpdateAggregateConfigDeliveryChannelRequest is the request struct for api UpdateAggregateConfigDeliveryChannel
type UpdateAggregateConfigDeliveryChannelRequest struct {
	*requests.RpcRequest
	NonCompliantNotification            requests.Boolean `position:"Query" name:"NonCompliantNotification"`
	ClientToken                         string           `position:"Query" name:"ClientToken"`
	ConfigurationSnapshot               requests.Boolean `position:"Query" name:"ConfigurationSnapshot"`
	Description                         string           `position:"Query" name:"Description"`
	AggregatorId                        string           `position:"Query" name:"AggregatorId"`
	DeliveryChannelTargetArn            string           `position:"Query" name:"DeliveryChannelTargetArn"`
	DeliveryChannelCondition            string           `position:"Query" name:"DeliveryChannelCondition"`
	ConfigurationItemChangeNotification requests.Boolean `position:"Query" name:"ConfigurationItemChangeNotification"`
	DeliveryChannelName                 string           `position:"Query" name:"DeliveryChannelName"`
	DeliverySnapshotTime                string           `position:"Query" name:"DeliverySnapshotTime"`
	DeliveryChannelId                   string           `position:"Query" name:"DeliveryChannelId"`
	OversizedDataOSSTargetArn           string           `position:"Query" name:"OversizedDataOSSTargetArn"`
	Status                              requests.Integer `position:"Query" name:"Status"`
}

// UpdateAggregateConfigDeliveryChannelResponse is the response struct for api UpdateAggregateConfigDeliveryChannel
type UpdateAggregateConfigDeliveryChannelResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	DeliveryChannelId string `json:"DeliveryChannelId" xml:"DeliveryChannelId"`
}

// CreateUpdateAggregateConfigDeliveryChannelRequest creates a request to invoke UpdateAggregateConfigDeliveryChannel API
func CreateUpdateAggregateConfigDeliveryChannelRequest() (request *UpdateAggregateConfigDeliveryChannelRequest) {
	request = &UpdateAggregateConfigDeliveryChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "UpdateAggregateConfigDeliveryChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAggregateConfigDeliveryChannelResponse creates a response to parse from UpdateAggregateConfigDeliveryChannel response
func CreateUpdateAggregateConfigDeliveryChannelResponse() (response *UpdateAggregateConfigDeliveryChannelResponse) {
	response = &UpdateAggregateConfigDeliveryChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
