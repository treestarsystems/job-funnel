package load

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// HTML page data is not returned because it is too large.
type DbRetrieveJobPost struct {
	JobTitle       string                      `json:"title,omitempty"`
	Description    string                      `json:"description,omitempty"`
	CodingLanguage datatypes.JSONSlice[string] `json:"codinglanguage,omitempty"`
	Database       datatypes.JSONSlice[string] `json:"database,omitempty"`
	CompanyName    string                      `json:"companyname,omitempty"`
	Pay            datatypes.JSONSlice[string] `json:"pay,omitempty"`
	Location       datatypes.JSONSlice[string] `json:"location,omitempty"`
	Links          datatypes.JSONSlice[string] `json:"link,omitempty"`
}

type DbLoadJobPost struct {
	gorm.Model
	JobTitle       string
	Description    string
	CodingLanguage datatypes.JSONSlice[string]
	Database       datatypes.JSONSlice[string]
	CompanyName    string
	Pay            datatypes.JSONSlice[string]
	Location       datatypes.JSONSlice[string]
	Links          datatypes.JSONSlice[string]
}
