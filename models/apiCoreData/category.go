package models

import (
	"log"
	"server_go/api_warehouse_system/helpers"
	"server_go/api_warehouse_system/initializers"

	"strconv"

	modelsStruc "server_go/api_warehouse_system/models/struc"

	_ "github.com/go-sql-driver/mysql"
)

func ListCategory(act string) []modelsStruc.Category {
	var Categorys []modelsStruc.Category
	var err error
	sql_statement := "SELECT id_prod_category, category_name, is_active_flag, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name FROM prod_category "
	if act == "/active_only" {
		sql_statement = sql_statement + " where is_active_flag=1 "
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Category
		err = list.Scan(
			&row.Id_prod_category,
			&row.Category_name,
			&row.Is_active_flag,
			&row.Record_update_time,
			&row.Record_update_by_name,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.Id_prod_category))

			row.Id_prod_category_enc = id_enc

		}
		Categorys = append(Categorys, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Categorys
}

func GetCategory(id string) modelsStruc.Category {
	var Category modelsStruc.Category
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT id_prod_category, category_name, is_active_flag, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name FROM prod_category WHERE id_prod_category=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	rows := statement.QueryRow(id_int)
	err = rows.Scan(
		&Category.Id_prod_category,
		&Category.Category_name,
		&Category.Is_active_flag,
		&Category.Record_update_time,
		&Category.Record_update_by_name,
	)
	if err != nil {
		return Category
	} else {
		id_enc := helpers.Encrypt(strconv.Itoa(Category.Id_prod_category))

		Category.Id_prod_category_enc = id_enc

	}
	return Category
}
