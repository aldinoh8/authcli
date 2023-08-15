package test

import (
	"database/sql"
	"testing"

	"github.com/aldinoh8/authcli/config"
	"github.com/aldinoh8/authcli/handler"

	"github.com/stretchr/testify/assert"
)

var (
	db          *sql.DB
	userHandler handler.User
)

func TestMain(m *testing.M) {
	db = config.InitDatabase("mysql:mysql@tcp(localhost:3306)/ftgo_p1_bcypt_demo_test")
	userHandler = handler.User{DB: db}
	// handler

	m.Run()

	// rapihin semua test nya
	db.Exec("TRUNCATE TABLE users;")
}

func TestRegister(t *testing.T) {
	t.Run("success register", func(t *testing.T) {
		err := userHandler.Register("test124", "test124", 12)
		assert.Nil(t, err)
	})

	t.Run("failed register because username is not unique", func(t *testing.T) {
		err := userHandler.Register("test124", "test124", 12)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	userHandler.Register("testLogin", "123456", 12)
	t.Run("success login", func(t *testing.T) {
		user, err := userHandler.Login("testLogin", "123456")

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, user.Username, "testLogin")
		assert.Equal(t, user.Age, 12)
	})

	t.Run("failed login username not found", func(t *testing.T) {
		_, err := userHandler.Login("usergaada", "salah")

		assert.NotNil(t, err)
		assert.Equal(t, "username/password invalid", err.Error())
	})

	t.Run("failed login invalid password", func(t *testing.T) {
		_, err := userHandler.Login("testLogin", "salah")

		assert.NotNil(t, err)
	})
}

func TestAddMenu(t *testing.T) {
	t.Run("success", func(t *testing.T) {})
	t.Run("failed", func(t *testing.T) {})
}
