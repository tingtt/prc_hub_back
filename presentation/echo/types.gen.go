// Package echo provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package echo

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
)

const (
	BearerScopes = "Bearer.Scopes"
)

// CreateEventBody defines model for CreateEventBody.
type CreateEventBody struct {
	Completed   *bool                  `json:"completed,omitempty"`
	Datetimes   *[]CreateEventDatetime `json:"datetimes,omitempty"`
	Description *string                `json:"description,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        string                 `json:"name"`
	Published   *bool                  `json:"published,omitempty"`
}

// CreateEventDatetime defines model for CreateEventDatetime.
type CreateEventDatetime struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

// CreateEventDocumentBody defines model for CreateEventDocumentBody.
type CreateEventDocumentBody struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// CreateUserBody defines model for CreateUserBody.
type CreateUserBody struct {
	Email          openapi_types.Email `json:"email"`
	GithubUsername *string             `json:"github_username,omitempty"`
	Name           string              `json:"name"`
	Password       string              `json:"password"`
	TwitterId      *string             `json:"twitter_id,omitempty"`
}

// Event defines model for Event.
type Event struct {
	Completed   bool             `json:"completed"`
	Datetimes   []EventDatetime  `json:"datetimes"`
	Description *string          `json:"description,omitempty"`
	Documents   *[]EventDocument `json:"documents,omitempty"`
	Id          int64            `json:"id"`
	Location    *string          `json:"location,omitempty"`
	Name        string           `json:"name"`
	Published   bool             `json:"published"`
	User        *User            `json:"user,omitempty"`
}

// EventDatetime defines model for EventDatetime.
type EventDatetime struct {
	End   string `json:"end"`
	Start string `json:"start"`
}

// EventDocument defines model for EventDocument.
type EventDocument struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

// LoginBody defines model for LoginBody.
type LoginBody struct {
	Email    openapi_types.Email `json:"email"`
	Password string              `json:"password"`
}

// Token defines model for Token.
type Token struct {
	Token string `json:"token"`
}

// UpdateEventBody defines model for UpdateEventBody.
type UpdateEventBody struct {
	Completed   *bool                  `json:"completed,omitempty"`
	Datetimes   *[]CreateEventDatetime `json:"datetimes,omitempty"`
	Description *string                `json:"description,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	Published   *bool                  `json:"published,omitempty"`
}

// UpdateEventDocumentBody defines model for UpdateEventDocumentBody.
type UpdateEventDocumentBody struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// UpdateUserBody defines model for UpdateUserBody.
type UpdateUserBody struct {
	Email               *string `json:"email,omitempty"`
	GithubUsername      *string `json:"github_username,omitempty"`
	Manage              *bool   `json:"manage,omitempty"`
	Name                *string `json:"name,omitempty"`
	Password            *string `json:"password,omitempty"`
	PostEventAvailabled *bool   `json:"post_event_availabled,omitempty"`
	TwitterId           *string `json:"twitter_id,omitempty"`
}

// User defines model for User.
type User struct {
	Admin               bool                `json:"admin"`
	Email               openapi_types.Email `json:"email"`
	GithubUsername      *string             `json:"github_username,omitempty"`
	Id                  int64               `json:"id"`
	Manage              bool                `json:"manage"`
	Name                string              `json:"name"`
	PostEventAvailabled bool                `json:"post_event_availabled"`
	TwitterId           *string             `json:"twitter_id,omitempty"`
}

// UserWithToken defines model for UserWithToken.
type UserWithToken struct {
	Admin               bool                `json:"admin"`
	Email               openapi_types.Email `json:"email"`
	GithubUsername      *string             `json:"github_username,omitempty"`
	Id                  int64               `json:"id"`
	Name                string              `json:"name"`
	PostEventAvailabled bool                `json:"post_event_availabled"`
	Token               string              `json:"token"`
	TwitterId           *string             `json:"twitter_id,omitempty"`
}

// DocumentId defines model for document_id.
type DocumentId = int64

// Embed defines model for embed.
type Embed = []string

// Id defines model for id.
type Id = int64

// Location defines model for location.
type Location = string

// LocationContain defines model for location_contain.
type LocationContain = string

// Name defines model for name.
type Name = string

// NameContain defines model for name_contain.
type NameContain = string

// GetEventsParams defines parameters for GetEvents.
type GetEventsParams struct {
	Name            *Name            `form:"name,omitempty" json:"name,omitempty"`
	NameContain     *NameContain     `form:"name_contain,omitempty" json:"name_contain,omitempty"`
	Location        *Location        `form:"location,omitempty" json:"location,omitempty"`
	LocationContain *LocationContain `form:"location_contain,omitempty" json:"location_contain,omitempty"`
	Embed           *Embed           `form:"embed,omitempty" json:"embed,omitempty"`
}

// GetEventsParamsEmbed defines parameters for GetEvents.
type GetEventsParamsEmbed string

// GetEventsIdParams defines parameters for GetEventsId.
type GetEventsIdParams struct {
	Embed *Embed `form:"embed,omitempty" json:"embed,omitempty"`
}

// GetEventsIdParamsEmbed defines parameters for GetEventsId.
type GetEventsIdParamsEmbed string

// GetEventsIdDocumentsParams defines parameters for GetEventsIdDocuments.
type GetEventsIdDocumentsParams struct {
	Name        *Name        `form:"name,omitempty" json:"name,omitempty"`
	NameContain *NameContain `form:"name_contain,omitempty" json:"name_contain,omitempty"`
}

// PostEventsJSONRequestBody defines body for PostEvents for application/json ContentType.
type PostEventsJSONRequestBody = CreateEventBody

// PatchEventsIdJSONRequestBody defines body for PatchEventsId for application/json ContentType.
type PatchEventsIdJSONRequestBody = UpdateEventBody

// PostEventsIdDocumentsJSONRequestBody defines body for PostEventsIdDocuments for application/json ContentType.
type PostEventsIdDocumentsJSONRequestBody = CreateEventDocumentBody

// PatchEventsIdDocumentsDocumentIdJSONRequestBody defines body for PatchEventsIdDocumentsDocumentId for application/json ContentType.
type PatchEventsIdDocumentsDocumentIdJSONRequestBody = UpdateEventDocumentBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = CreateUserBody

// PostUsersSignInJSONRequestBody defines body for PostUsersSignIn for application/json ContentType.
type PostUsersSignInJSONRequestBody = LoginBody

// PatchUsersIdJSONRequestBody defines body for PatchUsersId for application/json ContentType.
type PatchUsersIdJSONRequestBody = UpdateUserBody
