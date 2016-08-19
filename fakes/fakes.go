/*
* Copyright 2016 Comcast Cable Communications Management, LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package fakes

import (
	"io"
	"net/http"

	"github.com/unrolled/render"
)

//FakeInfo - Fake info struct
type FakeInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
}

// FakeRenderer - a fake unrolled Renderer
type FakeRenderer struct {
	render.Render
	SpyStatusCode     int
	SpyValue          interface{}
	SpyResponseWriter http.ResponseWriter
}

// Text - fake text Renderer
func (s *FakeRenderer) Text(res http.ResponseWriter, statusCode int, value string) (err error) {
	s.SpyStatusCode = statusCode
	s.SpyValue = value
	s.SpyResponseWriter = res
	return
}

// JSON - fake json renderer
func (s *FakeRenderer) JSON(res http.ResponseWriter, statusCode int, value interface{}) (err error) {
	s.SpyStatusCode = statusCode
	s.SpyValue = value
	s.SpyResponseWriter = res
	return
}

//FakeResponseBody - a fake response body object
type FakeResponseBody struct {
	io.Reader
}

//Close - close fake body
func (FakeResponseBody) Close() error { return nil }

//FakeRequestBody - a fake response body object
type FakeRequestBody struct {
	io.Reader
}

//Close - close fake body
func (FakeRequestBody) Close() error { return nil }

//FakeHTTPClient - a fake http client
type FakeHTTPClient struct {
	http.Client
	SpyRequest   *http.Request
	FakeResponse *http.Response
	FakeError    error
}

// Do - Fake HTTP client do method
func (s *FakeHTTPClient) Do(fakeRequest *http.Request) (*http.Response, error) {
	s.SpyRequest = fakeRequest
	return s.FakeResponse, s.FakeError
}
