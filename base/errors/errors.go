package errors

type Errors struct {
	errors map[string][]string
}

//创建错误
func New() *Errors  {
	return &Errors{
		errors: make(map[string][]string),
	}
}

//添加错误
func (e *Errors) AddError(key,err string)  {
	if _,ok := e.errors[key];!ok {
		e.errors[key] = make([]string,0,5)
	}
	e.errors[key] = append(e.errors[key],err)
}

//获取错误
func (e *Errors) Errors() map[string][]string  {
	return e.errors
}

//通过key 获取错误
func (e *Errors) ErrorByKey(key string) []string  {
	return e.errors[key]
}

//判断是否还有错误
func (e *Errors) HasErrors() bool  {
	return len(e.errors) != 0   // 不等于 0 有错误
}

