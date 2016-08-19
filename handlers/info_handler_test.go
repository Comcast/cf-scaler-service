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
	"fmt"
	"net/http"

	"github.com/comcast/cf-scaler-service/fakes"
	. "github.com/comcast/cf-scaler-service/handlers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InfoHandler", func() {
	Describe("given an InfoHandler method", func() {
		Context("when called with a valid request", func() {
			var (
				fakeRequest        *http.Request
				fakeRenderer       *fakes.FakeRenderer
				fakeResponseWriter http.ResponseWriter
				ctrlResponseBody   = `{"name":"CFScalerService","version":"","arch":""}`
			)

			BeforeEach(func() {
				fakeRequest, _ = http.NewRequest("GET", "http://google.com", nil)
				fakeRenderer = new(fakes.FakeRenderer)
				InfoHandler(fakeResponseWriter, fakeRequest, fakeRenderer)
			})
			It("should not return an error", func() {
				Ω(fakeRenderer.SpyStatusCode).Should(Equal(http.StatusOK))
			})
			XIt("should return the correct JSON response", func() {
				Ω(ctrlResponseBody).Should(Equal(fmt.Sprintf("%v", fakeRenderer.SpyValue)))
			})
		})
	})
})
