package freshservice

import (
	"fmt"
	"time"
)

//LocationItem holds one location
type LocationItem struct {
	Location Location `json:"location,omitempty"`
}

//LocationList holds a list of locations
type LocationList struct {
	Locations []Location `json:"locations,omitempty"`
}

//Address ...
type Address struct {
	//Address line 1.
	Line1 string `json:"line1,omitempty"`
	//Address line 2.
	Line2 string `json:"line2,omitempty"`
	//	Name of the City.
	City string `json:"city,omitempty"`
	//Name of the State.
	State string `json:"state,omitempty"`
	//Name of the Country.
	Country string `json:"country,omitempty"`
	//Zipcode of the location.
	Zipcode string `json:"zipcode,omitempty"`
}

//Location location
type Location struct {
	//Unique ID of the location.
	ID int `json:"id,omitempty"`
	//Name of the location.
	Name string `json:"name,omitempty"`
	//ID of the parent location.
	ParentLocationID int `json:"parent_location_id,omitempty"`
	//User ID of the primary contact (must be a user in Freshservice).
	PrimaryContactID int `json:"primary_contact_id,omitempty"`
	//Address ...
	Address Address `json:"address,omitempty"`
	//Date and time when the location was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//Date and time when the location was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

//ListLocations get locations by page
func (c *Client) ListLocations(page int) ([]Location, error) {
	var list LocationList
	err := c.ReadObject(fmt.Sprintf("/api/v2/locations?page=%v", page), &list)
	return list.Locations, err
}

//ListAllLocations read all locations
func (c *Client) ListAllLocations() ([]Location, error) {
	all := make([]Location, 0)
	i := 1
	for true {
		page, err := c.ListLocations(i)
		if err != nil {
			return nil, err
		}
		if len(page) == 0 {
			break
		}
		all = append(all, page...)
		i++
	}
	return all, nil
}

//CreateLocation creates a location
func (c *Client) CreateLocation(loc *Location) (*Location, error) {
	var out LocationItem
	err := c.WriteObject("/api/v2/locations", "POST", loc, &out)
	if err != nil {
		return nil, err
	}
	return &out.Location, nil
}
