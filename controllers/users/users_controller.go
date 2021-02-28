package users

import (
	"github.com/danielgom/bookstore_usersapi/domain/users"
	"github.com/danielgom/bookstore_usersapi/services"
	"github.com/danielgom/bookstore_usersapi/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getUserId(uIdParam string) (int64, *errors.RestErr) {
	pInt, err := strconv.ParseInt(uIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("User id should be a valid number")
	}
	return pInt, nil
}

func GetById(c *gin.Context) {

	id, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	r, getErr := services.UserService.GetUserById(id)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, r.Marshaller(c.GetHeader("X-Public") == "true"))

}

func Create(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	r, saveErr := services.UserService.CreateUser(&user)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, r.Marshaller(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	r, updateErr := services.UserService.UpdateUser(&user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}

	c.JSON(http.StatusOK, r.Marshaller(c.GetHeader("X-Public") == "true"))
}

func UpdatePartial(c *gin.Context) {
	userId, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	r, partUpdateErr := services.UserService.UpdateUserPartial(&user)
	if partUpdateErr != nil {
		c.JSON(partUpdateErr.Status, partUpdateErr)
		return
	}

	c.JSON(http.StatusOK, r.Marshaller(c.GetHeader("X-Public") == "true"))

}

func Delete(c *gin.Context) {
	id, idErr := getUserId(c.Param("userId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if deleteErr := services.UserService.DeleteUser(id); deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	userList, err := services.UserService.Search(status)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, userList.Marshaller(c.GetHeader("X-Public") == "true"))
}
