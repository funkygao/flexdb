package entity

// PermModels are predefined permission models for each app.
var PermModels = []Model{
	{
		Name:        "用户",
		Kind:        ModelPerm,
		IdentName:   "pin",
		StoreEngine: StoreEngineUSF,
		Ver:         1,
		Slots: []*Column{
			{Name: "pin", Kind: ColumnText, Required: true, Slot: 1, Ordinal: 1},
			{Name: "用户类型", Kind: ColumnChoice, Choices: "普通用户,超级管理员", Required: true, Slot: 2, Ordinal: 2},
			{Name: "所属部门", Kind: ColumnText, Required: true, Slot: 3, Ordinal: 3},
		},
	},

	{
		Name:        "用户组",
		Kind:        ModelPerm,
		IdentName:   "用户组编码",
		StoreEngine: StoreEngineUSF,
		Ver:         1,
		Slots: []*Column{
			{Name: "用户组编码", Kind: ColumnText, Slot: 1, Ordinal: 1},
			{Name: "用户组名称", Kind: ColumnText, Slot: 2, Ordinal: 2},
		},
	},

	{
		Name:        "菜单",
		Remark:      "功能表现为菜单资源，页面中的按钮或页面里其他模块",
		Kind:        ModelPerm,
		StoreEngine: StoreEngineUSF,
		Ver:         1,
		Slots: []*Column{
			{Name: "菜单编码", Kind: ColumnText, Required: true, Slot: 1, Ordinal: 1},
			{Name: "类型", Kind: ColumnChoice, Choices: "菜单,按钮,其他", Required: true, Slot: 2, Ordinal: 2},
		},
	},

	{
		Name:        "菜单角色",
		Remark:      "菜单资源的集合，可以将一类操作功能设置成一个角色，将角色分配给用户",
		Kind:        ModelPerm,
		IdentName:   "角色编码",
		StoreEngine: StoreEngineUSF,
		Ver:         1,
		Slots: []*Column{
			{Name: "角色编码", Kind: ColumnText, Required: true, Slot: 1, Ordinal: 1},
			{Name: "角色名称", Kind: ColumnChoice, Choices: "菜单,按钮,其他", Required: true, Slot: 2, Ordinal: 2},
		},
	},

	{
		Name:        "数据角色",
		Kind:        ModelPerm,
		IdentName:   "角色编码",
		StoreEngine: StoreEngineUSF,
		Ver:         1,
		Slots: []*Column{
			{Name: "角色编码", Kind: ColumnText, Required: true, Slot: 1, Ordinal: 1},
			{Name: "角色名称", Kind: ColumnText, Required: true, Slot: 2, Ordinal: 2},
		},
	},
}
