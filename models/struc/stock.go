package models

type Product struct {
	Id_product            int     `json:"id_product" db:"id_product"`
	Product_name          string  `json:"product_name" db:"product_name"`
	Product_brand         string  `json:"product_brand" db:"product_brand"`
	Product_model         string  `json:"product_model" db:"product_model"`
	Product_detail        string  `json:"product_detail" db:"product_detail"`
	Is_fifo_356           int     `json:"is_fifo_356" db:"is_fifo_356"`
	Id_prod_category      int     `json:"id_prod_category" db:"id_prod_category"`
	Qty_per_pack          int     `json:"qty_per_pack" db:"qty_per_pack"`
	Unit_name             string  `json:"unit_name" db:"unit_name"`
	Unit_name_pack        string  `json:"unit_name_pack" db:"unit_name_pack"`
	Is_main_356           int     `json:"is_main_356" db:"is_main_356"`
	Is_volume_flag        int     `json:"is_volume_flag" db:"is_volume_flag"`
	Avg_price_vol         float64 `json:"avg_price_vol" db:"avg_price_vol"`
	Sale_price_per_unit   float64 `json:"sale_price_per_unit" db:"sale_price_per_unit"`
	Is_active_flag        int     `json:"is_active_flag" db:"is_active_flag"`
	Product_barcode       string  `json:"product_barcode" db:"product_barcode"`
	Include_vat_flag      int     `json:"include_vat_flag" db:"include_vat_flag"`
	Last_cost_price       float64 `json:"last_cost_price" db:"last_cost_price"`
	Id_supplier           int     `json:"id_supplier" db:"id_supplier"`
	Is_use_car_356        int     `json:"is_use_car_356" db:"is_use_car_356"`
	Record_update_time    int     `json:"record_update_time" db:"record_update_time"`
	Record_update_by_name string  `json:"record_update_by_name" db:"record_update_by_name"`
	Key_search            string  `json:"key_search" db:"key_search"`
	Id_product_enc        string  `json:"id_product_enc" db:"id_product_enc"`
}

type Promotion struct {
	Id_promo_detail       int     `json:"id_promo_detail" db:"id_promo_detail"`
	Promotion_name        string  `json:"promotion_name" db:"promotion_name"`
	Id_supplier           int     `json:"id_supplier" db:"id_supplier"`
	Id_product            int     `json:"id_product" db:"id_product"`
	Tmp_product_name      string  `json:"tmp_product_name" db:"tmp_product_name"`
	Promo_start_ddmm20yy  string  `json:"promo_start_ddmm20yy" db:"promo_start_ddmm20yy"`
	Promo_end_ddmm20yy    string  `json:"promo_end_ddmm20yy" db:"promo_end_ddmm20yy"`
	Promo_cost            float64 `json:"promo_cost" db:"promo_cost"`
	Cost_multiply         int     `json:"cost_multiply" db:"cost_multiply"`
	Pro_expired_0356_flag int     `json:"pro_expired_0356_flag" db:"pro_expired_0356_flag"`
	Actor_name            string  `json:"actor_name" db:"actor_name"`
	Show_promo_356_flag   int     `json:"show_promo_356_flag" db:"show_promo_356_flag"`
	Record_update_time    int     `json:"record_update_time" db:"record_update_time"`
	Record_update_by_name string  `json:"record_update_by_name" db:"record_update_by_name"`
	Id_promo_detail_enc   string  `json:"id_promo_detail_enc" db:"id_promo_detail_enc"`
}

type Stock struct {
	Id_stock_core          int     `json:"id_stock_core" db:"id_stock_core"`
	Id_branch              int     `json:"id_branch" db:"id_branch"`
	Id_product             int     `json:"id_product" db:"id_product"`
	Id_supplier            int     `json:"id_supplier" db:"id_supplier"`
	Id_prod_category       int     `json:"id_prod_category" db:"id_prod_category"`
	Product_name           string  `json:"product_name" db:"product_name"`
	Product_brand          string  `json:"product_brand" db:"product_brand"`
	Product_model          string  `json:"product_model" db:"product_model"`
	Product_detail         string  `json:"product_detail" db:"product_detail"`
	Part_serial_no         string  `json:"part_serial_no" db:"part_serial_no"`
	Part_engine_no         string  `json:"part_engine_no" db:"part_engine_no"`
	Last_cost_price        int     `json:"last_cost_price" db:"last_cost_price"`
	Total_qty_cost         int     `json:"total_qty_cost" db:"total_qty_cost"`
	Total_qty_free         int     `json:"total_qty_free" db:"total_qty_free"`
	Qty_out_cost           int     `json:"qty_out_cost" db:"qty_out_cost"`
	Qty_out_free           int     `json:"qty_out_free" db:"qty_out_free"`
	Total_qty_balance      int     `json:"total_qty_balance" db:"total_qty_balance"`
	Price_per_unit         float64 `json:"price_per_unit" db:"price_per_unit"`
	Price_per_pack         float64 `json:"price_per_pack" db:"price_per_pack"`
	Barcode                string  `json:"barcode" db:"barcode"`
	Key_search             string  `json:"key_search" db:"key_search"`
	Record_update_time     int     `json:"record_update_time" db:"record_update_time"`
	Record_update_by_name  string  `json:"record_update_by_name" db:"record_update_by_name"`
	Id_loca_store          int     `json:"id_loca_store" db:"id_loca_store"`
	Tmp_loca_store_name    string  `json:"tmp_loca_store_name" db:"tmp_loca_store_name"`
	Has_promotion_356_flag int     `json:"has_promotion_356_flag" db:"has_promotion_356_flag"`
	Id_promotion           int     `json:"id_promotion" db:"id_promotion"`
	Id_stock_core_enc      string  `json:"id_stock_core_enc" db:"id_stock_core_enc"`
}
