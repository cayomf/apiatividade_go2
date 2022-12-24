package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const dbuser = "root"
const dbpass = "root"
const dbname = "atividade_db"

func GetEntregas() []Entrega {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	// if there is an error opening the connection, handle it
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()
	results, err := db.Query("SELECT * FROM entrega")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	entregas := []Entrega{}
	for results.Next() {
		var entrega Entrega
		// for each row, scan into the Entrega struct
		err = results.Scan(&entrega.Id, &entrega.AlunoId, &entrega.AtividadeId, &entrega.Nota, &entrega.Disciplina)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// append the entrega into entregas array
		entregas = append(entregas, entrega)
	}

	return entregas

}

func GetEntrega(id int) *Entrega {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	entrega := &Entrega{}
	if err != nil {
		// simply print the error to the console
		fmt.Println("Err", err.Error())
		// returns nil on error
		return nil
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM entrega where id=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&entrega.Id, &entrega.AlunoId, &entrega.AtividadeId, &entrega.Nota, &entrega.Disciplina)
		if err != nil {
			return nil
		}
	} else {

		return nil
	}

	return entrega
}

func AddEntrega(entrega Entrega) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	insert, err := db.Query(
		"INSERT INTO entrega (id,aluno_id,atividade_id,nota,disciplina) VALUES (?,?,?,?,?)",
		entrega.Id, entrega.AlunoId, entrega.AtividadeId, entrega.Nota, entrega.Disciplina)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

}

func UpdateEntrega(id int, nota string) {

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)

	if err != nil {
		panic(err.Error())
	}

	// defer the close till after this function has finished
	// executing
	defer db.Close()

	update, err := db.Query(
		"UPDATE entrega SET nota=? WHERE id=?",
		nota, id)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	defer update.Close()
}
