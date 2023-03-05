package models

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
)

var DB *sql.DB

func ListPosition() []modelsStruc.Position {
	DB = initializers.DB
	var Positions []modelsStruc.Position
	var err error

	lists, err := DB.Query("SELECT id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag FROM staff_position")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer lists.Close()
	for lists.Next() {
		var row modelsStruc.Position
		err = lists.Scan(
			&row.Id_staff_position,
			&row.Position_name,
			&row.Row_order_position,
			&row.Is_active_flag,
		)
		if err != nil {
			log.Fatal(err.Error())
		}
		Positions = append(Positions, row)
	}

	err = lists.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Positions
}

func GetPosition(id string) modelsStruc.Position {
	DB = initializers.DB
	var Position modelsStruc.Position
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	rows := DB.QueryRow("SELECT * FROM staff_position WHERE id_staff_position=?", id_int)
	err = rows.Scan(
		&Position.Id_staff_position,
		&Position.Position_name,
		&Position.Row_order_position,
		&Position.Is_active_flag,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return Position
}
