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
	"github.com/jjideenschmiede/gohood/products"
)

// ProductsRequest is to structure the data
type ProductsRequest struct {
	AccountName string               `xml:"accountName"`
	AccountPass string               `xml:"accountPass"`
	Items       ProductsRequestItems `xml:"items"`
}

type ProductsRequestItems struct {
	Item []ProductsRequestItem `xml:"item"`
}

type ProductsRequestItem struct {
	ItemId         int                           `xml:"itemID,omitempty"`
	ItemMode       string                        `xml:"itemMode"`
	CategoryId     int                           `xml:"categoryID"`
	ItemName       string                        `xml:"itemName"`
	Quantity       int                           `xml:"quantity"`
	Condition      string                        `xml:"condition"`
	Description    string                        `xml:"description"`
	Shipmethods    ProductsRequestShipmethods    `xml:"shipmethods"`
	Price          string                        `xml:"price"`
	SalesTax       string                        `xml:"salesTax"`
	PackagingSize  string                        `xml:"packagingSize"`
	PackagingUnit  string                        `xml:"packagingUnit"`
	ProdCatId      string                        `xml:"prodCatID"`
	ProdCatId2     string                        `xml:"prodCatID2"`
	ProdCatId3     string                        `xml:"prodCatID3"`
	ShortDesc      string                        `xml:"shortDesc"`
	IfIsSoldOut    string                        `xml:"ifIsSoldOut"`
	ProductOptions ProductsRequestProductOptions `xml:"productOptions"`
	Ean            string                        `xml:"ean"`
	Isbn           string                        `xml:"isbn"`
	Mpn            string                        `xml:"mpn"`
	Manufacturer   string                        `xml:"manufacturer"`
	Weight         string                        `xml:"weight"`
	Images         ProductsRequestImages         `xml:"images"`
}

type ProductsRequestShipmethods struct {
	Shipmethod []ProductsRequestShipmethod `xml:"shipmethod"`
}

type ProductsRequestShipmethod struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value"`
}

type ProductsRequestProductOptions struct {
	ProductOption []ProductsRequestProductOption `xml:"productOption"`
}

type ProductsRequestProductOption struct {
	OptionPrice      string                       `xml:"optionPrice"`
	OptionQuantity   int                          `xml:"optionQuantity"`
	OptionItemNumber int                          `xml:"optionItemNumber"`
	Mpn              string                       `xml:"mpn"`
	Ean              string                       `xml:"ean"`
	PackagingSize    string                       `xml:"PackagingSize"`
	OptionsDetails   ProductsRequestOptionDetails `xml:"optionDetails"`
}

type ProductsRequestOptionDetails struct {
	NameValueList []ProductsRequestNameValueList `xml:"nameValueList"`
}

type ProductsRequestNameValueList struct {
	Name  string `xml:"name"`
	Value string `xml:"value"`
}

type ProductsRequestImages struct {
	ImageUrl []string `xml:"imageURL"`
}

// AddProduct are to set a new product
func AddProduct(request ProductsRequest) (products.Return, error) {

	// Hash the password
	hash := fmt.Sprintf("%x", md5.Sum([]byte(request.AccountPass)))

	// Define body data
	body := products.Api{
		"public",
		"2.0",
		request.AccountName,
		hash,
		"itemInsert",
		request.AccountName,
		hash,
		request.Items,
	}

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return products.Return{}, err
	}

	// Config new request
	c := Config{convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return products.Return{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode products.Return

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return products.Return{}, err
	}

	// Return data
	return decode, nil

}

// UpdateProduct is to update a product
func UpdateProduct(request ProductsRequest) (products.Return, error) {

	// Hash the password
	hash := fmt.Sprintf("%x", md5.Sum([]byte(request.AccountPass)))

	// Define body data
	body := products.Api{
		"public",
		"2.0",
		request.AccountName,
		hash,
		"itemUpdate",
		request.AccountName,
		hash,
		request.Items,
	}

	// Convert body
	convert, err := xml.Marshal(body)
	if err != nil {
		return products.Return{}, err
	}

	// Config new request
	c := Config{convert}

	// Send new request
	response, err := c.Send()
	if err != nil {
		return products.Return{}, err
	}

	// Close request body
	defer response.Body.Close()

	// Decode data
	var decode products.Return

	err = xml.NewDecoder(response.Body).Decode(&decode)
	if err != nil {
		return products.Return{}, err
	}

	// Return data
	return decode, nil

}
