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

package main

import (
	"fmt"
	"os"

	"github.com/comcast/cf-scaler-service"
)

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

// GetInfo - Method to return info struct.
func GetInfo() Info {
	info := Info{
		Name:    "CFScalarService",
		Version: Versionflag,
		Arch:    Archflag,
	}
	return info
}

func main() {
	cfscalerservice.StartNegroni().Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
