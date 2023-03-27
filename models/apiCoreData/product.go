package models

import (
	"fmt"
	"log"
	"server_go/api_warehouse_system/helpers"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ListProduct(act string) []modelsStruc.Product {
	var Products []modelsStruc.Product
	var err error
	sql_statement := "SELECT id_product,product_name, IFNULL(product_brand,'') as product_brand , IFNULL(product_model,'') as product_model, IFNULL(product_detail,'') as product_detail, IFNULL(is_fifo_356,	0) as is_fifo_356, IFNULL(id_prod_category,0) as id_prod_category, IFNULL(qty_per_pack,0) as qty_per_pack, IFNULL(unit_name,'') as unit_name, IFNULL(unit_name_pack,'') as unit_name_pack, IFNULL(is_main_356,0) as is_main_356, IFNULL(is_volume_flag,0) as is_volume_flag, IFNULL(avg_price_vol,0) as avg_price_vol, IFNULL(sale_price_per_unit,0) as sale_price_per_unit, IFNULL(is_active_flag,0) as is_active_flag, IFNULL(product_barcode,'') AS product_barcode, IFNULL(include_vat_flag,0) AS include_vat_flag, IFNULL(last_cost_price,0) AS last_cost_price , IFNULL(id_supplier,0) AS id_supplier, IFNULL(is_use_car_356,0) AS is_use_car_356,IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name , IFNULL(key_search,'') AS key_search FROM product  "
	if act == "/active_only" {
		sql_statement = sql_statement + " WHERE is_active_flag=1 ORDER BY record_create_time DESC"
	} else {
		sql_statement = sql_statement + " ORDER BY record_create_time DESC"
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Product
		err = list.Scan(
			&row.Id_product,
			&row.Product_name,
			&row.Product_brand,
			&row.Product_model,
			&row.Product_detail,
			&row.Is_fifo_356,
			&row.Id_prod_category,
			&row.Qty_per_pack,
			&row.Unit_name,
			&row.Unit_name_pack,
			&row.Is_main_356,
			&row.Is_volume_flag,
			&row.Avg_price_vol,
			&row.Sale_price_per_unit,
			&row.Is_active_flag,
			&row.Product_barcode,
			&row.Include_vat_flag,
			&row.Last_cost_price,
			&row.Id_supplier,
			&row.Is_use_car_356,
			&row.Record_update_time,
			&row.Record_update_by_name,
			&row.Key_search,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.Id_product))
			row.Id_product_enc = id_enc
		}
		Products = append(Products, row)
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Products
}

func GetProduct(id_enc string) modelsStruc.Product {
	var Product modelsStruc.Product
	var err error
	id_str := helpers.Decrypt(id_enc)
	id_int, conv_err := strconv.ParseInt(id_str, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	fmt.Println(id_int)
	statement, err := initializers.DB.Prepare("SELECT id_product,product_name, IFNULL(product_brand,'') as product_brand , IFNULL(product_model,'') as product_model, IFNULL(product_detail,'') as product_detail, IFNULL(is_fifo_356,	0) as is_fifo_356, IFNULL(id_prod_category,0) as id_prod_category, IFNULL(qty_per_pack,0) as qty_per_pack, IFNULL(unit_name,'') as unit_name, IFNULL(unit_name_pack,'') as unit_name_pack, IFNULL(is_main_356,0) as is_main_356, IFNULL(is_volume_flag,0) as is_volume_flag, IFNULL(avg_price_vol,0) as avg_price_vol, IFNULL(sale_price_per_unit,0) as sale_price_per_unit, IFNULL(is_active_flag,0) as is_active_flag, IFNULL(product_barcode,'') AS product_barcode, IFNULL(include_vat_flag,0) AS include_vat_flag, IFNULL(last_cost_price,0) AS last_cost_price , IFNULL(id_supplier,0) AS id_supplier, IFNULL(is_use_car_356,0) AS is_use_car_356,IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name , IFNULL(key_search,'') AS key_search FROM product WHERE id_product=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	row := statement.QueryRow(id_int)
	err = row.Scan(
		&Product.Id_product,
		&Product.Product_name,
		&Product.Product_brand,
		&Product.Product_model,
		&Product.Product_detail,
		&Product.Is_fifo_356,
		&Product.Id_prod_category,
		&Product.Qty_per_pack,
		&Product.Unit_name,
		&Product.Unit_name_pack,
		&Product.Is_main_356,
		&Product.Is_volume_flag,
		&Product.Avg_price_vol,
		&Product.Sale_price_per_unit,
		&Product.Is_active_flag,
		&Product.Product_barcode,
		&Product.Include_vat_flag,
		&Product.Last_cost_price,
		&Product.Id_supplier,
		&Product.Is_use_car_356,
		&Product.Record_update_time,
		&Product.Record_update_by_name,
		&Product.Key_search,
	)
	if err != nil {
		return Product
	} else {
		Product.Id_product_enc = id_enc
	}
	return Product
}
