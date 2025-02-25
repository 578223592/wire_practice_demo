package conf

type DbConf struct {
	//	xxx
	StartYear int
}

func NewDbConf() *DbConf {
	return &DbConf{}
}
