package entities

import "mime/multipart"

type RequirementFile struct {
	RequirementId int                   `form:"requirement_id"`
	UserId        int                   `form:"user_id"`
	File          *multipart.FileHeader `form:"file"`
}
