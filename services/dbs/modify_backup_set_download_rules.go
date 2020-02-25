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

// ModifyBackupSetDownloadRules invokes the dbs.ModifyBackupSetDownloadRules API synchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupsetdownloadrules.html
func (client *Client) ModifyBackupSetDownloadRules(request *ModifyBackupSetDownloadRulesRequest) (response *ModifyBackupSetDownloadRulesResponse, err error) {
	response = CreateModifyBackupSetDownloadRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackupSetDownloadRulesWithChan invokes the dbs.ModifyBackupSetDownloadRules API asynchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupsetdownloadrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupSetDownloadRulesWithChan(request *ModifyBackupSetDownloadRulesRequest) (<-chan *ModifyBackupSetDownloadRulesResponse, <-chan error) {
	responseChan := make(chan *ModifyBackupSetDownloadRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackupSetDownloadRules(request)
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

// ModifyBackupSetDownloadRulesWithCallback invokes the dbs.ModifyBackupSetDownloadRules API asynchronously
// api document: https://help.aliyun.com/api/dbs/modifybackupsetdownloadrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBackupSetDownloadRulesWithCallback(request *ModifyBackupSetDownloadRulesRequest, callback func(response *ModifyBackupSetDownloadRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackupSetDownloadRulesResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackupSetDownloadRules(request)
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

// ModifyBackupSetDownloadRulesRequest is the request struct for api ModifyBackupSetDownloadRules
type ModifyBackupSetDownloadRulesRequest struct {
	*requests.RpcRequest
	FullDataFormat              string           `position:"Query" name:"FullDataFormat"`
	BackupGatewayId             requests.Integer `position:"Query" name:"BackupGatewayId"`
	ClientToken                 string           `position:"Query" name:"ClientToken"`
	BackupSetDownloadTargetType string           `position:"Query" name:"BackupSetDownloadTargetType"`
	BackupPlanId                string           `position:"Query" name:"BackupPlanId"`
	OwnerId                     string           `position:"Query" name:"OwnerId"`
	OpenAutoDownload            requests.Boolean `position:"Query" name:"OpenAutoDownload"`
	IncrementDataFormat         string           `position:"Query" name:"IncrementDataFormat"`
	BackupSetDownloadDir        string           `position:"Query" name:"BackupSetDownloadDir"`
}

// ModifyBackupSetDownloadRulesResponse is the response struct for api ModifyBackupSetDownloadRules
type ModifyBackupSetDownloadRulesResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	BackupPlanId   string `json:"BackupPlanId" xml:"BackupPlanId"`
}

// CreateModifyBackupSetDownloadRulesRequest creates a request to invoke ModifyBackupSetDownloadRules API
func CreateModifyBackupSetDownloadRulesRequest() (request *ModifyBackupSetDownloadRulesRequest) {
	request = &ModifyBackupSetDownloadRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "ModifyBackupSetDownloadRules", "", "")
	return
}

// CreateModifyBackupSetDownloadRulesResponse creates a response to parse from ModifyBackupSetDownloadRules response
func CreateModifyBackupSetDownloadRulesResponse() (response *ModifyBackupSetDownloadRulesResponse) {
	response = &ModifyBackupSetDownloadRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
