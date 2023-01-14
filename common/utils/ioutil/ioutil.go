/**
 * @author:       wangxuebing
 * @fileName:     ioutil.go
 * @date:         2023/1/14 16:54
 * @description:
 */

package ioutil

import "os"

// PathExist 判断文件夹是否存在,不存在则创建
func PathExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err == nil {
			return nil
		} else {
			return err
		}
	}
	return err
}
