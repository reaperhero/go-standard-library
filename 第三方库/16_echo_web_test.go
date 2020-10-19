package third

// # 搭建框架
// 一、main
//func initLogger() {
//	logFile, err := os.OpenFile("main.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err == nil {
//		logrus.SetOutput(logFile)
//	} else {
//		logrus.Fatalln("Failed to write log to file", err)
//	}
//	logrus.SetFormatter(&logrus.JSONFormatter{})
//	return
//}
//
//func GetDb() *sqlx.DB {
//	databaseHost := utils.GetEnvWithDefault("DBHOST", "192.168.40.136")
//	databaseName := utils.GetEnvWithDefault("DBNAME", "operationsplatform")
//	databaseUser := utils.GetEnvWithDefault("DBUSER", "root")
//	databasePort := utils.GetEnvWithDefault("DBPORT", "3306")
//	databasePass := utils.GetEnvWithDefault("DBPASS", "123456")
//
//	var err error
//	connStr := databaseUser + ":" + databasePass + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
//	DB, err := sqlx.Open("mysql", connStr)
//	logrus.Info(connStr, err)
//	DB.SetConnMaxLifetime(time.Minute * 10)
//	DB.SetMaxIdleConns(5)
//	DB.SetMaxOpenConns(20)
//	if err == nil {
//		err = DB.Ping()
//	}
//
//	if err != nil {
//		logrus.Fatalf("database connect error: %s", err)
//		return nil
//	}
//	return DB
//}
//
//func httpRun(ucase usecase.Usecase) {
//
//	e := echo.New()
//
//	e.Use(middleware.Recover())
//
//	dataPath := os.Getenv("DATA_PATH")
//	accessLogFile, err := os.OpenFile(dataPath+"access.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
//	if err != nil {
//		log.Fatalln("Failed create access log file")
//	}
//
//	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
//		Output: accessLogFile,
//		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
//			`"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
//			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
//			`"bytes_out":${bytes_out},"user_agent":"${user_agent}"}` + "\n",
//	}))
//
//	http.NewHTTPHandler(e, ucase)
//	err = e.Start(":8000")
//}
//
//func main() {
//	initLogger()
//	repo := respository.NewRepository(GetDb())
//	usecase := usecase.NewUsecase(repo)
//	httpRun(usecase)
//}




// 二、http
//type Hander struct {
//	usecase usecase.Usecase
//}
//
//func NewHTTPHandler(e *echo.Echo, u usecase.Usecase) {
//	hander := &Hander{
//		usecase: u,
//	}
//	e.GET("/health", func(c echo.Context) error {
//		return c.String(200, time.Now().Format(time.RFC3339Nano))
//	})
//
//	usersetRouter(e.Group("/user"), hander)
//
//}
//
//func usersetRouter(e *echo.Group, hander *Hander) {
//	e.POST("/api/create_user", hander.RegisterUser)
//}
//
//
//func (h *Hander) RegisterUser(c echo.Context) error {
//	request := struct {
//		Username string `json:"username"`
//		Passwrod string `json:"password"`
//	}{}
//	if err := c.Bind(&request); err != nil {
//		return c.JSON(200, model.GetErrorMap(model.ErrDbInvalid))
//	}
//	if request.Username == "" || request.Passwrod == "" {
//		return c.JSON(200, model.GetErrorMap(model.ErrInvalidParam))
//	}
//	err := h.usecase.CreateUser(request.Username, request.Passwrod)
//	return c.JSON(200, model.GetErrorMap(err))
//}



//三、usecase
//type UsecaseImpl struct {
//	repo respository.DbRepository
//}
//
//func NewUsecase(repository respository.DbRepository) Usecase {
//	return &UsecaseImpl{repo: repository}
//}
//
//func (u *UsecaseImpl) CreateUser(username, password string) error {
//	password = utils.EncodePassword(password)
//	if !u.repo.DbExistUser(username) && password != "" {
//		return u.repo.DbCreateUser(username, password)
//	}
//	return model.ErrUserExist
//}
//
//type Usecase interface {
//	CreateUser(username,password string) error
//}





//四、repository
//type RepositoryImpl struct {
//	db *sqlx.DB
//}
//
//func NewRepository(conn *sqlx.DB) DbRepository {
//	return &RepositoryImpl{
//		conn,
//	}
//}
//
//func (r *RepositoryImpl) DbCreateUser(username, password string) error {
//	sql := "insert into user(name,password,status) value(?,?,?)"
//	result, err := r.db.Exec(sql, username, password, 0)
//	if err != nil {
//		return model.ErrDbInvalid
//	}
//	userId, _ := result.LastInsertId()
//	logrus.WithField("user id", userId)
//	return nil
//}
//
//func (r *RepositoryImpl) DbExistUser(username string) bool {
//	sql := "select * from user where username=?"
//	var user model.User
//	err := r.db.Get(&user, sql, username)
//	if err != nil {
//		return false
//	}
//	return true
//}
//type DbRepository interface {
//	DbCreateUser(username, password string) error
//	DbExistUser(username string) bool
//}







// 二、echo 常用方法

//Create a Cookie
//func writeCookie(c echo.Context) error {
//	cookie := new(http.Cookie)
//	cookie.Name = "username"
//	cookie.Value = "jon"
//	cookie.Expires = time.Now().Add(24 * time.Hour)
//	c.SetCookie(cookie)
//	return c.String(http.StatusOK, "write a cookie")
//}

//Read a Cookie
//func readCookie(c echo.Context) error {
//	cookie, err := c.Cookie("username")
//	if err != nil {
//		return err
//	}
//	fmt.Println(cookie.Name)
//	fmt.Println(cookie.Value)
//	return c.String(http.StatusOK, "read a cookie")
//}


// 自定义请求绑定
//type User struct {
//  Name  string `json:"name" form:"name" query:"name"`
//  Email string `json:"email" form:"email" query:"email"`
//}
//func(c echo.Context) (err error) {
//  u := new(User)
//  if err = c.Bind(u); err != nil {
//    return
//  }
//  return c.JSON(http.StatusOK, u)
//}
//curl -X POST http://localhost:1323/users  -H 'Content-Type: application/json'  -d '{"name":"Joe","email":"joe@labstack"}'
//curl -X POST http://localhost:1323/users  -d 'name=Joe' -d 'email=joe@labstack.com'
//curl -X GET http://localhost:1323/users\?name\=Joe\&email\=joe@labstack.com


// 原生绑定
// func(c echo.Context) error {
//	name := c.FormValue("name")
//	return c.String(http.StatusOK, name)
//}
//curl -X POST http://localhost:1323 -d 'name=Joe'
//
//func(c echo.Context) error {
//	name := c.QueryParam("name")
//	return c.String(http.StatusOK, name)
//})
//curl -X GET http://localhost:1323\?name\=Joe

// 路径参数
//e.GET("/users/:name", func(c echo.Context) error {
//	name := c.Param("name")
//	return c.String(http.StatusOK, name)
//})



// 参数校验
//type User struct {
//	Name  string `json:"name" validate:"required"`
//	Email string `json:"email" validate:"required,email"`
//}
//
//CustomValidator struct {
//	validator *validator.Validate
//}
//func (cv *CustomValidator) Validate(i interface{}) error {
//	return cv.validator.Struct(i)
//}
//
//func main() {
//	e := echo.New()
//	e.Validator = &CustomValidator{validator: validator.New()}
//	e.POST("/users", func(c echo.Context) (err error) {
//		u := new(User)
//		if err = c.Bind(u); err != nil {
//			return
//		}
//		if err = c.Validate(u); err != nil {
//			return
//		}
//		return c.JSON(http.StatusOK, u)
//	})
//	e.Logger.Fatal(e.Start(":1323"))
//}



// 三、echo响应

//Send JSON
//func(c echo.Context) error {
//  u := &User{
//    Name:  "Jon",
//    Email: "jon@labstack.com",
//  }
//  return c.JSON(http.StatusOK, u)
//}


//Send File
//func(c echo.Context) error {
//  return c.File("<PATH_TO_YOUR_FILE>")
//}


//Send Stream
//func(c echo.Context) error {
//  f, err := os.Open("<PATH_TO_IMAGE>")
//  if err != nil {
//    return err
//  }
//  return c.Stream(http.StatusOK, "image/png", f)
//}



// JWT
//https://echo.labstack.com/cookbook/jwt



//websocket
//https://echo.labstack.com/cookbook/websocket