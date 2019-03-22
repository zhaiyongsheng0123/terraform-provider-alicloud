package dds

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

// DescribeReplicaUsage invokes the dds.DescribeReplicaUsage API synchronously
// api document: https://help.aliyun.com/api/dds/describereplicausage.html
func (client *Client) DescribeReplicaUsage(request *DescribeReplicaUsageRequest) (response *DescribeReplicaUsageResponse, err error) {
	response = CreateDescribeReplicaUsageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReplicaUsageWithChan invokes the dds.DescribeReplicaUsage API asynchronously
// api document: https://help.aliyun.com/api/dds/describereplicausage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReplicaUsageWithChan(request *DescribeReplicaUsageRequest) (<-chan *DescribeReplicaUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeReplicaUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReplicaUsage(request)
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

// DescribeReplicaUsageWithCallback invokes the dds.DescribeReplicaUsage API asynchronously
// api document: https://help.aliyun.com/api/dds/describereplicausage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReplicaUsageWithCallback(request *DescribeReplicaUsageRequest, callback func(response *DescribeReplicaUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReplicaUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeReplicaUsage(request)
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

// DescribeReplicaUsageRequest is the request struct for api DescribeReplicaUsage
type DescribeReplicaUsageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SourceDBInstanceId      string           `position:"Query" name:"SourceDBInstanceId"`
	DestinationDBInstanceId string           `position:"Query" name:"DestinationDBInstanceId"`
	SecurityToken           string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	ReplicaId               string           `position:"Query" name:"ReplicaId"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeReplicaUsageResponse is the response struct for api DescribeReplicaUsage
type DescribeReplicaUsageResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	StartTime       string          `json:"StartTime" xml:"StartTime"`
	EndTime         string          `json:"EndTime" xml:"EndTime"`
	ReplicaId       string          `json:"ReplicaId" xml:"ReplicaId"`
	PerformanceKeys PerformanceKeys `json:"PerformanceKeys" xml:"PerformanceKeys"`
}

// CreateDescribeReplicaUsageRequest creates a request to invoke DescribeReplicaUsage API
func CreateDescribeReplicaUsageRequest() (request *DescribeReplicaUsageRequest) {
	request = &DescribeReplicaUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "DescribeReplicaUsage", "dds", "openAPI")
	return
}

// CreateDescribeReplicaUsageResponse creates a response to parse from DescribeReplicaUsage response
func CreateDescribeReplicaUsageResponse() (response *DescribeReplicaUsageResponse) {
	response = &DescribeReplicaUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
