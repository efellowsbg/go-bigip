package bigip

import (
	"context"
	"encoding/json"
)

type (
	WebtopType        string
	CustomizationType string
	InitialState      string
	LinkType          string
)

const (
	Portal      WebtopType        = "portal-access"
	Full        WebtopType        = "full"
	Network     WebtopType        = "network-access"
	Modern      CustomizationType = "Modern"
	Standard    CustomizationType = "Standard"
	Collapsed   InitialState      = "Collapsed"
	Expanded    InitialState      = "Expanded"
	LinkTypeUri LinkType          = "uri"
)

type BooledString bool

// Some endpoints have a "booledString" a boolean value that is represented as a string in the json payload
func (b BooledString) MarshalJSON() ([]byte, error) {
	str := "false"
	if b {
		str = "true"
	}
	return json.Marshal(str)
}

func (b BooledString) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	b = str == "true"
	return nil
}

type Webtop struct {
	Name               string            `json:"name,omitempty"`
	Partition          string            `json:"partition,omitempty"`
	Description        string            `json:"description,omitempty"`
	Type               WebtopType        `json:"webtopType,omitempty"`
	LinkType           LinkType          `json:"linkType,omitempty"`
	TMPartition        string            `json:"tmPartition,omitempty"`
	CustomizationType  CustomizationType `json:"customizationType,omitempty"`
	CustomizationGroup string            `json:"customizationGroup"`
	LocationSpecific   BooledString      `json:"locationSpecific"`
	MinimizeToTray     BooledString      `json:"minimizeToTray"`         // Updateable
	ShowSearch         BooledString      `json:"showSearch"`             // Updateable
	WarningOnClose     BooledString      `json:"warningOnClose"`         // Updateable
	UrlEntryField      BooledString      `json:"urlEntryField"`          // Updateable
	ResourceSearch     BooledString      `json:"resourceSearch"`         // Updateable
	InitialState       InitialState      `json:"initialState,omitempty"` // Updateable
}

type WebtopRead struct {
	Webtop
	FullPath                    string `json:"fullPath,omitempty"`
	Generation                  int    `json:"generation,omitempty"`
	SelfLink                    string `json:"selfLink,omitempty"`
	CustomizationGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"customizationGroupReference,omitempty"`
}

func (b *BigIP) CreateWebtop(ctx context.Context, webtop Webtop) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return b.post(webtop, uriMgmt, uriTm, uriApm, uriResource, uriWebtop)
}

func (b *BigIP) DeleteWebtop(ctx context.Context, name string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return b.delete(uriMgmt, uriTm, uriApm, uriResource, uriWebtop, name)
}

func (b *BigIP) GetWebtop(ctx context.Context, name string) (*WebtopRead, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	var webtop WebtopRead
	err, _ := b.getForEntity(&webtop, uriMgmt, uriTm, uriApm, uriResource, uriWebtop, name)
	return &webtop, err
}

func (b *BigIP) ModifyWebtop(ctx context.Context, name string, webtop Webtop) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return b.patch(webtop, uriMgmt, uriTm, uriApm, uriResource, uriWebtop, name)
}
