package entitys

import (
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type EntityUser interface {
	EntityRegister(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError)
	EntityLogin(input *schemas.SchemaUser) (*models.ModelUser, schemas.SchemaDatabaseError)
}
