package models

import (
	"log"
	"server_go/api_warehouse_system/initializers"

	"strconv"

	modelsStruc "server_go/api_warehouse_system/models/struc"

	_ "github.com/go-sql-driver/mysql"
)

func ListBranch() []modelsStruc.Branch {
	var Branchs []modelsStruc.Branch
	var err error

	list, err := initializers.DB.Query("SELECT id_branch, branch_name, IFNULL(address,'') AS address, IFNULL(id_addr_tambon,0) AS id_addr_tambon, IFNULL(id_addr_amphure, 0) AS id_addr_amphure, IFNULL(id_addr_province, 0) AS id_addr_province, IFNULL(branch_tel, '') as branch_tel,  is_active_flag  FROM branch")
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
			//&row.Branch_line_id,
			&row.Is_active_flag,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			row.Id_branch_enc = "string_enc"
		}
		Branchs = append(Branchs, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Branchs
}

func GetBranch(id string) modelsStruc.Branch {
	var Branch modelsStruc.Branch
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT * FROM branch WHERE id_branch=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	rows := statement.QueryRow(id_int)
	err = rows.Scan(
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
		Branch.Id_branch_enc = "string_enc"
	}
	return Branch
}
