package load

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// HTML page data is not returned because it is too large.
type LoadDbRetrieveJobPost struct {
	JobTitle       string                      `json:"title" binding:"required"`
	Description    string                      `json:"description" binding:"required"`
	CodingLanguage datatypes.JSONSlice[string] `json:"codinglanguage" binding:"required"`
	Database       datatypes.JSONSlice[string] `json:"database" binding:"required"`
	CompanyName    string                      `json:"companyname" binding:"required"`
	Pay            datatypes.JSONSlice[string] `json:"pay" binding:"required"`
	Location       datatypes.JSONSlice[string] `json:"location" binding:"required"`
	Links          datatypes.JSONSlice[string] `json:"link" binding:"required"`
}

type LoadDbInsertJobPost struct {
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
