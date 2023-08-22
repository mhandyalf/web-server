package database

import (
	"database/sql"
	"web-server/entity"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/avenger")
	if err != nil {
		panic(err)
	}
	return db
}

func GetHeroesFromDB(db *sql.DB) []entity.Hero {
	rows, err := db.Query("SELECT * FROM Heroes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	heroes := []entity.Hero{}
	for rows.Next() {
		var hero entity.Hero
		err := rows.Scan(&hero.ID, &hero.Name, &hero.Universe, &hero.Skill, &hero.ImageURL)
		if err != nil {
			panic(err)
		}
		heroes = append(heroes, hero)
	}

	return heroes
}

func GetVillainsFromDB(db *sql.DB) []entity.Villain {
	rows, err := db.Query("SELECT * FROM Villains")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	villains := []entity.Villain{}
	for rows.Next() {
		var villain entity.Villain
		err := rows.Scan(&villain.ID, &villain.Name, &villain.Universe, &villain.ImageURL)
		if err != nil {
			panic(err)
		}
		villains = append(villains, villain)
	}

	return villains
}
