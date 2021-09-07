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
)

// ProductsBody is to structure the data
type ProductsBody struct {
	Api Api `xml:"api"`
}

type Api struct {
	Type        string `xml:"type,attr,omitempty"`
	Version     string `xml:"version,attr,omitempty"`
	User        string `xml:"user,attr,omitempty"`
	Password    string `xml:"password,attr,omitempty"`
	Function    string `xml:"function,omitempty"`
	AccountName string `xml:"accountName"`
	AccountPass string `xml:"accountPass"`
	Items       Items  `xml:"items"`
}

type Items struct {
	Item []Item `xml:"item"`
}

type Item struct {
	ItemId         int            `xml:"itemID,omitempty"`
	ItemMode       string         `xml:"itemMode"`
	CategoryId     int            `xml:"categoryID"`
	ItemName       string         `xml:"itemName"`
	Quantity       int            `xml:"quantity"`
	Condition      string         `xml:"condition"`
	Description    string         `xml:"description"`
	Shipmethods    Shipmethods    `xml:"shipmethods"`
	Price          string         `xml:"price"`
	SalesTax       string         `xml:"salesTax"`
	ProductOptions ProductOptions `xml:"productOptions"`
	Ean            string         `xml:"ean"`
	Isbn           string         `xml:"isbn"`
	Mpn            string         `xml:"mpn"`
	Manufacturer   string         `xml:"manufacturer"`
	Weight         string         `xml:"weight"`
	Images         Images         `xml:"images"`
}

type Shipmethods struct {
	Shipmethod []Shipmethod `xml:"shipmethod"`
}

type Shipmethod struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

type ProductOptions struct {
	ProductOption []ProductOption `xml:"productOption"`
}

type ProductOption struct {
	OptionPrice      string        `xml:"optionPrice"`
	OptionQuantity   int           `xml:"optionQuantity"`
	OptionItemNumber int           `xml:"optionItemNumber"`
	Mpn              string        `xml:"mpn"`
	Ean              string        `xml:"ean"`
	PackagingSize    string        `xml:"PackagingSize"`
	OptionsDetails   OptionDetails `xml:"optionDetails"`
}

type OptionDetails struct {
	NameValueList []NameValueList `xml:"nameValueList"`
}

type NameValueList struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type Images struct {
	ImageUrl []string `xml:"imageURL"`
}

// ProductsReturn is to decode the xml response
type ProductsReturn struct {
	XmlName xml.Name           `xml:"xml_name"`
	Item    ProductsReturnItem `xml:"item"`
}

type ProductsReturnItem struct {
	XmlName     xml.Name `xml:"item"`
	ReferenceId int      `xml:"referenceID"`
	Status      string   `xml:"status"`
	Costs       int      `xml:"costs"`
	ItemId      int      `xml:"itemID"`
}

// AddProduct are to set a new product
func AddProduct(body ProductsBody) (ProductsReturn, error) {

	// Define body data
	body.Api.Type = "public"
	body.Api.Version = "2.0"
	body.Api.User = body.Api.AccountName
	body.Api.Password = body.Api.AccountPass
	body.Api.Function = "itemInsert"

	// Convert body
	convert, err := xml.Marshal(body.Api)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Config new request
	c := Config{convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return ProductsReturn{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode ProductsReturn

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return ProductsReturn{}, err
	}

	// Return data
	return decode, nil

}
