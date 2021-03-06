package smartag

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

// GrantRule is a nested struct in smartag response
type GrantRule struct {
	CcnUid        int64  `json:"CcnUid" xml:"CcnUid"`
	GrantRuleId   string `json:"GrantRuleId" xml:"GrantRuleId"`
	VbrRegionId   string `json:"VbrRegionId" xml:"VbrRegionId"`
	CcnInstanceId string `json:"CcnInstanceId" xml:"CcnInstanceId"`
	CreateTime    int64  `json:"CreateTime" xml:"CreateTime"`
	Bound         bool   `json:"Bound" xml:"Bound"`
	SmartAGUid    int64  `json:"SmartAGUid" xml:"SmartAGUid"`
	GmtCreate     int64  `json:"GmtCreate" xml:"GmtCreate"`
	VbrUid        int64  `json:"VbrUid" xml:"VbrUid"`
	RegionId      string `json:"RegionId" xml:"RegionId"`
	CenInstanceId string `json:"CenInstanceId" xml:"CenInstanceId"`
	InstanceId    string `json:"InstanceId" xml:"InstanceId"`
	GmtModified   int64  `json:"GmtModified" xml:"GmtModified"`
	CenUid        int64  `json:"CenUid" xml:"CenUid"`
	VbrInstanceId string `json:"VbrInstanceId" xml:"VbrInstanceId"`
	SmartAGId     string `json:"SmartAGId" xml:"SmartAGId"`
}
