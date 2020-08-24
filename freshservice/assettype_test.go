package freshservice

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestClient_CreateAssetType(t *testing.T) {
// domain, apiKey := GetCredentials()
// assert.NotEmpty(t, domain)
// assert.NotEmpty(t, apiKey)
// 	c := NewClient(domain, apiKey)

// 	resp, err := c.CreateAssetType(&AssetType{
// 		Name: "Shield",
// 	})

// 	assert.NoError(t, err)
// 	log.Println(resp)
// }

func TestClient_ListAllAssetTypes(t *testing.T) {
	domain, apiKey := GetCredentials()
	c := NewClient(domain, apiKey)

	list, err := c.ListAllAssetTypes()
	assert.NoError(t, err)
	for _, at := range list {
		log.Printf("%+v", at)
	}
}
