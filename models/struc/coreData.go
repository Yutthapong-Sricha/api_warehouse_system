package models

type Branch struct {
	Id_branch        int    `json:"id_branch" db:"id_branch"`
	Branch_name      string `json:"branch_name" db:"branch_name"`
	Address          string `json:"address" db:"address"`
	Id_addr_tambon   int    `json:"id_addr_tambon" db:"id_addr_tambon"`
	Id_addr_amphure  int    `json:"id_addr_amphure" db:"id_addr_amphure"`
	Id_addr_province int    `json:"id_addr_province" db:"id_addr_province"`
	Branch_tel       string `json:"branch_tel" db:"branch_tel"`
	Branch_line_id   string `json:"branch_line_id" db:"branch_line_id"`
	Is_active_flag   int    `json:"is_active_flag" db:"is_active_flag"`
	Id_branch_enc    string `json:"id_branch_enc" db:"id_branch_enc"`
}

type Position struct {
	Id_staff_position  int    `json:"id_staff_position" db:"id_staff_position"`
	Position_name      string `json:"position_name" db:"position_name"`
	Row_order_position int    `json:"row_order_position" db:"row_order_position"`
	Is_active_flag     int    `json:"is_active_flag" db:"is_active_flag"`
	Id_position_enc    string `json:"id_position_enc" db:"id_position_enc"`
}

type Category struct {
	Id_prod_category      int    `json:"id_prod_category" db:"id_prod_category"`
	Category_name         string `json:"category_name" db:"category_name"`
	Is_active_flag        int    `json:"is_active_flag" db:"is_active_flag"`
	Record_update_time    int    `json:"record_update_time" db:"record_update_time"`
	Record_update_by_name string `json:"record_update_by_name" db:"record_update_by_name"`
	Id_prod_category_enc  string `json:"id_prod_category_enc" db:"id_prod_category_enc"`
}
