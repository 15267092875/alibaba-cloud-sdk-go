package cloudfw

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

// LocalVpc is a nested struct in cloudfw response
type LocalVpc struct {
	VpcName             string      `json:"VpcName" xml:"VpcName"`
	EniPrivateIpAddress string      `json:"EniPrivateIpAddress" xml:"EniPrivateIpAddress"`
	NetworkInstanceType string      `json:"NetworkInstanceType" xml:"NetworkInstanceType"`
	RouteMode           string      `json:"RouteMode" xml:"RouteMode"`
	RegionNo            string      `json:"RegionNo" xml:"RegionNo"`
	NetworkInstanceName string      `json:"NetworkInstanceName" xml:"NetworkInstanceName"`
	OwnerId             string      `json:"OwnerId" xml:"OwnerId"`
	EniId               string      `json:"EniId" xml:"EniId"`
	TransitRouterId     string      `json:"TransitRouterId" xml:"TransitRouterId"`
	RouterInterfaceId   string      `json:"RouterInterfaceId" xml:"RouterInterfaceId"`
	TransitRouterType   string      `json:"TransitRouterType" xml:"TransitRouterType"`
	VpcId               string      `json:"VpcId" xml:"VpcId"`
	SupportManualMode   string      `json:"SupportManualMode" xml:"SupportManualMode"`
	NetworkInstanceId   string      `json:"NetworkInstanceId" xml:"NetworkInstanceId"`
	ManualVSwitchId     string      `json:"ManualVSwitchId" xml:"ManualVSwitchId"`
	AttachmentId        string      `json:"AttachmentId" xml:"AttachmentId"`
	AttachmentName      string      `json:"AttachmentName" xml:"AttachmentName"`
	DefendCidrList      []string    `json:"DefendCidrList" xml:"DefendCidrList"`
	EniList             []EniInfo   `json:"EniList" xml:"EniList"`
	VpcCidrTableList    []CidrTable `json:"VpcCidrTableList" xml:"VpcCidrTableList"`
}
