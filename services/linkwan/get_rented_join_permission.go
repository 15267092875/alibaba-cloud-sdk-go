package linkwan

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

// GetRentedJoinPermission invokes the linkwan.GetRentedJoinPermission API synchronously
func (client *Client) GetRentedJoinPermission(request *GetRentedJoinPermissionRequest) (response *GetRentedJoinPermissionResponse, err error) {
	response = CreateGetRentedJoinPermissionResponse()
	err = client.DoAction(request, response)
	return
}

// GetRentedJoinPermissionWithChan invokes the linkwan.GetRentedJoinPermission API asynchronously
func (client *Client) GetRentedJoinPermissionWithChan(request *GetRentedJoinPermissionRequest) (<-chan *GetRentedJoinPermissionResponse, <-chan error) {
	responseChan := make(chan *GetRentedJoinPermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRentedJoinPermission(request)
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

// GetRentedJoinPermissionWithCallback invokes the linkwan.GetRentedJoinPermission API asynchronously
func (client *Client) GetRentedJoinPermissionWithCallback(request *GetRentedJoinPermissionRequest, callback func(response *GetRentedJoinPermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRentedJoinPermissionResponse
		var err error
		defer close(result)
		response, err = client.GetRentedJoinPermission(request)
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

// GetRentedJoinPermissionRequest is the request struct for api GetRentedJoinPermission
type GetRentedJoinPermissionRequest struct {
	*requests.RpcRequest
	JoinPermissionId string `position:"Query" name:"JoinPermissionId"`
	ApiProduct       string `position:"Body" name:"ApiProduct"`
	ApiRevision      string `position:"Body" name:"ApiRevision"`
}

// GetRentedJoinPermissionResponse is the response struct for api GetRentedJoinPermission
type GetRentedJoinPermissionResponse struct {
	*responses.BaseResponse
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	Success   bool                          `json:"Success" xml:"Success"`
	Data      DataInGetRentedJoinPermission `json:"Data" xml:"Data"`
}

// CreateGetRentedJoinPermissionRequest creates a request to invoke GetRentedJoinPermission API
func CreateGetRentedJoinPermissionRequest() (request *GetRentedJoinPermissionRequest) {
	request = &GetRentedJoinPermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "GetRentedJoinPermission", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetRentedJoinPermissionResponse creates a response to parse from GetRentedJoinPermission response
func CreateGetRentedJoinPermissionResponse() (response *GetRentedJoinPermissionResponse) {
	response = &GetRentedJoinPermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}