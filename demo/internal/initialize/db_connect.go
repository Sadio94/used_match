package initialize

//
//import (
//	"context"
//	"crypto/tls"
//	"errors"
//	"fmt"
//	"github.com/go-sql-driver/mysql"
//	_ "github.com/go-sql-driver/mysql"
//	_ "github.com/mattn/go-sqlite3"
//	"github.com/sirupsen/logrus"
//	"github.com/flora/greed"
//	"github.com/flora/greed/config"
//	"github.com/flora/greed/exec"
//	"github.com/flora/greed/file"
//	"github.com/flora/greed/orm"
//	"github.com/flora/greed/protect"
//	"os"
//	"path"
//	"strings"
//	"time"
//	"xorm.io/xorm"
//	log2 "xorm.io/xorm/log"
//)
//
//type DBOption func(*DBOptions)
//
//type DBOptions struct {
//	DBName  string
//	User    string
//	Host    string
//	Port    int
//	Pwd     string
//	SDBAPwd string
//	SSL     bool
//	//ping超时时间(ms)
//	TimeOut     int
//	Loc         string
//	CharSet     string
//	MaxOpenConn int
//	MaxIdleConn int
//	MaxLifeTime time.Duration
//	MaxIdleTime time.Duration
//	//重连最小延迟(根据网络短暂故障可能发生的最小间隔合理设置)(ms)
//	ReConnectMinDelay int
//	//重连最大延迟（根据网络短暂故障可能发生的最大间隔合理设置)(s)
//	ReConnectMaxDelay int
//	//重连随机因子
//	ReConnectFactor float64
//	//重连Jitter开启标志,默认开启
//	ReConnectJitterFlag bool
//	//db服务名称
//	DBServiceName string
//	//是否安装DB-install-param
//	IsInstallDB string
//	//unix sock path
//	SockPath string
//
//	DBOrm *orm.Orm
//	//shell cmd
//	Cmd     *exec.Cmd
//	Protect *protect.Protect
//	Logger  *logrus.Logger
//}
//
//type DBClient struct {
//	Opts *DBOptions
//}
//
//func NewDBClient(opts ...DBOption) *DBClient {
//	options := newDBOptions(opts...)
//	return &DBClient{
//		Opts: options,
//	}
//}
//
//var defaultLog = logrus.StandardLogger()
//
//func newDBOptions(opts ...DBOption) *DBOptions {
//	options := &DBOptions{
//		User:                "root",
//		Host:                "127.0.0.1",
//		Port:                3306,
//		SSL:                 false,
//		TimeOut:             5000,
//		Loc:                 "Local",
//		CharSet:             "utf8mb4",
//		ReConnectJitterFlag: true,
//		ReConnectFactor:     1.2,
//		ReConnectMinDelay:   10,
//		ReConnectMaxDelay:   100,
//		DBServiceName:       "DBService",
//		IsInstallDB:         "yes",
//		MaxOpenConn:         10_000,
//		//MaxIdleConn:         1000,
//		MaxLifeTime: 300 * time.Second,
//		MaxIdleTime: 300 * time.Second,
//		DBOrm:       orm.NewOrm(defaultLog),
//		Cmd:         exec.NewCmd(defaultLog),
//		Protect:     protect.NewProtect(defaultLog),
//		Logger:      defaultLog,
//	}
//	for _, o := range opts {
//		o(options)
//	}
//	return options
//}
//
////region settings
//
//func DBUser(userName string) DBOption {
//	return func(o *DBOptions) {
//		o.User = userName
//	}
//}
//
//func DBName(dbName string) DBOption {
//	return func(o *DBOptions) {
//		o.DBName = dbName
//	}
//}
//
//func DBHost(host string) DBOption {
//	return func(o *DBOptions) {
//		o.Host = host
//	}
//}
//func DBLoc(loc string) DBOption {
//	return func(o *DBOptions) {
//		o.Loc = loc
//	}
//}
//func DBCharSet(charset string) DBOption {
//	return func(o *DBOptions) {
//		o.CharSet = charset
//	}
//}
//func DBPwd(pwd string) DBOption {
//	return func(o *DBOptions) {
//		o.Pwd = pwd
//	}
//}
//
//func DBSDBAPwd(sDBAPwd string) DBOption {
//	return func(o *DBOptions) {
//		o.SDBAPwd = sDBAPwd
//	}
//}
//
//func DBSSL(ssl bool) DBOption {
//	return func(o *DBOptions) {
//		o.SSL = ssl
//	}
//}
//
//func DBPort(port int) DBOption {
//	return func(o *DBOptions) {
//		o.Port = port
//	}
//}
//
//func IsInstallDB(isInstallDB string) DBOption {
//	return func(o *DBOptions) {
//		o.IsInstallDB = isInstallDB
//	}
//}
//
//func SockPath(sockPath string) DBOption {
//	return func(o *DBOptions) {
//		o.SockPath = sockPath
//	}
//}
//
//func DBServiceName(dbServiceName string) DBOption {
//	return func(o *DBOptions) {
//		o.DBServiceName = dbServiceName
//	}
//}
//
//func DBTimeOut(timeout int) DBOption {
//	return func(o *DBOptions) {
//		o.TimeOut = timeout
//	}
//}
//
//func DBMaxOpenConn(maxOpenConn int) DBOption {
//	return func(o *DBOptions) {
//		o.MaxOpenConn = maxOpenConn
//	}
//}
//
//func DBMaxIdleConn(maxIdleConn int) DBOption {
//	return func(o *DBOptions) {
//		o.MaxIdleConn = maxIdleConn
//	}
//}
//
////func DBMaxIdleTime(maxIdleTime time.Duration) DBOption {
////	return func(o *DBOptions) {
////		o.MaxIdleTime = maxIdleTime
////	}
////}
//
//func DBMaxLifeTime(maxLifeTime time.Duration) DBOption {
//	return func(o *DBOptions) {
//		o.MaxLifeTime = maxLifeTime
//	}
//}
//
//func DBLog(log ...*logrus.Logger) DBOption {
//	return func(o *DBOptions) {
//		if log != nil {
//			defaultLog = log[0]
//		}
//		o.Logger = defaultLog
//		o.DBOrm = orm.NewOrm(defaultLog)
//		o.Cmd = exec.NewCmd(defaultLog)
//		o.Protect = protect.NewProtect(defaultLog)
//	}
//}
//
////endregion
//
//func getDBCfg(dbConfigPath, appName string, serverIni config.Configer) (*DBConfig, error) {
//	dbCfg := &DBConfig{}
//	if !file.FileExist(dbConfigPath) {
//		return nil, errors.New(fmt.Sprintf("%s not exist", dbConfigPath))
//	}
//
//	db, err := config.NewConfig("ini", dbConfigPath)
//	if err != nil {
//		return nil, err
//	}
//
//	dbCfg.MysqlHost = serverIni.String("Network::dbip")
//	if dbCfg.MysqlHost == "" {
//		dbCfg.MysqlHost = "127.0.0.1"
//	}
//	dbCfg.DBVIP = dbCfg.MysqlHost
//	//兼容ipv6
//	if strings.Count(dbCfg.MysqlHost, ":") >= 2 {
//		dbCfg.MysqlHost = "[" + dbCfg.MysqlHost + "]"
//	}
//	dbCfg.MysqlPort, err = serverIni.Int("DBService::dbport")
//	if err != nil {
//		return nil, err
//	}
//
//	dbCfg.MysqlUser = db.String(appName + "::user")
//	dbCfg.MysqlPwd = db.String(appName + "::pwd")
//	dbCfg.MysqlDBAPwd = db.String("db_config::sdba")
//
//	return dbCfg, nil
//}
//
////获取数据库实例
//func (db *DBClient) GetDBInstance() *xorm.Engine {
//	//db.waitDBInstalled()
//	connectionString := ""
//	//增加单向ssl
//	if db.Opts.SSL {
//		// Verify that registering / using a custom cfg works
//		err := mysql.RegisterTLSConfig("custom-skip-verify", &tls.Config{
//			InsecureSkipVerify: true,
//		})
//		if err != nil {
//			db.Opts.DBOrm.Logger.Errorf("[DB.RegisterTLSConfig]ex:%v", err)
//			os.Exit(1)
//		}
//		connectionString = fmt.Sprintf("%s:%s@(%s:%d)/%s?loc=%s&tls=%s", db.Opts.User, db.Opts.Pwd, db.Opts.Host, db.Opts.Port, db.Opts.DBName, db.Opts.Loc, "custom-skip-verify")
//	} else {
//		connectionString = fmt.Sprintf("%s:%s@(%s:%d)/%s?loc=%s", db.Opts.User, db.Opts.Pwd, db.Opts.Host, db.Opts.Port, db.Opts.DBName, db.Opts.Loc)
//	}
//
//	engine, err := xorm.NewEngine("mysql", connectionString) //test.db
//	if err != nil {
//		db.Opts.DBOrm.Logger.Errorf("[DB.NewEngine]ex:%v", err)
//		os.Exit(1)
//	}
//	db.initDBConn(engine)
//	go db.dbMonitor(engine)
//	return engine
//}
//
//// 获取数据库实例
//// 对于使用非logrus日志库的服务屏蔽最后同步成功的日志
//func (db *DBClient) GetDBInstanceNoLog() *xorm.Engine {
//	connectionString := ""
//	//增加单向ssl
//	if db.Opts.SSL {
//		// Verify that registering / using a custom cfg works
//		err := mysql.RegisterTLSConfig("custom-skip-verify", &tls.Config{
//			InsecureSkipVerify: true,
//		})
//		if err != nil {
//			db.Opts.DBOrm.Logger.Errorf("[DB.RegisterTLSConfig]ex:%v", err)
//			os.Exit(1)
//		}
//		connectionString = fmt.Sprintf("%s:%s@(%s:%d)/%s?loc=%s&tls=%s", db.Opts.User, db.Opts.Pwd, db.Opts.Host, db.Opts.Port, db.Opts.DBName, db.Opts.Loc, "custom-skip-verify")
//	} else {
//		connectionString = fmt.Sprintf("%s:%s@(%s:%d)/%s?loc=%s", db.Opts.User, db.Opts.Pwd, db.Opts.Host, db.Opts.Port, db.Opts.DBName, db.Opts.Loc)
//	}
//
//	engine, err := xorm.NewEngine("mysql", connectionString) //test.db
//	if err != nil {
//		db.Opts.DBOrm.Logger.Errorf("[DB.NewEngine]ex:%v", err)
//		os.Exit(1)
//	}
//	db.initDBConnNoLog(engine)
//	go db.dbMonitor(engine)
//	return engine
//}
//
////sqlite3
//func (db *DBClient) GetSQLite3DBInstance(dbPath string) *xorm.Engine {
//	engine, err := xorm.NewEngine("sqlite3", dbPath)
//	if err != nil {
//		db.Opts.DBOrm.Logger.Errorf("[DB.NewEngine]ex:%v", err)
//		os.Exit(1)
//	}
//	db.initDBConn(engine)
//	//go db.dbMonitor(engine)
//	return engine
//}
//
////数据库迁移
//func (db *DBClient) MigrationDB(beans ...interface{}) error {
//	connectionStr := fmt.Sprintf("sdba:%s@unix(%s)/%s?loc=Local", db.Opts.SDBAPwd, db.Opts.SockPath, db.Opts.DBName)
//	if db.Opts.SSL {
//		// Verify that registering / using a custom cfg works
//		err := mysql.RegisterTLSConfig("custom-skip-verify", &tls.Config{
//			InsecureSkipVerify: true,
//		})
//		if err != nil {
//			db.Opts.DBOrm.Logger.Errorf("[DB.RegisterTLSConfig]ex:%v", err)
//			os.Exit(1)
//		}
//		connectionStr = fmt.Sprintf("sdba:%s@unix(%s)/%s?loc=Local&tls=%s", db.Opts.SDBAPwd, db.Opts.SockPath, db.Opts.DBName, "custom-skip-verify")
//	}
//	engine, err := xorm.NewEngine("mysql", connectionStr)
//	if err != nil {
//		db.Opts.DBOrm.Logger.Errorf("[MigrationDB.NewEngine]error:%v", err)
//		return err
//	}
//	//engine.ShowSQL(false)
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	err = engine.PingContext(ctx)
//	if err != nil {
//		db.Opts.DBOrm.Logger.Errorf("[MigrationDB.initConn]无法访问数据库 %s", db.Opts.DBName) //屏蔽数据敏感信息
//		return errors.New(fmt.Sprintf("[MigrationDB.initConn]无法访问数据库 %s", db.Opts.DBName))
//	}
//	//var database []Database
//	database, err := engine.QueryString("SELECT SCHEMA_NAME AS `Database` FROM INFORMATION_SCHEMA.SCHEMATA")
//	if err != nil {
//		db.Opts.DBOrm.Logger.Errorf("[MigrationDB]get databases error:%v", err)
//		return err
//	}
//	isExist := false
//	for _, m := range database {
//		for _, name := range m {
//			if name == db.Opts.DBName {
//				isExist = true
//			}
//		}
//	}
//	if isExist {
//		err = engine.Sync2(beans)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
////创建数据库
//func CreateDB(mysqlPath, sockPath, dbName, dbaPwd, charset string, ssl bool) error {
//	mysqlExec := path.Join(mysqlPath, "mysql")
//	createDbSql := fmt.Sprintf(" 'create database IF NOT EXISTS %s DEFAULT CHARACTER SET %s COLLATE %s_general_ci' ", dbName, charset, charset)
//	if !ssl {
//		_, err := exec.ExecuteSimpleCmdOut(fmt.Sprintf("%s -usdba -S %s -p'%s' -e %s", mysqlExec, sockPath, dbaPwd, createDbSql))
//		if err != nil {
//			return err
//		}
//	} else {
//		_, err := exec.ExecuteSimpleCmdOut(fmt.Sprintf("%s -usdba -S %s -p'%s' -e %s --ssl=1", mysqlExec, sockPath, dbaPwd, createDbSql))
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//
//}
//
////region private method
//
////等待数据库启动后再执行
//func (db *DBClient) dbIsOrNotInstalled() bool {
//	//判断数据库是否安装
//	dbInstalled := false
//	//查看是否安装db
//	if db.Opts.IsInstallDB == "yes" {
//		dbInstalled = true
//	} else {
//		res, err := db.Opts.Cmd.ExecuteCmdOut(fmt.Sprintf("systemctl list-units| grep %s.service", db.Opts.DBServiceName))
//		if err != nil {
//			db.Opts.DBOrm.Logger.Errorf("[CodeFirstMigration]:execute systemctl db status error:%v", err)
//		}
//		if res != "" {
//			dbInstalled = true
//		}
//	}
//	return dbInstalled
//}
//
//func (db *DBClient) waitDBInstalled() {
//	for {
//		//判断数据库是否安装
//		if db.dbIsOrNotInstalled() {
//			break
//		}
//		db.Opts.DBOrm.Logger.Warn("[waitDBInstalled]数据库未安装，等待安装完毕后启动")
//		//5s后再检测数据库是否已经安装
//		<-time.After(5 * time.Second)
//	}
//}
//
////db connect initialize
//func (db *DBClient) initDBConn(engine *xorm.Engine) {
//	db.dbSet(engine)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db.Opts.TimeOut)*time.Millisecond)
//	defer cancel()
//	err := engine.PingContext(ctx)
//	if err != nil {
//		//屏蔽数据敏感信息
//		db.Opts.DBOrm.Logger.Errorf("[DB.initConn]无法访问数据库 %s", db.Opts.DBName) //屏蔽数据敏感信息
//		os.Exit(1)
//	}
//
//	db.Opts.DBOrm.Logger.Info("[DB.initConn]Connect Success!")
//}
//
////db connect initialize no log
//func (db *DBClient) initDBConnNoLog(engine *xorm.Engine) {
//	db.dbSet(engine)
//	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(db.Opts.TimeOut)*time.Millisecond)
//	defer cancel()
//	err := engine.PingContext(ctx)
//	if err != nil {
//		//屏蔽数据敏感信息
//		db.Opts.DBOrm.Logger.Errorf("[DB.initConn]无法访问数据库 %s", db.Opts.DBName) //屏蔽数据敏感信息
//		os.Exit(1)
//	}
//}
//
////db monitor  request conn and deal center
//func (db *DBClient) dbMonitor(engine *xorm.Engine) {
//	db.Opts.Protect.RecoverPanic("dbMonitor", func() {
//		b := greed.NewBackOff(
//			greed.WithMinDelay(time.Duration(db.Opts.ReConnectMinDelay)*time.Millisecond),
//			greed.WithMaxDelay(time.Duration(db.Opts.ReConnectMaxDelay)*time.Second),
//			greed.WithFactor(db.Opts.ReConnectFactor),
//			greed.WithJitterFlag(db.Opts.ReConnectJitterFlag),
//		)
//		if b.DrivePingOff {
//			db.Opts.DBOrm.ReadPingChan()
//		}
//		for {
//			db.dbHeartBreak(b, engine)
//			time.Sleep(b.Duration())
//			db.Opts.DBOrm.Logger.Infof("[DB.BackOff]%v", b.Duration())
//		}
//	})
//}
//
////db ping
//func (db *DBClient) dbHeartBreak(b *greed.Backoff, engine *xorm.Engine) {
//	ctx := context.Background()
//	done := make(chan struct{}, 1)
//	go func(ctx context.Context) {
//		db.Opts.Protect.RecoverPanic("dbHeartBreak", func() {
//			err := engine.PingContext(ctx)
//			if err == nil {
//				db.dbSet(engine)
//				done <- struct{}{}
//			} else {
//				db.Opts.DBOrm.Logger.Errorf("[DB.HeartBreak]ex:%s", err)
//			}
//		})
//	}(ctx)
//	select {
//	case <-done:
//		db.Opts.DBOrm.Logger.Info("[DB.HeartBreak]connect heart-break is active!")
//		if b.DrivePingOff {
//			db.Opts.DBOrm.ReadPingChan()
//			b.Reset()
//		}
//		return
//	case <-time.After(time.Duration(db.Opts.TimeOut) * time.Millisecond):
//		db.Opts.DBOrm.Logger.Error("[DB.HeartBreak]connect timeout or failed,try connect again!")
//		return
//	}
//}
//
////db set
//func (db *DBClient) dbSet(engine *xorm.Engine) {
//	engine.Charset(db.Opts.CharSet)
//	engine.SetMaxOpenConns(db.Opts.MaxOpenConn)
//	engine.SetMaxIdleConns(db.Opts.MaxIdleConn)
//	engine.SetConnMaxLifetime(db.Opts.MaxLifeTime)
//	engine.ShowSQL(false)
//	logger := log2.NewSimpleLogger(db.Opts.Logger.Out)
//	logger.SetLevel(log2.LOG_ERR)
//	engine.SetLogger(logger)
//	//engine.ShowExecTime(true)
//}
//
////endregion
