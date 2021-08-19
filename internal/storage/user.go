package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dannydd88/gobase/pkg/base"
)

type UserEntry struct {
	Email string `json:"email"`
	Uuid  string `json:"uuid"`
}

type UserData struct {
	Users []UserEntry `json:"users"`
}

func (u *UserData) Add(email, uuid string) error {
	if exists, _ := u.exists(email); exists {
		return fmt.Errorf("duplicate user email -> %s", email)
	}
	u.Users = append(u.Users, UserEntry{Email: email, Uuid: uuid})
	return nil
}

func (u *UserData) Remove(email string) error {
	exists, index := u.exists(email)
	if !exists {
		return fmt.Errorf("do not exists user email -> %s", email)
	}
	u.Users = append(u.Users[:index], u.Users[index+1:]...)
	return nil
}

func (u *UserData) Load(filename string) error {
	filename, err := fpath(filename)
	if err != nil {
		return err
	}
	if base.FileExists(base.String(filename)) {
		raw, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		err = json.Unmarshal(raw, &u)
		if err != nil {
			return err
		}
	} else {
		u.Users = []UserEntry{}
	}
	return nil
}

func (u *UserData) Save(filename string) error {
	filename, err := fpath(filename)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserData) exists(email string) (bool, int) {
	var entry *UserEntry
	var index int
	for i, e := range u.Users {
		if e.Email == email {
			entry = &e
			index = i
			break
		}
	}
	return entry != nil, index
}

func fpath(filename string) (string, error) {
	if !filepath.IsAbs(filename) {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		filename = filepath.Join(dir, filename)
	}
	return filename, nil
}
