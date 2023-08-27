package repositories

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	shared "rhSystem_server/app/application/error"
	"rhSystem_server/app/domain/appointments/entities"
	"rhSystem_server/app/infrastructure/database"

	"github.com/lib/pq"
)

type PostgresAppointmentsRepository struct {
	db *sql.DB
}

func New() *PostgresAppointmentsRepository {
	return &PostgresAppointmentsRepository{
		db: database.GetDB(),
	}
}

func (repository *PostgresAppointmentsRepository) Create(a *entities.Appointment, candidateId string) (bool, *shared.AppError) {

	// VERIFICA SE SLOT FOI UTILIZADO. SE SIM, VERIFICA STATUS:
	// "canceled": proceder
	// "pending": verificar timestamp (mais de 25 minutos proceder) * caso algum usuário tenha criado um appointment
	//e não tenha confirmado com o link enviado pro email
	// "confirmed": interromper execução e retornar resposta coerente para o usuário
	result, err := repository.db.Exec(
		`INSERT INTO appointments (id, datetime, slot, candidate, status)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (datetime) DO UPDATE
		SET candidate = $6, status = $7, updated_at = $8
		WHERE (
			appointments.status = 3
			OR (
				appointments.status = 1
				AND NOW() > appointments.created_at + INTERVAL '25 minutes'
				AND NOW() > appointments.updated_at + INTERVAL '25 minutes'
			)
		);`,
		a.Id,
		a.Datetime,
		a.Slot,
		candidateId,
		a.Status,
		candidateId,
		a.Status,
		time.Now(),
	)

	if err != nil {

		fmt.Println("err: ", err)

		if pgError, ok := err.(*pq.Error); ok {
			if pgError.Code == "23505" {
				return false, &shared.AppError{Err: pgError, Message: "Ocorreu um problema no servidor/Registro duplicado", StatusCode: http.StatusInternalServerError}
			}
		}

		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("rowsAffected: ", rowsAffected)

	if rowsAffected < 1 {
		return false, &shared.AppError{Err: err, Message: "Este horário não está mais disponível", StatusCode: http.StatusBadRequest}
	}

	return true, nil
}

func (repository *PostgresAppointmentsRepository) FindByCandidateId(candidateId string) (map[string]interface{}, *shared.AppError) { // if slot is valid or not

	rows, err := repository.db.Query(
		`SELECT id, status, created_at FROM appointments WHERE candidate = $1`,
		candidateId,
	)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer rows.Close()

	if rows.Next() {
		var id string
		var status int
		var createdAt time.Time
		if err := rows.Scan(
			&id,
			&status,
			&createdAt,
		); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		if err := rows.Err(); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		response := make(map[string]interface{})
		response["id"] = id
		response["status"] = status
		response["created_at"] = createdAt

		return response, nil
	}

	// no appointment from that candidate found and no error
	return nil, nil
}

func (repository *PostgresAppointmentsRepository) FindByDatetime(datetime string) ([]int, *shared.AppError) { // if slot is valid or not

	rows, err := repository.db.Query(
		`SELECT slot FROM appointments 
		WHERE DATE(datetime) = DATE($1) 
		AND status <> 2
		AND NOW() < appointments.created_at + INTERVAL '25 minutes'
		AND NOW() < appointments.updated_at + INTERVAL '25 minutes'`,
		datetime,
	)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int

		if err := rows.Scan(
			&id,
		); err != nil {
			return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
		}

		ids = append(ids, id)
	}

	if err := rows.Err(); err != nil {
		return nil, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	return ids, nil
}

func (repository *PostgresAppointmentsRepository) UpdateStatus(id int, status int) (bool, *shared.AppError) { // if slot is valid or not

	result, err := repository.db.Exec(
		`UPDATE appointments SET status = $1 WHERE id = $2`,
		status,
		id,
	)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	if affected > 0 {
		return true, nil
	}

	return false, &shared.AppError{Err: err, Message: "Não foi possível confirmar o agendamento. Tente novamente", StatusCode: http.StatusBadRequest}
}

func (repository *PostgresAppointmentsRepository) UpdateStatusToConfirmed(id string) (bool, *shared.AppError) {

	result, err := repository.db.Exec(
		`UPDATE appointments SET status = 2 WHERE id = $1
		AND NOW() < appointments.created_at + INTERVAL '25 minutes'`,
		id, // problems with id?
	)

	// database error during query processing & nothing to do with business logic
	if err != nil {
		fmt.Println("Error during query, ", err)

		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	affected, err := result.RowsAffected()

	if err != nil {
		return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
	}

	if affected > 0 {
		return true, nil
	}

	return false, &shared.AppError{Err: err, Message: "Ocorreu um problema interno no servidor", StatusCode: http.StatusInternalServerError}
}
