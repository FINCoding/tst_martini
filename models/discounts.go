package models

type Discounts struct{
    id string
    title string
    description string
    is_sc string
    id_car string
}

func NewDiscounts(id, title, description, id_sc, id_car string)*Discounts{
    return &Discounts{id, title, description, id_sc, id_car}
}