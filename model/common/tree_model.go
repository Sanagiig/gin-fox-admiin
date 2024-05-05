package common

type TreeModelNode interface {
	Is(id string) bool
	IsRoot() bool
	GetParentNodeID() string
	GetNodeID() string
	IsChildrenOf(model TreeModelNode) bool
	IsParentOf(model TreeModelNode) bool
	GetChildren() []TreeModelNode
	Append(model TreeModelNode)
	InitChildren()
}

type TreeModel struct {
	UUIDModel
	ParentModel
	Children []TreeModelNode `json:"children" gorm:"-"`
}

func (t *TreeModel) Is(id string) bool {
	return t.ID == id
}

func (t *TreeModel) IsRoot() bool {
	return t.ParentID == ""
}

func (t *TreeModel) IsChildrenOf(model TreeModelNode) bool {
	return t.ParentID == model.GetNodeID()
}

func (t *TreeModel) IsParentOf(model TreeModelNode) bool {
	return t.ID == model.GetParentNodeID()
}

func (t *TreeModel) HasChild(model TreeModelNode) bool {
	for _, child := range t.Children {
		if child.IsParentOf(model) {
			return true
		}
	}
	return false
}

func (t *TreeModel) GetParentNodeID() string {
	return t.ParentID
}

func (t *TreeModel) GetNodeID() string {
	return t.ID
}

func (t *TreeModel) GetChildren() []TreeModelNode {
	return t.Children
}

func (t *TreeModel) Append(model TreeModelNode) {
	t.Children = append(t.Children, model)
}

func (t *TreeModel) InitChildren() {
	if t.Children == nil {
		t.Children = make([]TreeModelNode, 0, 0)
	}
}
