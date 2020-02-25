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

// DescribePreCheckProgressList invokes the dbs.DescribePreCheckProgressList API synchronously
// api document: https://help.aliyun.com/api/dbs/describeprecheckprogresslist.html
func (client *Client) DescribePreCheckProgressList(request *DescribePreCheckProgressListRequest) (response *DescribePreCheckProgressListResponse, err error) {
	response = CreateDescribePreCheckProgressListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribePreCheckProgressListWithChan invokes the dbs.DescribePreCheckProgressList API asynchronously
// api document: https://help.aliyun.com/api/dbs/describeprecheckprogresslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePreCheckProgressListWithChan(request *DescribePreCheckProgressListRequest) (<-chan *DescribePreCheckProgressListResponse, <-chan error) {
	responseChan := make(chan *DescribePreCheckProgressListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribePreCheckProgressList(request)
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

// DescribePreCheckProgressListWithCallback invokes the dbs.DescribePreCheckProgressList API asynchronously
// api document: https://help.aliyun.com/api/dbs/describeprecheckprogresslist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribePreCheckProgressListWithCallback(request *DescribePreCheckProgressListRequest, callback func(response *DescribePreCheckProgressListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribePreCheckProgressListResponse
		var err error
		defer close(result)
		response, err = client.DescribePreCheckProgressList(request)
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

// DescribePreCheckProgressListRequest is the request struct for api DescribePreCheckProgressList
type DescribePreCheckProgressListRequest struct {
	*requests.RpcRequest
	ClientToken   string `position:"Query" name:"ClientToken"`
	BackupPlanId  string `position:"Query" name:"BackupPlanId"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	RestoreTaskId string `position:"Query" name:"RestoreTaskId"`
}

// DescribePreCheckProgressListResponse is the response struct for api DescribePreCheckProgressList
type DescribePreCheckProgressListResponse struct {
	*responses.BaseResponse
	Status         string                              `json:"Status" xml:"Status"`
	Progress       int                                 `json:"Progress" xml:"Progress"`
	Success        bool                                `json:"Success" xml:"Success"`
	ErrCode        string                              `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string                              `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int                                 `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                              `json:"RequestId" xml:"RequestId"`
	Items          ItemsInDescribePreCheckProgressList `json:"Items" xml:"Items"`
}

// CreateDescribePreCheckProgressListRequest creates a request to invoke DescribePreCheckProgressList API
func CreateDescribePreCheckProgressListRequest() (request *DescribePreCheckProgressListRequest) {
	request = &DescribePreCheckProgressListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribePreCheckProgressList", "", "")
	return
}

// CreateDescribePreCheckProgressListResponse creates a response to parse from DescribePreCheckProgressList response
func CreateDescribePreCheckProgressListResponse() (response *DescribePreCheckProgressListResponse) {
	response = &DescribePreCheckProgressListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
