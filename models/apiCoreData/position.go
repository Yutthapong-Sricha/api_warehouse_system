package models

import (
	"log"
	"strconv"

	"server_go/api_warehouse_system/helpers"
	"server_go/api_warehouse_system/initializers"

	_ "github.com/go-sql-driver/mysql"

	modelsStruc "server_go/api_warehouse_system/models/struc"
)

func ListPosition(act string) []modelsStruc.Position {
	var Positions []modelsStruc.Position
	var err error
	sql_statement := "SELECT id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag FROM staff_position "
	if act == "/active_only" {
		sql_statement = sql_statement + " where is_active_flag=1 "
	}
	list, err := initializers.DB.Query(sql_statement)
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
			id_enc := helpers.Encrypt(strconv.Itoa(row.Id_staff_position))

			row.Id_position_enc = id_enc

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
	statement, err := initializers.DB.Prepare("SELECT id_staff_position, IFNULL(position_name,'') AS position_name, IFNULL(row_order_position,'') AS row_order_position, is_active_flag FROM staff_position WHERE id_staff_position=?")
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
		id_enc := helpers.Encrypt(strconv.Itoa(Position.Id_staff_position))

		Position.Id_position_enc = id_enc

	}
	return Position
}
