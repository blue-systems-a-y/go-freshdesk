package freshservice

import (
	"fmt"
	"time"
)

//AssetItem holds one asset
type AssetItem struct {
	Asset Asset `json:"asset,omitempty"`
}

//AssetList holds a list of assets
type AssetList struct {
	Assets []Asset `json:"assets,omitempty"`
}

//Asset ...
type Asset struct {
	ID        int    `json:"id,omitempty"`
	DisplayID int    `json:"display_id,omitempty"`
	Name      string `json:"name,omitempty"`
	//Description of the asset.
	Description string `json:"description,omitempty"`
	//ID of the asset type.
	AssetTypeID int `json:"asset_type_id,omitempty"`
	//Asset tag of the asset.
	AssetTag string `json:"asset_tag,omitempty"`
	//Impact of the asset.
	Impact string `json:"impact,omitempty"`
	//	Indicates whether the asset was created by a user or discovery tools (Probe or Agent).
	AuthorType string `json:"author_type,omitempty"`
	//Usage type of the asset (Loaner / Permanent).
	UsageType string `json:"usage_type,omitempty"`
	//ID of the associated user (Used By).
	UserID int `json:"user_id,omitempty"`
	//ID of the associated location.
	LocationID int `json:"location_id,omitempty"`
	//ID of the associated department.
	DepartmentID int `json:"department_id,omitempty"`
	//ID of the associated agent (Managed By).
	AgentID int `json:"agent_id,omitempty"`
	//ID of the associated agent group (Managed By Group).
	GroupID int `json:"group_id,omitempty"`
	//Date and time when the asset was assigned.
	AssignedOn *time.Time `json:"assigned_on,omitempty"`
	//Date and time when the asset was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	//Date and time when the asset was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	//TypeFields type specific fields
	TypeFields map[string]interface{} `json:"type_fields,omitempty"`
}

//ListAssets ...
func (c *Client) ListAssets(page int) ([]Asset, error) {
	var list AssetList
	err := c.ReadObject(fmt.Sprintf("/api/v2/assets?page=%v", page), &list)
	return list.Assets, err
}

//ListAlAssets ...
func (c *Client) ListAlAssets() ([]Asset, error) {
	all := make([]Asset, 0)
	i := 1
	for true {
		page, err := c.ListAssets(i)
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

//CreateAsset creates an asset
func (c *Client) CreateAsset(asset *Asset) (*Asset, error) {
	var out AssetItem
	err := c.WriteObject("/api/v2/assets", "POST", asset, &out)
	if err != nil {
		return nil, err
	}
	return &out.Asset, nil
}

//DeleteAsset deletes an asset
func (c *Client) DeleteAsset(displayID int) error {
	var out AssetItem
	err := c.WriteObject(fmt.Sprintf("/api/v2/assets/%v", displayID), "DELETE", "", &out)
	if err != nil {
		return err
	}
	return nil
}
