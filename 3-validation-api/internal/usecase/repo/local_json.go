package repo

import (
	"encoding/json"
	"errors"
	"os"
)

type LocalJSONRepo struct {
	filename string
}

func NewLocalJSONRepo(filename string) (*LocalJSONRepo, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		emptyMap := make(map[string]string)
		data, err := json.Marshal(emptyMap)
		if err != nil {
			return nil, err
		}
		if err := os.WriteFile(filename, data, 0664); err != nil {
			return nil, err
		}
	}

	return &LocalJSONRepo{filename: filename}, nil
}

func (repo LocalJSONRepo) SaveEmailHash(email, hash string) error {
	data, err := os.ReadFile(repo.filename)
	if err != nil {
		return err
	}

	hashMap := make(map[string]string)
	if len(data) > 0 {
		if err := json.Unmarshal(data, &hashMap); err != nil {
			return err
		}
	}

	hashMap[hash] = email

	updatedData, err := json.MarshalIndent(hashMap, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile(repo.filename, updatedData, 0664)
}

func (repo LocalJSONRepo) GetEmailByHash(hash string) (string, error) {
	hashMap, err := repo.getHashMap()
	if err != nil {
		return "", err
	}

	if email, ok := hashMap[hash]; ok {
		return email, nil
	}

	return "", errors.New("unknown hash")
}

func (repo LocalJSONRepo) getHashMap() (map[string]string, error) {
	data, err := os.ReadFile(repo.filename)
	if err != nil {
		return nil, err
	}

	hashMap := make(map[string]string)
	if len(data) > 0 {
		if err := json.Unmarshal(data, &hashMap); err != nil {
			return nil, err
		}
	}

	return hashMap, nil
}
