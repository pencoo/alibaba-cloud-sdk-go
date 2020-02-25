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

// DescribeFullBackupList invokes the dbs.DescribeFullBackupList API synchronously
// api document: https://help.aliyun.com/api/dbs/describefullbackuplist.html
func (client *Client) DescribeFullBackupList(request *DescribeFullBackupListRequest) (response *DescribeFullBackupListResponse, err error) {
	response = CreateDescribeFullBackupListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFullBackupListWithChan invokes the dbs.DescribeFullBackupList API asynchronously
// api document: https://help.aliyun.com/api/dbs/describefullbackuplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFullBackupListWithChan(request *DescribeFullBackupListRequest) (<-chan *DescribeFullBackupListResponse, <-chan error) {
	responseChan := make(chan *DescribeFullBackupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFullBackupList(request)
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

// DescribeFullBackupListWithCallback invokes the dbs.DescribeFullBackupList API asynchronously
// api document: https://help.aliyun.com/api/dbs/describefullbackuplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFullBackupListWithCallback(request *DescribeFullBackupListRequest, callback func(response *DescribeFullBackupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFullBackupListResponse
		var err error
		defer close(result)
		response, err = client.DescribeFullBackupList(request)
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

// DescribeFullBackupListRequest is the request struct for api DescribeFullBackupList
type DescribeFullBackupListRequest struct {
	*requests.RpcRequest
	ClientToken     string           `position:"Query" name:"ClientToken"`
	BackupPlanId    string           `position:"Query" name:"BackupPlanId"`
	PageNum         requests.Integer `position:"Query" name:"PageNum"`
	OwnerId         string           `position:"Query" name:"OwnerId"`
	ShowStorageType requests.Boolean `position:"Query" name:"ShowStorageType"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeFullBackupListResponse is the response struct for api DescribeFullBackupList
type DescribeFullBackupListResponse struct {
	*responses.BaseResponse
	Success        bool                          `json:"Success" xml:"Success"`
	ErrCode        string                        `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string                        `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int                           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	TotalPages     int                           `json:"TotalPages" xml:"TotalPages"`
	PageSize       int                           `json:"PageSize" xml:"PageSize"`
	PageNum        int                           `json:"PageNum" xml:"PageNum"`
	TotalElements  int                           `json:"TotalElements" xml:"TotalElements"`
	Items          ItemsInDescribeFullBackupList `json:"Items" xml:"Items"`
}

// CreateDescribeFullBackupListRequest creates a request to invoke DescribeFullBackupList API
func CreateDescribeFullBackupListRequest() (request *DescribeFullBackupListRequest) {
	request = &DescribeFullBackupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeFullBackupList", "", "")
	return
}

// CreateDescribeFullBackupListResponse creates a response to parse from DescribeFullBackupList response
func CreateDescribeFullBackupListResponse() (response *DescribeFullBackupListResponse) {
	response = &DescribeFullBackupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
