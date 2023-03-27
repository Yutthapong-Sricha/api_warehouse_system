package models

import (
	"log"
	"server_go/api_warehouse_system/helpers"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func ListBranch(act string) []modelsStruc.Branch {
	var Branchs []modelsStruc.Branch
	var err error
	sql_statement := "SELECT id_branch, branch_name, IFNULL(address,'') AS address, IFNULL(id_addr_tambon,0) AS id_addr_tambon, IFNULL(id_addr_amphure, 0) AS id_addr_amphure, IFNULL(id_addr_province, 0) AS id_addr_province, IFNULL(branch_tel, '') as branch_tel,  is_active_flag  FROM branch "
	if act == "/active_only" {
		sql_statement = sql_statement + " where is_active_flag=1 "
	}
	list, err := initializers.DB.Query(sql_statement)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Branch
		err = list.Scan(
			&row.Id_branch,
			&row.Branch_name,
			&row.Address,
			&row.Id_addr_tambon,
			&row.Id_addr_amphure,
			&row.Id_addr_province,
			&row.Branch_tel,
			&row.Is_active_flag,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			id_enc := helpers.Encrypt(strconv.Itoa(row.Id_branch))
			row.Id_branch_enc = id_enc
		}
		Branchs = append(Branchs, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Branchs
}

func GetBranch(id_enc string) modelsStruc.Branch {
	var Branch modelsStruc.Branch
	var err error
	id_str := helpers.Decrypt(id_enc)
	id_int, conv_err := strconv.ParseInt(id_str, 10, 0)
	if conv_err != nil {
		panic(conv_err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT id_branch, branch_name, IFNULL(address,'') AS address, IFNULL(id_addr_tambon,0) AS id_addr_tambon, IFNULL(id_addr_amphure,0) AS id_addr_amphure , IFNULL(id_addr_province,0) AS id_addr_province, IFNULL(branch_tel,'') AS branch_tel, IFNULL(branch_line_id,'') AS branch_line_id , IFNULL(is_active_flag,0) AS is_active_flag FROM branch WHERE id_branch=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	row := statement.QueryRow(id_int)
	err = row.Scan(
		&Branch.Id_branch,
		&Branch.Branch_name,
		&Branch.Address,
		&Branch.Id_addr_tambon,
		&Branch.Id_addr_amphure,
		&Branch.Id_addr_province,
		&Branch.Branch_tel,
		&Branch.Branch_line_id,
		&Branch.Is_active_flag,
	)
	if err != nil {
		return Branch
	} else {
		Branch.Id_branch_enc = id_enc
	}

	return Branch
}
