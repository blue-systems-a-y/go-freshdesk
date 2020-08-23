package freshservice

import (
	"fmt"
	"time"
)

//AssetTypeItem ....
type AssetTypeItem struct {
	AssetType AssetType `json:"asset_type,omitempty"`
}

//AssetTypeList ...
type AssetTypeList struct {
	AssetTypes []AssetType `json:"asset_types,omitempty"`
}

//AssetType Asset types are a broad classification of different collection of assets.
type AssetType struct {
	//ID Unique ID of the asset type.
	ID int `json:"id,omitempty"`
	//Name Name of the asset type.
	Name string `json:"name,omitempty"`
	//Description Short description of the asset type.
	Description string `json:"description,omitempty"`
	//IParentAssetTypeID D of the parent asset type.
	ParentAssetTypeID int `json:"parent_asset_type_id,omitempty"`
	//Visible Visibility of the default asset type. Set to true if the asset type is visible. Custom asset types are set to true by default and cannot be modified.
	Visible bool `json:"visible,omitempty"`
	//Date and time when the asset type was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//Date and time when the asset type was last updated
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

//ListAssetTypes ...
func (c *Client) ListAssetTypes(page int) ([]AssetType, error) {
	var list AssetTypeList
	err := c.ReadObject(fmt.Sprintf("/api/v2/asset_types?page=%v", page), &list)
	return list.AssetTypes, err
}

//ListAllAssetTypes ...
func (c *Client) ListAllAssetTypes() ([]AssetType, error) {
	all := make([]AssetType, 0)
	i := 1
	for true {
		page, err := c.ListAssetTypes(i)
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

//CreateAssetType creates am asset type
func (c *Client) CreateAssetType(at *AssetType) (*AssetType, error) {
	var out AssetTypeItem
	err := c.WriteObject("/api/v2/asset_types", "POST", at, &out)
	if err != nil {
		return nil, err
	}
	return &out.AssetType, nil
}
