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

// UpdateSynonymsDicts invokes the elasticsearch.UpdateSynonymsDicts API synchronously
func (client *Client) UpdateSynonymsDicts(request *UpdateSynonymsDictsRequest) (response *UpdateSynonymsDictsResponse, err error) {
	response = CreateUpdateSynonymsDictsResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSynonymsDictsWithChan invokes the elasticsearch.UpdateSynonymsDicts API asynchronously
func (client *Client) UpdateSynonymsDictsWithChan(request *UpdateSynonymsDictsRequest) (<-chan *UpdateSynonymsDictsResponse, <-chan error) {
	responseChan := make(chan *UpdateSynonymsDictsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSynonymsDicts(request)
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

// UpdateSynonymsDictsWithCallback invokes the elasticsearch.UpdateSynonymsDicts API asynchronously
func (client *Client) UpdateSynonymsDictsWithCallback(request *UpdateSynonymsDictsRequest, callback func(response *UpdateSynonymsDictsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSynonymsDictsResponse
		var err error
		defer close(result)
		response, err = client.UpdateSynonymsDicts(request)
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

// UpdateSynonymsDictsRequest is the request struct for api UpdateSynonymsDicts
type UpdateSynonymsDictsRequest struct {
	*requests.RoaRequest
	InstanceId  string `position:"Path" name:"InstanceId"`
	ClientToken string `position:"Query" name:"clientToken"`
	Body        string `position:"Body" name:"body"`
}

// UpdateSynonymsDictsResponse is the response struct for api UpdateSynonymsDicts
type UpdateSynonymsDictsResponse struct {
	*responses.BaseResponse
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Result    []DictList `json:"Result" xml:"Result"`
}

// CreateUpdateSynonymsDictsRequest creates a request to invoke UpdateSynonymsDicts API
func CreateUpdateSynonymsDictsRequest() (request *UpdateSynonymsDictsRequest) {
	request = &UpdateSynonymsDictsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "UpdateSynonymsDicts", "/openapi/instances/[InstanceId]/synonymsDict", "elasticsearch", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateSynonymsDictsResponse creates a response to parse from UpdateSynonymsDicts response
func CreateUpdateSynonymsDictsResponse() (response *UpdateSynonymsDictsResponse) {
	response = &UpdateSynonymsDictsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
