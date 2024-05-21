package authorization

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gomessage/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

// Register 创建或注册用户
func Register(g *gin.Context) {
	var user Account
	err := g.ShouldBindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}

	user.PasswordHash = HashAndSalt(user.Password)
	createUser, err := CreateUser(user.Users)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("数据入库时出现错误", err))
		return
	}

	g.JSON(http.StatusOK, utils.ResponseSuccessful("注册成功", createUser))
}

// Login 登录功能
func Login(g *gin.Context) {
	var user Account
	err := g.ShouldBindJSON(&user)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}

	userObject, err := QueryUserByUsername(user.Username)
	if err != nil {
		g.JSON(http.StatusUnauthorized, utils.ResponseFailure("账号不存在", err))
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
			g.JSON(http.StatusInternalServerError, utils.ResponseFailure("服务器内部错误", err))
			return
		}

		CreateSession(userObject.Username, tokenString)

		g.JSON(http.StatusOK, utils.ResponseSuccessful("登录成功", map[string]any{
			"id":    userObject.ID,
			"token": tokenString,
		}))
		return
	} else {
		g.JSON(http.StatusUnauthorized, utils.ResponseFailure("密码错误", nil))
		return
	}
}

// Logout 注销和登出
func Logout(g *gin.Context) {
	var tokenString string
	tokenStringList, _ := g.Request.Header["Authorization"]
	tokenString = tokenStringList[0]

	fmt.Println(g.Request.Header["Authorization"])

	//前端会有特殊情况下携带空字符串过来，这里判断一下，只有token不为空的情况下，才会删除session表中的token
	if len(tokenString) != 0 {
		DeleteSession(tokenString)
	}

	g.JSON(http.StatusOK, utils.ResponseSuccessful("登出成功", true))
}

type Pass struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// UpdatePassword 修改密码
func UpdatePassword(g *gin.Context) {
	//获取密码（这里没有校验老密码，是因为方便用户真的记不起来了，依然可以修改密码成功）
	var password Pass
	err := g.ShouldBindJSON(&password)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}

	//获取用户对象
	userID, _ := strconv.Atoi(g.Param("id"))
	user, err := GetUserById(userID)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("请求内容错误", err))
		return
	}

	//直接修改用户密码（这里没有校验老密码，是因为方便用户真的记不起来了，依然可以修改密码成功）
	//普通人直接调用接口无法访问，只有GoMessage前端加密过来，才能调用成功
	user.PasswordHash = HashAndSalt(password.NewPassword)
	updateUser, err := UpdateUserPassword(user.ID, user)
	if err != nil {
		g.JSON(http.StatusBadRequest, utils.ResponseFailure("数据入库时出现错误", err))
		return
	}
	g.JSON(http.StatusOK, utils.ResponseSuccessful("修改成功", updateUser))
}

//func PostUser(g *gin.Context) {
//	var body *Users
//	err := g.ShouldBindJSON(&body)
//	if err != nil {
//		g.JSON(http.StatusBadRequest, "提交的数据错误")
//		return
//	}
//
//	user, err := CreateUser(body)
//	if err != nil {
//		return
//	}
//
//	g.JSON(http.StatusOK, user)
//}
//
//func DelUser(g *gin.Context) {
//	id := g.Param("id")
//	id2, _ := strconv.Atoi(id)
//	delUser, err := DeleteUser(id2)
//	if err != nil {
//		return
//	}
//	var date string
//	if delUser == 1 {
//		date = fmt.Sprintf("已成功删除用户，id:%s", id)
//	} else {
//		date = fmt.Sprintf("未查询到用户，id:%s", id)
//	}
//	g.JSON(http.StatusOK, date)
//}
//
//func ListUsers(g *gin.Context) {
//	var resp *[]Users
//	resp, _ = ListUser()
//	g.JSON(http.StatusOK, resp)
//}
//
//func GetUsers(g *gin.Context) {
//	id := g.Param("id")
//	id2, _ := strconv.Atoi(id)
//	user, err := GetUserById(id2)
//	if err != nil {
//		return
//	}
//	g.JSON(http.StatusOK, user)
//}
//func CreateUserView(g *gin.Context) {
//	var user Users
//	err := g.ShouldBindJSON(&user)
//	if err != nil {
//		g.JSON(http.StatusBadRequest, gin.H{"error": "请求内容错误"})
//		return
//	}
//	createdUser, err := CreateUser(&user)
//	if err != nil {
//		g.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
//		return
//	}
//	g.JSON(http.StatusOK, createdUser)
//}
//
//func DeleteUserView(g *gin.Context) {
//	id := g.Param("id")
//	id2, _ := strconv.Atoi(id)
//	rowsAffected, err := DeleteUser(id2)
//	if err != nil || rowsAffected == 0 {
//		g.JSON(http.StatusInternalServerError, gin.H{"error": "删除用户失败"})
//		return
//	}
//	g.JSON(http.StatusOK, gin.H{"message": "用户删除成功"})
//}
//
//func UpdateUserView(g *gin.Context) {
//	var user Users
//	id := g.Param("id")
//	id2, _ := strconv.Atoi(id)
//	err := g.ShouldBindJSON(&user)
//	if err != nil {
//		g.JSON(http.StatusBadRequest, gin.H{"error": "请求内容错误"})
//		return
//	}
//	updatedUser, err := UpdateUser(id2, &user)
//	if err != nil {
//		g.JSON(http.StatusInternalServerError, gin.H{"error": "更新用户失败"})
//		return
//	}
//	g.JSON(http.StatusOK, updatedUser)
//}
//
//func GetUserView(g *gin.Context) {
//	id := g.Param("id")
//	id2, _ := strconv.Atoi(id)
//	user, err := GetUserById(id2)
//	if err != nil {
//		g.JSON(http.StatusInternalServerError, gin.H{"error": "获取用户信息失败"})
//		return
//	}
//	g.JSON(http.StatusOK, user)
//}
