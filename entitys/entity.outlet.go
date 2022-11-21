package entitys

import (
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type EntityOutlet interface {
	EntityCreate(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaOutlet) (*models.ModelOutlet, schemas.SchemaDatabaseError)
}
