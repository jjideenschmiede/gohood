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

package categories

import "encoding/xml"

// Api is to set the request
type Api struct {
	Type        string `xml:"type,attr,omitempty"`
	Version     string `xml:"version,attr,omitempty"`
	User        string `xml:"user,attr,omitempty"`
	Password    string `xml:"password,attr,omitempty"`
	Function    string `xml:"function,omitempty"`
	AccountName string `xml:"accountName,omitempty"`
	AccountPass string `xml:"accountPass,omitempty"`
	CategoryId  int    `xml:"categoryID"`
}

// Return is to decode the xml response
type Return struct {
	XmlName    xml.Name         `xml:"response"`
	Category   ReturnCategory   `xml:"category,omitempty"`
	Categories ReturnCategories `xml:"categories"`
}

type ReturnCategory struct {
	CategoryId    int    `xml:"categoryID"`
	ParentId      int    `xml:"parentID"`
	CategoryName  string `xml:"categoryName"`
	ChildCount    int    `xml:"childCount"`
	InsertProduct int    `xml:"insertProduct"`
}

type ReturnCategories struct {
	Category []ReturnCategory `xml:"category"`
}
