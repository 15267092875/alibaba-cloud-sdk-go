package sgw

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

// BucketCache is a nested struct in sgw response
type BucketCache struct {
	GatewayId   string `json:"GatewayId" xml:"GatewayId"`
	GatewayName string `json:"GatewayName" xml:"GatewayName"`
	RegionId    string `json:"RegionId" xml:"RegionId"`
	Location    string `json:"Location" xml:"Location"`
	Category    string `json:"Category" xml:"Category"`
	Type        string `json:"Type" xml:"Type"`
	BucketName  string `json:"BucketName" xml:"BucketName"`
	Protocol    string `json:"Protocol" xml:"Protocol"`
	CacheMode   string `json:"CacheMode" xml:"CacheMode"`
	CacheStats  string `json:"CacheStats" xml:"CacheStats"`
	ShareName   string `json:"ShareName" xml:"ShareName"`
	VpcId       string `json:"VpcId" xml:"VpcId"`
	VpcCidr     string `json:"VpcCidr" xml:"VpcCidr"`
	MountPoint  string `json:"MountPoint" xml:"MountPoint"`
}