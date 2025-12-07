package repositories

import (
	"errors"

	"github.com/CaioDS/fli/internal/infrastructure/context"
)

type LocalRepository struct {
	localDbContext *context.LocalDbContext
}

func NewLocalRepository(localDbContext *context.LocalDbContext) *LocalRepository {
	return &LocalRepository{
		localDbContext: localDbContext,
	}
}

func (r *LocalRepository) GetStoredDefaultPath() (string, error) {
	path, err := r.localDbContext.Get("config", []byte("defaultPath"))
	if err != nil {
		return "", err
	}

	return string(path), nil
}

func (r *LocalRepository) SaveVersionRegistry(version string, path string) error {
	err := r.localDbContext.CreateBucket("versions")
	if err != nil {
		return errors.New("Failed to create versions bucket")
	}

	err = r.localDbContext.Put("versions", []byte(version), []byte(path))
	if err != nil {
		return errors.New("failed to save version occurency on database!")
	}

	return nil
}

func (r *LocalRepository) SaveDefaultPath(path string) error {
	err := r.localDbContext.CreateBucket("config")
	if err != nil {
		return errors.New("Failed to create config bucket")
	}

	err = r.localDbContext.Put("config", []byte("defaultPath"), []byte(path))
	if err != nil {
		return errors.New("failed to save default path on database!")
	}

	return nil
}