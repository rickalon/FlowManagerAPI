package domain

import (
	"errors"

	"github.com/rickalon/FlowManagerAPI/internal/repositories"
)

type Proyect struct {
	Proyect_id int    `json:"proyect_id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
}

func ValidateProyect(p *Proyect) error {
	if p.Name == "" {
		return errors.New("you have to name the proyect")
	}
	return nil
}

func CreateProyect(db *repositories.PqDB, p *Proyect) error {
	_, err := db.DB.Exec("INSERT INTO PROYECTS(name) VALUES ($1);", p.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetProyectByName(db *repositories.PqDB, p *Proyect) error {
	err := db.DB.QueryRow("SELECT proyect_id,created_at from PROYECTS where name=$1", p.Name).Scan(&p.Proyect_id, &p.Created_at)
	if err != nil {
		return err
	}
	return nil
}

func GetProyectById(db *repositories.PqDB, p *Proyect) error {
	err := db.DB.QueryRow("SELECT proyect_id,name,created_at from PROYECTS where proyect_id=$1", p.Proyect_id).Scan(&p.Proyect_id, &p.Name, &p.Created_at)
	if err != nil {
		return err
	}
	return nil
}

func RemoveProyect(db *repositories.PqDB, p *Proyect) error {
	_, err := db.DB.Exec("DELETE FROM PROYECTS where proyect_id=$1", p.Proyect_id)
	if err != nil {
		return err
	}
	return nil
}
