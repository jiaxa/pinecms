package tables

type Page struct {
	Catid       int64 `xorm:"pk"`
	Title       string
	Keywords    string
	Description string
	Content     string
	Updatetime  int64
	Position    string `xorm:"-"`
}
