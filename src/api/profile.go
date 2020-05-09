package api

import (
	"context"
	"fmt"

	protobuf "github.com/oojob/protobuf"
	profile "github.com/oojob/protorepo-profile-go"
	model "github.com/oojob/service-profile/src/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateProfile cretaes a profile
func (c *API) CreateProfile(ctx context.Context, in *profile.Profile) (*profile.Profile, error) {
	identity := in.GetIdentity()

	profileData := model.Profile{
		GivenName: "dododucck",
		Identity: model.IdentifierModel{
			Identifier:                identity.GetIdentifier(),
			Name:                      identity.GetName(),
			AlternateName:             identity.GetAlternateName(),
			Type:                      identity.GetType(),
			AdditionalType:            identity.GetAdditionalType(),
			Description:               identity.GetDescription(),
			DisambiguatingDescription: identity.GetDisambiguatingDescription(),
			Headline:                  identity.GetHeadline(),
			Slogan:                    identity.GetSlogan(),
		},
	}
	// fmt.Printf(profileData.GivenName)

	context := c.App.NewContext()
	_, err := context.CreateProfile(&profileData)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Invalid Document: %v", err),
		)
	}

	return nil, nil
}

func (c *API) ConfirmProfile(ctx context.Context, in *profile.ConfirmProfileRequest) (*protobuf.DefaultResponse, error) {
	return nil, nil
}

func (c *API) ReadProfile(ctx context.Context, in *profile.ReadProfileRequest) (*profile.Profile, error) {
	return nil, nil
}

func (c *API) UpdateProfile(ctx context.Context, in *profile.Profile) (*profile.Profile, error) {
	return nil, nil
}

func (c *API) ValidateUsername(ctx context.Context, in *profile.ValidateUsernameRequest) (*protobuf.DefaultResponse, error) {
	return nil, nil
}

func (c *API) ValidateEmail(ctx context.Context, in *profile.ValidateEmailRequest) (*protobuf.DefaultResponse, error) {
	return nil, nil
}

// Check check the context
func (c *API) Check(ctx context.Context, in *protobuf.HealthCheckRequest) (*protobuf.HealthCheckResponse, error) {
	return &protobuf.HealthCheckResponse{
		Status: protobuf.HealthCheckResponse_SERVING,
	}, nil
}

// Watch watch the serving status
func (c *API) Watch(_ *protobuf.HealthCheckRequest, stream profile.ProfileService_WatchServer) error {
	stream.Send(&protobuf.HealthCheckResponse{
		Status: protobuf.HealthCheckResponse_SERVING,
	})

	return nil
}
