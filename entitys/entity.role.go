package entitys

import (
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type EntityRole interface {
	EntityCreate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelRole, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaRole) (*models.ModelRole, schemas.SchemaDatabaseError)
}
