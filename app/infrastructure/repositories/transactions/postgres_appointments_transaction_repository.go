package repositories

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lib/pq"

	"stupid-vacancy/app/application/interfaces"
	"stupid-vacancy/app/domain/company/entities"
	"stupid-vacancy/app/infrastructure/database"

	shared "stupid-vacancy/app/application/error"
)

type PostgresCompaniesTransactionRepository struct {
	db *sql.DB
}

func New() *PostgresCompaniesTransactionRepository {
	return &PostgresCompaniesTransactionRepository{
		db: database.GetDB(),
	}
}

func (repository *PostgresCompaniesTransactionRepository) Run(
	cm *entities.CompanyMetadata,
	c *entities.Company,
	indexingEngine *interfaces.IndexingEngineInterface) (*entities.CompanyMetadata, *shared.AppError) {

	tx, err := repository.db.Begin()

	if err != nil {
		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer tx.Rollback()

	_, err = tx.Exec(
		`INSERT INTO companies (id, name) VALUES ($1, $2)`,
		c.Id,
		c.Name,
	)

	hasError := fail(err)

	if hasError != nil {
		return nil, hasError
	}

	_, err = tx.Exec(
		`INSERT INTO "companiesMetadata" (id, cnpj, "companyId", nome, fantasia, tipo, "atividadePrin", abertura, email, telefone, cep, numero, complemento) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		cm.Id,
		cm.CNPJ,
		cm.CompanyId,
		cm.Nome,
		cm.Fantasia,
		cm.Tipo,
		cm.AtividadePrin,
		cm.Abertura,
		cm.Email,
		cm.Telefone,
		cm.CEP,
		cm.Numero,
		cm.Complemento,
	)

	hasError = fail(err)

	if hasError != nil {
		return nil, hasError
	}

	success := (*indexingEngine).Push(
		cm.Id.String(),
		cm.Fantasia,
		cm.Nome,
		cm.Email,
		cm.CNPJ,
	)

	if !success {

		fmt.Println("Error during engine indexing")

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema na indexação de resultados", StatusCode: http.StatusInternalServerError}

	}

	if err = tx.Commit(); err != nil {

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	return cm, nil
}

func fail(err error) *shared.AppError {

	if err != nil {

		fmt.Println("err: ", err)

		pgErr, ok := err.(*pq.Error)

		// id is a unique constraint (?)
		if ok && pgErr.Code == "23505" { // Unique contraint validation error [custom error to client]
			return &shared.AppError{Err: pgErr, Message: "Empresa já existe", StatusCode: http.StatusBadRequest}
		}

		return &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	return nil
}
