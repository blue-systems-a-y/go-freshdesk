package freshservice

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListAllLocations(t *testing.T) {
	domain, apiKey := GetCredentials()
	assert.NotEmpty(t, domain)
	assert.NotEmpty(t, apiKey)
	c := NewClient(domain, apiKey)

	locations, err := c.ListAllLocations()
	assert.NoError(t, err)
	assert.NotEmpty(t, locations)
	for _, location := range locations {
		log.Printf("%+v", location)
	}
}

// func TestClient_CreateLocation(t *testing.T) {
// 	domain, apiKey := GetCredentials()
// assert.NotEmpty(t, domain)
// assert.NotEmpty(t, apiKey)
// 	c := NewClient(domain, apiKey)

// 	resp, err := c.CreateLocation(&Location{
// 		Name: "Haoreg",
// 	})

// 	assert.NoError(t, err)
// 	log.Println(resp)
// }
