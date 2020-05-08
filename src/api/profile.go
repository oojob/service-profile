package api

import (
	"context"

	protobuf "github.com/oojob/protobuf"
	profile "github.com/oojob/protorepo-profile-go"
)

func (c *API) CreateProfile(ctx context.Context, in *profile.Profile) (*profile.Profile, error) {
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
