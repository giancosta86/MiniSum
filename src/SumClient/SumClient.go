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
	"fmt"
	"os"
	"bufio"	
	"net/url"	
	"net/http"
	"bytes"
	"strings"
	"log"
)

var consoleReader = bufio.NewReader(os.Stdin)


//Interactively reequests SumService's coordinates until a port is correctly returned by PortAgent
//
func askForSumService(portAgentBinding *ServiceBinding) (sumServiceBinding ServiceBinding, err error) {
	for {
		portAgentParams := PortAgentParams{}
		
		fmt.Println()
		fmt.Println()
		fmt.Println("== SUMSERVICE COORDINATES ==")
		fmt.Println()		
		
		fmt.Print("Docker host: ")	
		if portAgentParams.DockerHost, err = consoleReader.ReadString('\n'); err != nil {			
			return 
		}		
		
		fmt.Print("Docker port (usually, 2375): ")	
		if portAgentParams.DockerPort, err = consoleReader.ReadString('\n'); err != nil {
			return
		}		
	
		
		fmt.Print("Container name: ")
		if portAgentParams.ContainerName, err = consoleReader.ReadString('\n'); err != nil {
			return
		}
		
		fmt.Print("Container port: ")
		if portAgentParams.ContainerPort, err = consoleReader.ReadString('\n'); err != nil {
			return
		}
		
		
		if sumServiceBinding.Port, err = getPort(portAgentBinding, &portAgentParams); err == nil {
			sumServiceBinding.Host = strings.TrimSpace(portAgentParams.DockerHost)
			return			
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
			fmt.Println() 
		}	
	}
} 


//Application's entry point
//
func main() {		
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %v <port agent's host> <port agent's port>\n", os.Args[0])
		os.Exit(1)
	}
	
	portAgentBinding := ServiceBinding {
		Host: os.Args[1],
		Port: os.Args[2],
	}
		
	
	sumServiceBinding, err := askForSumService(&portAgentBinding)	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("SumService is listening on %v\n", sumServiceBinding)
	
	for {		
		fmt.Println()
		
		fmt.Print("Op1: ")
		op1, err := consoleReader.ReadString('\n')
		if err != nil {			
			return
		}
		
		fmt.Print("Op2: ")
		op2, err := consoleReader.ReadString('\n')
		if err != nil {
			return
		}
		
		
		queryParams := url.Values{}
		queryParams.Set("op1", op1)
		queryParams.Set("op2", op2)
		 
		client := &http.Client{}
		requestUrl := fmt.Sprintf("http://%v/sum?%v", sumServiceBinding, queryParams.Encode())		
		
		
		serviceResponse, err := client.Get(requestUrl)
		if err != nil {
			log.Fatalln("Error while running the service:", err)			
		}
		
		
		responseBuffer := bytes.Buffer{}
		responseBuffer.ReadFrom(serviceResponse.Body)
		responseString := responseBuffer.String()
		
		fmt.Println("Result:",  responseString)
	}	
}