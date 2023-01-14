/**
 * @author:       wangxuebing
 * @fileName:     commonmodel.go
 * @date:         2023/1/14 16:37
 * @description:
 */

package core

type CommonId struct {
	ID string `json:"id"`
}

type CommonKv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
