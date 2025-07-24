package controller

import (
	"go-todo-app/initializers"
	"go-todo-app/model"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController interface {
	// Define methods for user operations
	Signup(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUserProfile(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	db *gorm.DB
}

// New creates a new instance of UserController - Constructor function
func NewUserController() UserController {
	return &userController{
		db: initializers.DB,
	}
}

// Implementing UserController Interface methods
func (c *userController) Signup(ctx *gin.Context) {
	// Logic for user signup
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user data"})
		return
	}

	// Check if user already exists
	var existingUser model.User
	if err := c.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		ctx.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	// Create new user
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hash)

	// Save user to the database
	if err := c.db.Create(&user).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(201, gin.H{"message": "User created successfully"})
}

func (c *userController) Login(ctx *gin.Context) {
	// Logic for user login
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid login data"})
		return
	}

	// Check if user exists
	var existingUser model.User
	if err := c.db.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		ctx.JSON(401, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate a jwt token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": existingUser.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
		"iat": time.Now().Unix(),                     // Issued at time
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to generate token - " + err.Error()})
		return
	}

	// Send as a cookie or in the response body
	ctx.SetSameSite(http.SameSiteLaxMode) // Set SameSite attribute for the cookie
	ctx.SetCookie("token", tokenString, 3600*24, "/", "", false, true)

	ctx.JSON(200, gin.H{"message": "Login successful"})
}

func (c *userController) GetUserProfile(ctx *gin.Context) {
	// Logic to get user profile
	userID := ctx.MustGet("userID").(uint) // Get user ID from context
	var user model.User
	if err := c.db.First(&user, userID).Error; err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(200, user)
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	// Logic to delete user
	userID := ctx.MustGet("userID").(uint) // Get user ID from context

	if err := c.db.Delete(&model.User{}, userID).Error; err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}
