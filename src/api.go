package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
)

var (
	//MongoSession s
	MongoSession *mgo.Session
	//UsersCollection s
	UsersCollection *mgo.Collection
)

//User object user data
type User struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

//SaveToDB s
func (u *User) SaveToDB() error {
	err := UsersCollection.Insert(&u)

	if err != nil {
		return err
	}
	return nil
}

//ReadFromDB r
func (u *User) ReadFromDB() ([]User, error) {
	result := []User{}
	err := UsersCollection.Find(nil).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func getUsers(c echo.Context) error {

	user := new(User)
	result, err := user.ReadFromDB()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func getUserByID(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	user := new(User)
	err := c.Bind(user)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	user.SaveToDB()
	return c.NoContent(http.StatusCreated)
}

func init() {
	MongoSession, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	MongoSession.SetMode(mgo.Monotonic, true)
	UsersCollection = MongoSession.DB("maejo").C("users")
}

func main() {
	defer MongoSession.Close()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", index)
	e.GET("/users", getUsers)
	e.GET("/users/:id", getUserByID)
	e.POST("/users", saveUser)

	e.Logger.Fatal(e.Start(":1323"))
}
