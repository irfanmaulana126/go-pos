package entitys

import (
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type EntityProduct interface {
	EntityCreate(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityResult(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityResults() (*[]models.ModelProduct, schemas.SchemaDatabaseError)
	EntityDelete(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityUpdate(input *schemas.SchemaProduct) (*models.ModelProduct, schemas.SchemaDatabaseError)
	EntityProductByOutlet(input *schemas.SchemaProduct) (*[]models.ModelProduct, schemas.SchemaDatabaseError)
}
