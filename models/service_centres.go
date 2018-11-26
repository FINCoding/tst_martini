package models

type ServiceCentres struct{
    id string
    adress string
    id_car string
}

func NewSC(id, adress, id_car string)*ServiceCentres{
    return &ServiceCentres{id, adress, id_car}
}