package helper

// Binding from JSON
type Staff struct {
	Id_customer   string `json:"id_customer" db:"id_customer"`
	Cust_fullname string `json:"cust_fullname" db:"cust_fullname"`
	Cust_mobile   string `json:"cust_mobile" db:"cust_mobile"`
}

type Supplier struct {
	Id_supplier      string `json:"id_supplier" db:"id_supplier"`
	Supp_name        string `json:"supp_name" db:"supp_name"`
	Supp_address     string `json:"supp_address" db:"supp_address"`
	Id_addr_tambon   string `json:"id_addr_tambon" db:"id_addr_tambon"`
	Id_addr_amphure  string `json:"id_addr_amphure" db:"id_addr_amphure"`
	Id_addr_province string `json:"id_addr_province" db:"id_addr_province"`
}
