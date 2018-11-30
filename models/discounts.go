package models

type Discount struct{
    Id string
    Title string
    Description string
    //Id_sc string
    //Id_car string
}

//func NewDiscount(id, title, description, id_sc, id_car string)*Discount{
//    return &Discount{id, title, description, id_sc, id_car}
//}

func NewDiscount(id, title, description string)*Discount{
    return &Discount{id, title, description}
}