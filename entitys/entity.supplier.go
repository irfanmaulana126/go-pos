package entitys

import (
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type EntitySupplier interface {
	EntityCreate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaSupplier) (*models.ModelSupplier, schemas.SchemaDatabaseError)
}
