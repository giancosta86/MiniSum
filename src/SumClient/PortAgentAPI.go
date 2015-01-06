/*
The MIT License (MIT)

Copyright (c) 2015 Gianluca Costa

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"net/http"
	"net/url"
	"errors"
	"fmt"
	"bytes"
	"strconv"	
)

//A <host>:<port> service binding
//
type ServiceBinding struct {
	Host string
	Port string
}

func (serviceBinding ServiceBinding) String() string {
	return fmt.Sprintf("%v:%v", serviceBinding.Host, serviceBinding.Port)
}


//PortAgent's parameters for retrieving SumService's host port
//
type PortAgentParams struct {
	DockerHost string
	DockerPort string
	ContainerName string
	ContainerPort string
}


func (portAgentParams *PortAgentParams) toUrlQuery() (result url.Values) {	
	result = url.Values{}
	
	result.Add("dockerHost", portAgentParams.DockerHost)
	result.Add("dockerPort", portAgentParams.DockerPort)
	result.Add("containerName", portAgentParams.ContainerName)
	result.Add("containerPort", portAgentParams.ContainerPort)	
	
	return
}


//Calls PortAgent via HTTP to retrieve SumService's port
//
func getPort(portAgentBinding *ServiceBinding, portAgentParams *PortAgentParams) (servicePort string, err error) {		
	portAgentClient := http.Client{}
	portAgentRequestUrl := fmt.Sprintf("http://%v/getPort?%v", portAgentBinding, portAgentParams.toUrlQuery().Encode())	
	
	if portAgentResponse, err := portAgentClient.Get(portAgentRequestUrl); err == nil {			
		responseBuffer := bytes.Buffer{}
		responseBuffer.ReadFrom(portAgentResponse.Body)
		responseString := responseBuffer.String()
		
		if _, err := strconv.Atoi(responseString); err == nil {
			servicePort = responseString 	
										 	
			return servicePort, nil				
		} else {
			return "", errors.New(responseString)
		}
	} else {
		return "", err
	}
}