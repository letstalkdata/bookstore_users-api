package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/letstalkdata/bookstore_users-api/domain/users"
	"github.com/letstalkdata/bookstore_users-api/services"
	"github.com/letstalkdata/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User

	//fmt.Println(user)
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//TODO: Handle Error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	// TODO: Handle JSON error
	//	fmt.Println(err.Error())
	//	return
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
