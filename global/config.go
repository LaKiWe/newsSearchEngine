package global

type Config struct {
	Addr       string `yaml:"addr"`       //listen addr
	Data       string `json:"data"`       //path of data
	Auth       string `json:"auth"`       //`{name}:{passwd}`
	Debug      bool   `yaml:"debug"`      //debug mode
	Dictionary string `json:"dictionary"` //path of dictionary
	Shard      int    `yaml:"shard"`      //shards num
	Timeout    int64  `json:"timeout"`    //time out ~ ms
	BufferNum  int    `yaml:"bufferNum"`  //bufferd shard num
}
