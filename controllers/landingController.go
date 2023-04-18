package controllers

import (
	"log"
	"net/http"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"

	"github.com/gin-gonic/gin"
)

func Landing(c *gin.Context) {
	var Branchs []modelsStruc.Branch
	var err error
	sql_statement := "SELECT id_branch, branch_name, IFNULL(address,'') AS address, IFNULL(id_addr_tambon,0) AS id_addr_tambon, IFNULL(id_addr_amphure, 0) AS id_addr_amphure, IFNULL(id_addr_province, 0) AS id_addr_province, IFNULL(branch_tel, '') as branch_tel,  is_active_flag  FROM branch where is_active_flag=1 "
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
		}
		Branchs = append(Branchs, row)
	}
	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}

	c.IndentedJSON(http.StatusOK, gin.H{"branchs": Branchs})
}
