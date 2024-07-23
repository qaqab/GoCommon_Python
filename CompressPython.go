package GoCommon_Python

import "strings"

func Python_in(target string, str_array []string) bool {
	// 遍历字符串数组
	for _, element := range str_array {
		// 如果目标字符串与当前元素相等
		if target == element {
			// 返回 true
			return true
		}
	}
	// 遍历结束后，仍未找到目标字符串，返回 false
	return false
}
func Python_split(target string, sep string, num int) []string {
	// 使用指定的分隔符将目标字符串分割成字符串数组
	data_list := strings.Split(target, sep)

	// 如果分割后的字符串数组长度大于指定的数量
	if len(data_list) > num {
		// 取前num个字符串放入success_list中
		success_list := data_list[:num]

		// 将剩余的字符串用指定的分隔符连接起来，并添加到success_list的末尾
		success_list = append(success_list, (strings.Join(data_list[num:], sep)))

		// 返回success_list
		return success_list
	}

	// 如果分割后的字符串数组长度小于等于指定的数量，则直接返回data_list
	return data_list
}
