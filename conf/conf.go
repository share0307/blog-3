package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/zxysilent/logs"
)

/**
	应用配置结构体，使用的是 toml 配置
 */
type appconf struct {
	Title   string `toml:"title"`
	Explain string `toml:"explain"`
	Mode    string `toml:"mode"`
	Addr    string `toml:"addr"`
	Srv     string `toml:"srv"`
	Jwtkey  string `toml:"jwtkey"`
	Jwtexp  int    `toml:"jwtexp"`
	Author  struct {
		Name    string `toml:"name"`
		Website string `toml:"website"`
	} `toml:"author"`	// 结构体也是可以使用 tag 的
	Wechat struct {
		Appid  string `toml:"appid"`
		Secret string `toml:"secret"`
	} `toml:"wechat"`
	Database struct {
		Host   string `toml:"host"`
		Port   int    `toml:"port"`
		User   string `toml:"user"`
		Passwd string `toml:"passwd"`
		Dbname string `toml:"dbname"`
		Params string `toml:"params"`
	} `toml:"database"`
	Xorm struct {
		Idle        int  `toml:"idle"`
		Open        int  `toml:"open"`
		Show        bool `toml:"show"`
		Sync        bool `toml:"sync"`
		CacheEnable bool `toml:"cache_enable"`
		CacheCount  int  `toml:"cache_count"`
		ConnMaxLifetime  int  `toml:"conn_max_lifetime"`
	} `toml:"xorm"`
}

/*
	是否生成环境
 */
func (app *appconf) IsProd() bool {
	return app.Mode == "prod"
}
/**
	是否测试环境
 */
func (app *appconf) IsDev() bool {
	return app.Mode == "dev"
}

const _dsn = "%s:%s@tcp(%s:%d)/%s?%s"

func (app *appconf) Dsn() string {
	return fmt.Sprintf(_dsn, app.Database.User, app.Database.Passwd, app.Database.Host, app.Database.Port, app.Database.Dbname, app.Database.Params)
}

var (
	// 全局变量
	App       *appconf
	// 应用配置目录
	defConfig = "./conf/conf.toml"
)

/**
	初始化，此Init是需要手工调用得，非 init 方法
 */
func Init() {
	var err error
	App, err = initConf()
	if err != nil {
		logs.Fatal("config init error : ", err.Error())
	}
	logs.Debug("conf init")
}

func initConf() (*appconf, error) {
	// app := &appconf{}
	app := new(appconf)	//与上面效果一样
	_, err := toml.DecodeFile(defConfig, &app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
