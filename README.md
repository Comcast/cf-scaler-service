# cf-scaler-service

This Cloudfoundry helper service is to allow the remote scaling of application instances without having to access the console of CLI.

The service works by interacting with the Cloudfoundry REST API to effect changes to the instances assigned to an application. A possible use case would be a monitoring event triggering a call to increase capacity.

### Prerequisites
1. Golang 1.5.3+
2. Godep
3. cloudcontroller-client

### Build from Source
```sh
git clone git@github.com:comcast/cf-scaler-service.git
go get github.com/tools/godep
go get github.com/codegangsta/negroni
go get github.com/unrolled/render
go get github.com/xchapter7x/cloudcontroller-client
godep go build server/main.go
```

The binary will be in the repository directory. Ensure the PORT environment variable is set before starting.

### Usage
Post a request with the below JSON structure, credentials for the Cloudfoundry foundation are passed through as Basic Auth in the header.
```json
{
	"loginurl": "https://login.mycfplatform.comcast.net",
	"apiurl": "https://mycfplatform.comcast.net/v2",
	"org": "MY-ORG",
	"space": "my-space",
    "appname": "my-application-generic#1.0.0.81-f4e05e",
	"scalefactor": 0.5
}
```
**loginurl** - authentication endpoint for Cloudfoundry foundation.  
**apiurl** - api endpoint for the Cloudfoundry foundation.  
**org** - name of the target organization.  
**space** - name of the target space.  
**appname** - name of the application to be scaled.  
**scalefactor** - float multiplier, >1 grows the number of instances, <1 shrinks.  

### Todo
1. Improve test coverage and fakes for cloudcontroller-client.
2. Finish implementing Concourse.ci pipeline.
