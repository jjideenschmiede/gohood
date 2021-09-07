//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of gohood.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package gohood

import (
	"bytes"
	"net/http"
)

const (
	baseUrl = "https://www.hood.de/api.htm"
	method  = "POST"
)

// Config is to define the request data
type Config struct {
	Body []byte
}

// Send is to send a new request
func (c *Config) Send() (*http.Response, error) {

	// Define client
	client := &http.Client{}

	// Define request
	request, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(c.Body))
	if err != nil {
		return nil, err
	}

	// Set header
	request.Header.Set("Content-Type", "text/xml; charset=UTF-8")

	// Send request & get response
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	// Return data
	return response, nil

}
