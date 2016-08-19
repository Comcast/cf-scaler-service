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

package cfscalerservice

import (
	"net/http"

	"github.com/codegangsta/negroni"

	"github.com/comcast/cf-scaler-service/handlers"
	"github.com/unrolled/render"
)

//StartNegroni - spins up a negroni instance, ready to be Run
func StartNegroni() (n *negroni.Negroni) {

	r := render.New(render.Options{
		IndentJSON: true,
	})

	router := http.NewServeMux()
	n = negroni.Classic()

	router.HandleFunc("/info", func(res http.ResponseWriter, req *http.Request) {
		handlers.InfoHandler(res, req, r)
	})
	router.HandleFunc("/scale", func(res http.ResponseWriter, req *http.Request) {
		handlers.ScaleHandler(res, req, r)
	})

	n.UseHandler(router)
	return n
}
