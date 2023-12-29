package config

import (
	"database/sql"
	"fmt"
	"main/models"
	"main/secretm"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {

	//SecretModel,err := secretmGetSecret()
	_, err := secretm.GetSecret(os.Getenv("SecretName"))

	return err
}

func DBConnection() error {
	Db, err = sql.Open("Mysql")
}

func ConstStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = ""

	dbUser = keys.UserName
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "shop-gm"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswordstrue", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
