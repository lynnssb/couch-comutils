/**
 * @author:       wangxuebing
 * @fileName:     characterutil.go
 * @date:         2023/1/14 16:39
 * @description:
 */

package characterutil

import (
	"strings"
)

// StitchingBuilderStr 字符串拼接
func StitchingBuilderStr(args ...string) string {
	var build strings.Builder
	for _, v := range args {
		build.WriteString(v)
	}
	return build.String()
}
