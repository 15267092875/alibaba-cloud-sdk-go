package mse

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

// GatewaySwimmingLaneRoute is a nested struct in mse response
type GatewaySwimmingLaneRoute struct {
	GatewayId                        int64                                `json:"GatewayId" xml:"GatewayId"`
	GatewayUniqueId                  string                               `json:"GatewayUniqueId" xml:"GatewayUniqueId"`
	CanaryModel                      int                                  `json:"CanaryModel" xml:"CanaryModel"`
	Percentage                       int                                  `json:"Percentage" xml:"Percentage"`
	RouteIndependentPercentageEnable string                               `json:"RouteIndependentPercentageEnable" xml:"RouteIndependentPercentageEnable"`
	RouteIdList                      []int64                              `json:"RouteIdList" xml:"RouteIdList"`
	Conditions                       []ConditionsItem                     `json:"Conditions" xml:"Conditions"`
	RouteIndependentPercentageList   []RouteIndependentPercentageListItem `json:"RouteIndependentPercentageList" xml:"RouteIndependentPercentageList"`
}
