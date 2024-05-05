package common

import (
	"gin-one/model/common"
	"gin-one/utils/helper"
	"github.com/jinzhu/copier"
)

// AssembleTree
//
//	@Description: 将 list 组装成树
//	@param list
//	@return res
//	@return err
func AssembleTree[T common.TreeModelNode](list []T, rootIds []string) (res []T, err error) {
	rawList := make([]T, 0, len(list))
	err = copier.Copy(&rawList, &list)

	if err != nil {
		return res, err
	}
	for i := 0; i < len(rawList); i++ {
		for j := 0; j < len(rawList); j++ {
			p := rawList[j]
			c := rawList[i]
			if p.IsParentOf(c) {
				p.Append(c)
				break
			}
		}
	}

	for i := 0; i < len(rawList); i++ {
		node := rawList[i]
		if helper.HasEle(rootIds, node.GetNodeID()) {
			res = append(res, node)
		}
	}

	return res, nil
}
