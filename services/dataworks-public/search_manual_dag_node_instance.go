package dataworks_public

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

// SearchManualDagNodeInstance invokes the dataworks_public.SearchManualDagNodeInstance API synchronously
func (client *Client) SearchManualDagNodeInstance(request *SearchManualDagNodeInstanceRequest) (response *SearchManualDagNodeInstanceResponse, err error) {
	response = CreateSearchManualDagNodeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// SearchManualDagNodeInstanceWithChan invokes the dataworks_public.SearchManualDagNodeInstance API asynchronously
func (client *Client) SearchManualDagNodeInstanceWithChan(request *SearchManualDagNodeInstanceRequest) (<-chan *SearchManualDagNodeInstanceResponse, <-chan error) {
	responseChan := make(chan *SearchManualDagNodeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchManualDagNodeInstance(request)
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

// SearchManualDagNodeInstanceWithCallback invokes the dataworks_public.SearchManualDagNodeInstance API asynchronously
func (client *Client) SearchManualDagNodeInstanceWithCallback(request *SearchManualDagNodeInstanceRequest, callback func(response *SearchManualDagNodeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchManualDagNodeInstanceResponse
		var err error
		defer close(result)
		response, err = client.SearchManualDagNodeInstance(request)
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

// SearchManualDagNodeInstanceRequest is the request struct for api SearchManualDagNodeInstance
type SearchManualDagNodeInstanceRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Query" name:"ProjectName"`
	DagId       requests.Integer `position:"Query" name:"DagId"`
}

// SearchManualDagNodeInstanceResponse is the response struct for api SearchManualDagNodeInstance
type SearchManualDagNodeInstanceResponse struct {
	*responses.BaseResponse
	RequestId string                            `json:"RequestId" xml:"RequestId"`
	ErrCode   string                            `json:"ErrCode" xml:"ErrCode"`
	ErrMsg    string                            `json:"ErrMsg" xml:"ErrMsg"`
	Success   bool                              `json:"Success" xml:"Success"`
	Data      DataInSearchManualDagNodeInstance `json:"Data" xml:"Data"`
}

// CreateSearchManualDagNodeInstanceRequest creates a request to invoke SearchManualDagNodeInstance API
func CreateSearchManualDagNodeInstanceRequest() (request *SearchManualDagNodeInstanceRequest) {
	request = &SearchManualDagNodeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2018-06-01", "SearchManualDagNodeInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchManualDagNodeInstanceResponse creates a response to parse from SearchManualDagNodeInstance response
func CreateSearchManualDagNodeInstanceResponse() (response *SearchManualDagNodeInstanceResponse) {
	response = &SearchManualDagNodeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}