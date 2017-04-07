package models

import (
	"time"
)

// Meta has metadata
type Meta struct {
	ID          uint `gorm:"primary_key"`
	Timestamp   time.Time
	Family      string
	Release     string
	Definitions []Definition
}

//TODO ALAS

// Definition : >definitions>definition
type Definition struct {
	ID     uint `gorm:"primary_key"`
	MetaID uint `json:"-" xml:"-"`

	Title         string
	Description   string
	Advisory      Advisory
	AffectedPacks []Package
	References    []Reference
}

// Package affedted
type Package struct {
	ID           uint `gorm:"primary_key"`
	DefinitionID uint `json:"-" xml:"-"`

	Name    string
	Version string // affected earlier than this version
}

// Reference : >definitions>definition>metadata>reference
type Reference struct {
	ID           uint `gorm:"primary_key"`
	DefinitionID uint `json:"-" xml:"-"`

	Source string
	RefID  string
	RefURL string
}

// Advisory : >definitions>definition>metadata>advisory
type Advisory struct {
	ID           uint `gorm:"primary_key"`
	DefinitionID uint `json:"-" xml:"-"`

	Severity        string
	Cves            []Cve
	Bugzillas       []Bugzilla
	AffectedCPEList []Cpe
}

// Cve : >definitions>definition>metadata>advisory>cve
// RedHat OVAL
type Cve struct {
	ID         uint `gorm:"primary_key"`
	AdvisoryID uint `json:"-" xml:"-"`

	CveID  string
	Cvss2  string
	Cvss3  string
	Cwe    string
	Href   string
	Public string
}

// Bugzilla : >definitions>definition>metadata>advisory>bugzilla
// RedHat OVAL
type Bugzilla struct {
	ID         uint `gorm:"primary_key"`
	AdvisoryID uint `json:"-" xml:"-"`

	BugzillaID string
	URL        string
	Title      string
}

// Cpe : >definitions>definition>metadata>advisory>affected_cpe_list
type Cpe struct {
	ID         uint `gorm:"primary_key"`
	AdvisoryID uint `json:"-" xml:"-"`

	Cpe string
}
