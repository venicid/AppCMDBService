package resourceTree

import (
	"AppCMDBService/logger"
	"AppCMDBService/model"
	"AppCMDBService/view"
	"fmt"
)

func GetGlobalTree() (result []*view.GlobalTreeResponse, err error) {

	// 获取所有的product
	products, err := ListMiniProductRecords()
	if err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("[service.resourceTree.GetGlobalTree], result: %v, count:%d, ", products, len(products))
	logger.Logger.Warn(msg)

	/**
	[
	{id:12, parentId:1},
	{id:13, parentId:1},
	{id:14, parentId:1},
	]
	*/
	// 转换为map
	var productParentMap = make(map[int][]*model.ProductMini)
	for _, item := range products {
		parentId := item.ParentId
		productParentMap[*parentId] = append(productParentMap[*parentId], item)
	}
	/**
	{
	1: [12,13,14],
	12: [121, 122, 123],
	121: [1211, 1212],
	}
	*/

	// 递归查询
	return GetNodeChildren(0, "product", productParentMap), nil

}

// GetNodeChildren 递归查询子节点列表
func GetNodeChildren(parentId int, resourceType string, productParentMap map[int][]*model.ProductMini) []*view.GlobalTreeResponse {

	var nodeChildren []*view.GlobalTreeResponse

	productChildren, ok := productParentMap[parentId]
	if !ok {
		return nodeChildren
	}
	for _, item := range productChildren {

		nodeData := make(map[string]int)
		nodeData["id"] = *item.Id

		row := &view.GlobalTreeResponse{}
		row.Id = item.Id
		row.ParentId = item.ParentId
		row.Text = item.ProductName
		row.Type = item.ProductType
		row.Data = nodeData

		deptChildren := GetNodeChildren(*item.Id, "product", productParentMap)
		row.Children = deptChildren

		nodeChildren = append(nodeChildren, row)
	}

	msg := fmt.Sprintf("[service.resourceTree.tree], result: %v, count:%d, ", nodeChildren, len(nodeChildren))
	logger.Logger.Warn(msg)

	return nodeChildren

}
