package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/internal/templates"
	"github.com/blaze-d83/blog-app/types"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RenderLoginPage(c echo.Context, errMsg string) error {
	w := c.Response().Writer
	if errMsg != "" {
		w.Write([]byte(fmt.Sprintf(`<p class="error-message">%s</p>`, errMsg)))
		return nil
	}
	loginPage := templates.LoginPage(errMsg)
	err := loginPage.Render(context.Background(), w)
	if err != nil {
		http.Error(w, "Failed to render login page", http.StatusInternalServerError)
		return err
	}
	return nil
}

func GetLoginPageHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return RenderLoginPage(c, "")
	}
}

func ProcessAdminLoginHandler(dbInstance *db.Database) echo.HandlerFunc {
	return func(c echo.Context) error {

		username := c.FormValue("username")
		password := c.FormValue("password")

		if username == "" || password == "" {
			return RenderLoginPage(c, "Missing username or password")
		}

		var admin types.Admin

		if err := dbInstance.DB.Where("username = ?", username).First(&admin).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return RenderLoginPage(c, "Invalid username or password")
			}
			log.Println("Database error: ", err)
			http.Error(c.Response().Writer, "Internal Server Error", http.StatusInternalServerError)
			return err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(admin.Pass), []byte(password)); err != nil {
			return RenderLoginPage(c, "Invalid username or password")
		}

		sess, err := session.Get("session", c)
		if err != nil {
			log.Println("Session error: ", err)
			http.Error(c.Response().Writer, "Internal server error", http.StatusInternalServerError)
			return err
		}
		sess.Options = &sessions.Options{
			Path: "/",
			MaxAge: 86400 * 7,
			HttpOnly: true,
		}
		sess.Values["admin_id"] = admin.ID
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			log.Println("Failed to save session: ", err)
			return c.String(http.StatusInternalServerError, "Failed to save session")
		}
		return c.Redirect(http.StatusSeeOther, "/admin/dashboard")
	}
}
