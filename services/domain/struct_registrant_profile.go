package domain

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

// RegistrantProfile is a nested struct in domain response
type RegistrantProfile struct {
	TelExt                   string `json:"TelExt" xml:"TelExt"`
	UpdateTime               string `json:"UpdateTime" xml:"UpdateTime"`
	ZhProvince               string `json:"ZhProvince" xml:"ZhProvince"`
	CreateTime               string `json:"CreateTime" xml:"CreateTime"`
	Telephone                string `json:"Telephone" xml:"Telephone"`
	RegistrantOrganization   string `json:"RegistrantOrganization" xml:"RegistrantOrganization"`
	City                     string `json:"City" xml:"City"`
	ZhCity                   string `json:"ZhCity" xml:"ZhCity"`
	TelArea                  string `json:"TelArea" xml:"TelArea"`
	Address                  string `json:"Address" xml:"Address"`
	RealNameStatus           string `json:"RealNameStatus" xml:"RealNameStatus"`
	PostalCode               string `json:"PostalCode" xml:"PostalCode"`
	RegistrantProfileType    string `json:"RegistrantProfileType" xml:"RegistrantProfileType"`
	RegistrantProfileId      int64  `json:"RegistrantProfileId" xml:"RegistrantProfileId"`
	ZhRegistrantOrganization string `json:"ZhRegistrantOrganization" xml:"ZhRegistrantOrganization"`
	DefaultRegistrantProfile bool   `json:"DefaultRegistrantProfile" xml:"DefaultRegistrantProfile"`
	Email                    string `json:"Email" xml:"Email"`
	ZhRegistrantName         string `json:"ZhRegistrantName" xml:"ZhRegistrantName"`
	RegistrantType           string `json:"RegistrantType" xml:"RegistrantType"`
	Country                  string `json:"Country" xml:"Country"`
	RegistrantName           string `json:"RegistrantName" xml:"RegistrantName"`
	EmailVerificationStatus  int    `json:"EmailVerificationStatus" xml:"EmailVerificationStatus"`
	ZhAddress                string `json:"ZhAddress" xml:"ZhAddress"`
	Province                 string `json:"Province" xml:"Province"`
	CredentialNo             string `json:"CredentialNo" xml:"CredentialNo"`
	CredentialType           string `json:"CredentialType" xml:"CredentialType"`
	Remark                   string `json:"Remark" xml:"Remark"`
}
