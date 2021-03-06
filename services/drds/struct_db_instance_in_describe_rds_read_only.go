package drds

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

// DbInstanceInDescribeRdsReadOnly is a nested struct in drds response
type DbInstanceInDescribeRdsReadOnly struct {
	DBInstanceId        string `json:"DBInstanceId" xml:"DBInstanceId"`
	DmInstanceId        string `json:"DmInstanceId" xml:"DmInstanceId"`
	ConnectUrl          string `json:"ConnectUrl" xml:"ConnectUrl"`
	Port                int    `json:"Port" xml:"Port"`
	DBInstanceStatus    string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	DbInstType          string `json:"DbInstType" xml:"DbInstType"`
	ReadWeight          int    `json:"ReadWeight" xml:"ReadWeight"`
	NetworkType         string `json:"NetworkType" xml:"NetworkType"`
	Engine              string `json:"Engine" xml:"Engine"`
	EngineVersion       string `json:"EngineVersion" xml:"EngineVersion"`
	RdsInstType         string `json:"RdsInstType" xml:"RdsInstType"`
	PayType             string `json:"PayType" xml:"PayType"`
	ExpireTime          string `json:"ExpireTime" xml:"ExpireTime"`
	RemainDays          int    `json:"RemainDays" xml:"RemainDays"`
	DBInstanceClassType string `json:"DBInstanceClassType" xml:"DBInstanceClassType"`
	DBInstanceCPU       string `json:"DBInstanceCPU" xml:"DBInstanceCPU"`
	DBInstanceMemory    int64  `json:"DBInstanceMemory" xml:"DBInstanceMemory"`
	DBInstanceStorage   int64  `json:"DBInstanceStorage" xml:"DBInstanceStorage"`
}
