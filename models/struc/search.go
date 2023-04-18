package models

type Keyword struct {
	Keysearch string `json:"Keysearch" db:"Keysearch"`
}

// Binding from JSON
type Searchcust struct {
	Value         int    `json:"value" db:"value"`
	Cust_code_spc string `json:"cust_code_spc" db:"cust_code_spc"`
	Title         string `json:"title" db:"title"`
	Cust_nickname string `json:"cust_nickname" db:"cust_nickname"`
	Cust_tax_id   string `json:"cust_tax_id" db:"cust_tax_id"`
	Cust_mobile   string `json:"cust_mobile" db:"cust_mobile"`
	Cust_tel      string `json:"cust_tel" db:"cust_tel"`
	Cust_email    string `json:"cust_email" db:"cust_email"`
}

// Binding from JSON
type Searchprod struct {
	Value               int     `json:"value" db:"value"`
	Title               string  `json:"title" db:"title"`
	Sale_price_per_unit float64 `json:"sale_price_per_unit" db:"sale_price_per_unit"`
	Id_supplier         int     `json:"id_supplier" db:"id_supplier"`
	Id_prod_category    int     `json:"id_prod_category" db:"id_prod_category"`
	Qty_per_pack        int     `json:"qty_per_pack" db:"qty_per_pack"`
	Unit_name           string  `json:"unit_name" db:"unit_name"`
}
