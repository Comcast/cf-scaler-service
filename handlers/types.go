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

//Application ---
type Application struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

//Applications ---
type Applications struct {
	Applications []Application
}

//Space ---
type Space struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

//Organization ---
type Organization struct {
	GUID string
	Name string
}

//Response ---
type Response struct {
	Resources []Resource `json:"resources"`
}

//Resource ---
type Resource struct {
	Metadata map[string]string      `json:"metadata"`
	Entity   map[string]interface{} `json:"entity"`
}

// SummaryResponse ---
type SummaryResponse struct {
	Instances int `json:"running_instances"`
}
