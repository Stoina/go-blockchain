package repository

import (
	"github.com/Stoina/go-blockchain/cmd/bc-backend/repository/result"
	db "github.com/Stoina/go-database"
)

// BlockRepository exported
// ...
type BlockRepository struct {
}

// NewBlockRepository exported
// ...
func NewBlockRepository() *BlockRepository {
	return &BlockRepository{}
}

// Create exported
// ...
func (br *BlockRepository) Create(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// Update exported
// ...
func (br *BlockRepository) Update(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// Delete exported
// ...
func (br *BlockRepository) Delete(jsonString string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadAll exported
// ...
func (br *BlockRepository) ReadAll(db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadByID exported
// ...
func (br *BlockRepository) ReadByID(id string, db *db.Connection) *result.RepositoryResult {
	return nil
}

// ReadByProperty exported
// ...
func (br *BlockRepository) ReadByProperty(propertyName string, value interface{}, db *db.Connection) *result.RepositoryResult {
	return nil
}
