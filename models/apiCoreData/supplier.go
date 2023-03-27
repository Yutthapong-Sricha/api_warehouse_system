package models

import (
	"log"
	"server_go/api_warehouse_system/helpers"
	"server_go/api_warehouse_system/initializers"
	"strconv"

	modelsStruc "server_go/api_warehouse_system/models/struc"

	_ "github.com/go-sql-driver/mysql"
)

func ListSupplier(act string) []modelsStruc.Supplier {
	var Suppliers []modelsStruc.Supplier
	var err error
	sql_statement := "SELECT id_supplier, IFNULL(supp_code,'') as supp_code, IFNULL(supp_tax_id,'') as supp_tax_id, supp_name, IFNULL(supp_address,'') as supp_address, id_addr_tambon, IFNULL(tmp_tambon_name,'') as tmp_tambon_name, id_addr_amphure, IFNULL(tmp_amphure_name,'') as tmp_amphure_name, id_addr_province, IFNULL(tmp_province_name,'') as tmp_province_name, IFNULL(tmp_zip_code,'') AS tmp_zip_code, IFNULL(supp_tel,'') AS supp_tel, IFNULL(supp_email,'') AS supp_email, IFNULL(supp_mobile,'') AS supp_mobile, IFNULL(contact1_name,'') AS contact1_name, IFNULL(contact1_mobile,'') AS contact1_mobile, IFNULL(contact1_email,'') AS contact1_email, IFNULL(contact1_note,'') AS contact1_note, IFNULL(contact2_name,'') AS contact2_name, IFNULL(contact2_mobile,'') AS contact2_mobile, IFNULL(contact2_email,'') AS contact2_email, IFNULL(contact2_note,'') AS contact2_note, IFNULL(due_date,'') AS due_date, IFNULL(acc_no,'') AS acc_no, IFNULL(bankname,'') AS bankname, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name, IFNULL(note,'') AS note , IFNULL(status_356_flag,0) AS status_356_flag FROM supplier  "
	if act == "/active_only" {
		sql_statement = sql_statement + " WHERE status_356_flag=5 ORDER BY record_create_time DESC"
	} else {
		sql_statement = sql_statement + " ORDER BY record_create_time DESC"
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Supplier
		err = list.Scan(
			&row.Id_supplier,
			&row.Supp_code,
			&row.Supp_tax_id,
			&row.Supp_name,
			&row.Supp_address,
			&row.Id_addr_tambon,
			&row.Tmp_tambon_name,
			&row.Id_addr_amphure,
			&row.Tmp_amphure_name,
			&row.Id_addr_province,
			&row.Tmp_province_name,
			&row.Tmp_zip_code,
			&row.Supp_tel,
			&row.Supp_email,
			&row.Supp_mobile,
			&row.Contact1_name,
			&row.Contact1_mobile,
			&row.Contact1_email,
			&row.Contact1_note,
			&row.Contact2_name,
			&row.Contact2_mobile,
			&row.Contact2_email,
			&row.Contact2_note,
			&row.Due_date,
			&row.Acc_no,
			&row.Bankname,
			&row.Record_update_time,
			&row.Record_update_by_name,
			&row.Note,
			&row.Status_356_flag,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.Id_supplier))

			row.Id_supplier_enc = id_enc

		}
		Suppliers = append(Suppliers, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Suppliers
}

func GetSupplier(id string) modelsStruc.Supplier {
	var Supplier modelsStruc.Supplier
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT id_supplier, IFNULL(supp_code,'') as supp_code, IFNULL(supp_tax_id,'') as supp_tax_id, supp_name, IFNULL(supp_address,'') as supp_address, id_addr_tambon, IFNULL(tmp_tambon_name,'') as tmp_tambon_name, id_addr_amphure, IFNULL(tmp_amphure_name,'') as tmp_amphure_name, id_addr_province, IFNULL(tmp_province_name,'') as tmp_province_name, IFNULL(tmp_zip_code,'') AS tmp_zip_code, IFNULL(supp_tel,'') AS supp_tel, IFNULL(supp_email,'') AS supp_email, IFNULL(supp_mobile,'') AS supp_mobile, IFNULL(contact1_name,'') AS contact1_name, IFNULL(contact1_mobile,'') AS contact1_mobile, IFNULL(contact1_email,'') AS contact1_email, IFNULL(contact1_note,'') AS contact1_note, IFNULL(contact2_name,'') AS contact2_name, IFNULL(contact2_mobile,'') AS contact2_mobile, IFNULL(contact2_email,'') AS contact2_email, IFNULL(contact2_note,'') AS contact2_note, IFNULL(due_date,'') AS due_date, IFNULL(acc_no,'') AS acc_no, IFNULL(bankname,'') AS bankname, IFNULL(record_update_time,0) AS record_update_time, IFNULL(record_update_by_name,'') AS record_update_by_name, IFNULL(note,'') AS note , IFNULL(status_356_flag,0) AS status_356_flag FROM supplier WHERE id_supplier=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	row := statement.QueryRow(id_int)
	err = row.Scan(
		&Supplier.Id_supplier,
		&Supplier.Supp_code,
		&Supplier.Supp_tax_id,
		&Supplier.Supp_name,
		&Supplier.Supp_address,
		&Supplier.Id_addr_tambon,
		&Supplier.Tmp_tambon_name,
		&Supplier.Id_addr_amphure,
		&Supplier.Tmp_amphure_name,
		&Supplier.Id_addr_province,
		&Supplier.Tmp_province_name,
		&Supplier.Tmp_zip_code,
		&Supplier.Supp_tel,
		&Supplier.Supp_email,
		&Supplier.Supp_mobile,
		&Supplier.Contact1_name,
		&Supplier.Contact1_mobile,
		&Supplier.Contact1_email,
		&Supplier.Contact1_note,
		&Supplier.Contact2_name,
		&Supplier.Contact2_mobile,
		&Supplier.Contact2_email,
		&Supplier.Contact2_note,
		&Supplier.Due_date,
		&Supplier.Acc_no,
		&Supplier.Bankname,
		&Supplier.Note,
		&Supplier.Record_update_time,
		&Supplier.Record_update_by_name,
	)
	if err != nil {
		return Supplier
	} else {
		id_enc := helpers.Encrypt(strconv.Itoa(Supplier.Id_supplier))

		Supplier.Id_supplier_enc = id_enc

	}
	return Supplier

}
