/**
 * @Author: Resynz
 * @Date: 2021/12/28 14:12
 */
package util

func CheckIsInStrArray(arr []string, obj string) bool {
	for _, v := range arr {
		if v == obj {
			return true
		}
	}
	return false
}
