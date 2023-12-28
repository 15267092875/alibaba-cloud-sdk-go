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

// Data is a nested struct in oceanbasepro response
type Data struct {
	Event                   string                        `json:"Event" xml:"Event"`
	IOWaitTime              float64                       `json:"IOWaitTime" xml:"IOWaitTime"`
	Executions              int64                         `json:"Executions" xml:"Executions"`
	ConcurrencyWaitTime     float64                       `json:"ConcurrencyWaitTime" xml:"ConcurrencyWaitTime"`
	ExecuteTime             float64                       `json:"ExecuteTime" xml:"ExecuteTime"`
	CpuTime                 float64                       `json:"CpuTime" xml:"CpuTime"`
	DbType                  string                        `json:"DbType" xml:"DbType"`
	Encoding                string                        `json:"Encoding" xml:"Encoding"`
	ProjectName             string                        `json:"ProjectName" xml:"ProjectName"`
	ProjectId               string                        `json:"ProjectId" xml:"ProjectId"`
	ElapsedTime             float64                       `json:"ElapsedTime" xml:"ElapsedTime"`
	OutlineId               int64                         `json:"OutlineId" xml:"OutlineId"`
	DeployMode              string                        `json:"DeployMode" xml:"DeployMode"`
	NodeIp                  string                        `json:"NodeIp" xml:"NodeIp"`
	OutlineData             string                        `json:"OutlineData" xml:"OutlineData"`
	Charset                 string                        `json:"Charset" xml:"Charset"`
	Message                 string                        `json:"Message" xml:"Message"`
	QuerySQL                string                        `json:"QuerySQL" xml:"QuerySQL"`
	ProjectOwner            string                        `json:"ProjectOwner" xml:"ProjectOwner"`
	RPCCount                int64                         `json:"RPCCount" xml:"RPCCount"`
	ScheduleTime            float64                       `json:"ScheduleTime" xml:"ScheduleTime"`
	EndpointId              string                        `json:"EndpointId" xml:"EndpointId"`
	DiagnosisRule           string                        `json:"DiagnosisRule" xml:"DiagnosisRule"`
	DefaultValue            string                        `json:"DefaultValue" xml:"DefaultValue"`
	Mem                     int                           `json:"Mem" xml:"Mem"`
	DbName                  string                        `json:"DbName" xml:"DbName"`
	DiskRead                int64                         `json:"DiskRead" xml:"DiskRead"`
	ZoneId                  string                        `json:"ZoneId" xml:"ZoneId"`
	Readonly                bool                          `json:"Readonly" xml:"Readonly"`
	MinMem                  int64                         `json:"MinMem" xml:"MinMem"`
	MaxMem                  int64                         `json:"MaxMem" xml:"MaxMem"`
	Status                  string                        `json:"Status" xml:"Status"`
	Collation               string                        `json:"Collation" xml:"Collation"`
	BlockIndexCacheHit      int64                         `json:"BlockIndexCacheHit" xml:"BlockIndexCacheHit"`
	Name                    string                        `json:"Name" xml:"Name"`
	TenantMode              string                        `json:"TenantMode" xml:"TenantMode"`
	TenantName              string                        `json:"TenantName" xml:"TenantName"`
	ValueType               string                        `json:"ValueType" xml:"ValueType"`
	CreateTime              string                        `json:"CreateTime" xml:"CreateTime"`
	ReturnRows              int64                         `json:"ReturnRows" xml:"ReturnRows"`
	HitCount                int                           `json:"HitCount" xml:"HitCount"`
	PlanId                  int                           `json:"PlanId" xml:"PlanId"`
	UnitCpu                 int                           `json:"UnitCpu" xml:"UnitCpu"`
	FirstLoadTimeUTCString  string                        `json:"FirstLoadTimeUTCString" xml:"FirstLoadTimeUTCString"`
	InstanceId              string                        `json:"InstanceId" xml:"InstanceId"`
	OutlineTimeUTCString    string                        `json:"OutlineTimeUTCString" xml:"OutlineTimeUTCString"`
	RemotePlans             int64                         `json:"RemotePlans" xml:"RemotePlans"`
	ZoneName                string                        `json:"ZoneName" xml:"ZoneName"`
	AppWaitTime             float64                       `json:"AppWaitTime" xml:"AppWaitTime"`
	RequestTime             float64                       `json:"RequestTime" xml:"RequestTime"`
	TenantId                string                        `json:"TenantId" xml:"TenantId"`
	AvgExecutionTimeMS      int64                         `json:"AvgExecutionTimeMS" xml:"AvgExecutionTimeMS"`
	MemstoreReadRowCount    int64                         `json:"MemstoreReadRowCount" xml:"MemstoreReadRowCount"`
	SQLType                 int64                         `json:"SQLType" xml:"SQLType"`
	BloomFilterCacheHit     int64                         `json:"BloomFilterCacheHit" xml:"BloomFilterCacheHit"`
	MissPlans               int64                         `json:"MissPlans" xml:"MissPlans"`
	NeedReboot              bool                          `json:"NeedReboot" xml:"NeedReboot"`
	RequiredSize            string                        `json:"RequiredSize" xml:"RequiredSize"`
	PlanUnionHash           string                        `json:"PlanUnionHash" xml:"PlanUnionHash"`
	MaxElapsedTime          float64                       `json:"MaxElapsedTime" xml:"MaxElapsedTime"`
	SQLId                   string                        `json:"SQLId" xml:"SQLId"`
	OrderId                 string                        `json:"OrderId" xml:"OrderId"`
	GetPlanTime             float64                       `json:"GetPlanTime" xml:"GetPlanTime"`
	TotalCount              int                           `json:"TotalCount" xml:"TotalCount"`
	CurrentValue            string                        `json:"CurrentValue" xml:"CurrentValue"`
	MaxCpuTime              float64                       `json:"MaxCpuTime" xml:"MaxCpuTime"`
	VpcId                   string                        `json:"VpcId" xml:"VpcId"`
	DatabaseName            string                        `json:"DatabaseName" xml:"DatabaseName"`
	OutlineTime             int64                         `json:"OutlineTime" xml:"OutlineTime"`
	FailTimes               int64                         `json:"FailTimes" xml:"FailTimes"`
	PlanFull                string                        `json:"PlanFull" xml:"PlanFull"`
	SafeMem                 string                        `json:"SafeMem" xml:"SafeMem"`
	LogicalRead             int64                         `json:"LogicalRead" xml:"LogicalRead"`
	SQLText                 string                        `json:"SQLText" xml:"SQLText"`
	Success                 bool                          `json:"Success" xml:"Success"`
	Suggestion              string                        `json:"Suggestion" xml:"Suggestion"`
	DataSize                string                        `json:"DataSize" xml:"DataSize"`
	AffectedRows            int64                         `json:"AffectedRows" xml:"AffectedRows"`
	DryRunResult            bool                          `json:"DryRunResult" xml:"DryRunResult"`
	BusinessName            string                        `json:"BusinessName" xml:"BusinessName"`
	Cpu                     int                           `json:"Cpu" xml:"Cpu"`
	SsstoreReadRowCount     int64                         `json:"SsstoreReadRowCount" xml:"SsstoreReadRowCount"`
	RetryCount              int64                         `json:"RetryCount" xml:"RetryCount"`
	ClientIp                string                        `json:"ClientIp" xml:"ClientIp"`
	UnitNum                 int                           `json:"UnitNum" xml:"UnitNum"`
	Unit                    string                        `json:"Unit" xml:"Unit"`
	Diagnosis               string                        `json:"Diagnosis" xml:"Diagnosis"`
	RowCacheHit             int64                         `json:"RowCacheHit" xml:"RowCacheHit"`
	DeployType              string                        `json:"DeployType" xml:"DeployType"`
	NetWaitTime             float64                       `json:"NetWaitTime" xml:"NetWaitTime"`
	UnitMem                 int                           `json:"UnitMem" xml:"UnitMem"`
	FirstLoadTime           int64                         `json:"FirstLoadTime" xml:"FirstLoadTime"`
	ResourceGroupId         string                        `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ExecPerSecond           float64                       `json:"ExecPerSecond" xml:"ExecPerSecond"`
	AvgExecutionMS          float64                       `json:"AvgExecutionMS" xml:"AvgExecutionMS"`
	QueueTime               float64                       `json:"QueueTime" xml:"QueueTime"`
	PageNumber              int                           `json:"PageNumber" xml:"PageNumber"`
	TotalWaitTime           float64                       `json:"TotalWaitTime" xml:"TotalWaitTime"`
	PrimaryZone             string                        `json:"PrimaryZone" xml:"PrimaryZone"`
	Description             string                        `json:"Description" xml:"Description"`
	BlockCacheHit           int64                         `json:"BlockCacheHit" xml:"BlockCacheHit"`
	PlanInfo                string                        `json:"PlanInfo" xml:"PlanInfo"`
	RequestTimeUTCString    string                        `json:"RequestTimeUTCString" xml:"RequestTimeUTCString"`
	Key                     int64                         `json:"Key" xml:"Key"`
	DecodeTime              float64                       `json:"DecodeTime" xml:"DecodeTime"`
	Series                  string                        `json:"Series" xml:"Series"`
	UserName                string                        `json:"UserName" xml:"UserName"`
	MergedVersion           int                           `json:"MergedVersion" xml:"MergedVersion"`
	UsedMem                 int64                         `json:"UsedMem" xml:"UsedMem"`
	UsedDiskSize            string                        `json:"UsedDiskSize" xml:"UsedDiskSize"`
	RejectedValue           []string                      `json:"RejectedValue" xml:"RejectedValue"`
	AcceptableValue         []string                      `json:"AcceptableValue" xml:"AcceptableValue"`
	DestConfig              DestConfig                    `json:"DestConfig" xml:"DestConfig"`
	SourceConfig            SourceConfig                  `json:"SourceConfig" xml:"SourceConfig"`
	TransferMapping         TransferMapping               `json:"TransferMapping" xml:"TransferMapping"`
	TransferStepConfig      TransferStepConfig            `json:"TransferStepConfig" xml:"TransferStepConfig"`
	Tables                  []TablesItem                  `json:"Tables" xml:"Tables"`
	AvailableZones          []AvailableZonesItem          `json:"AvailableZones" xml:"AvailableZones"`
	Labels                  []Label                       `json:"Labels" xml:"Labels"`
	Parameters              []ParametersItem              `json:"Parameters" xml:"Parameters"`
	Steps                   []Step                        `json:"Steps" xml:"Steps"`
	AvailableSpecifications []AvailableSpecificationsItem `json:"AvailableSpecifications" xml:"AvailableSpecifications"`
	Users                   []UsersItem                   `json:"Users" xml:"Users"`
}
