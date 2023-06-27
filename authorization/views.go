package authorization

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/general"
	"net/http"
	"strconv"
	"time"
)

func Register(g *gin.Context) {
	var user Account
	err := g.ShouldBindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, general.ResponseFailure("请求内容错误", err))
		return
	}

	user.PasswordHash = HashAndSalt(user.Password)
	createUser, err := CreateUser(user.Users)
	if err != nil {
		g.JSON(http.StatusBadRequest, general.ResponseFailure("数据入库时出现错误", err))
		return
	}

	g.JSON(http.StatusOK, general.ResponseSuccessful("注册成功", createUser))
}

func Login(g *gin.Context) {
	var user Account
	err := g.ShouldBindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, general.ResponseFailure("请求内容错误", err))
		return
	}

	userObject, err := QueryUserByUsername(user.Username)
	if err != nil {
		g.JSON(http.StatusUnauthorized, general.ResponseFailure("账号不存在", err))
		return
	}

	if ComparePassword(userObject.PasswordHash, user.Password) {
		expirationTime := time.Now().Add(60 * time.Minute) //到期时间
		claims := &Claims{
			Username: userObject.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(JwtKey)
		if err != nil {
			g.JSON(http.StatusInternalServerError, general.ResponseFailure("服务器内部错误", err))
			return
		}

		CreateSession(userObject.Username, tokenString)

		g.JSON(http.StatusOK, general.ResponseSuccessful("登录成功", map[string]string{
			"token": tokenString,
		}))
		return
	} else {
		g.JSON(http.StatusUnauthorized, general.ResponseFailure("密码错误", nil))
		return
	}
}

func Logout(g *gin.Context) {
	var tokenString string
	tokenStringList, _ := g.Request.Header["Authorization"]
	tokenString = tokenStringList[0]

	fmt.Println(g.Request.Header["Authorization"])

	//前端会有特殊情况下携带空字符串过来，这里判断一下，只有token不为空的情况下，才会删除session表中的token
	if len(tokenString) != 0 {
		DeleteSession(tokenString)
	}

	g.JSON(http.StatusOK, general.ResponseSuccessful("登出成功", true))
}

func PostUser(g *gin.Context) {
	var body *Users
	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, "提交的数据错误")
		return
	}

	user, err := CreateUser(body)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, user)
}

func DelUser(g *gin.Context) {
	id := g.Param("id")
	id2, _ := strconv.Atoi(id)
	delUser, err := DeleteUser(id2)
	if err != nil {
		return
	}
	var date string
	if delUser == 1 {
		date = fmt.Sprintf("已成功删除用户，id:%s", id)
	} else {
		date = fmt.Sprintf("未查询到用户，id:%s", id)
	}
	g.JSON(http.StatusOK, date)
}

func PutUser(g *gin.Context) {
	var body Users

	id := g.Param("id")
	id2, _ := strconv.Atoi(id)

	err := g.ShouldBindJSON(&body)
	if err != nil {
		return
	}

	user, err := UpdateUser(id2, &body)
	if err != nil {
		return
	}

	g.JSON(http.StatusOK, user)
}

func ListUsers(g *gin.Context) {
	var resp *[]Users
	resp, _ = ListUser()
	g.JSON(http.StatusOK, resp)
}

func GetUsers(g *gin.Context) {
	id := g.Param("id")
	id2, _ := strconv.Atoi(id)
	user, err := GetUserById(id2)
	if err != nil {
		return
	}
	g.JSON(http.StatusOK, user)
}
