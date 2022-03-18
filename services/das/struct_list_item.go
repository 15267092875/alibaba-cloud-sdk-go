package das

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

// ListItem is a nested struct in das response
type ListItem struct {
	SqlCount                    int64  `json:"SqlCount" xml:"SqlCount"`
	SumUpdatedRows              int64  `json:"SumUpdatedRows" xml:"SumUpdatedRows"`
	AvgRows                     int64  `json:"AvgRows" xml:"AvgRows"`
	ErrorCount                  int64  `json:"ErrorCount" xml:"ErrorCount"`
	AvgRowsExamined             int64  `json:"AvgRowsExamined" xml:"AvgRowsExamined"`
	Rows                        int64  `json:"Rows" xml:"Rows"`
	AvgRt                       string `json:"AvgRt" xml:"AvgRt"`
	LockWaitTime                string `json:"LockWaitTime" xml:"LockWaitTime"`
	ExaminedRows                int64  `json:"ExaminedRows" xml:"ExaminedRows"`
	AvgExaminedRows             string `json:"AvgExaminedRows" xml:"AvgExaminedRows"`
	TableList                   string `json:"TableList" xml:"TableList"`
	RtRate                      string `json:"RtRate" xml:"RtRate"`
	OriginHost                  string `json:"OriginHost" xml:"OriginHost"`
	Key                         string `json:"Key" xml:"Key"`
	AvgSqlCount                 int64  `json:"AvgSqlCount" xml:"AvgSqlCount"`
	PhysicalSyncRead            int64  `json:"PhysicalSyncRead" xml:"PhysicalSyncRead"`
	ErrorCnt                    int64  `json:"ErrorCnt" xml:"ErrorCnt"`
	TaskId                      string `json:"TaskId" xml:"TaskId"`
	FirstTime                   int64  `json:"FirstTime" xml:"FirstTime"`
	InstanceId                  string `json:"InstanceId" xml:"InstanceId"`
	UserId                      string `json:"UserId" xml:"UserId"`
	Cnt                         int64  `json:"Cnt" xml:"Cnt"`
	VpcId                       string `json:"VpcId" xml:"VpcId"`
	AvgReturnedRows             string `json:"AvgReturnedRows" xml:"AvgReturnedRows"`
	PhysicalAsyncRead           int64  `json:"PhysicalAsyncRead" xml:"PhysicalAsyncRead"`
	FetchRows                   int64  `json:"FetchRows" xml:"FetchRows"`
	Type                        string `json:"Type" xml:"Type"`
	Version                     int64  `json:"Version" xml:"Version"`
	AvgPhysicalSyncRead         string `json:"AvgPhysicalSyncRead" xml:"AvgPhysicalSyncRead"`
	Ip                          string `json:"Ip" xml:"Ip"`
	SqlId                       string `json:"SqlId" xml:"SqlId"`
	AvgPhysicalAsyncRead        int64  `json:"AvgPhysicalAsyncRead" xml:"AvgPhysicalAsyncRead"`
	Sql                         string `json:"Sql" xml:"Sql"`
	CountRate                   string `json:"CountRate" xml:"CountRate"`
	DbName                      string `json:"DbName" xml:"DbName"`
	AvgLockWaitTime             string `json:"AvgLockWaitTime" xml:"AvgLockWaitTime"`
	LogicalRead                 int64  `json:"LogicalRead" xml:"LogicalRead"`
	RtGreaterThanOneSecondCount int64  `json:"RtGreaterThanOneSecondCount" xml:"RtGreaterThanOneSecondCount"`
	SqlTextFeature              string `json:"SqlTextFeature" xml:"SqlTextFeature"`
	Port                        int64  `json:"Port" xml:"Port"`
	SumRt                       string `json:"SumRt" xml:"SumRt"`
	SqlType                     string `json:"SqlType" xml:"SqlType"`
	AvgUpdatedRows              string `json:"AvgUpdatedRows" xml:"AvgUpdatedRows"`
	Count                       int64  `json:"Count" xml:"Count"`
	AvgFetchRows                int64  `json:"AvgFetchRows" xml:"AvgFetchRows"`
	Database                    string `json:"Database" xml:"Database"`
	CntRate                     string `json:"CntRate" xml:"CntRate"`
	AvgLogicalRead              string `json:"AvgLogicalRead" xml:"AvgLogicalRead"`
}
