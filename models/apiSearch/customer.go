package models

import (
	"log"
	"server_go/api_warehouse_system/initializers"
	modelsStruc "server_go/api_warehouse_system/models/struc"

	"unicode/utf8"

	"github.com/gin-gonic/gin"
)

func Cust(c *gin.Context) []modelsStruc.Searchcust {
	var json modelsStruc.Keyword
	var listsearch []modelsStruc.Searchcust
	if err := c.ShouldBindJSON(&json); err != nil {
		//c.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return listsearch
	}
	var keysearch = json.Keysearch
	//fmt.Println(keysearch)
	//fmt.Println(utf8.RuneCountInString(keysearch))
	if utf8.RuneCountInString(keysearch) >= 4 {
		var err error
		sql_statement := "SELECT id_customer AS 'value', IFNULL(cust_code_spc,'') AS cust_code_spc, CONCAT(cust_fullname,' (',IFNULL(cust_mobile,''),' ',IFNULL(cust_tel,''),')') AS 'title', IFNULL(cust_nickname,'') AS cust_nickname, IFNULL(cust_tax_id,'') AS cust_tax_id, IFNULL(cust_mobile,'') AS cust_mobile, IFNULL(cust_tel,'') AS cust_tel, IFNULL(cust_email,'') AS cust_email FROM customer WHERE CONCAT(IFNULL(cust_code_spc,''),' ',cust_fullname,' ', IFNULL(cust_nickname,''), ' ',IFNULL(cust_tax_id,''),' ',IFNULL(cust_mobile,''),' ',IFNULL(cust_tel,''),' ',IFNULL(cust_email,'')) LIKE '%" + keysearch + "%' "
		list, query_err := initializers.DB.Query(sql_statement)
		if query_err != nil {
			log.Fatal(query_err.Error())
		}
		defer list.Close()
		for list.Next() {
			var row modelsStruc.Searchcust
			err = list.Scan(
				&row.Value,
				&row.Cust_code_spc,
				&row.Title,
				&row.Cust_nickname,
				&row.Cust_tax_id,
				&row.Cust_mobile,
				&row.Cust_tel,
				&row.Cust_email,
			)
			if err != nil {
				log.Fatal(err.Error())
			}
			listsearch = append(listsearch, row)
		}
		err = list.Err()
		if err != nil {
			log.Fatal(err.Error())
		}
		return listsearch
	}
	return listsearch
}
