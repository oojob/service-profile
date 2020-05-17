package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	protobuf "github.com/oojob/protobuf"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// EmailModel email
type EmailModel struct {
	Email       string                     `bson:"email,omitempty"`
	EmailStatus protobuf.Email_EmailStatus `bson:"verified,omitempty"`
	Show        bool                       `bson:"show,omitempty"`
}

// EducationModel education
type EducationModel struct {
	Education string `bson:"education,omitempty"`
	Show      bool   `bson:"show,omitempty"`
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
	PostalCode int64  `bson:"postal_code,omitempty"`
	Street     string `bson:"street,omitempty"`
}

// MetadataModel metadata
type MetadataModel struct {
	CreatedAt     timestamp.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt     timestamp.Timestamp `bson:"updated_at,omitempty"`
	PublishedDate timestamp.Timestamp `bson:"published_date,omitempty"`
	EndDate       timestamp.Timestamp `bson:"end_date,omitempty"`
	LastActive    timestamp.Timestamp `bson:"last_active,omitempty"`
}

// TokenDetails for token data
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// Profile profile model schema
type Profile struct {
	ID              primitive.ObjectID   `bson:"_id,omitempty"`
	Identity        IdentifierModel      `bson:"identity,omitempty"`
	GivenName       string               `bson:"given_name,omitempty"`
	MiddleName      string               `bson:"middle_name,omitempty"`
	FamilyName      string               `bson:"family_name,omitempty"`
	Username        string               `bson:"username,omitempty"`
	Email           EmailModel           `bson:"email,omitempty"`
	Gender          string               `bson:"gender,omitempty"`
	Birthdate       timestamp.Timestamp  `bson:"birthdate,omitempty"`
	CurrentPosition string               `bson:"current_position,omitempty"`
	Education       EducationModel       `bson:"education,omitempty"`
	Address         AddressModel         `bson:"address,omitempty"`
	Security        ProfileSecutiryModel `bson:"security,omitempty"`
	Metadata        MetadataModel        `bson:"metadata,omitempty"`
}
