package models

// Binding from JSON
type Customer struct {
	Id_customer   int    `json:"id_customer" db:"id_customer"`
	Cust_fullname string `json:"cust_fullname" db:"cust_fullname"`
	Cust_mobile   string `json:"cust_mobile" db:"cust_mobile"`
}

type Supplier struct {
	Id_supplier      int    `json:"id_supplier" db:"id_supplier"`
	Supp_name        string `json:"supp_name" db:"supp_name"`
	Supp_address     string `json:"supp_address" db:"supp_address"`
	Id_addr_tambon   int    `json:"id_addr_tambon" db:"id_addr_tambon"`
	Id_addr_amphure  int    `json:"id_addr_amphure" db:"id_addr_amphure"`
	Id_addr_province int    `json:"id_addr_province" db:"id_addr_province"`
}

type Staff struct {
	Id_staff_core         int    `json:"id_staff_core" db:"id_staff_core"`
	Id_branch             int    `json:"id_branch" db:"id_branch"`
	Staff_name            string `json:"staff_name" db:"staff_name"`
	Staff_lastname        string `json:"staff_lastname" db:"staff_lastname"`
	Staff_fullname        string `json:"staff_fullname" db:"staff_fullname"`
	Staff_nickname        string `json:"staff_nickname" db:"staff_nickname"`
	Staff_tax_id          string `json:"staff_tax_id" db:"staff_tax_id"`
	Staff_address         string `json:"staff_address" db:"staff_address"`
	Id_addr_tambon        int    `json:"id_addr_tambon" db:"id_addr_tambon"`
	Id_addr_amphure       int    `json:"id_addr_amphure" db:"id_addr_amphure"`
	Id_addr_province      int    `json:"id_addr_province" db:"id_addr_province"`
	Staff_mobile1         string `json:"staff_mobile1" db:"staff_mobile1"`
	Staff_mobile2         string `json:"staff_mobile2" db:"staff_mobile2"`
	Staff_tel             string `json:"staff_tel" db:"staff_tel"`
	Pwd_verification      string `json:"pwd_verification" db:"pwd_verification"`
	Is_active_flag        int    `json:"is_active_flag" db:"is_active_flag"`
	Staff_line_id         string `json:"staff_line_id" db:"staff_line_id"`
	Id_staff_position     int    `json:"id_staff_position" db:"id_staff_position"`
	Tmp_position_name     string `json:"tmp_position_name" db:"tmp_position_name"`
	Birth_day_ddmm20yy    string `json:"birth_day_ddmm20yy" db:"birth_day_ddmm20yy"`
	Record_create_time    int    `json:"record_create_time" db:"record_create_time"`
	Record_create_by_id   int    `json:"record_create_by_id" db:"record_create_by_id"`
	Record_create_by_name string `json:"record_create_by_name" db:"record_create_by_name"`
	Record_update_time    int    `json:"record_update_time" db:"record_update_time"`
	Record_update_by_id   int    `json:"record_update_by_id" db:"record_update_by_id"`
	Record_update_by_name string `json:"record_update_by_name" db:"record_update_by_name"`
	Note                  string `json:"note" db:"note"`
}
