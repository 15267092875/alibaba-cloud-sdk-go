package alb

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

// CreateRule invokes the alb.CreateRule API synchronously
func (client *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
	response = CreateCreateRuleResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRuleWithChan invokes the alb.CreateRule API asynchronously
func (client *Client) CreateRuleWithChan(request *CreateRuleRequest) (<-chan *CreateRuleResponse, <-chan error) {
	responseChan := make(chan *CreateRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRule(request)
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

// CreateRuleWithCallback invokes the alb.CreateRule API asynchronously
func (client *Client) CreateRuleWithCallback(request *CreateRuleRequest, callback func(response *CreateRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRuleResponse
		var err error
		defer close(result)
		response, err = client.CreateRule(request)
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

// CreateRuleRequest is the request struct for api CreateRule
type CreateRuleRequest struct {
	*requests.RpcRequest
	ClientToken    string                      `position:"Query" name:"ClientToken"`
	RuleName       string                      `position:"Query" name:"RuleName"`
	ListenerId     string                      `position:"Query" name:"ListenerId"`
	Tag            *[]CreateRuleTag            `position:"Query" name:"Tag"  type:"Repeated"`
	Direction      string                      `position:"Query" name:"Direction"`
	RuleActions    *[]CreateRuleRuleActions    `position:"Query" name:"RuleActions"  type:"Repeated"`
	RuleConditions *[]CreateRuleRuleConditions `position:"Query" name:"RuleConditions"  type:"Repeated"`
	DryRun         requests.Boolean            `position:"Query" name:"DryRun"`
	Priority       requests.Integer            `position:"Query" name:"Priority"`
}

// CreateRuleTag is a repeated param struct in CreateRuleRequest
type CreateRuleTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateRuleRuleActions is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActions struct {
	FixedResponseConfig CreateRuleRuleActionsFixedResponseConfig `name:"FixedResponseConfig" type:"Struct"`
	TrafficMirrorConfig CreateRuleRuleActionsTrafficMirrorConfig `name:"TrafficMirrorConfig" type:"Struct"`
	ForwardGroupConfig  CreateRuleRuleActionsForwardGroupConfig  `name:"ForwardGroupConfig" type:"Struct"`
	RemoveHeaderConfig  CreateRuleRuleActionsRemoveHeaderConfig  `name:"RemoveHeaderConfig" type:"Struct"`
	InsertHeaderConfig  CreateRuleRuleActionsInsertHeaderConfig  `name:"InsertHeaderConfig" type:"Struct"`
	TrafficLimitConfig  CreateRuleRuleActionsTrafficLimitConfig  `name:"TrafficLimitConfig" type:"Struct"`
	CorsConfig          CreateRuleRuleActionsCorsConfig          `name:"CorsConfig" type:"Struct"`
	RedirectConfig      CreateRuleRuleActionsRedirectConfig      `name:"RedirectConfig" type:"Struct"`
	Type                string                                   `name:"Type"`
	Order               string                                   `name:"Order"`
	RewriteConfig       CreateRuleRuleActionsRewriteConfig       `name:"RewriteConfig" type:"Struct"`
}

// CreateRuleRuleConditions is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditions struct {
	MethodConfig             CreateRuleRuleConditionsMethodConfig             `name:"MethodConfig" type:"Struct"`
	SourceIpConfig           CreateRuleRuleConditionsSourceIpConfig           `name:"SourceIpConfig" type:"Struct"`
	HostConfig               CreateRuleRuleConditionsHostConfig               `name:"HostConfig" type:"Struct"`
	QueryStringConfig        CreateRuleRuleConditionsQueryStringConfig        `name:"QueryStringConfig" type:"Struct"`
	ResponseStatusCodeConfig CreateRuleRuleConditionsResponseStatusCodeConfig `name:"ResponseStatusCodeConfig" type:"Struct"`
	PathConfig               CreateRuleRuleConditionsPathConfig               `name:"PathConfig" type:"Struct"`
	CookieConfig             CreateRuleRuleConditionsCookieConfig             `name:"CookieConfig" type:"Struct"`
	Type                     string                                           `name:"Type"`
	HeaderConfig             CreateRuleRuleConditionsHeaderConfig             `name:"HeaderConfig" type:"Struct"`
	ResponseHeaderConfig     CreateRuleRuleConditionsResponseHeaderConfig     `name:"ResponseHeaderConfig" type:"Struct"`
}

// CreateRuleRuleActionsFixedResponseConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsFixedResponseConfig struct {
	HttpCode    string `name:"HttpCode"`
	Content     string `name:"Content"`
	ContentType string `name:"ContentType"`
}

// CreateRuleRuleActionsTrafficMirrorConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsTrafficMirrorConfig struct {
	MirrorGroupConfig CreateRuleRuleActionsTrafficMirrorConfigMirrorGroupConfig `name:"MirrorGroupConfig" type:"Struct"`
	TargetType        string                                                    `name:"TargetType"`
}

// CreateRuleRuleActionsForwardGroupConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsForwardGroupConfig struct {
	ServerGroupStickySession CreateRuleRuleActionsForwardGroupConfigServerGroupStickySession `name:"ServerGroupStickySession" type:"Struct"`
	ServerGroupTuples        *[]CreateRuleRuleActionsForwardGroupConfigServerGroupTuplesItem `name:"ServerGroupTuples" type:"Repeated"`
}

// CreateRuleRuleActionsRemoveHeaderConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsRemoveHeaderConfig struct {
	Key string `name:"Key"`
}

// CreateRuleRuleActionsInsertHeaderConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsInsertHeaderConfig struct {
	ValueType    string `name:"ValueType"`
	CoverEnabled string `name:"CoverEnabled"`
	Value        string `name:"Value"`
	Key          string `name:"Key"`
}

// CreateRuleRuleActionsTrafficLimitConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsTrafficLimitConfig struct {
	QPS      string `name:"QPS"`
	PerIpQps string `name:"PerIpQps"`
}

// CreateRuleRuleActionsCorsConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsCorsConfig struct {
	AllowCredentials string    `name:"AllowCredentials"`
	AllowOrigin      *[]string `name:"AllowOrigin" type:"Repeated"`
	MaxAge           string    `name:"MaxAge"`
	AllowMethods     *[]string `name:"AllowMethods" type:"Repeated"`
	AllowHeaders     *[]string `name:"AllowHeaders" type:"Repeated"`
	ExposeHeaders    *[]string `name:"ExposeHeaders" type:"Repeated"`
}

// CreateRuleRuleActionsRedirectConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsRedirectConfig struct {
	Path     string `name:"Path"`
	Protocol string `name:"Protocol"`
	Port     string `name:"Port"`
	Query    string `name:"Query"`
	Host     string `name:"Host"`
	HttpCode string `name:"HttpCode"`
}

// CreateRuleRuleActionsRewriteConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsRewriteConfig struct {
	Path  string `name:"Path"`
	Query string `name:"Query"`
	Host  string `name:"Host"`
}

// CreateRuleRuleConditionsMethodConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsMethodConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsSourceIpConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsSourceIpConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsHostConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsHostConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsQueryStringConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsQueryStringConfig struct {
	Values *[]CreateRuleRuleConditionsQueryStringConfigValuesItem `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsResponseStatusCodeConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsResponseStatusCodeConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsPathConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsPathConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsCookieConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsCookieConfig struct {
	Values *[]CreateRuleRuleConditionsCookieConfigValuesItem `name:"Values" type:"Repeated"`
}

// CreateRuleRuleConditionsHeaderConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsHeaderConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
	Key    string    `name:"Key"`
}

// CreateRuleRuleConditionsResponseHeaderConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsResponseHeaderConfig struct {
	Values *[]string `name:"Values" type:"Repeated"`
	Key    string    `name:"Key"`
}

// CreateRuleRuleActionsTrafficMirrorConfigMirrorGroupConfig is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsTrafficMirrorConfigMirrorGroupConfig struct {
	ServerGroupTuples *[]CreateRuleRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuplesItem `name:"ServerGroupTuples" type:"Repeated"`
}

// CreateRuleRuleActionsForwardGroupConfigServerGroupStickySession is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsForwardGroupConfigServerGroupStickySession struct {
	Enabled string `name:"Enabled"`
	Timeout string `name:"Timeout"`
}

// CreateRuleRuleActionsForwardGroupConfigServerGroupTuplesItem is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsForwardGroupConfigServerGroupTuplesItem struct {
	ServerGroupId string `name:"ServerGroupId"`
	Weight        string `name:"Weight"`
}

// CreateRuleRuleConditionsQueryStringConfigValuesItem is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsQueryStringConfigValuesItem struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateRuleRuleConditionsCookieConfigValuesItem is a repeated param struct in CreateRuleRequest
type CreateRuleRuleConditionsCookieConfigValuesItem struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateRuleRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuplesItem is a repeated param struct in CreateRuleRequest
type CreateRuleRuleActionsTrafficMirrorConfigMirrorGroupConfigServerGroupTuplesItem struct {
	ServerGroupId string `name:"ServerGroupId"`
}

// CreateRuleResponse is the response struct for api CreateRule
type CreateRuleResponse struct {
	*responses.BaseResponse
	JobId     string `json:"JobId" xml:"JobId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	RuleId    string `json:"RuleId" xml:"RuleId"`
}

// CreateCreateRuleRequest creates a request to invoke CreateRule API
func CreateCreateRuleRequest() (request *CreateRuleRequest) {
	request = &CreateRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alb", "2020-06-16", "CreateRule", "alb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateRuleResponse creates a response to parse from CreateRule response
func CreateCreateRuleResponse() (response *CreateRuleResponse) {
	response = &CreateRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
