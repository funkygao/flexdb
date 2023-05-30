package perm

const (
	model = `
[request_definition]
# 一个基本的请求是一个元组对象，至少包含subject（访问实体）, object（访问的资源）和 action（访问方法）
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
# some|any|max|min
# some表示括号中的表达式个数大于等于1就行
#
# 这句话的意思就是将 request 和所有 policy 比对完之后，所有 policy 的策略结果（p.eft）为allow的个数 >=1，整个请求的策略就是为 true。
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`

	policy = `
# policy与policy_definition一一对应
p, alice, data1, read
p, bob, data2, write
`
)
