package models

type GetProfileRequest struct {
	UUID string `url:"uuid" validate:"required"`
}
