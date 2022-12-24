package engine

import (
	"fmt"
	"log"
	"newsSearchEngine/engine/tokenizer"
	"os"
	"runtime"
)

var (
	Dir       string               //文件夹
	OpEngine  *Engine              //引擎
	Debug     bool                 //调试
	Tokenizer *tokenizer.Tokenizer //分词器
	Shard     int                  //分片
	Timeout   int64                //超时关闭数据库
	BufferNum int                  //分片缓冲数
)

func Init() error {
	//读取当前路径下的所有目录，就是数据库名称
	dirs, err := os.ReadDir(Dir)
	if err != nil {
		if os.IsNotExist(err) {
			//创建
			err := os.MkdirAll(Dir, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	if len(dirs) != 0 {
		if dirs[0].IsDir() {
			OpEngine = GetDataBase(dirs[0].Name())
			log.Println("db:", dirs[0].Name())
		}
	}
	return nil
}

func NewEngine(name string) *Engine {
	var engine = &Engine{
		IndexPath:    fmt.Sprintf("%s%c%s", Dir, os.PathSeparator, name),
		DatabaseName: name,
		Tokenizer:    Tokenizer,
		Shard:        Shard,
		Timeout:      Timeout,
		BufferNum:    BufferNum,
	}
	option := engine.GetOptions()

	engine.InitOption(option)
	engine.IsDebug = Debug

	return engine
}

func GetDataBase(name string) *Engine {
	if name == "" {
		name = "default"
	}
	if OpEngine != nil && name == OpEngine.DatabaseName {
		return OpEngine
	}
	engine := NewEngine(name)
	return engine
}

func GetIndexCount() int64 {
	return OpEngine.GetIndexCount()
}

func GetDocumentCount() int64 {
	return OpEngine.GetDocumentCount()
}

func DropDataBase() error {
	OpEngine = nil
	//释放资源
	runtime.GC()

	return nil
}
