package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/lib/pq"

	// "rhSystem_server/app/application/interfaces"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/infrastructure/database"

	shared "rhSystem_server/app/application/error"
)

type PostgresAppointmentsTransactionRepository struct {
	db *sql.DB
}

func New() *PostgresAppointmentsTransactionRepository {
	return &PostgresAppointmentsTransactionRepository{
		db: database.GetDB(),
	}
}

func (repository *PostgresAppointmentsTransactionRepository) Run(c *entities.Candidate, a *entities.Appointment) (bool, *shared.AppError) {

	tx, err := repository.db.Begin()

	if err != nil {
		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer tx.Rollback()

	_, err = tx.Exec(
		`INSERT INTO candidates (id, name, email, phone) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING`,
		c.Id,
		c.Name,
		c.Email,
		c.Phone,
	)

	hasError := fail(err)

	if hasError != nil {
		return false, hasError
	}

	// VERIFICA SE SLOT FOI UTILIZADO. SE SIM, VERIFICA STATUS:
	// "canceled": proceder
	// "pending": verificar timestamp (mais de 15 minutos proceder) * caso algum usuário tenha criado um appointment e não tenha confirmado no
	// prazo combinado com o link enviado pro email
	// "confirmed": interromper execução e retornar resposta coerente para o usuário
	result, err := tx.Exec(
		`INSERT INTO appointments (id, datetime, slot, candidate, status)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (datetime) DO UPDATE
		SET candidate = $6, status = $7, updated_at = $8
		WHERE (
			appointments.status = 3
			OR (
				appointments.status = 1
				AND appointments.created_at >= NOW() - INTERVAL '15 minutes'
				OR appointments.updated_at >= NOW() - INTERVAL '15 minutes'
			)
		);`,
		a.Id,
		a.Datetime,
		a.Slot,
		c.Id,
		a.Status,
		c.Id,
		a.Status,
		time.Now(),
	)

	hasError = fail(err)

	// what happens if appointment scheduling is not succeeded

	if hasError != nil {
		return false, hasError
	}

	if err = tx.Commit(); err != nil {
		fmt.Println("erro no commit")
		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	rowsAffected, _ := result.RowsAffected()

	fmt.Println("rowsAffected: ", rowsAffected)

	return (rowsAffected > 0), nil
}

func fail(err error) *shared.AppError {

	if err != nil {

		fmt.Println("err: ", err)

		pgErr, ok := err.(*pq.Error)

		if ok && pgErr.Code == "23505" { // Unique contraint validation error [custom error to client]
			return &shared.AppError{Err: pgErr, Message: "Registro já existe", StatusCode: http.StatusBadRequest}
		}

		return &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	return nil
}
