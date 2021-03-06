package freshservice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAssets(t *testing.T) {
	domain, apiKey := GetCredentials()
	assert.NotEmpty(t, domain)
	assert.NotEmpty(t, apiKey)
	c := NewClient(domain, apiKey)
	i := 1
	assets, err := c.ListAssets(i)
	assert.NoError(t, err)
	for len(assets) > 0 {
		assert.NotEmpty(t, assets)
		i++
		assets, err = c.ListAssets(i)
		assert.NoError(t, err)
	}
}

func TestClient_GetAllAssets(t *testing.T) {
	domain, apiKey := GetCredentials()
	assert.NotEmpty(t, domain)
	assert.NotEmpty(t, apiKey)

	c := NewClient(domain, apiKey)
	assets, err := c.ListAllAssets()
	assert.NoError(t, err)
	assert.NotEmpty(t, assets)
}

// func TestClient_CreateAsset(t *testing.T) {
// domain, apiKey := GetCredentials()
// assert.NotEmpty(t, domain)
// assert.NotEmpty(t, apiKey)
// 	c := NewClient(domain, apiKey)

// 	asset := &Asset{
// 		Name:        "Shield 1",
// 		AssetTypeID: 17000226416,
// 		TypeFields:  make(map[string]interface{}),
// 	}
// 	asset.TypeFields["asset_state_17000220160"] = "In Use"
// 	asset.TypeFields["product_17000220160"] = 17000019183
// 	resp, err := c.CreateAsset(asset)

// 	assert.NoError(t, err)
// 	log.Println(resp)
// }
