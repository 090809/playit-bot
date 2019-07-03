package user

import (
	"encoding/json"
	"log"
	"os"
)

const defaultStoragePath = "./storage/"

type Repository struct {
	Type        string
	StoragePath string
}

func NewRepository(Type *string, storagePath *string) *Repository {
	var rType string
	var rPath string
	if Type == nil {
		rType = "file"
	} else {
		rType = *Type
	}
	if storagePath == nil {
		rPath = defaultStoragePath
	} else {
		rPath = *storagePath
	}
	return &Repository{Type: rType, StoragePath: rPath}
}

func (r *Repository) Find(id string) *User {
	log.Printf("Finding user: %s", id)

	switch r.Type {
	case "file":
		return r.findFile(id)
	default:
		log.Fatal("[FATAL] Non existing type of repository set!")
	}
	return nil
}

func (r *Repository) findFile(id string) *User {
	file, err := os.Open(r.StoragePath + id + ".json")
	if err != nil {
		log.Printf("[WARNING] Ошибка при попытке открытия файла: %v", err)
		return nil
	}
	defer file.Close()

	user := &User{}

	if err := json.NewDecoder(file).Decode(user); err != nil {
		log.Printf("[ERROR] Ошибка при расшифровке файла пользователя \"%s\": %v", id, err)
	}

	return user
}

func (r *Repository) Save(u *User) {
	switch r.Type {
	case "file":
		r.saveFile(u)
	default:
		log.Fatal("[FATAL] Non existing type of repository set!")
	}
}

func (r *Repository) saveFile(u *User) {
	data, err := json.Marshal(u)
	if err != nil {
		log.Printf("[ERROR] Не получилось обработать пользователя \"%s\"!", u.Id)
		return
	}

	var file *os.File
	file, err = os.OpenFile(r.StoragePath+u.Id+".json", os.O_RDWR|os.O_TRUNC, 0755)
	if err != nil {
		file, err = os.Create(r.StoragePath + u.Id + ".json")
		if err != nil {
			log.Fatalf("Фантастика какая-то, я не смог ни открыть, ни создать файл: %v", err)
		}
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		log.Fatal(err)
	}
}
