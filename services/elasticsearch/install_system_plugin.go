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

// InstallSystemPlugin invokes the elasticsearch.InstallSystemPlugin API synchronously
func (client *Client) InstallSystemPlugin(request *InstallSystemPluginRequest) (response *InstallSystemPluginResponse, err error) {
	response = CreateInstallSystemPluginResponse()
	err = client.DoAction(request, response)
	return
}

// InstallSystemPluginWithChan invokes the elasticsearch.InstallSystemPlugin API asynchronously
func (client *Client) InstallSystemPluginWithChan(request *InstallSystemPluginRequest) (<-chan *InstallSystemPluginResponse, <-chan error) {
	responseChan := make(chan *InstallSystemPluginResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallSystemPlugin(request)
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

// InstallSystemPluginWithCallback invokes the elasticsearch.InstallSystemPlugin API asynchronously
func (client *Client) InstallSystemPluginWithCallback(request *InstallSystemPluginRequest, callback func(response *InstallSystemPluginResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallSystemPluginResponse
		var err error
		defer close(result)
		response, err = client.InstallSystemPlugin(request)
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

// InstallSystemPluginRequest is the request struct for api InstallSystemPlugin
type InstallSystemPluginRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	Body        string `position:"Body" name:"body"`
}

// InstallSystemPluginResponse is the response struct for api InstallSystemPlugin
type InstallSystemPluginResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Result    []string `json:"Result" xml:"Result"`
}

// CreateInstallSystemPluginRequest creates a request to invoke InstallSystemPlugin API
func CreateInstallSystemPluginRequest() (request *InstallSystemPluginRequest) {
	request = &InstallSystemPluginRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "InstallSystemPlugin", "/openapi/instances/[InstanceId]/plugins/system/actions/install", "elasticsearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallSystemPluginResponse creates a response to parse from InstallSystemPlugin response
func CreateInstallSystemPluginResponse() (response *InstallSystemPluginResponse) {
	response = &InstallSystemPluginResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
