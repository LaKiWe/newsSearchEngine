package global

type Config struct {
	Addr      string `yaml:"addr"`      //listen addr
	Data      string `json:"data"`      //path of data
	Shard     int    `yaml:"shard"`     //shards num
	Timeout   int64  `json:"timeout"`   //time out ~ ms
	BufferNum int    `yaml:"bufferNum"` //bufferd shard num
}
