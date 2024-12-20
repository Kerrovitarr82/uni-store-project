package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userRole := c.GetString("user_role") //берется из токена
	err = nil
	if userRole != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

// Проверка уровня доступа для юзера
func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userRole := c.GetString("user_role") //берется из токена
	uid := c.GetString("uid")            //берется из токена
	err = nil

	if userRole == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	err = CheckUserType(c, userRole)
	return err
}
