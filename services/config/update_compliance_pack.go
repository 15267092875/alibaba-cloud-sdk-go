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

// UpdateCompliancePack invokes the config.UpdateCompliancePack API synchronously
func (client *Client) UpdateCompliancePack(request *UpdateCompliancePackRequest) (response *UpdateCompliancePackResponse, err error) {
	response = CreateUpdateCompliancePackResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateCompliancePackWithChan invokes the config.UpdateCompliancePack API asynchronously
func (client *Client) UpdateCompliancePackWithChan(request *UpdateCompliancePackRequest) (<-chan *UpdateCompliancePackResponse, <-chan error) {
	responseChan := make(chan *UpdateCompliancePackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateCompliancePack(request)
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

// UpdateCompliancePackWithCallback invokes the config.UpdateCompliancePack API asynchronously
func (client *Client) UpdateCompliancePackWithCallback(request *UpdateCompliancePackRequest, callback func(response *UpdateCompliancePackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateCompliancePackResponse
		var err error
		defer close(result)
		response, err = client.UpdateCompliancePack(request)
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

// UpdateCompliancePackRequest is the request struct for api UpdateCompliancePack
type UpdateCompliancePackRequest struct {
	*requests.RpcRequest
	TagKeyScope             string                             `position:"Body" name:"TagKeyScope"`
	CompliancePackName      string                             `position:"Body" name:"CompliancePackName"`
	ClientToken             string                             `position:"Body" name:"ClientToken"`
	Description             string                             `position:"Body" name:"Description"`
	TagValueScope           string                             `position:"Body" name:"TagValueScope"`
	RegionIdsScope          string                             `position:"Body" name:"RegionIdsScope"`
	CompliancePackId        string                             `position:"Body" name:"CompliancePackId"`
	ConfigRules             *[]UpdateCompliancePackConfigRules `position:"Body" name:"ConfigRules"  type:"Json"`
	RiskLevel               requests.Integer                   `position:"Body" name:"RiskLevel"`
	ResourceGroupIdsScope   string                             `position:"Body" name:"ResourceGroupIdsScope"`
	ExcludeResourceIdsScope string                             `position:"Body" name:"ExcludeResourceIdsScope"`
}

// UpdateCompliancePackConfigRules is a repeated param struct in UpdateCompliancePackRequest
type UpdateCompliancePackConfigRules struct {
	ManagedRuleIdentifier string                                                     `name:"ManagedRuleIdentifier"`
	ConfigRuleParameters  *[]UpdateCompliancePackConfigRulesConfigRuleParametersItem `name:"ConfigRuleParameters" type:"Repeated"`
	ConfigRuleId          string                                                     `name:"ConfigRuleId"`
	ConfigRuleName        string                                                     `name:"ConfigRuleName"`
	Description           string                                                     `name:"Description"`
	RiskLevel             string                                                     `name:"RiskLevel"`
}

// UpdateCompliancePackConfigRulesConfigRuleParametersItem is a repeated param struct in UpdateCompliancePackRequest
type UpdateCompliancePackConfigRulesConfigRuleParametersItem struct {
	ParameterValue string `name:"ParameterValue"`
	ParameterName  string `name:"ParameterName"`
}

// UpdateCompliancePackResponse is the response struct for api UpdateCompliancePack
type UpdateCompliancePackResponse struct {
	*responses.BaseResponse
	CompliancePackId string `json:"CompliancePackId" xml:"CompliancePackId"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateCompliancePackRequest creates a request to invoke UpdateCompliancePack API
func CreateUpdateCompliancePackRequest() (request *UpdateCompliancePackRequest) {
	request = &UpdateCompliancePackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "UpdateCompliancePack", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateCompliancePackResponse creates a response to parse from UpdateCompliancePack response
func CreateUpdateCompliancePackResponse() (response *UpdateCompliancePackResponse) {
	response = &UpdateCompliancePackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
