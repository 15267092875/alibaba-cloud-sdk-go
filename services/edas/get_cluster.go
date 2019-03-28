package edas

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

// GetCluster invokes the edas.GetCluster API synchronously
// api document: https://help.aliyun.com/api/edas/getcluster.html
func (client *Client) GetCluster(request *GetClusterRequest) (response *GetClusterResponse, err error) {
	response = CreateGetClusterResponse()
	err = client.DoAction(request, response)
	return
}

// GetClusterWithChan invokes the edas.GetCluster API asynchronously
// api document: https://help.aliyun.com/api/edas/getcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetClusterWithChan(request *GetClusterRequest) (<-chan *GetClusterResponse, <-chan error) {
	responseChan := make(chan *GetClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCluster(request)
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

// GetClusterWithCallback invokes the edas.GetCluster API asynchronously
// api document: https://help.aliyun.com/api/edas/getcluster.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetClusterWithCallback(request *GetClusterRequest, callback func(response *GetClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetClusterResponse
		var err error
		defer close(result)
		response, err = client.GetCluster(request)
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

// GetClusterRequest is the request struct for api GetCluster
type GetClusterRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// GetClusterResponse is the response struct for api GetCluster
type GetClusterResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	Cluster   Cluster `json:"Cluster" xml:"Cluster"`
}

// CreateGetClusterRequest creates a request to invoke GetCluster API
func CreateGetClusterRequest() (request *GetClusterRequest) {
	request = &GetClusterRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetCluster", "/pop/v5/resource/cluster", "", "")
	request.Method = requests.GET
	return
}

// CreateGetClusterResponse creates a response to parse from GetCluster response
func CreateGetClusterResponse() (response *GetClusterResponse) {
	response = &GetClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
