package services

import (
	"github.com/irfanmaulana126/go-pos/entitys"
	"github.com/irfanmaulana126/go-pos/models"
	"github.com/irfanmaulana126/go-pos/schemas"
)

type serviceTransaction struct {
	transaction entitys.EntityTransaction
}

func NewServiceTransaction(transaction entitys.EntityTransaction) *serviceTransaction {
	return &serviceTransaction{transaction: transaction}
}

/**
* ==========================================
* Service Create New Transaction Teritory
*===========================================
 */

func (s *serviceTransaction) EntityCreate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction schemas.SchemaTransaction
	transaction.CustomerID = input.CustomerID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products

	res, err := s.transaction.EntityCreate(&transaction)
	return res, err
}

/**
* ============================================
* Service Results All Transaction Teritory
*=============================================
 */

func (s *serviceTransaction) EntityResults() (*[]models.ModelTransaction, schemas.SchemaDatabaseError) {
	res, err := s.transaction.EntityResults()
	return res, err
}

/**
* ==========================================
* Service Result Transaction By ID Teritory
*===========================================
 */

func (s *serviceTransaction) EntityResult(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction schemas.SchemaTransaction
	transaction.ID = input.ID

	res, err := s.transaction.EntityResult(&transaction)
	return res, err
}

/**
* ==========================================
* Service Delete Transaction By ID Teritory
*===========================================
 */

func (s *serviceTransaction) EntityDelete(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction schemas.SchemaTransaction
	transaction.ID = input.ID

	res, err := s.transaction.EntityDelete(&transaction)
	return res, err
}

/**
* ==========================================
* Service Update Transaction By ID Teritory
*===========================================
 */

func (s *serviceTransaction) EntityUpdate(input *schemas.SchemaTransaction) (*models.ModelTransaction, schemas.SchemaDatabaseError) {
	var transaction schemas.SchemaTransaction
	transaction.CustomerID = input.CustomerID
	transaction.OutletID = input.OutletID
	transaction.Products = input.Products

	res, err := s.transaction.EntityUpdate(&transaction)
	return res, err
}
