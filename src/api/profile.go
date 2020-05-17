package api

import (
	"context"
	"fmt"

	protobuf "github.com/oojob/protobuf"
	profile "github.com/oojob/protorepo-profile-go"
	model "github.com/oojob/service-profile/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getIdentity(identity *protobuf.Identifier) model.IdentifierModel {
	identityModel := model.IdentifierModel{
		Identifier:                identity.GetIdentifier(),
		Name:                      identity.GetName(),
		AlternateName:             identity.GetAlternateName(),
		Type:                      identity.GetType(),
		AdditionalType:            identity.GetAdditionalType(),
		Description:               identity.GetDescription(),
		DisambiguatingDescription: identity.GetDisambiguatingDescription(),
		Headline:                  identity.GetHeadline(),
		Slogan:                    identity.GetSlogan(),
	}

	return identityModel
}

func getEmail(email *protobuf.Email) model.EmailModel {
	emailModel := model.EmailModel{
		Email:       email.GetEmail(),
		Show:        email.GetShow(),
		EmailStatus: email.GetStatus(),
	}

	return emailModel
}

func getEducation(education *profile.Education) model.EducationModel {
	educationModel := model.EducationModel{
		Education: education.GetEducation(),
		Show:      education.GetShow(),
	}

	return educationModel
}

func getAddress(address *protobuf.Address) model.AddressModel {
	addressModel := model.AddressModel{
		Country:    address.GetCountry(),
		Locality:   address.GetLocality(),
		Region:     address.GetRegion(),
		PostalCode: address.GetPostalCode(),
		Street:     address.GetStreet(),
	}

	return addressModel
}

func getSecurity(security *profile.ProfileSecurity) model.ProfileSecutiryModel {
	securityModel := model.ProfileSecutiryModel{
		Password:     security.GetPassword(),
		PasswordSalt: security.GetPasswordSalt(),
		PasswordHash: security.GetPasswordHash(),
		Code:         security.GetCode(),
		CodeType:     security.GetCodeType(),
		AccountType:  "basic", // profile.ProfileSecurity_AccountType_name[security.GetAccountType()]
		Verified:     security.GetVerified(),
	}

	return securityModel
}

func getProfile(in *profile.Profile) model.Profile {
	profileModel := model.Profile{
		Identity:        getIdentity(in.GetIdentity()),
		GivenName:       in.GetGivenName(),
		MiddleName:      in.GetMiddleName(),
		FamilyName:      in.GetFamilyName(),
		Username:        in.GetUsername(),
		Email:           getEmail(in.GetEmail()),
		Gender:          in.GetGender(),
		CurrentPosition: in.GetCurrentPosition(),
		Education:       getEducation(in.GetEducation()),
		Address:         getAddress(in.GetAddress()),
		Security:        getSecurity(in.GetSecurity()),
	}

	return profileModel
}

// CreateProfile cretaes a profile
func (c *API) CreateProfile(ctx context.Context, in *profile.Profile) (*protobuf.Id, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	context := c.App.NewContext()
	profileData := getProfile(in)

	res, err := context.CreateProfile(&profileData)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Document: %v", err),
		)
	}

	return &protobuf.Id{
		Id: res,
	}, nil
}

// ConfirmProfile :- confirm acccount
func (c *API) ConfirmProfile(ctx context.Context, in *profile.ConfirmProfileRequest) (*protobuf.DefaultResponse, error) {
	return nil, nil
}

// ReadProfile :- read profile
func (c *API) ReadProfile(ctx context.Context, in *profile.ReadProfileRequest) (*profile.Profile, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	context := c.App.NewContext()

	id, err := primitive.ObjectIDFromHex(in.GetAccountId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	_, err = context.ReadProfile(&id)
	if err != nil {
		// span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Invalid Email Value: %v", err),
		)
	}

	return nil, nil
}

// UpdateProfile :- update account
func (c *API) UpdateProfile(ctx context.Context, in *profile.Profile) (*protobuf.Id, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	context := c.App.NewContext()
	profileData := getProfile(in)

	res, err := context.UpdateProfile(&profileData)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Document: %v", err),
		)
	}

	return &protobuf.Id{
		Id: res,
	}, nil
}

// ValidateUsername :- validate username
func (c *API) ValidateUsername(ctx context.Context, in *profile.ValidateUsernameRequest) (*protobuf.DefaultResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	context := c.App.NewContext()

	success, err := context.ValidateUsername(in.Username)
	if err != nil {
		// span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Invalid Username Value: %v", err),
		)
	}

	return &protobuf.DefaultResponse{
		Status: success,
		Code:   int64(codes.OK),
		Error:  "",
	}, nil
}

// ValidateEmail :- validate email
func (c *API) ValidateEmail(ctx context.Context, in *profile.ValidateEmailRequest) (*protobuf.DefaultResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	context := c.App.NewContext()

	success, err := context.ValidateEmail(in.Email)
	if err != nil {
		// span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Invalid Email Value: %v", err),
		)
	}

	return &protobuf.DefaultResponse{
		Status: success,
		Code:   int64(codes.OK),
		Error:  "",
	}, nil
}

// Auth :- authenticates and generate jwt token
func (c *API) Auth(ctx context.Context, in *profile.AuthRequest) (*profile.AuthResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	context := c.App.NewContext()

	token, err := context.Auth(in)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Invalid Email Value: %v", err),
		)
	}

	return &profile.AuthResponse{
		AccessToken: token,
		Valid:       true,
	}, nil
}

// Logout help's us to logout from site
func (c *API) Logout(ctx context.Context, in *profile.TokenRequest) (*protobuf.DefaultResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return nil, nil
}

// RefreshToken help's us to refresh the access token without signing out
func (c *API) RefreshToken(ctx context.Context, in *profile.TokenRequest) (*profile.AuthResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return nil, nil
}

// VerifyToken help's us to verify the auth token
func (c *API) VerifyToken(ctx context.Context, in *profile.TokenRequest) (*profile.AccessDetails, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	return nil, nil
}

// Check check the context
func (c *API) Check(ctx context.Context, in *protobuf.HealthCheckRequest) (*protobuf.HealthCheckResponse, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if in.Service == "" {
		// check the server overall health status.
		return &protobuf.HealthCheckResponse{
			Status: protobuf.HealthCheckResponse_NOT_SERVING,
		}, nil
	}

	statusMap := make(map[string]protobuf.HealthCheckResponse_ServingStatus)
	if status, ok := statusMap[in.Service]; ok {
		return &protobuf.HealthCheckResponse{
			Status: status,
		}, nil
	}

	return nil, status.Errorf(codes.Internal, "unknown service")
}

// Watch watch the serving status
func (c *API) Watch(_ *protobuf.HealthCheckRequest, stream profile.ProfileService_WatchServer) error {
	stream.Send(&protobuf.HealthCheckResponse{
		Status: protobuf.HealthCheckResponse_SERVING,
	})

	return nil
}
