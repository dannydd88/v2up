package user

import (
	"fmt"

	"github.com/dannydd88/v2up/internal/infra"
	"github.com/urfave/cli/v2"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) Add(c *cli.Context) error {
	infra.GetLogger().Log("[USER]", "add")
	return nil
}

func (u *User) Remove(c *cli.Context) error {
	fmt.Println("remove")
	return nil
}
