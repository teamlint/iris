package user

import (
	"errors"
	"strconv"
	"strings"

	"github.com/teamlint/iris"
	"github.com/teamlint/iris/mvc"
	"github.com/teamlint/iris/sessions"
)

const sessionIDKey = "UserID"

// paths
var (
	PathLogin  = mvc.Response{Path: "/user/login"}
	PathLogout = mvc.Response{Path: "/user/logout"}
)

// AuthController is the user authentication controller, a custom shared controller.
type AuthController struct {
	// context is auto-binded if struct depends on this,
	// in this controller we don't we do everything with mvc-style,
	// and that's neither the 30% of its features.
	Ctx iris.Context

	Source  *DataSource
	Session *sessions.Session

	// the whole controller is request-scoped because we already depend on Session, so
	// this will be new for each new incoming request, BeginRequest sets that based on the session.
	UserID int64
}

// BeginRequest saves login state to the context, the user id.
func (c *AuthController) BeginRequest(ctx iris.Context) {
	c.UserID, _ = c.Session.GetInt64(sessionIDKey)
}

// EndRequest is here just to complete the BaseController
// in order to be tell iris to call the `BeginRequest` before the main method.
func (c *AuthController) EndRequest(ctx iris.Context) {}

type Resp struct {
	Success   bool
	Name      string
	Content   string
	Data      interface{}
	MVCResult mvc.Result `json:"-"`
}

func (r Resp) Dispatch(ctx iris.Context) {
	if !r.Success {
		ctx.StatusCode(iris.StatusBadRequest)
	}
	if ctx.IsAjax() {
		ctx.JSON(r)
	} else {
		ctx.Values().Set("Title", "通用错误处理")
		ctx.Values().Set("Message", r.Content)
		if r.MVCResult != nil {
			r.MVCResult.Dispatch(ctx)
		}
		// ctx.EndRequest()
		// ctx.View("shared/error.html")
		// return mvc.View{
		// 	Code: iris.StatusBadRequest,
		// 	Name: "shared/error.html",
		// 	Data: iris.Map{"Title": "User Error", "Message": strings.ToUpper(err.Error())},
		// }
	}
}
func (c *AuthController) fireError(err error) mvc.Result {
	// if c.Ctx.IsAjax() {
	if err != nil {
		return &Resp{Name: "请求错误", Content: err.Error(), Data: err}
	}
	return Resp{Success: true, Name: "请求成功", Content: err.Error(), Data: err}

	// }
	// return mvc.View{
	// 	Code: iris.StatusBadRequest,
	// 	Name: "shared/error.html",
	// 	Data: iris.Map{"Title": "User Error", "Message": strings.ToUpper(err.Error())},
	// }
}

func (c *AuthController) redirectTo(id int64) mvc.Response {
	return mvc.Response{Path: "/user/" + strconv.Itoa(int(id))}
}

func (c *AuthController) createOrUpdate(firstname, username, password string) (user Model, err error) {
	username = strings.Trim(username, " ")
	if username == "" || password == "" || firstname == "" {
		return user, errors.New("empty firstname, username or/and password")
	}

	userToInsert := Model{
		Firstname: firstname,
		Username:  username,
		password:  password,
	} // password is hashed by the Source.

	newUser, err := c.Source.InsertOrUpdate(userToInsert)
	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (c *AuthController) isLoggedIn() bool {
	// we don't search by session, we have the user id
	// already by the `BeginRequest` middleware.
	return c.UserID > 0
}

func (c *AuthController) verify(username, password string) (user Model, err error) {
	if username == "" || password == "" {
		return user, errors.New("please fill both username and password fields")
	}

	u, found := c.Source.GetByUsername(username)
	if !found {
		// if user found with that username not found at all.
		return user, errors.New("用户不存在 user with that username does not exist")
	}

	if ok, err := ValidatePassword(password, u.HashedPassword); err != nil || !ok {
		// if user found but an error occurred or the password is not valid.
		return user, errors.New("please try to login with valid credentials")
	}

	return u, nil
}

// if logged in then destroy the session
// and redirect to the login page
// otherwise redirect to the registration page.
func (c *AuthController) logout() mvc.Response {
	if c.isLoggedIn() {
		c.Session.Destroy()
	}
	return PathLogin
}
