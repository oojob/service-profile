package model

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	protobuf "github.com/oojob/protobuf"
	profile "github.com/oojob/protorepo-profile-go"
)

// TimeModel time
type TimeModel struct {
	Opens        time.Time             `bson:"opens,omitempty"`
	Closes       time.Time             `bson:"closes,omitempty"`
	DaysOfWeek   []protobuf.DaysOfWeek `bson:"days_of_week,omitempty"`
	ValidFrom    time.Time             `bson:"valid_from,omitempty"`
	ValidThrough time.Time             `bson:"valid_through,omitempty"`
}

// IdentifierModel identifier
type IdentifierModel struct {
	Identifier                string `bson:"identifier,omitempty"`
	Name                      string `bson:"name,omitempty"`
	AlternateName             string `bson:"alternate_name,omitempty"`
	Type                      string `bson:"type,omitempty"`
	AdditionalType            string `bson:"additional_type,omitempty"`
	Description               string `bson:"description,omitempty"`
	DisambiguatingDescription string `bson:"disambiguating_description,omitempty"`
	Headline                  string `bson:"headline,omitempty"`
	Slogan                    string `bson:"slogan,omitempty"`
}

// ProfileSecutiryModel profile security
type ProfileSecutiryModel struct {
	Password     string `bson:"password,omitempty"`
	PasswordSalt string `bson:"password_salt,omitempty"`
	PasswordHash string `bson:"password_hash,omitempty"`
	Code         string `bson:"code,omitempty"`
	CodeType     string `bson:"code_type,omitempty"`
	AccountType  string `bson:"account_type,omitempty"`
	Verified     bool   `bson:"verified,omitempty"`
}

// AddressModel address
type AddressModel struct {
	Country    string `bson:"country,omitempty"`
	Locality   string `bson:"locality,omitempty"`
	Region     string `bson:"region,omitempty"`
	PostalCode string `bson:"postal_code,omitempty"`
	Street     string `bson:"street,omitempty"`
}

// MetadataModel metadata
type MetadataModel struct {
	CreatedAt     time.Time `bson:"created_at,omitempty"`
	UpdatedAt     time.Time `bson:"updated_at,omitempty"`
	PublishedDate time.Time `bson:"published_date,omitempty"`
	EndDate       time.Time `bson:"end_date,omitempty"`
	LastActive    time.Time `bson:"last_active,omitempty"`
}

// Profile profile
type Profile struct {
	Identity        *protobuf.Identifier `bson:"identity,omitempty"`
	GivenName       string               `bson:"given_name,omitempty"`
	MiddleName      string               `bson:"middle_name,omitempty"`
	FamilyName      string               `bson:"family_name,omitempty"`
	Username        string               `bson:"username,omitempty"`
	Email           *protobuf.Email      `bson:"email,omitempty"`
	Gender          string               `bson:"gender,omitempty"`
	Birthdate       *timestamp.Timestamp `bson:"birthdate,omitempty"`
	CurrentPosition string               `bson:"current_position,omitempty"`
	Education       *profile.Education   `bson:"education,omitempty"`
	Address         *protobuf.Address    `bson:"address,omitempty"`
	Metadata        *protobuf.Metadata   `bson:"metadata,omitempty"`
}
