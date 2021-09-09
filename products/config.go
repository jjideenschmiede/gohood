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

package products

import (
	"encoding/xml"
)

type Api struct {
	Type        string      `xml:"type,attr"`
	Version     string      `xml:"version,attr"`
	User        string      `xml:"user,attr"`
	Password    string      `xml:"password,attr"`
	Function    string      `xml:"function"`
	AccountName string      `xml:"accountName"`
	AccountPass string      `xml:"accountPass"`
	Items       interface{} `xml:"items"`
}

// Return is to decode the xml response
type Return struct {
	XmlName xml.Name   `xml:"response"`
	Item    ReturnItem `xml:"item"`
}

type ReturnItem struct {
	XmlName     xml.Name `xml:"item"`
	ReferenceId int      `xml:"referenceID"`
	Status      string   `xml:"status"`
	Costs       int      `xml:"costs,omitempty"`
	ItemId      int      `xml:"itemID,omitempty"`
	Message     string   `xml:"message,omitempty"`
}
