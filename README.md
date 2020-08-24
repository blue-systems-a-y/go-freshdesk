# go-freshdesk
Go client library for Freshdesk
https://developers.freshdesk.com/api/

---
## Overview
go-freshdesk is a client package written in go, used to ease the operations on Freshdesk services, by providing many functions and tools to deal with them, especially the inventory managements, and tickiting system.

---
**Badges**

[![Build Status](http://img.shields.io/travis/badges/badgerbadgerbadger.svg?style=flat-square)](https://travis-ci.org/badges/badgerbadgerbadger)   [![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)
---

## Table of Contents

> Choose the section you are interested in
- [Overview](##Overview)
- [Prerequisites](##Prerequisites)
- [Integration](##Integration)
- [Example](##Example)
- [License](##license)



## Prerequisites
To be able to use the client:
- A valid freshdesk account --> Valid ```domain``` & ```apiKey```.


---
## Integration

- Import the package in your target project
```go
import github.com/EDIT THIS
```
- Create a freshdesk client
```go
fc := freshservice.NewClient("DOMAIN_NAME", "API_KEY")
```
---

## Usage
### Freshdesk Services
- Assets Operations
    - >```ListAssets``` --> lists all assets on a certin page.
    - >```ListAllAssets``` --> lists all assets on all pages.
    - >```CreateAsset``` --> creates an asset in the inventory.
    - >```DeleteAsset``` --> Deletes an asset from the inventory.
- Asset Types Operations
    - >```ListAssetTypes``` --> lists all asset types on a certin page.
    - >```ListAllAssetTypes``` --> lists all asset types on all pages.
    - >```CreateAssetType``` --> creates an asset type in the inventory.

- Locations Operations
    - >```ListLocations``` --> lists all locations on a certin page.
    - >```ListAllLocations``` --> lists all locations on all pages.
    - >```CreateLocation``` --> creates a new location.

- Products Operations
    - >```ListProducts``` --> lists all locations on a certin page.
    - >```ListAllProducts``` --> lists all locations on all pages.
    - >```CreateProduct``` --> creates a new location.

- Tickets Operations
    - >```ListTickets``` --> lists all tickets on a certin page.
    - >```ListAllTickets``` --> lists all tickets on all pages.
    - >```CreateTickets``` --> creates a new ticket.   

***For Further informations about the apis, see Freshdesk APIs docs: https://api.freshservice.com/v2/***    

---

## Example 1
```go
// Create an asset
c := NewClient(MY_DOMIN, MY_APIKEY)
newAsset := &Asset{
    Name:        "HP Labtop",
    AssetTypeID: 11568,
}
res,err := c.CreateAsset(newAsset)
```
---
## Example 2
```go
//List all tickets
c := NewClient(MY_DOMIN, MY_APIKEY)
tickets, err := c.ListAllTickets()
for _, ticket := range tickets {
    log.Printf("%+v", ticket)
    }
}
``` 
## License

- **[MIT license](http://opensource.org/licenses/mit-license.php)**

## Authors

* **Danny Tylman** - *Gabriel Systems* 