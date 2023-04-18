package models

import (
	"errors"
	"log"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func InstShortCust(c *gin.Context) (string, error) {
	now := time.Now()
	var json modelsStruc.ShortCust
	var ShortCust modelsStruc.ShortCust
	var err error
	if err := c.ShouldBindJSON(&json); err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	checkdup, err_chk := initializers.DB.Prepare("SELECT  cust_name, cust_lastname, cust_mobile, cust_tax_id FROM customer WHERE Cust_tax_id=?")
	if err_chk != nil {
		log.Fatal(err_chk.Error())
		return "", err_chk
	}
	row_dup := checkdup.QueryRow(json.Cust_tax_id)
	err = row_dup.Scan(
		&ShortCust.Cust_name,
		&ShortCust.Cust_lastname,
		&ShortCust.Cust_mobile,
		&ShortCust.Cust_tax_id,
	)
	if err == nil {
		myError := errors.New("เลขที่บัตรประชาชน ซ้ำในระบบ")
		return ShortCust.Cust_tax_id, myError
	}

	var timestamp = strconv.FormatInt(now.Unix(), 10)
	var fullname string = json.Cust_name + " " + json.Cust_lastname

	statement := "INSERT INTO `customer` ( cust_name, cust_lastname, cust_fullname, cust_mobile, cust_tax_id, record_create_time, record_update_time ) VALUES (?,?,?,?,?,?,?)"
	inst, err := initializers.DB.Prepare(statement)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	result, err := inst.Exec(json.Cust_name, json.Cust_lastname, fullname, json.Cust_mobile, json.Cust_tax_id, timestamp, timestamp)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	defer inst.Close()
	_, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	return fullname, nil
}
