package entitys

import (
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type EntityTransaction interface {
	EntityCreate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError)
}
