package base

var Users map[string]map[string]float64 = map[string]map[string]float64{}

/**
 * 判断用户是否存在
 * @param name 用户名称
 * @return exists true存在 false不存在
 */
func UserExists(name string) bool {
	_,exists := Users[name]
	return exists
}

/**
 * 添加用户
 * @param name 用户名称
 * @param position 用户位置信息
 */
func AddUser(name string,position map[string]float64)  {
	Users[name] = position
}

/**
 * 修改用户
 * @param name 用户名称
 * @param position 用户位置信息
 */
func SetUser(name string,position map[string]float64)  {
	Users[name] = position
}

/**
 * 查询所有用户
 * @return Users 所有用户列表
 */
func GetAllUser() map[string]map[string]float64 {
	return Users
}