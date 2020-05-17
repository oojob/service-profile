package db

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	profile_repo "github.com/oojob/protorepo-profile-go"
	"github.com/oojob/service-profile/src/model"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

var (
	accessSecret  = []byte(viper.GetString("accesssecret"))
	refreshSecret = []byte(viper.GetString("refreshsecret"))
)

// PersonData for storing data
type PersonData struct {
	Authorized  bool
	AccessUUID  string
	Username    string
	Email       string
	UserID      string
	Identifier  string
	AccountType string
	Verified    bool
	Exp         int64
}

// AccessTokenClaim for storing jwt claim data
type AccessTokenClaim struct {
	Person *PersonData
	jwt.StandardClaims
}

// RefreshData used for storing refresh data
type RefreshData struct {
	Sub         string
	RefreshUUID string
	UserID      string
	Exp         int64
}

// RefreshTokenClaim :- claims
type RefreshTokenClaim struct {
	Refresh *RefreshData
	jwt.StandardClaims
}

// Authable :- interface for auth
type Authable interface {
	Decode(token string) (*AccessTokenClaim, error)
	Encode(profile *model.Profile) (string, error)
}

// Encode a claim into a JWT
func (db *Database) Encode(profile *model.Profile) (*profile_repo.AuthResponse, error) {
	ts, err := db.CreateToken(profile)
	if err != nil {
		return nil, err
	}
	userid := profile.ID.Hex()

	at := time.Unix(ts.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(ts.RtExpires, 0)
	now := time.Now()
	errAccess := db.redis.Set(ts.AccessUUID, userid, at.Sub(now)).Err()
	if errAccess != nil {
		return nil, errAccess
	}
	errRefresh := db.redis.Set(ts.RefreshUUID, userid, rt.Sub(now)).Err()
	if errRefresh != nil {
		return nil, errRefresh
	}

	return &profile_repo.AuthResponse{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
		Valid:        true,
	}, nil
}

// Decode a token string into a token object
// func (db *Database) Decode(tokenString string) (*CustomClaims, error) {

// 	// Parse the token
// 	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return key, nil
// 	})

// 	// Validate the token and return the custom claims
// 	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
// 		return claims, nil
// 	} else {
// 		return nil, err
// 	}
// }

// CreateToken generates anew token
func (db *Database) CreateToken(profile *model.Profile) (*model.TokenDetails, error) {
	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	// store the error
	var err error

	//Creating Access Token
	Person := PersonData{}
	Person.Authorized = true
	Person.AccessUUID = td.AccessUUID
	Person.Username = profile.Username
	Person.Email = profile.Email.Email
	Person.UserID = profile.ID.Hex()
	Person.Identifier = profile.Identity.Identifier
	Person.AccountType = profile.Security.AccountType
	Person.Verified = profile.Security.Verified
	Person.Exp = td.AtExpires
	authTokenClaim := AccessTokenClaim{
		&Person,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "account.service",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, authTokenClaim)
	td.AccessToken, err = token.SignedString(accessSecret)
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := RefreshData{}
	rtClaims.RefreshUUID = td.RefreshUUID
	rtClaims.UserID = profile.ID.Hex()
	rtClaims.Exp = td.RtExpires
	refreshTokenClaim := RefreshTokenClaim{
		&rtClaims,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "account.service",
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim)
	td.RefreshToken, err = rt.SignedString(refreshSecret)
	if err != nil {
		return nil, err
	}

	return td, nil
}
