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
	"encoding/xml"
	"github.com/jjideenschmiede/gohood/categories"
)

// Categories are to get all child categories by id
func Categories(body categories.Api) (categories.Return, error) {

	// Define body data
	body.Type = "public"
	body.Version = "2.0"
	body.User = body.AccountName
	body.Password = body.AccountPass
	body.Function = "categoriesBrowse"

	// Remove fields
	body.AccountName = ""
	body.AccountPass = ""

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return categories.Return{}, err
	}

	// Config new request
	c := Config{convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return categories.Return{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode categories.Return

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return categories.Return{}, err
	}

	// Return data
	return decode, nil

}
