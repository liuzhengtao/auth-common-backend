package consts

type MenuTypeEnum int

// 定义枚举值及其对应的字符串标签
const (
	NULL MenuTypeEnum = iota
	MENU
	CATALOG
	EXTLINK
	BUTTON
)

// 定义一个结构体来存储枚举值和标签的映射
type menuTypeEnumValue struct {
	value int
	label string
	key   string
}

// 定义一个全局变量来存储所有的枚举值
var menuTypeEnumMap = map[MenuTypeEnum]menuTypeEnumValue{
	NULL:    {value: 0, label: ""},
	MENU:    {value: 1, key: "MENU", label: "菜单"},
	CATALOG: {value: 2, key: "CATALOG", label: "目录"},
	EXTLINK: {value: 3, key: "EXTLINK", label: "外链"},
	BUTTON:  {value: 4, key: "BUTTON", label: "按钮"},
}

// Value 实现 Value 方法
func (e MenuTypeEnum) Value() int {
	return menuTypeEnumMap[e].value
}

// Label 实现 Label 方法
func (e MenuTypeEnum) Label() string {
	return menuTypeEnumMap[e].label
}
func (e MenuTypeEnum) Key() string {
	return menuTypeEnumMap[e].key
}
