package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pierangelo1982/go-experiment/delbarbaStoreJsonInDB/model"
)

func ParseData(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		x := "1900-01-01"
		tm, err := time.Parse("2006-01-02", x)
		if err != nil {
			fmt.Println(err)
		}
		return tm
	}
	return t
}

func returnIdCustomerDB(codice int64) int {
	var id int
	db2, err := sql.Open("mysql", "root:alnitek82@tcp(0.0.0.0:3310)/delbarba_backend_development")
	if err != nil {
		panic(err.Error())
	}
	defer db2.Close()
	sqlStatement := fmt.Sprintf("SELECT id FROM customers WHERE customer_code=%d", codice)
	fmt.Println(sqlStatement)
	_ = db2.QueryRow(sqlStatement).Scan(&id)
	fmt.Println(id)
	return id
}

func GetCustomerDetail(id string, token string) {
	url := fmt.Sprintf("http://api.fintyreclub.it/gommista/users/%s?token=%s", id, token)
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(body)
	var customers = model.Customer{}
	err = json.Unmarshal([]byte(body), &customers)
	if err != nil {
		fmt.Println(err)
	}
	db, err := sql.Open("mysql", "root:alnitek82@tcp(0.0.0.0:3310)/delbarba_backend_development")
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(customers)
	isAPP := 1
	for i := range customers.Veicoli {
		fmt.Println(customers.Veicoli[i].Marca)
		var x string = customers.Veicoli[i].KmAttuali
		kmTot, err := strconv.ParseInt(x, 10, 64)
		myID, err := strconv.ParseInt(customers.ID, 10, 64)
		tmpId := returnIdCustomerDB(myID)
		_, err = db.Exec("INSERT INTO vehicles (customer_id, vehicle_code, id_cliente, brand, model, annual_km, total_km, plate, type_vehicle, is_app) VALUES" + fmt.Sprintf("('%d', '%s', '%s', '%s', '%s', '%s', '%d', '%s', '%s', '%d')", tmpId, customers.Veicoli[i].ID, customers.Veicoli[i].IDCliente, customers.Veicoli[i].Marca, customers.Veicoli[i].Modello, customers.Veicoli[i].PercorrenzaAnnua, kmTot, customers.Veicoli[i].Targa, customers.Veicoli[i].TipoVeicolo, isAPP))
		if err != nil {
			fmt.Println(err)
		}
	}
	db.Close()
}
