package database

import (
	"bytes"
	"context"
	"io"

	"github.com/ExpertizeEafit/Api/src/api/domain/requirement/entities"
)

const (
	InsertUserRequirement = `INSERT INTO user_requirement(requirement_id, user_id, status,file) VALUES (?,?,?,?);`
)

func (repository *requirementRepositoryDatabase) UploadRequirement(ctx context.Context, data entities.RequirementFile) error {
	file, err := data.File.Open()
	defer file.Close()
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, file); err != nil {
		return err
	}
	_, err = repository.database.Exec(InsertUserRequirement, data.RequirementId, data.UserId, "PENDING", buff.Bytes())
	if err != nil {
		return err
	}
	return nil
}
