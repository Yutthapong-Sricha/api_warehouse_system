package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
)

var DB *sql.DB

func ListPosition() []modelsStruc.Position {
	DB = initializers.DB
	var Positions []modelsStruc.Position
	var err error

	lists, err := DB.Query("select id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag from staff_position")
	if err != nil {
		log.Fatal(err)
	}
	defer lists.Close()
	for lists.Next() {
		var row modelsStruc.Position
		err := lists.Scan(
			&row.Id_staff_position,
			&row.Position_name,
			&row.Row_order_position,
			&row.Is_active_flag,
		)
		if err != nil {
			log.Fatal(err)
		}
		Positions = append(Positions, row)
	}

	err = lists.Err()
	if err != nil {
		log.Fatal(err)
	}
	return Positions
}
