package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/internal/templates"
	"github.com/blaze-d83/blog-app/types"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RenderLoginPage(c echo.Context, errorMessage, username string) error  {
	w := c.Response().Writer

	loginPage := templates.LoginPage(errorMessage, username, "")
	err := loginPage.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "Failed to render login page", http.StatusInternalServerError)
		return  err
	}
	return nil
}

func AdminLoginHandler(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == http.MethodGet {
			return RenderLoginPage(c, "", "")
		}
		if c.Request().Method == http.MethodPost {
			return ProcessAdminLogin(c, dbInstance)
		}
		http.Error(c.Response().Writer, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}
}

func ProcessAdminLogin(c echo.Context, dbInstance *db.Database) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	
	if username == "" || password == "" {
		return RenderLoginPage(c, "Username or password missing", username)
	}

	var admin types.Admin

	if err := dbInstance.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return RenderLoginPage(c, "Invalid username or password", username)
		}
		log.Println("Database error: ", err)
		http.Error(c.Response().Writer, "Internal Server Error", http.StatusInternalServerError)
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Pass), []byte(password)); err != nil {
		return RenderLoginPage(c, "Invalid username or password", username)
	}

	return c.Redirect(http.StatusSeeOther, "/admin/dashboard")

}
