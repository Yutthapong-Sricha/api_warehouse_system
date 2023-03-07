package models

import (
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
)

func ListPosition() []modelsStruc.Position {
	var Positions []modelsStruc.Position
	var err error

	list, err := initializers.DB.Query("SELECT id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag FROM staff_position")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer list.Close()
	for list.Next() {
		var row modelsStruc.Position
		err = list.Scan(
			&row.Id_staff_position,
			&row.Position_name,
			&row.Row_order_position,
			&row.Is_active_flag,
		)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			row.Id_position_enc = "string_enc"
		}
		Positions = append(Positions, row)
	}

	err = list.Err()
	if err != nil {
		log.Fatal(err.Error())
	}
	return Positions
}

func GetPosition(id string) modelsStruc.Position {
	var Position modelsStruc.Position
	var err error
	id_int, conv_err := strconv.ParseInt(id, 10, 0)
	if conv_err != nil {
		panic(err.Error())
	}
	statement, err := initializers.DB.Prepare("SELECT * FROM staff_position WHERE id_staff_position=?")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer statement.Close()
	rows := statement.QueryRow(id_int)
	err = rows.Scan(
		&Position.Id_staff_position,
		&Position.Position_name,
		&Position.Row_order_position,
		&Position.Is_active_flag,
	)
	if err != nil {
		return Position
	} else {
		Position.Id_position_enc = "string_enc"
	}
	return Position
}
