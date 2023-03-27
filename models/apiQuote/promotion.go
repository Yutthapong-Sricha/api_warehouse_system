package models

import (
	"log"
	"server_go/api_warehouse_system/helpers"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ListPromotion(act string) []modelsStruc.Promotion {
	var Promotions []modelsStruc.Promotion
	var err error
	sql_statement := "SELECT id_promo_detail, promotion_name, IFNULL(id_supplier,0) AS id_supplier, IFNULL(id_product,0) AS id_product , IFNULL(tmp_product_name,'') AS tmp_product_name, IFNULL(promo_start_ddmm20yy,'') AS promo_start_ddmm20yy, IFNULL(promo_end_ddmm20yy,'') AS promo_end_ddmm20yy, IFNULL(promo_cost,0) AS promo_cost, IFNULL(cost_multiply,0) AS cost_multiply, IFNULL(pro_expired_0356_flag,0) AS pro_expired_0356_flag, IFNULL(actor_name,'') AS actor_name, IFNULL(show_promo_356_flag,0) AS show_promo_356_flag, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name FROM promo_detail "
	if act == "/active_only" {
		sql_statement = sql_statement + " where pro_expired_0356_flag=5 "
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Promotion
		err = list.Scan(
			&row.Id_promo_detail,
			&row.Promotion_name,
			&row.Id_supplier,
			&row.Id_product,
			&row.Tmp_product_name,
			&row.Promo_start_ddmm20yy,
			&row.Promo_end_ddmm20yy,
			&row.Promo_cost,
			&row.Cost_multiply,
			&row.Pro_expired_0356_flag,
			&row.Actor_name,
			&row.Show_promo_356_flag,
			&row.Record_update_time,
			&row.Record_update_by_name,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.Id_promo_detail))
			row.Id_promo_detail_enc = id_enc
		}
		Promotions = append(Promotions, row)
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Promotions
}

func GetPromotion(id_enc string) modelsStruc.Promotion {
	var Promotion modelsStruc.Promotion
	var err error
	id_str := helpers.Decrypt(id_enc)
	id_int, conv_err := strconv.ParseInt(id_str, 10, 0)
	if conv_err != nil {
		panic(conv_err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT id_promo_detail, promotion_name, IFNULL(id_supplier,0) AS id_supplier, IFNULL(id_product,0) AS id_product , IFNULL(tmp_product_name,'') AS tmp_product_name, IFNULL(promo_start_ddmm20yy,'') AS promo_start_ddmm20yy, IFNULL(promo_end_ddmm20yy,'') AS promo_end_ddmm20yy, IFNULL(promo_cost,0) AS promo_cost, IFNULL(cost_multiply,0) AS cost_multiply, IFNULL(pro_expired_0356_flag,0) AS pro_expired_0356_flag, IFNULL(actor_name,'') AS actor_name, IFNULL(show_promo_356_flag,0) AS show_promo_356_flag, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name FROM promo_detail WHERE id_promo_detail=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	row := statement.QueryRow(id_int)
	err = row.Scan(
		&Promotion.Id_promo_detail,
		&Promotion.Promotion_name,
		&Promotion.Id_supplier,
		&Promotion.Id_product,
		&Promotion.Tmp_product_name,
		&Promotion.Promo_start_ddmm20yy,
		&Promotion.Promo_end_ddmm20yy,
		&Promotion.Promo_cost,
		&Promotion.Cost_multiply,
		&Promotion.Pro_expired_0356_flag,
		&Promotion.Actor_name,
		&Promotion.Show_promo_356_flag,
		&Promotion.Record_update_time,
		&Promotion.Record_update_by_name,
	)
	if err != nil {
		return Promotion
	} else {
		Promotion.Id_promo_detail_enc = id_enc
	}
	return Promotion
}
