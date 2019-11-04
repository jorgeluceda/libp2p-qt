package main

import (
	"github.com/therecipe/qt/core"
)

const (
	NodeName = int(core.Qt__UserRole) + 1<<iota
	NodeMessage
)

// func init() { CustomListModel_QmlRegisterType2("CustomQmlTypes", 1, 0, "CustomListModel") }
type NodeModel struct {
	core.QAbstractListModel

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*Node                  `property:"nodes"`

	_ func(*Node)                                 `slot:"addNode"`
	_ func(row int, nodeName, NodeMessage string) `slot:"editNode"`
	_ func(row int)                               `slot:"removeNode"`
}

type Node struct {
	core.QObject
	_ string `property:"nodeName"`
	_ string `property:"nodeMessage"`
}

func init() {
	Node_QRegisterMetaType()
}

func (m *NodeModel) init() {
	m.SetRoles(map[int]*core.QByteArray{
		NodeName:    core.NewQByteArray2("nodeName", -1),
		NodeMessage: core.NewQByteArray2("nodeMessage", -1),
	})

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
}

func (m *NodeModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Nodes()) {
		return core.NewQVariant()
	}

	var n = m.Nodes()[index.Row()]

	switch role {
	case NodeName:
		{
			return core.NewQVariant1(n.NodeName())
		}

	case NodeMessage:
		{
			return core.NewQVariant1(n.NodeMessage())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (m *NodeModel) rowCount(*core.QModelIndex) int {
	return len(m.Nodes())
}

func (m *NodeModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *NodeModel) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *NodeModel) addNode(p *Node) {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Nodes()), len(m.Nodes()))
	m.SetNodes(append(m.Nodes(), p))
	m.EndInsertRows()

	var nIndex = m.Index(len(m.Nodes()), 0, core.NewQModelIndex())
	m.DataChanged(nIndex, nIndex, []int{NodeName, NodeMessage})

}

func (m *NodeModel) editNode(row int, nodeName string, nodeMessage string) {
	var n = m.Nodes()[row]

	if nodeName != "" {
		n.SetNodeName(nodeName)
	}

	if nodeMessage != "" {
		n.SetNodeMessage(nodeMessage)
	}

	var nIndex = m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(nIndex, nIndex, []int{NodeName, NodeMessage})
}

func (m *NodeModel) removeNode(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetNodes(append(m.Nodes()[:row], m.Nodes()[row+1:]...))
	m.EndRemoveRows()
}
