package helper

type Position struct {
	Id_staff_position  string `json:"id_staff_position" db:"id_staff_position"`
	Position_name      string `json:"position_name" db:"position_name"`
	Row_order_position string `json:"row_order_position" db:"row_order_position"`
	Is_active_flag     string `json:"is_active_flag" db:"is_active_flag"`
}
