package excel

/*
excel的扩展采用的是https://github.com/360EntSecGroup-Skylar/excelize
*/

type Read struct{}

func NewRead() *Read {
	return &Read{}
}
