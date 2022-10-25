package v2ray

import (
	"context"
	"fmt"

	"v2up/internal"
	"v2up/internal/infra"
	"v2up/internal/storage"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	handler "github.com/v2fly/v2ray-core/v4/app/proxyman/command"
	"github.com/v2fly/v2ray-core/v4/common/protocol"
	"github.com/v2fly/v2ray-core/v4/common/serial"
	"github.com/v2fly/v2ray-core/v4/proxy/vmess"
	"google.golang.org/grpc"
)

type User struct {
	ready    bool
	conn     *grpc.ClientConn
	client   handler.HandlerServiceClient
	userData *storage.UserData
}

var u *User

func init() {
	u = &User{
		ready: false,
	}
}

func UserHandler() *User {
	return u
}

func (u *User) Add(c *cli.Context) error {
	infra.GetLogger().Log("[USER]", "do add user...")

	// ). setup
	err := u.setup()
	if err != nil {
		infra.GetLogger().Error("[USER]", "setup error ->", err)
		return err
	}
	defer u.tearDown()

	// ). check uuid
	var id string
	if c.IsSet(internal.FLAG_USER_UUID) {
		id = c.String(internal.FLAG_USER_UUID)
	} else {
		id = uuid.New().String()
	}

	// ). do add user to v2ray process
	email := c.String(internal.FLAG_USER_EMAIL)
	err = doAddUser(u.client, email, id)
	if err != nil {
		return err
	}
	infra.GetLogger().Log("[USER]", "v2ray add user success")

	// ). add & save user data
	err = u.userData.Add(email, id)
	if err != nil {
		return err
	}
	err = u.userData.Save(infra.GetConfig().User.Storage)
	if err != nil {
		return err
	}
	infra.GetLogger().Log("[USER]", "user data add user success")

	// ). notify
	switch infra.GetConfig().User.Notify.Type {
	case internal.USER_NOTIFY_TYPE_NONE:
		// do nothing
	case internal.USER_NOTIFY_TYPE_SMTP:
		if c.Bool(internal.FLAG_USER_SILENT) {
			break
		}
		tpl := infra.GetConfig().User.Notify.Template
		msg := fmt.Sprintf(tpl, id, email)
		infra.GetLogger().Log("[USER]", "email user ->", msg)
		err = infra.GetMailer().SendMail(
			email,
			infra.GetConfig().User.Notify.Title,
			msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *User) Remove(c *cli.Context) error {
	infra.GetLogger().Log("[USER]", "do remove user...")

	// ). setup
	err := u.setup()
	if err != nil {
		infra.GetLogger().Error("[USER]", "setup error ->", err)
		return err
	}
	defer u.tearDown()

	// ). do remove user to v2ray process
	email := c.String(internal.FLAG_USER_EMAIL)
	err = doRemoveUser(u.client, email)
	if err != nil {
		return err
	}
	infra.GetLogger().Log("[USER]", "v2ray remove user success")

	// ). remove & save user data
	err = u.userData.Remove(email)
	if err != nil {
		return err
	}
	err = u.userData.Save(infra.GetConfig().User.Storage)
	if err != nil {
		return err
	}
	infra.GetLogger().Log("[USER]", "user data remove user success")

	return nil
}

func (u *User) Restore(c *cli.Context) error {
	infra.GetLogger().Log("[USER]", "do restore user...")

	// ). setup
	err := u.setup()
	if err != nil {
		infra.GetLogger().Error("[USER]", "setup error ->", err)
		return err
	}
	defer u.tearDown()

	// ). do add all user to v2ray
	for _, e := range u.userData.Users {
		err = doAddUser(u.client, e.Email, e.Uuid)
		if err != nil {
			return err
		}
	}
	infra.GetLogger().Log("[USER]", "v2ray restore user success")

	return nil
}

func (u *User) setup() error {
	// ). check if is ready
	if u.ready {
		return nil
	}

	// ). setup grpc
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d",
			infra.GetConfig().Api.Host,
			infra.GetConfig().Api.Port),
		grpc.WithInsecure())
	if err != nil {
		return err
	}
	u.conn = conn

	// ). setup client
	u.client = handler.NewHandlerServiceClient(conn)

	// ). load user data
	u.userData = &storage.UserData{}
	err = u.userData.Load(infra.GetConfig().User.Storage)
	if err != nil {
		return err
	}

	// ).  change ready flag
	u.ready = true

	return nil
}

func (u *User) tearDown() {
	// ). check if is ready
	if !u.ready {
		return
	}

	// ). release all
	u.conn.Close()
	u.conn = nil
	u.client = nil
	u.userData = nil
	u.ready = false
}

func doAddUser(c handler.HandlerServiceClient, email, uuid string) error {
	_, err := c.AlterInbound(context.Background(), &handler.AlterInboundRequest{
		Tag: infra.GetConfig().User.Default.Tag,
		Operation: serial.ToTypedMessage(&handler.AddUserOperation{
			User: &protocol.User{
				Level: uint32(infra.GetConfig().User.Default.Level),
				Email: email,
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:               uuid,
					AlterId:          uint32(infra.GetConfig().User.Default.AlterId),
					SecuritySettings: &protocol.SecurityConfig{Type: protocol.SecurityType_AUTO},
				}),
			},
		}),
	})
	if err != nil {
		return err
	}

	infra.GetLogger().Log("[USER]", "Added user ->", email, uuid)
	return nil
}

func doRemoveUser(c handler.HandlerServiceClient, email string) error {
	_, err := c.AlterInbound(context.Background(), &handler.AlterInboundRequest{
		Tag: infra.GetConfig().User.Default.Tag,
		Operation: serial.ToTypedMessage(&handler.RemoveUserOperation{
			Email: email,
		}),
	})

	if err != nil {
		return err
	}

	infra.GetLogger().Log("[USER]", "Removed user ->", email)
	return nil
}
