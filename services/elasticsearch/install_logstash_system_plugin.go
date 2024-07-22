package elasticsearch

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

// InstallLogstashSystemPlugin invokes the elasticsearch.InstallLogstashSystemPlugin API synchronously
func (client *Client) InstallLogstashSystemPlugin(request *InstallLogstashSystemPluginRequest) (response *InstallLogstashSystemPluginResponse, err error) {
	response = CreateInstallLogstashSystemPluginResponse()
	err = client.DoAction(request, response)
	return
}

// InstallLogstashSystemPluginWithChan invokes the elasticsearch.InstallLogstashSystemPlugin API asynchronously
func (client *Client) InstallLogstashSystemPluginWithChan(request *InstallLogstashSystemPluginRequest) (<-chan *InstallLogstashSystemPluginResponse, <-chan error) {
	responseChan := make(chan *InstallLogstashSystemPluginResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallLogstashSystemPlugin(request)
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

// InstallLogstashSystemPluginWithCallback invokes the elasticsearch.InstallLogstashSystemPlugin API asynchronously
func (client *Client) InstallLogstashSystemPluginWithCallback(request *InstallLogstashSystemPluginRequest, callback func(response *InstallLogstashSystemPluginResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallLogstashSystemPluginResponse
		var err error
		defer close(result)
		response, err = client.InstallLogstashSystemPlugin(request)
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

// InstallLogstashSystemPluginRequest is the request struct for api InstallLogstashSystemPlugin
type InstallLogstashSystemPluginRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	Body        string `position:"Body" name:"body"`
}

// InstallLogstashSystemPluginResponse is the response struct for api InstallLogstashSystemPlugin
type InstallLogstashSystemPluginResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Result    []string `json:"Result" xml:"Result"`
}

// CreateInstallLogstashSystemPluginRequest creates a request to invoke InstallLogstashSystemPlugin API
func CreateInstallLogstashSystemPluginRequest() (request *InstallLogstashSystemPluginRequest) {
	request = &InstallLogstashSystemPluginRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "InstallLogstashSystemPlugin", "/openapi/logstashes/[InstanceId]/plugins/system/actions/install", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallLogstashSystemPluginResponse creates a response to parse from InstallLogstashSystemPlugin response
func CreateInstallLogstashSystemPluginResponse() (response *InstallLogstashSystemPluginResponse) {
	response = &InstallLogstashSystemPluginResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
