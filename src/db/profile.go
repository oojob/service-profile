package db

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"io"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	profile "github.com/oojob/protorepo-profile-go"
	"github.com/oojob/service-profile/src/model"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// EncodeToString :- encode string to form a valid code
func EncodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// HashPassword :- encrypt password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash :- check for hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidateEmail :- check for existing email value
func (db *Database) ValidateEmail(email string) (bool, error) {
	profileCollection := db.mongo.Collection("profile")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return false, err
	}
	defer session.EndSession(context.Background())

	var profile *model.Profile
	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result := profileCollection.FindOne(sessionContext, &bson.M{"email.email": email})
		if err := result.Decode(&profile); err != nil {
			return false, err
		}

		return false, nil
	})

	if profile != nil {
		return true, nil
	}
	return false, err
}

// ValidateUsername :- check for existing email value
func (db *Database) ValidateUsername(username string) (bool, error) {
	profileCollection := db.mongo.Collection("profile")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return false, err
	}
	defer session.EndSession(context.Background())

	var profile *model.Profile
	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result := profileCollection.FindOne(sessionContext, &bson.M{"username": username})
		if err := result.Decode(&profile); err != nil {
			return false, err
		}

		return false, nil
	})

	if profile != nil {
		return true, nil
	}
	return false, err
}

// CreateProfile create profile entity
func (db *Database) CreateProfile(in *model.Profile) (string, error) {
	var inerstionID string

	companyCollection := db.mongo.Collection("profile")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(context.Background())

	// modify input request object to fill few fields
	u1 := uuid.NewV4()

	in.Identity.Identifier = u1.String()
	in.Security.Code = EncodeToString(6)
	in.Security.CodeType = "ACCOUNT_CONFIRM"
	in.Security.Verified = false

	in.Metadata = model.MetadataModel{
		CreatedAt:     *ptypes.TimestampNow(),
		UpdatedAt:     *ptypes.TimestampNow(),
		PublishedDate: *ptypes.TimestampNow(),
		EndDate:       *ptypes.TimestampNow(),
		LastActive:    *ptypes.TimestampNow(),
	}

	// hash the password
	hashedPass, err := HashPassword(in.Security.Password)
	if err != nil {
		err := status.Error(codes.DataLoss, fmt.Sprintf("error hashing password: %v", err))
		return "", err
	}
	in.Security.Password = hashedPass
	in.Security.PasswordHash = hashedPass

	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result, err := companyCollection.InsertOne(sessionContext, in)
		if err != nil {
			return "", err
		}

		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			inerstionID = oid.Hex()
		}

		return "", nil
	})

	return inerstionID, err
}

// UpdateProfile create profile entity
func (db *Database) UpdateProfile(in *model.Profile) (string, error) {
	companyCollection := db.mongo.Collection("profile")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(context.Background())

	in.Metadata = model.MetadataModel{
		UpdatedAt:     *ptypes.TimestampNow(),
		PublishedDate: *ptypes.TimestampNow(),
		EndDate:       *ptypes.TimestampNow(),
		LastActive:    *ptypes.TimestampNow(),
	}

	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		_, err := companyCollection.UpdateOne(sessionContext, &bson.M{"_id": in.ID}, in)
		if err != nil {
			return "", err
		}

		return in.ID.Hex(), nil
	})

	return in.ID.Hex(), err
}

// Auth :- authentication
func (db *Database) Auth(in *profile.AuthRequest) (*profile.AuthResponse, error) {
	profileCollection := db.mongo.Collection("profile")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	var profile model.Profile
	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result := profileCollection.FindOne(sessionContext, bson.M{"username": in.GetUsername()})
		if err := result.Decode(&profile); err != nil {
			return nil, err
		}

		return nil, nil
	})

	if profile.Username == "" {
		return nil, errors.New("no account found for this usernames")
	}

	if matched := CheckPasswordHash(in.GetPassword(), profile.Security.Password); !matched {
		return nil, errors.New("password do not matched")
	}

	authResponse, err := db.Encode(&profile)
	if err != nil {
		return nil, err
	}

	return authResponse, nil
}

// VerifyToken help's us to verify and extract identity
func (db *Database) VerifyToken(tokenString string) (*profile.AccessDetails, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return accessSecret, nil
	})

	if claim, ok := token.Claims.(*AccessTokenClaim); !ok && !token.Valid {
		return &profile.AccessDetails{
			AccessUuid:  claim.Person.AccessUUID,
			AccountType: claim.Person.AccountType,
			Authorized:  claim.Person.Authorized,
			Username:    claim.Person.Username,
			Email:       claim.Person.Email,
			UserId:      claim.Person.UserID,
			Verified:    claim.Person.Verified,
			Identifier:  claim.Person.Identifier,
			Exp:         claim.Person.Exp,
		}, err
	}

	return nil, nil
}

// ReadProfile : -read a single profile
func (db *Database) ReadProfile(id *primitive.ObjectID) (*model.Profile, error) {
	profileCollection := db.mongo.Collection("profile")
	session, err := db.mongo.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(context.Background())

	var profile model.Profile
	_, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
		result := profileCollection.FindOne(sessionContext, &bson.M{"_id": id})
		if err := result.Decode(&profile); err != nil {
			return false, err
		}

		return false, nil
	})

	return &profile, err
}
