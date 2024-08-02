package repositories

import "github.com/techhub-jf/farmacia-back/app/gateway/postgres"

type MedicinesRepository struct {
	*postgres.Client
}

func NewMedicinesRepository(client *postgres.Client) *MedicinesRepository {
	return &MedicinesRepository{client}
}
