package dataworks_public

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

// NodesItemInListNodes is a nested struct in dataworks_public response
type NodesItemInListNodes struct {
	SchedulerType  string `json:"SchedulerType" xml:"SchedulerType"`
	RepeatInterval int64  `json:"RepeatInterval" xml:"RepeatInterval"`
	Repeatability  bool   `json:"Repeatability" xml:"Repeatability"`
	ProjectId      int64  `json:"ProjectId" xml:"ProjectId"`
	ProgramType    string `json:"ProgramType" xml:"ProgramType"`
	Priority       int    `json:"Priority" xml:"Priority"`
	OwnerId        string `json:"OwnerId" xml:"OwnerId"`
	Connection     string `json:"Connection" xml:"Connection"`
	ParamValues    string `json:"ParamValues" xml:"ParamValues"`
	RelatedFlowId  int64  `json:"RelatedFlowId" xml:"RelatedFlowId"`
	DqcType        int    `json:"DqcType" xml:"DqcType"`
	BaselineId     int64  `json:"BaselineId" xml:"BaselineId"`
	Description    string `json:"Description" xml:"Description"`
	NodeName       string `json:"NodeName" xml:"NodeName"`
	ResGroupName   string `json:"ResGroupName" xml:"ResGroupName"`
	BusinessId     int64  `json:"BusinessId" xml:"BusinessId"`
	DqcDescription string `json:"DqcDescription" xml:"DqcDescription"`
	CronExpress    string `json:"CronExpress" xml:"CronExpress"`
	NodeId         int64  `json:"NodeId" xml:"NodeId"`
}
