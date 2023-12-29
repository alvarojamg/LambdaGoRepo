package config

import (
	"fmt"
	"main/models"
	"main/tools"

	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SiginUp) error {
	fmt.Println("Comienza Registro")

	err := DBConnection()
	if err != nil {
		return err
	}
	defer Db.Close()
	sentence := "INSERT INTO users (User_Email, User_Uuid, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUId + "','" + tools.DateMySql() + "')"
	fmt.Println(sentence)
	_, err = Db.Exec(sentence)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("SigInup Succesful")
	return nil
}
