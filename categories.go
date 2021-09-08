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
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"github.com/jjideenschmiede/gohood/categories"
)

// CategoriesRequest is to structure the request data
type CategoriesRequest struct {
	AccountName string `xml:"accountName"`
	AccountPass string `xml:"accountPass"`
	CategoryId  int    `xml:"categoryID"`
}

// Categories are to get all child categories by id
func Categories(request CategoriesRequest) (categories.Return, error) {

	// Hash the password
	hash := fmt.Sprintf("%x", md5.Sum([]byte(request.AccountPass)))

	// Define body data
	body := categories.Api{
		"public",
		"2.0",
		request.AccountName,
		hash,
		"categoriesBrowse",
		request.CategoryId,
	}

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
