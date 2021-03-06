package ga

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

// EndpointConfigurationsItem is a nested struct in ga response
type EndpointConfigurationsItem struct {
	Endpoint                   string `json:"Endpoint" xml:"Endpoint"`
	EnableProxyProtocol        bool   `json:"EnableProxyProtocol" xml:"EnableProxyProtocol"`
	ProbeProtocol              string `json:"ProbeProtocol" xml:"ProbeProtocol"`
	EnableClientIPPreservation bool   `json:"EnableClientIPPreservation" xml:"EnableClientIPPreservation"`
	Weight                     int    `json:"Weight" xml:"Weight"`
	ProbePort                  int    `json:"ProbePort" xml:"ProbePort"`
	Type                       string `json:"Type" xml:"Type"`
}
