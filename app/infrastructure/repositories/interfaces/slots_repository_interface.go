package interfaces

import shared "rhSystem_server/app/application/error"

type SlotsRepositoryInterface interface {
	//Index()
	FindById(id int) (bool, *shared.AppError)
}
