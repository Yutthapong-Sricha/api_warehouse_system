package models

type Product struct {
	Id_product            int    `json:"id_product" db:"id_product"`
	Product_name          string `json:"product_name" db:"product_name"`
	Product_brand         string `json:"product_brand" db:"product_brand"`
	Product_model         string `json:"product_model" db:"product_model"`
	Product_detail        string `json:"product_detail" db:"product_detail"`
	Is_fifo_356           int    `json:"is_fifo_356" db:"is_fifo_356"`
	Id_prod_category      int    `json:"id_prod_category" db:"id_prod_category"`
	Qty_per_pack          int    `json:"qty_per_pack" db:"qty_per_pack"`
	Unit_name             string `json:"unit_name" db:"unit_name"`
	Unit_name_pack        string `json:"unit_name_pack" db:"unit_name_pack"`
	Is_main_356           int    `json:"is_main_356" db:"is_main_356"`
	Is_volume_flag        int    `json:"is_volume_flag" db:"is_volume_flag"`
	Avg_price_vol         int    `json:"avg_price_vol" db:"avg_price_vol"`
	Sale_price_per_unit   int    `json:"sale_price_per_unit" db:"sale_price_per_unit"`
	Is_active_flag        int    `json:"is_active_flag" db:"is_active_flag"`
	Product_barcode       string `json:"product_barcode" db:"product_barcode"`
	Last_cost_price       int    `json:"last_cost_price" db:"last_cost_price"`
	Id_supplier           int    `json:"id_supplier" db:"id_supplier"`
	Is_use_car_356        int    `json:"is_use_car_356" db:"is_use_car_356"`
	Record_update_time    int    `json:"record_update_time" db:"record_update_time"`
	Record_update_by_name string `json:"record_update_by_name" db:"record_update_by_name"`
	Id_product_enc        string `json:"id_product_enc" db:"id_product_enc"`
}
