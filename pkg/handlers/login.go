package handlers

 import (
// 	"context"
// 	"log"
// 	"net/http"
//
 	"github.com/blaze-d83/blog-app/internal/db"
 	// "github.com/blaze-d83/blog-app/types"
// 	"github.com/blaze-d83/blog-app/utils"
// 	"github.com/gorilla/sessions"
// 	"github.com/labstack/echo-contrib/session"
// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
 )

type LoginRepository struct {
	db *db.Database
}

func NewLoginHandler(db *db.Database) *LoginRepository {
	return &LoginRepository{db: db}
}

// func RenderLoginPage(c echo.Context, errMsg string) error {
// 	if c.Request().Header.Get("HX-Request") == "true"{
// 		w := c.Response().Writer
// 		loginForm := components.LoginForm(errMsg)
// 		err := loginForm.Render(context.Background(), w)
// 		if err != nil {
// 			http.Error(w, "Failed to render login form", http.StatusInternalServerError)
// 			return err
// 		}
// 		return nil
// 	}
// 	w := c.Response().Writer
// 	loginPage := components.LoginPage()
// 	err := loginPage.Render(context.Background(), w)
// 	if err != nil {
// 		http.Error(w, "Failed to render login page", http.StatusInternalServerError)
// 		return err
// 	}
// 	return nil
// }

// func (repo *LoginRepository) GetLoginPageHandler() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return RenderLoginPage(c, "")
// 	}
// }
//
// func (repo *LoginRepository) ProcessAdminLoginHandler() echo.HandlerFunc {
// 	return func(c echo.Context) error {
//
// 		username := c.FormValue("username")
// 		password := c.FormValue("password")
//
// 		if username == "" || password == "" {
// 			c.Response().WriteHeader(http.StatusBadRequest)
// 			return RenderLoginPage(c, "Missing username or password")
// 		}
//
// 		var admin types.Admin
//
// 		if err := repo.db.DB.Where("username = ?", username).First(&admin).Error; err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				c.Response().WriteHeader(http.StatusBadRequest)
// 				return RenderLoginPage(c, "Invalid username or password")
// 			}
// 			log.Println("Database error: ", err)
// 			http.Error(c.Response().Writer, "Internal Server Error", http.StatusInternalServerError)
// 			return err
// 		}
//
// 		if err := utils.CompareHashPassword(&admin, password); err != nil {
// 			c.Response().WriteHeader(http.StatusBadRequest)
// 			return RenderLoginPage(c, "Invalid username or password")
// 		}
//
// 		sess, err := session.Get("session", c)
// 		if err != nil {
// 			log.Println("Session error: ", err)
// 			c.Response().WriteHeader(http.StatusBadRequest)
// 			http.Error(c.Response().Writer, "Internal server error", http.StatusInternalServerError)
// 			return err
// 		}
// 		sess.Options = &sessions.Options{
// 			Path:     "/",
// 			MaxAge:   86400 * 7,
// 			HttpOnly: true,
// 		}
// 		sess.Values["admin_id"] = admin.ID
// 		if err := sess.Save(c.Request(), c.Response()); err != nil {
// 			log.Println("Failed to save session: ", err)
// 			return c.String(http.StatusInternalServerError, "Failed to save session")
// 		}
// 		return c.NoContent(http.StatusOK) 
// 	}
// }
