/**
 * @author:       wangxuebing
 * @fileName:     encodeutil.go
 * @date:         2023/1/14 16:56
 * @description:
 */

package edcode

import "net/url"

// EncodeURL URL编码
func EncodeURL(api string, params map[string]string) (string, error) {
	reqUrl, err := url.Parse(api)
	if err != nil {
		return "", err
	}
	query := reqUrl.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	reqUrl.RawQuery = query.Encode()
	result, _ := url.QueryUnescape(reqUrl.String())
	return result, nil
}
