package dds

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

// NetworkAddress is a nested struct in dds response
type NetworkAddress struct {
	ExpiredTime    string `json:"ExpiredTime" xml:"ExpiredTime"`
	VPCId          string `json:"VPCId" xml:"VPCId"`
	Port           string `json:"Port" xml:"Port"`
	VSwitchId      string `json:"VSwitchId" xml:"VSwitchId"`
	IPAddress      string `json:"IPAddress" xml:"IPAddress"`
	ConnectionType string `json:"ConnectionType" xml:"ConnectionType"`
	NodeId         string `json:"NodeId" xml:"NodeId"`
	TxtRecord      string `json:"TxtRecord" xml:"TxtRecord"`
	Role           string `json:"Role" xml:"Role"`
	VswitchId      string `json:"VswitchId" xml:"VswitchId"`
	NetworkType    string `json:"NetworkType" xml:"NetworkType"`
	NodeType       string `json:"NodeType" xml:"NodeType"`
	NetworkAddress string `json:"NetworkAddress" xml:"NetworkAddress"`
}
