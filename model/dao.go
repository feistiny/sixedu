package model

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/feistiny/sixedu/etc"
	"io"
	"os"
	"reflect"
	"strings"
	"sync"
)

// 统一的模型接口
type Model interface {
	ToString() string        // 格式化输出数据信息
	ToStringSlice() []string // 格式化输出数据信息
	Save() bool
}

type nameModels map[string]func()Model
type PrimaryModels map[string]Model

var (
	path     string
	suffix   = ".sql"
	models   nameModels // 记录标识 user =》 结构体
	sqlLocks map[string]*sync.Mutex
)

// "User" => User{} 不支持
func init() {
	// 标识绑定注册
	models = make(nameModels)
	models["user"] = func() Model {
		return &User{}
	}
	sqlLocks = make(map[string]*sync.Mutex)
}

func loadConfig() {
	logs.Debug("loadConfig")
	// etc.Cfg 直接在 var 或者 init 里赋值, 单元测试直接报错
	// 那时候 etc.Cfg 还没有初始化
	path = etc.Cfg.String("SQL_PATH")
	_suffix := etc.Cfg.String("SQL_SUFFIX")
	if _suffix != "" {
		suffix = _suffix
	}
}

// 数据库文件 -> 通过配置设置
// name 数据库名称   user,admin
// primary 查询主键
// models 存放数据
func rfdata(name, primary string) (datas PrimaryModels, err error) {
	logs.Debug("rfdata start")
	// 1. 读取数据库文件 -》读取那个文件？
	f, err := openSql(name, os.O_RDONLY)
	if err != nil {
		return
	}
	// 关闭文件的资源流
	defer f.Close()
	// 创建读取的文件的缓冲区
	buf := bufio.NewReader(f)
	// 2. 遍历数据  每一行的数据 字段根据 , 分割；数据通过 \n 分割
	var titles []string
	logs.Debug("开始读取 sql 文件")
	for {
		var row []byte
		row, err = buf.ReadBytes('\n') // 根据换行读取文件信息 , 返回的是byte[]
		if err != nil {
			if err == io.EOF { // 是否文件读取结束
				err = nil
			} else {
				fmt.Println("读取文件异常", err)
			}
			break
		}
		// 去掉字符串，并分割数据
		fields := strings.Split(strings.TrimSuffix(string(row), "\n"), ",")
		// fmt.Println("读取到的数据", fields)
		if len(titles) == 0 {
			// 字段名称, 需要根据这个动态赋值 model
			titles = fields
			continue
		}

		if _, ok := models[name]; !ok {
			err = errors.New(name + " 对应的 model 没有找到")
			return
		}

		var m Model
		var modelKey string
		m = models[name]()
		// logs.Debug("m address %p", m)
		rm := reflect.ValueOf(m)
		for i, field := range fields {
			fieldName := titles[i]
			if fieldName == primary {
				modelKey = field
			}
			rmm := rm.MethodByName("Set" + strings.Title(fieldName))
			rmm.Call([]reflect.Value{reflect.ValueOf(field)})
		}
		// logs.Debug("%+v", m)
		if datas == nil {
			datas = make(PrimaryModels)
		}
		datas[modelKey] = m
	}
	return
}

func openSql(name string, mode int) (f *os.File, err error) {
	logs.Debug("openSql", name)
	var mu *sync.Mutex
	var ok bool
	if mu, ok = sqlLocks[name]; !ok {
		mu = &sync.Mutex{}
		sqlLocks[name] = mu
	}
	mu.Lock()
	defer mu.Unlock()

	loadConfig()
	path := path + name + suffix
	f, err = os.OpenFile(path, mode, 0666) // 根据路径读取文件信息
	if err != nil {
		fmt.Println("文件读取异常,", err)
		return nil, errors.New(fmt.Sprintf("sql 文件打开失败 %s", path))
	}
	logs.Debug("opened file ", path)
	return
}

// 把模型写入 sql 文件
func rwdata(name string, m Model) (err error) {
	f, err := openSql(name, os.O_WRONLY|os.O_APPEND)
	if err != nil {
		return
	}
	n, err := f.WriteString(m.ToString() + "\n")
	if err != nil {
		logs.Error("数据写入异常 %v %T\n", err, f)
		return errors.New("写入失败")
	}
	logs.Debug("写入字节数 ", n)
	return
}
