package voicenavigator

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

// DebugAuthorized invokes the voicenavigator.DebugAuthorized API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/debugauthorized.html
func (client *Client) DebugAuthorized(request *DebugAuthorizedRequest) (response *DebugAuthorizedResponse, err error) {
	response = CreateDebugAuthorizedResponse()
	err = client.DoAction(request, response)
	return
}

// DebugAuthorizedWithChan invokes the voicenavigator.DebugAuthorized API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/debugauthorized.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugAuthorizedWithChan(request *DebugAuthorizedRequest) (<-chan *DebugAuthorizedResponse, <-chan error) {
	responseChan := make(chan *DebugAuthorizedResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DebugAuthorized(request)
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

// DebugAuthorizedWithCallback invokes the voicenavigator.DebugAuthorized API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/debugauthorized.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DebugAuthorizedWithCallback(request *DebugAuthorizedRequest, callback func(response *DebugAuthorizedResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DebugAuthorizedResponse
		var err error
		defer close(result)
		response, err = client.DebugAuthorized(request)
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

// DebugAuthorizedRequest is the request struct for api DebugAuthorized
type DebugAuthorizedRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// DebugAuthorizedResponse is the response struct for api DebugAuthorized
type DebugAuthorizedResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateDebugAuthorizedRequest creates a request to invoke DebugAuthorized API
func CreateDebugAuthorizedRequest() (request *DebugAuthorizedRequest) {
	request = &DebugAuthorizedRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "DebugAuthorized", "voicebot", "openAPI")
	return
}

// CreateDebugAuthorizedResponse creates a response to parse from DebugAuthorized response
func CreateDebugAuthorizedResponse() (response *DebugAuthorizedResponse) {
	response = &DebugAuthorizedResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
