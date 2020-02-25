package dbs

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// RenewBackupPlan invokes the dbs.RenewBackupPlan API synchronously
// api document: https://help.aliyun.com/api/dbs/renewbackupplan.html
func (client *Client) RenewBackupPlan(request *RenewBackupPlanRequest) (response *RenewBackupPlanResponse, err error) {
	response = CreateRenewBackupPlanResponse()
	err = client.DoAction(request, response)
	return
}

// RenewBackupPlanWithChan invokes the dbs.RenewBackupPlan API asynchronously
// api document: https://help.aliyun.com/api/dbs/renewbackupplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewBackupPlanWithChan(request *RenewBackupPlanRequest) (<-chan *RenewBackupPlanResponse, <-chan error) {
	responseChan := make(chan *RenewBackupPlanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenewBackupPlan(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// RenewBackupPlanWithCallback invokes the dbs.RenewBackupPlan API asynchronously
// api document: https://help.aliyun.com/api/dbs/renewbackupplan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RenewBackupPlanWithCallback(request *RenewBackupPlanRequest, callback func(response *RenewBackupPlanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenewBackupPlanResponse
		var err error
		defer close(result)
		response, err = client.RenewBackupPlan(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// RenewBackupPlanRequest is the request struct for api RenewBackupPlan
type RenewBackupPlanRequest struct {
	*requests.RpcRequest
	Period       string           `position:"Query" name:"Period"`
	ClientToken  string           `position:"Query" name:"ClientToken"`
	BackupPlanId string           `position:"Query" name:"BackupPlanId"`
	OwnerId      string           `position:"Query" name:"OwnerId"`
	UsedTime     requests.Integer `position:"Query" name:"UsedTime"`
}

// RenewBackupPlanResponse is the response struct for api RenewBackupPlan
type RenewBackupPlanResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
	OrderId        string `json:"OrderId" xml:"OrderId"`
}

// CreateRenewBackupPlanRequest creates a request to invoke RenewBackupPlan API
func CreateRenewBackupPlanRequest() (request *RenewBackupPlanRequest) {
	request = &RenewBackupPlanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "RenewBackupPlan", "", "")
	return
}

// CreateRenewBackupPlanResponse creates a response to parse from RenewBackupPlan response
func CreateRenewBackupPlanResponse() (response *RenewBackupPlanResponse) {
	response = &RenewBackupPlanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
