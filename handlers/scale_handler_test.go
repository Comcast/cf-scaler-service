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

package handlers_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/comcast/cf-scaler-service/fakes"
	. "github.com/comcast/cf-scaler-service/handlers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ScaleHandler", func() {
	Describe("given a ScaleHandler method", func() {
		XContext("when called with a valid scale up request", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fixtureRequest, _ := ioutil.ReadFile("./fixtures/sample_good.request.json")
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBuffer(fixtureRequest)})
				fakeRequest.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})

			It("should not return an error", func() {
				//Ω(err).ShouldNot(HaveOccurred())
			})
			It("should return the correct success response", func() {

			})
			It("should return the correct success status code", func() {
				Ω(fakeRenderer.SpyStatusCode).Should(Equal(http.StatusOK))
			})
		})

		Context("when called with an empty request", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBufferString(``)})
				fakeRequest.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})
			It("should return the correct bad request response", func() {
				Ω(fmt.Sprintf("%v", fakeRenderer.SpyValue)).Should(Equal("{Bad Request}"))
			})
			It("should return the correct bad request status code", func() {
				Ω(fakeRenderer.SpyStatusCode).Should(Equal(http.StatusBadRequest))
			})
		})

		Context("when called with a scale up request without the auth header", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fixtureRequest, _ := ioutil.ReadFile("./fixtures/sample_good_request.json")
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBuffer(fixtureRequest)})
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})

			It("should return the correct invalid response", func() {
				Ω(fmt.Sprintf("%v", fakeRenderer.SpyValue)).Should(Equal("{Unauthorized}"))
			})
			It("should return the correct unauthorized status code", func() {
				Ω(fakeRenderer.SpyStatusCode).Should(Equal(http.StatusUnauthorized))
			})
		})
		Context("when called with a scale up request with bad credentials", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fixtureRequest, _ := ioutil.ReadFile("./fixtures/sample_request.xml")
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBuffer(fixtureRequest)})
				fakeRequest.Header.Add("Authorization", "Basic c3OlUzZGk1b2FDTUsyZw==")
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})

			It("should return the correct invalid response", func() {

			})
			It("should return the correct unauthorized status code", func() {
				Ω(fakeRenderer.SpyStatusCode).Should(Equal(http.StatusUnauthorized))
			})
		})
		XContext("when called with a scale up request with invalid org", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fixtureRequest, _ := ioutil.ReadFile("../fixtures/sample_badorg_request.json")
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBuffer(fixtureRequest)})
				fakeRequest.Header.Add("Authorization", "Basic YWRtaW46YWRtaW4=")
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)

			})
			It("should return the correct invalid response", func() {

			})
			It("should return the correct bad request status code", func() {
				Ω(fakeRenderer.SpyStatusCode).Should(Equal(http.StatusInternalServerError))
			})
		})
		XContext("when called with a scale up request with invalid space", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fixtureRequest, _ := ioutil.ReadFile("./fixtures/sample_request.xml")
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBuffer(fixtureRequest)})
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})
			It("should return the correct invalid response", func() {

			})
			It("should return the correct bad request status code", func() {

			})
		})
		XContext("when called with a scale up request with invalid app", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
			)

			BeforeEach(func() {
				fakeRenderer = new(fakes.FakeRenderer)
				fixtureRequest, _ := ioutil.ReadFile("./fixtures/sample_request.xml")
				fakeRequest, _ = http.NewRequest("POST", "http://google.com", fakes.FakeResponseBody{bytes.NewBuffer(fixtureRequest)})
				ScaleHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})

			It("should return an error", func() {

			})
			It("should return the correct invalid response", func() {

			})
			It("should return the correct bad request status code", func() {

			})
		})
	})
})
