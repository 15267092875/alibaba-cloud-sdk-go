package oceanbasepro

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

// DataInDescribeTenants is a nested struct in oceanbasepro response
type DataInDescribeTenants struct {
	VpcId                 string `json:"VpcId" xml:"VpcId"`
	Status                string `json:"Status" xml:"Status"`
	PrimaryZone           string `json:"PrimaryZone" xml:"PrimaryZone"`
	DeployType            string `json:"DeployType" xml:"DeployType"`
	DeployMode            string `json:"DeployMode" xml:"DeployMode"`
	CreateTime            string `json:"CreateTime" xml:"CreateTime"`
	TenantName            string `json:"TenantName" xml:"TenantName"`
	Mem                   int    `json:"Mem" xml:"Mem"`
	Cpu                   int    `json:"Cpu" xml:"Cpu"`
	Description           string `json:"Description" xml:"Description"`
	TenantMode            string `json:"TenantMode" xml:"TenantMode"`
	TenantId              string `json:"TenantId" xml:"TenantId"`
	UnitCpu               int    `json:"UnitCpu" xml:"UnitCpu"`
	UnitMem               int    `json:"UnitMem" xml:"UnitMem"`
	UnitNum               int    `json:"UnitNum" xml:"UnitNum"`
	UsedDiskSize          string `json:"UsedDiskSize" xml:"UsedDiskSize"`
	Charset               string `json:"Charset" xml:"Charset"`
	Collation             string `json:"Collation" xml:"Collation"`
	EnableReadOnlyReplica bool   `json:"EnableReadOnlyReplica" xml:"EnableReadOnlyReplica"`
}
