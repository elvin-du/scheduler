package task

/*
TODO:
现在只支持GET方法.
Target对象应该都是自己写的服务，和Scheduler放在同一个内网来保证api的安全性．
以后有可能需要更新本服务来支持对第三方有验证要求的API调用
*/
type Target struct {
	Addr       string
	Path       string
	Params map[string]string
}
