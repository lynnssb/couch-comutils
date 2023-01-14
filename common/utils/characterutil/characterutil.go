/**
 * @author:       wangxuebing
 * @fileName:     characterutil.go
 * @date:         2023/1/14 16:39
 * @description:
 */

package characterutil

import "bytes"

// StitchingBufStr 字符串拼接
func StitchingBufStr(args ...string) (str string) {
	var buf bytes.Buffer
	for _, v := range args {
		buf.WriteString(v)
	}
	str = buf.String()
	return
}
