package initializer

import (
	"flag"
	"fmt"
	"newsSearchEngine/global"
	"os"

	"gopkg.in/yaml.v2"
)

// Parser 解析器
func Parser() *global.Config {
	var addr = flag.String("addr", "127.0.0.1:9134", "设置监听地址和端口")
	//兼容windows
	dir := fmt.Sprintf(".%sdata", string(os.PathSeparator))

	var dataDir = flag.String("data", dir, "设置数据存储目录")

	var debug = flag.Bool("debug", true, "设置是否开启调试模式")

	var dictionaryPath = flag.String("dictionary", "./db/dictionary.txt", "设置词典路径")

	var timeout = flag.Int64("timeout", 10*60, "数据库超时关闭时间(秒)")

	var auth = flag.String("auth", "root:passwd", "开启认证，例如: admin:123456")

	var bufferNum = flag.Int("bufferNum", 1000, "分片缓冲数量")

	var configPath = flag.String("config", "", "配置文件路径，配置此项其他参数忽略")

	flag.Parse()

	config := &global.Config{}

	if *configPath != "" {
		file, err := os.ReadFile(*configPath)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(file, config)
		if err != nil {
			panic(err)
		}
		return config
	}

	config = &global.Config{
		Addr:       *addr,
		Data:       *dataDir,
		Debug:      *debug,
		Dictionary: *dictionaryPath,
		Timeout:    *timeout,
		Auth:       *auth,
		BufferNum:  *bufferNum,
	}
	return config
}
