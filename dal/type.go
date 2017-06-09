package dal

type OneTimeJob struct {
//	Id          string `json:"id" db:"id"`
//	Name        string `json:"name" db:"name"`
//	Description string `json:"description" db:"description"`
	Retry       int    `json:"retry" db:"retry"` //失败后重试次数，默认不重试
//	Method      string `json:"method" db:"method"`
	Addr        string `json:"addr" db:"addr"`
	Path        string `json:"path" db:"path"`
	Params      string `json:"params" db:"params"`
	StartAt     int64 `json:"start_at" db:"start_at"`
//	CreatedAt   int64 `json:"created_at" db:"created_at"`
//	retErr      string `json:"ret_err" db:"ret_err"`
}

type EveryJob struct {
	Id          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Retry       int    `json:"retry" db:"retry"` //失败后重试次数，默认不重试
	Method      string `json:"method" db:"method"`
	Addr        string `json:"addr" db:"addr"`
	Path        string `json:"path" db:"path"`
	Params      string `json:"params" db:"params"`
	spec        string `json:"spec" db:"spec"`
	CreatedAt   int64 `json:"created_at" db:"created_at"`
	retErr      string `json:"ret_err" db:"ret_err"`
}
