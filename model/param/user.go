package param

type User struct {
	Id   []int64 `param:"id,100"`
	Name string  `param:"name,Tomcat"`
}
