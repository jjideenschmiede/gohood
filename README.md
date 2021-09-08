# gohood

With this library you can access the endpoints of the hood.de API and read, update or write data. We will continue to develop this library. If you have any questions or comments, please contact info@jj-development.de.

# Install

```console
go get github.com/jjideenschmiede/gohood
```

# How to use?

## Add a product

In order to add a new product, some data is needed. The whole thing can look like this.

**Important! The password must be hashed.**

```go
// Create product body
body := gohood.ProductsBody{
    gohood.Api{
        AccountName: "",
        AccountPass: "",
        Items: gohood.Items{
            Item: []gohood.Item{},
        },
    },
}

// Add new item
body.Api.Items.Item = append(body.Api.Items.Item, gohood.Item{
    ItemMode:    "shopProduct",
    CategoryId:  17055,
    ItemName:    "J&J Marketplace",
    Quantity:    250,
    Condition:   "new",
    Description: "Unsere Software, damit Sie ihr Warenwirtschaftssystem direkt mit einem Marktplatz verbinden können.",
    Shipmethods: gohood.Shipmethods{
        Shipmethod: []gohood.Shipmethod{},
    },
    Price:          "0,75",
    SalesTax:       "19",
    PackagingSize:  "",
    PackagingUnit:  "",
    ProdCatId:      "",
    ProdCatId2:     "",
    ProdCatId3:     "",
    ShortDesc:      "",
    IfIsSoldOut:    "",
    ProductOptions: gohood.ProductOptions{},
    Ean:            "4251209995017",
    Isbn:           "",
    Mpn:            "JJMRTPL2021",
    Manufacturer:   "JJ Ideenschmiede GmbH",
    Weight:         "0,00",
    Images: gohood.Images{
        ImageUrl: []string{},
    },
})

// Add shipping methods
body.Api.Items.Item[0].Shipmethods.Shipmethod = append(body.Api.Items.Item[0].Shipmethods.Shipmethod, gohood.Shipmethod{
    Name:  "DHLPacket_nat",
    Value: "5",
})

// Add an image
body.Api.Items.Item[0].Images.ImageUrl = append(body.Api.Items.Item[0].Images.ImageUrl, "https://lh3.googleusercontent.com/glsgmb/AJtb4XCbjWmNDgametDQMVJo6Oh6Kok2GuoRs59ozCAQMmFl9G2f2PTq6PI0_9GZROeNzO0w13M5A91gArGZ6u-GGHsjPw=w304-h899-rw-no-sc0x00ffffff")

// Add a new product
product, err := gohood.AddProduct(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(product)
}
```

## Update a product

In order to update a product, the ID of the product is required. Otherwise you can use the entire data identically to the anlgegen of a product.

```go
// Create product body
body := gohood.ProductsBody{
    gohood.Api{
        AccountName: "",
        AccountPass: "",
        Items: gohood.Items{
            Item: []gohood.Item{},
        },
    },
}

// Add new item
body.Api.Items.Item = append(body.Api.Items.Item, gohood.Item{
    ItemId:      12321873,
    ItemMode:    "shopProduct",
    CategoryId:  17055,
    ItemName:    "J&J Marketplace",
    Quantity:    250,
    Condition:   "new",
    Description: "Unsere Software, damit Sie ihr Warenwirtschaftssystem direkt mit einem Marktplatz verbinden können.",
    Shipmethods: gohood.Shipmethods{
        Shipmethod: []gohood.Shipmethod{},
    },
    Price:          "0,75",
    SalesTax:       "19",
    PackagingSize:  "",
    PackagingUnit:  "",
    ProdCatId:      "",
    ProdCatId2:     "",
    ProdCatId3:     "",
    ShortDesc:      "",
    IfIsSoldOut:    "",
    ProductOptions: gohood.ProductOptions{},
    Ean:            "4251209995017",
    Isbn:           "",
    Mpn:            "JJMRTPL2021",
    Manufacturer:   "JJ Ideenschmiede GmbH",
    Weight:         "0,00",
    Images: gohood.Images{
        ImageUrl: []string{},
    },
})

// Add shipping methods
body.Api.Items.Item[0].Shipmethods.Shipmethod = append(body.Api.Items.Item[0].Shipmethods.Shipmethod, gohood.Shipmethod{
    Name:  "DHLPacket_nat",
    Value: "5",
})

// Add an image
body.Api.Items.Item[0].Images.ImageUrl = append(body.Api.Items.Item[0].Images.ImageUrl, "https://lh3.googleusercontent.com/glsgmb/AJtb4XCbjWmNDgametDQMVJo6Oh6Kok2GuoRs59ozCAQMmFl9G2f2PTq6PI0_9GZROeNzO0w13M5A91gArGZ6u-GGHsjPw=w304-h899-rw-no-sc0x00ffffff")

// Update a product
product, err := gohood.UpdateProduct(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(product)
}
```

## Browse categories

With the following function it is possible to search in the categories of Hood.de.

```go
// Define request body
body := gohood.CategoriesRequest{
    AccountName: "",
    AccountPass: "",
    CategoryId:  0,
}

// Get child categories
categories, err := gohood.Categories(body)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(categories)
}
```