package waf_openapi

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

// ModifyDomainClusterType invokes the waf_openapi.ModifyDomainClusterType API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifydomainclustertype.html
func (client *Client) ModifyDomainClusterType(request *ModifyDomainClusterTypeRequest) (response *ModifyDomainClusterTypeResponse, err error) {
	response = CreateModifyDomainClusterTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDomainClusterTypeWithChan invokes the waf_openapi.ModifyDomainClusterType API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifydomainclustertype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDomainClusterTypeWithChan(request *ModifyDomainClusterTypeRequest) (<-chan *ModifyDomainClusterTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDomainClusterTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDomainClusterType(request)
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

// ModifyDomainClusterTypeWithCallback invokes the waf_openapi.ModifyDomainClusterType API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/modifydomainclustertype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDomainClusterTypeWithCallback(request *ModifyDomainClusterTypeRequest, callback func(response *ModifyDomainClusterTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDomainClusterTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDomainClusterType(request)
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

// ModifyDomainClusterTypeRequest is the request struct for api ModifyDomainClusterType
type ModifyDomainClusterTypeRequest struct {
	*requests.RpcRequest
	ClusterType requests.Integer `position:"Query" name:"ClusterType"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Domain      string           `position:"Query" name:"Domain"`
}

// ModifyDomainClusterTypeResponse is the response struct for api ModifyDomainClusterType
type ModifyDomainClusterTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDomainClusterTypeRequest creates a request to invoke ModifyDomainClusterType API
func CreateModifyDomainClusterTypeRequest() (request *ModifyDomainClusterTypeRequest) {
	request = &ModifyDomainClusterTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "ModifyDomainClusterType", "waf", "openAPI")
	return
}

// CreateModifyDomainClusterTypeResponse creates a response to parse from ModifyDomainClusterType response
func CreateModifyDomainClusterTypeResponse() (response *ModifyDomainClusterTypeResponse) {
	response = &ModifyDomainClusterTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
