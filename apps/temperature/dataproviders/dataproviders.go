package dataproviders

import (
	"temperature/common/pg"
	"temperature/dataproviders/temperature_repository"
)

type CoreProviders interface {
	GetUserRepository() temperature_repository.ITemperature
}

type coreProviders struct {
	temperatureRepository temperature_repository.ITemperature
}

func NewCoreProviders(newPg *pg.DB) (CoreProviders, error) {

	tr := temperature_repository.NewTemperatureRepository(newPg)

	return coreProviders{
		temperatureRepository: tr,
	}, nil
}

func (p coreProviders) GetUserRepository() temperature_repository.ITemperature {
	return p.temperatureRepository
}
