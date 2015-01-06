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
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"bytes"
	"strings"
	"path/filepath"
)




//Application's entry point
//
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %v <port>\n", filepath.Base(os.Args[0]))
		os.Exit(1) 
	}	
	
	port := os.Args[1]
	
	http.HandleFunc("/getPort", getPort)
	
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal("Cannot start agent:", err)
	}		 	
}


//Data structure used to unmarshal Docker API's JSON response
//
type reducedContainerInfo struct {	
	HostConfig struct {	
		PortBindings map[string][]struct {
			HostPort string
		}	
	}
}


//Function handling /getPort requests
//
func getPort(writer http.ResponseWriter, request *http.Request) {		
	requestQuery := request.URL.Query()		
	dockerHost := strings.TrimSpace(requestQuery.Get("dockerHost"))
	dockerPort := strings.TrimSpace(requestQuery.Get("dockerPort"))
	containerName := strings.TrimSpace(requestQuery.Get("containerName"))
	containerPort := strings.TrimSpace(requestQuery.Get("containerPort"))
	
	
	dockerRequestUrl := fmt.Sprintf("http://%v:%v/containers/%v/json",
		dockerHost,
		dockerPort,
		containerName)
	
	dockerClient := http.Client{}
	if dockerResponse, err := dockerClient.Get(dockerRequestUrl); err == nil {
		buffer := bytes.Buffer{}
		buffer.ReadFrom(dockerResponse.Body)
		
		containerInfo := reducedContainerInfo{}		
		
		if err := json.Unmarshal(buffer.Bytes(), &containerInfo); err == nil {
			containerBindingKey := fmt.Sprintf("%v/tcp", containerPort)
			
			if portBindings, found := containerInfo.HostConfig.PortBindings[containerBindingKey]; found {
				//To simplify the program, only the first binding will be considered
				chosenBinding := portBindings[0]
								
				containerPort := chosenBinding.HostPort
				
				writer.Write([]byte(containerPort))			
			} else {
				writer.Write([]byte("The container does not expose the requested TCP port"))
			}												
		} else {
			writer.Write([]byte(err.Error()))
		}
		
	} else {
		writer.Write([]byte(err.Error()))
	}
}