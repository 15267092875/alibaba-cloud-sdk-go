package sae

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

// ApplicationScalingRule is a nested struct in sae response
type ApplicationScalingRule struct {
	UpdateTime            int64  `json:"UpdateTime" xml:"UpdateTime"`
	AppId                 string `json:"AppId" xml:"AppId"`
	CreateTime            int64  `json:"CreateTime" xml:"CreateTime"`
	LastDisableTime       int64  `json:"LastDisableTime" xml:"LastDisableTime"`
	ScaleRuleEnabled      bool   `json:"ScaleRuleEnabled" xml:"ScaleRuleEnabled"`
	ScaleRuleType         string `json:"ScaleRuleType" xml:"ScaleRuleType"`
	ScaleRuleName         string `json:"ScaleRuleName" xml:"ScaleRuleName"`
	MinReadyInstances     int    `json:"MinReadyInstances" xml:"MinReadyInstances"`
	MinReadyInstanceRatio int    `json:"MinReadyInstanceRatio" xml:"MinReadyInstanceRatio"`
	Timer                 Timer  `json:"Timer" xml:"Timer"`
	Metric                Metric `json:"Metric" xml:"Metric"`
}
