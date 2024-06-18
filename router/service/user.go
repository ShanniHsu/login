package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	create_user "login/content/create-user"
	deposit_amount "login/content/deposit-amount"
	forget_password "login/content/forget-password"
	get_user_info "login/content/get-user-info"
	login_by_account "login/content/login-by-account"
	reset_forget_password "login/content/reset-forget-password"
	token_change "login/content/token-change"
	withdraw_amount "login/content/withdraw-amount"
	"login/model"
	"login/pkg/jwt"
	"login/pkg/mail"
	"login/pkg/rule"
	"login/router/middleware"
	"login/router/repository"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UserService interface {
	Register(req *create_user.Request) (err error)
	ForgetPassword(req *forget_password.Request) (err error)
	ResetForgetPassword(req *reset_forget_password.Request) (err error)
	Login(ctx *gin.Context, req *login_by_account.Request) (err error)
	//auth
	BuildTempToken(ctx *gin.Context) (url string, err error)
	TokenChange(req *token_change.Request) (token string, err error)
	GetUserInfo(req *get_user_info.Request) (res get_user_info.Response, err error)
	DepositAmount(req *deposit_amount.Request) (res deposit_amount.Response, err error)
	WithdrawAmount(req *withdraw_amount.Request) (res withdraw_amount.Response, err error)
	Logout(ctx *gin.Context) (err error)
}

type userService struct {
	repo repository.Repo
}

func NewUserService(repo repository.Repo) UserService {
	return userService{
		repo: repo,
	}
}

func (s userService) Register(req *create_user.Request) (err error) {
	//取註冊資料
	//lastName := ctx.PostForm("last_name")
	//firstName := ctx.PostForm("first_name")
	//nickName := ctx.PostForm("nick_name")
	//account := ctx.PostForm("account")
	//password := ctx.PostForm("password")
	//email := ctx.PostForm("email")
	//gender := ctx.PostForm("gender")
	fmt.Println("test: ", req)

	//檢查此帳號是否存在
	user, err := s.repo.UserRepository.GetByAccount(req.Account)
	if user.ID != 0 {
		err = errors.New("The account is existed!")
		return
	}

	//檢查此信箱是否存在
	user, err = s.repo.UserRepository.GetByEmail(req.Email)
	if user.ID != 0 {
		err = errors.New("The E-mail is existed!")
		return
	}

	//model.User型態
	var data = new(model.User)
	data.LastName = req.LastName
	data.FirstName = req.FirstName
	data.NickName = req.NickName
	data.Account = req.Account
	data.Password = rule.HashPassword(req.Password)
	data.Email = req.Email
	data.Gender = req.Gender
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	//Insert into Users
	err = s.repo.UserRepository.Create(data)
	if err != nil {
		err = errors.New("Create Failed!")
		return
	}
	return

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "Register successfully!",
	//})

	//確定註冊資料無誤
	//req := new(create_user.Request)
	//err := ctx.BindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, req)
	//	return
	//}
	//
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, "Parameter Error")
	//	return
	//}
	//user := new(model.User)
	////檢查此帳號是否存在
	//user, err = repository.GetByAccount(req.Account)
	//if user.ID != 0 {
	//	err = errors.New("The account is existed!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	////檢查此信箱是否存在
	//user, err = repository.GetByEmail(req.Email)
	//if user.ID != 0 {
	//	err = errors.New("The e-mail is existed!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//password := DomainLogic.HashPassword(req.Password)
	////model.User型態
	//var data = new(model.User)
	//data.LastName = req.LastName
	//data.FirstName = req.FirstName
	//data.NickName = req.NickName
	//data.Account = req.Account
	//data.Password = password
	//data.Email = req.Email
	//data.Gender = req.Gender
	//data.CreatedAt = time.Now()
	//data.UpdatedAt = time.Now()
	//
	////Insert into Users
	//err = repository.Create(data)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//ctx.JSON(http.StatusOK, "Register successfully!")
}

func (s userService) ForgetPassword(req *forget_password.Request) (err error) {
	//email := ctx.PostForm("email")
	//檢查Email是否存在
	user, err := s.repo.UserRepository.GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The email isn't existed!")
		}
		return
	}

	var token string
	token, err = jwt.GenerateToken()
	if err != nil {
		err = errors.New("Token generate failed!")
		return
	}

	newData := map[string]interface{}{
		"token": token,
	}
	//存token至user資料中
	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		err = errors.New("Update failed!")
		return
	}

	emailReplace := strings.Replace(user.Email, "@", "%40", -1)
	url := "http://127.0.0.1:8080/reset-forget-password?code=" + token + "&" + "email=" + emailReplace
	subject := "ShanniTest- 請重置密碼!"
	body := "請點擊此連結: " + url + " ,並重新設置密碼。"

	//寄信
	err = mail.SendMail(subject, user.Email, body)
	if err != nil {
		err = errors.New("Send mail failed!")
		return
	}

	return

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "Send mail Successfully!",
	//})
	//req := new(forget_password.Request)
	//err := ctx.BindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, err)
	//	return
	//}
	////檢查帳號
	//user := new(model.User)
	//user, err = repository.GetByAccount(req.Account)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		err = errors.New("The account isn't existed!")
	//	}
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	////檢查Email是否存在
	//if user.Email != req.Email {
	//	err = errors.New("The email isn't existed!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
}

func (s userService) ResetForgetPassword(req *reset_forget_password.Request) (err error) {
	//code := ctx.PostForm("code")
	//email := ctx.PostForm("email")
	user, err := s.repo.UserRepository.GetByToken(req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The code isn't existed!")
		}
		return
	}

	if user.Email != req.Email {
		err = errors.New("The E-mail is not yours!")
		return
	}

	//password := ctx.PostForm("password")
	hashPassword := rule.HashPassword(req.Password)

	newData := map[string]interface{}{
		"password": hashPassword,
	}
	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		err = errors.New("Update failed!")
		return
	}

	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "You reset password successfully. Please back to login page!",
	//})
	return
}

func (s userService) Login(ctx *gin.Context, req *login_by_account.Request) (err error) {
	//account := ctx.PostForm("account")
	//password := ctx.PostForm("password")
	//var userLog = new(model.UserLoginLog)
	//檢查帳號
	user, err := s.repo.UserRepository.GetByAccount(req.Account)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The account isn't existed!")
			return
		}
		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})
		err = errors.New("Get Failed!")
		return
	}

	//檢查密碼
	if !rule.CheckPasswordHash(req.Password, user.Password) {
		err = errors.New("The password isn't correct!")

		//紀錄Log
		//userLog.UserID = user.ID
		//userLog.Account = user.Account
		//userLog.Result = "0"
		//userLog.Remark = err.Error()
		//userLog.CreatedAt = time.Now()
		//userLog.UpdatedAt = time.Now()
		//
		//err = USER_LOGIN_LOG.Create(userLog)
		//if err != nil {
		//	log.Fatal(err)
		//	return
		//}

		//ctx.JSON(http.StatusBadRequest, gin.H{
		//	"status":  http.StatusBadRequest,
		//	"message": err.Error(),
		//})
		return
	}

	//req := new(login_by_account.Request)
	//err := ctx.BindJSON(req)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, err)
	//	return
	//}
	//user := new(model.User)
	////檢查帳號
	//user, err = repository.GetByAccount(req.Account)
	//if err != nil {
	//	if errors.Is(err, gorm.ErrRecordNotFound) {
	//		err = errors.New("The account isn't existed!")
	//	}
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}
	//
	////檢查密碼
	//if !ApplicationLogic.CheckPasswordHash(req.Password, user.Password) {
	//	err = errors.New("The password isn't correct!")
	//	ctx.JSON(http.StatusBadRequest, err.Error())
	//	return
	//}

	middleware.ClearSession(ctx)
	middleware.SaveSession(ctx, user.ID)

	token, err := jwt.GenerateToken()
	if err != nil {
		err = errors.New("Token generate failed!")
		return
	}

	newData := map[string]interface{}{
		"token": token,
	}

	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		err = errors.New("Update failed!")
		return
	}
	//紀錄Log
	//userLog.UserID = user.ID
	//userLog.Account = user.Account
	//userLog.Result = "1"
	//userLog.CreatedAt = time.Now()
	//userLog.UpdatedAt = time.Now()
	//err = USER_LOGIN_LOG.Create(userLog)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	return
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "You login successfully!",
	//	"data":    user,
	//})
}

func (s userService) BuildTempToken(ctx *gin.Context) (url string, err error) {
	userID := middleware.GetSession(ctx)

	user, err := s.repo.UserRepository.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The user isn't existed!")
			return
		}
		err = errors.New("Get failed!")
		return
	}

	tempToken, err := jwt.GenerateToken()
	if err != nil {
		err = errors.New("Token generate failed!")
		return
	}

	err = s.repo.UserRepository.SetRedis("tempToken", tempToken, 0)
	if err != nil {
		err = errors.New("Set redis failed!")
		return
	}

	url = "http://127.0.0.1:8080/users-auth/token-change?code=" + tempToken

	newData := map[string]interface{}{
		"temp_token": tempToken,
	}
	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		err = errors.New("Update failed!")
		return
	}
	return
}

func (s userService) TokenChange(req *token_change.Request) (token string, err error) {
	//code := ctx.PostForm("code")
	//redis 相關要移至userRepository
	value, err := s.repo.UserRepository.GetRedis("tempToken")

	if err != nil {
		if err == redis.Nil {
			err = errors.New("The url is expired!")
			return
		}
		err = errors.New("Get redis failed!")
		return
	}

	fmt.Println("value: ", value)
	fmt.Println("req.Code: ", req.Code)
	if value != req.Code {
		err = errors.New("The url is error!")
		return
	}

	user, err := s.repo.UserRepository.GetByTempToken(req.Code)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The code isn't existed!")
			return
		}
		err = errors.New("Get failed!")
		return
	}

	newData := map[string]interface{}{
		"temp_token": "",
	}

	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		err = errors.New("Update failed!")
		return
	}
	token = user.Token
	return
}

func (s userService) GetUserInfo(req *get_user_info.Request) (res get_user_info.Response, err error) {
	//token := ctx.PostForm("token")
	user, err := s.repo.UserRepository.GetByToken(req.Token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The token isn't existed!")
			return
		}
		err = errors.New("Get failed!")
		return
	}
	res = get_user_info.Response{
		ID:        user.ID,
		LastName:  user.LastName,
		FirstName: user.FirstName,
		NickName:  user.NickName,
		Email:     user.Email,
		Gender:    user.Gender,
		Amount:    user.Amount,
		TempToken: user.TempToken,
	}
	return
}

func (s userService) DepositAmount(req *deposit_amount.Request) (res deposit_amount.Response, err error) {
	//token := ctx.PostForm("token")
	//depositAmount := ctx.PostForm("deposit_amount")

	user, err := s.repo.UserRepository.GetByToken(req.Token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The token isn't existed!")
			return
		}
		err = errors.New("Get failed!")
		return
	}

	amount, err := strconv.ParseInt(req.DepositAmount, 10, 64)
	if err != nil {
		err = errors.New("Please enter number!")
		return
	}

	amount = amount + user.Amount
	newData := map[string]interface{}{
		"amount": amount,
	}

	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		err = errors.New("Update failed!")
		return
	}

	totalAmount := strconv.FormatInt(amount, 10)
	res.TempToken = user.TempToken
	res.TotalAmount = totalAmount
	return
}

func (s userService) WithdrawAmount(req *withdraw_amount.Request) (res withdraw_amount.Response, err error) {
	//token := ctx.PostForm("token")
	//withdrawAmount := ctx.PostForm("withdraw_amount")

	user, err := s.repo.UserRepository.GetByToken(req.Token)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("The token isn't existed!")
			return
		}
		err = errors.New("Get failed!")
		return
	}

	amount, err := strconv.ParseInt(req.WithdrawAmount, 10, 64)
	if err != nil {
		err = errors.New("Please enter number!")
		return
	}

	amount = user.Amount - amount
	if amount < 0 {
		err = errors.New("Balance Amount isn't enough!")
		return
	}
	newData := map[string]interface{}{
		"amount": amount,
	}

	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		errors.New("Update failed!")
		return
	}

	totalAmount := strconv.FormatInt(amount, 10)
	res.TempToken = user.TempToken
	res.TotalAmount = totalAmount
	return
}

func (s userService) Logout(ctx *gin.Context) (err error) {
	userID := middleware.GetSession(ctx)
	user, err := s.repo.UserRepository.GetByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			errors.New("The user isn't existed!")
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	newData := map[string]interface{}{
		"token": "",
	}

	err = s.repo.UserRepository.Update(user, newData)
	if err != nil {
		errors.New("Update failed!")
		return
	}
	middleware.ClearSession(ctx)
	return
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status":  http.StatusOK,
	//	"message": "User Sign out successfully",
	//})
}
