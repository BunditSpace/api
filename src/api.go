package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"./models"

	db "./helper/db"
)

const (
	APISERVER          = ":1323"
	DATABASESERVER     = "localhost:27017"
	DATABASENAME       = "maejo"
	DATABASECOLLECTION = "users"
)

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func getUsers(c echo.Context) error {

	user := new(models.User)
	result, err := user.ReadFromDB()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func getUserByID(c echo.Context) error {
	user := new(models.User)
	id := c.Param("id")
	user.ID = bson.ObjectIdHex(id)
	result, _ := user.ReadByID()
	return c.JSON(http.StatusOK, result)
}

func deleteUserByID(c echo.Context) error {
	user := new(models.User)
	id := c.Param("id")
	user.ID = bson.ObjectIdHex(id)
	user.DeleteByID()
	return c.NoContent(http.StatusOK)
}

func saveUser(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.SaveToDB()
	return c.NoContent(http.StatusCreated)
}

func updateUserByID(c echo.Context) error {
	user := new(models.User)
	id := c.Param("id")
	user.ID = bson.ObjectIdHex(id)

	err := c.Bind(user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.UpdateByID()
	return c.NoContent(http.StatusOK)
}

func login(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	result, err := user.Login()
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, result)
}

func init() {
	mongoSession, err := mgo.Dial(DATABASESERVER)
	if err != nil {
		panic(err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)
	db.MongoSession = mongoSession
	db.UsersCollection = db.MongoSession.DB(DATABASENAME).C(DATABASECOLLECTION)
}

func main() {
	defer db.MongoSession.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/", index)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserByID)
	e.POST("/users", saveUser)
	e.DELETE("/users/:id", deleteUserByID)
	e.PUT("/users/:id", updateUserByID)
	e.POST("/login", login)

	e.Logger.Fatal(e.Start(APISERVER))
}
