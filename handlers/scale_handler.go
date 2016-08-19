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

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/xchapter7x/cloudcontroller-client"
)

// Smessage - variable to hold json deserialized
var (
	Smessage     ScaleMessage
	responseCode int
	respMessage  ResponseMessage
)

type ScaleMessage struct {
	LoginURL    string  `json:"loginurl"`
	Apiurl      string  `json:"apiurl"`
	Org         string  `json:"org"`
	Space       string  `json:"space"`
	Appname     string  `json:"appname"`
	ScaleFactor float32 `json:"scalefactor"`
}

type CfMessage struct {
	InstanceCount int `json:"instances"`
}

type ResponseMessage struct {
	Status string `json:"status"`
}

func ScaleHandler(res http.ResponseWriter, req *http.Request, r Renderer) {
	auth := req.Header.Get("Authorization")
	if userpass, err := decodeBasic(auth); err == nil {
		bodyBytes, _ := ioutil.ReadAll(req.Body)

		if err := json.Unmarshal(bodyBytes, &Smessage); err != nil {
			fmt.Printf("Failed to unmarshal message: %v\n", err.Error())
			responseCode = http.StatusBadRequest
			respMessage = getResponseMessage("Bad Request")
		} else {
			fmt.Printf("Creating new client with: %v\n", Smessage)
			if cfClient, err := ccclient.New(Smessage.LoginURL, userpass[0], userpass[1], new(http.Client)).Login(); err == nil {
				scaleApplication(Smessage, cfClient)
			} else {
				fmt.Printf("Failed to login and create client")
				responseCode = http.StatusBadRequest
				respMessage = getResponseMessage(err.Error())
			}
		}
	}
	r.JSON(res, responseCode, respMessage)
}

func scaleApplication(message ScaleMessage, cfClient *ccclient.Client) {
	var (
		err       error
		orgGUID   string
		spaceGUID string
		appGUID   string
	)

	if orgGUID, err = getGUID(Smessage.Org, cfClient, "organizations"); err == nil {
		if spaceGUID, err = getGUID(Smessage.Space, cfClient, "spaces"); err == nil {
			if appGUID, err = getAppGUID(Smessage.Appname, orgGUID, spaceGUID, cfClient); err == nil {
				instanceCount, err := getAppInstanceCount(appGUID, cfClient)
				if err = updateAppInstanceCount(int(float32(instanceCount)*Smessage.ScaleFactor), appGUID, cfClient); err == nil {
					responseCode = http.StatusOK
					respMessage = getResponseMessage("complete")
				} else {
					responseCode = http.StatusInternalServerError
					respMessage = getResponseMessage(err.Error())
				}
			} else {
				responseCode = http.StatusInternalServerError
				respMessage = getResponseMessage(err.Error())
			}
		} else {
			responseCode = http.StatusInternalServerError
			respMessage = getResponseMessage(err.Error())
		}
	} else {
		responseCode = http.StatusInternalServerError
		respMessage = getResponseMessage(err.Error())
	}
}

func getAppInstanceCount(appGUID string, cfClient *ccclient.Client) (count int, err error) {
	var (
		responseMap *SummaryResponse
	)
	baseURL, err := url.Parse(strings.Join([]string{Smessage.Apiurl, "apps", appGUID, "summary"}, "/"))

	urlStr := fmt.Sprintf("%v", baseURL)
	fmt.Printf("Request: %s \n", urlStr)

	guidRequest, _ := http.NewRequest("GET", urlStr, nil)
	cfClient.AccessTokenDecorate(guidRequest)

	if guidResponse, err := cfClient.HttpClient().Do(guidRequest); err == nil {
		content, _ := ioutil.ReadAll(guidResponse.Body)
		fmt.Printf("Response received: %s \n", string(content[:]))
		responseMap = new(SummaryResponse)
		if err := json.Unmarshal(content, responseMap); err == nil {
			count = responseMap.Instances
			fmt.Printf("Running Instances: %v \n", count)
		}
	}
	return
}

func updateAppInstanceCount(instanceCount int, appGUID string, cfClient *ccclient.Client) (err error) {
	baseURL, err := url.Parse(strings.Join([]string{Smessage.Apiurl, "apps", appGUID}, "/"))

	urlStr := fmt.Sprintf("%v", baseURL)
	fmt.Printf("Request: %s \n", urlStr)

	cfmessage := CfMessage{
		InstanceCount: instanceCount,
	}

	content, _ := json.Marshal(cfmessage)

	guidRequest, _ := http.NewRequest("PUT", urlStr, bytes.NewBuffer(content))
	cfClient.AccessTokenDecorate(guidRequest)

	if guidResponse, err := cfClient.HttpClient().Do(guidRequest); err == nil {
		content, _ := ioutil.ReadAll(guidResponse.Body)
		fmt.Printf("Response received: %s \n", string(content[:]))
		if guidResponse.StatusCode != http.StatusOK {
			responseCode = guidResponse.StatusCode
			respMessage = getResponseMessage(guidResponse.Status)
		}
	}
	return
}

func getAppGUID(appName string, orgGUID string, spaceGUID string, cfClient *ccclient.Client) (appGUID string, err error) {
	var (
		responseMap *Response
	)
	baseURL, err := url.Parse(strings.Join([]string{Smessage.Apiurl, "apps"}, "/"))

	fmt.Printf("Query Params: %s, %s, %s\n", appName, orgGUID, spaceGUID)

	params := url.Values{}
	params.Add("q", strings.Join([]string{"name", appName}, ":"))
	params.Add("q", strings.Join([]string{"space_guid", spaceGUID}, ":"))
	params.Add("q", strings.Join([]string{"organization_guid", orgGUID}, ":"))
	baseURL.RawQuery = params.Encode()

	urlStr := fmt.Sprintf("%v", baseURL)
	fmt.Printf("Request: %s \n", urlStr)

	guidRequest, _ := http.NewRequest("GET", urlStr, nil)
	cfClient.AccessTokenDecorate(guidRequest)

	if guidResponse, err := cfClient.HttpClient().Do(guidRequest); err == nil {
		content, _ := ioutil.ReadAll(guidResponse.Body)
		fmt.Printf("Response received: %s \n", string(content[:]))
		responseMap = new(Response)
		if err := json.Unmarshal(content, responseMap); err == nil {
			appGUID = responseMap.Resources[0].Metadata["guid"]
		}
	}
	return
}

func getGUID(name string, cfClient *ccclient.Client, target string) (guid string, err error) {
	var (
		responseMap *Response
	)
	baseURL, err := url.Parse(strings.Join([]string{Smessage.Apiurl, target}, "/"))

	params := url.Values{}
	params.Add("q", strings.Join([]string{"name", name}, ":"))
	baseURL.RawQuery = params.Encode()

	urlStr := fmt.Sprintf("%v", baseURL)
	fmt.Printf("Request: %s \n", urlStr)

	guidRequest, _ := http.NewRequest("GET", urlStr, nil)
	cfClient.AccessTokenDecorate(guidRequest)

	if guidResponse, err := cfClient.HttpClient().Do(guidRequest); err == nil {
		content, _ := ioutil.ReadAll(guidResponse.Body)
		fmt.Printf("Response received: %s \n", string(content[:]))
		responseMap = new(Response)
		if err := json.Unmarshal(content, responseMap); err == nil {
			guid = responseMap.Resources[0].Metadata["guid"]
			fmt.Printf("Found GUID: %s for target: %s\n", guid, target)
		} else {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
	return
}

func getResponseMessage(message string) ResponseMessage {
	resmessage := ResponseMessage{
		Status: message,
	}
	return resmessage
}

func decodeBasic(auth string) (creds []string, err error) {
	fmt.Printf("Auth passed in %s \n", auth)
	if len(auth) < 6 || auth[:6] != "Basic " {
		responseCode = http.StatusUnauthorized
		respMessage = getResponseMessage("Unauthorized")
		err = errors.New("invalid credentials")
		return
	}

	b, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		responseCode = http.StatusUnauthorized
		respMessage = getResponseMessage("Unauthorized")
		err = errors.New("invalid credentials")
		return
	}

	creds = strings.SplitN(string(b), ":", 2)
	if len(creds) != 2 {
		responseCode = http.StatusUnauthorized
		respMessage = getResponseMessage("Unauthorized")
		err = errors.New("invalid credentials")
		return
	}
	return
}
