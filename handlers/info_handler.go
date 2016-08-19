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

package handlers

import "net/http"

// Var - Two variables set by ldflags at build time.
var (
	Archflag    string
	Versionflag string
)

// Info - Struct to contain application information.
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
}

//Renderer -- interface to define the required render object
type Renderer interface {
	JSON(w http.ResponseWriter, status int, v interface{}) error
}

// GetInfo - Method to return info struct.
func getInfo() Info {
	info := Info{
		Name:    "CFScalerService",
		Version: Versionflag,
		Arch:    Archflag,
	}
	return info
}

// InfoHandler - Returns info.
func InfoHandler(res http.ResponseWriter, req *http.Request, r Renderer) {
	r.JSON(res, http.StatusOK, getInfo())
}
