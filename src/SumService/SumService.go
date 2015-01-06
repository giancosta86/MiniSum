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
	"net/http"
	"log"
	"strconv"		
	"fmt"
	"strings"
	"path/filepath"
)


//Application's entry point
//
func main() {	
	if len(os.Args) == 1 {		
		fmt.Printf("Usage: %v <port>\n", filepath.Base(os.Args[0]))
		os.Exit(1) 
	}
	
	port := os.Args[1]
	
	http.HandleFunc("/sum", sum)
	
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal("Cannot start the service:", err)
	}		 	
}


//Function handling /sum requests
//
func sum(writer http.ResponseWriter, request *http.Request) {
	requestQuery := request.URL.Query()
	
	op1String := strings.TrimSpace(requestQuery.Get("op1"))
	op2String := strings.TrimSpace(requestQuery.Get("op2"))	
		
	op1, err := strconv.Atoi(op1String)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("Invalid op1: '%s'", op1String)))
		return
	}
	
	op2, err := strconv.Atoi(op2String)
	if err != nil {
		writer.Write([]byte(fmt.Sprintf("Invalid op2: '%s'", op2String)))
		return
	}
	
	
	result := op1 + op2	
	
	writer.Write([]byte(strconv.Itoa(result)))
}