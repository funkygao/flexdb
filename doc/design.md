# Tech Design

产品能力、集成能力、扩展能力、二开能力、生态化。

元数据是FlexDB基础架构的核心，一切的数据存储、页面交互、插件和扩展都是通过元数据引擎进行解释执行的。元数据使得业务人员描述业务然后系统自动生成成为可能。

元数据的存储与Salesforce完全相同，但采取了抄近的方式使得我们实现成本大大降低：不暴露SOQL，而是暴露API。

底层插件和扩展/集成架构，参考了Odoo：微内核架构，所有的字段类型/触发器/连接器都通过插件实现。

由于元数据机制，使得FlexDB拥有更高的整合度和DDL-free：它的复杂性提高了，但复杂性被封装到了FlexDB内部。



<details>
<summary><b>Table of content</b></summary>

## Table of content
   * [Assumption](#assumption)
      * [Model](#model)
      * [Page](#page)
      * [Instance](#instance)
      * [Sharding](#sharding)
   * [Details](#details)
      * [Constraints](#constraints)
      * [Native Columns](#native-columns)
      * [Validator DSL](#validator-dsl)
      * [OData](#odata)
   * [vs alternatives](#vs-alternatives)
      * [EAV](#eav)
   * [Capacity Planning](#capacity-planning)
      * [DB](#db)
   * [Deployment](#deployment)

</details>

----

## Layered Architecture

```
    amis  odata
     |     |
     +-----+
       |
    controller
       |
    entity
       |
    store
       |
    driver
```

## Assumption

We design based upon the following assumptions.

### Model

- 一个字段，只能属于一个模型：one model has many columns
   - 因此，是在字段表里加了一个`model_id`，而不是通过中间映射表匹配
- 每个模型有一个主识别字段: primary field
   - it can be referenced
   - its data type can be defined
   - above all, it can be searched
      - indexed

### Page

- A page must belong to 1 model
   - it cannot be shared among models
   - it can be generated from a model

### Instance

在定义好元信息后，即模型建立后，就可以在该模型上进行实例化(具体化)
- 模型：元信息
- 数据：实例化
   - 模型(元信息)的组成属性的各个取值能够构成一个实例
   - 包含具体的业务规则

### Sharding

- all data of an org must reside in 1 POD
- all data of an org must reside in 1 database instance
- each org has dedidated Data/Clob/Indexes tables for all models
- each table can use auto_increment

## Details

### Constraints

- model
   - a model can have at most {maxIndexesPerModel} indexes
   - a model can have at most {maxTextAreas} TextArea columns
   - a model can have at most {maxSlots} slots
- column
   - textarea column cannot be indexed
   - slot readonly, calculated
   - ref model, check it exists

### Native Columns

用户不能改动的字段(predefined columns)，包括2个层面：
- 所有模型的默认字段
   - 是在`Data`数据表里定义的字段，自动索引，索引是建立在数据表里的
   - create_time, created_by, update_time, updated_by, deleted
- 某一个模型的默认字段(模板模型)
   - Account: a, b c
   - Task: d, e
   - based on predefined config

### Validator DSL

>required, length>5, >0

### OData

OData (Open Data Protocol) is SQL for the Web.

## vs alternatives

### EAV

Entity-Attribute-Value，the what we call：列转行。

EAV的缺点：
- 无法索引
- 无法执行聚合运算
- 行放大
   - 查询效率低
- NoSQL并不能解决EAV的问题
   - MongoDB? Redis?
   - DynamoDB?

``` sql
SELECT a.entity AS id,
    a.value AS title,
    y.value AS production_year,
    r.value AS rating,
    b.value AS budget
FROM Attributes AS a
JOIN Attributes AS y USING (entity)
JOIN Attributes AS r USING (entity)
JOIN Attributes AS b USING (entity)
WHERE a.attribute = 'title'
    AND y.attribute = 'production_year'
    AND r.attribute = 'rating'
    AND b.attribute = 'budget'
```

### Odoo

各自租户拥有各自的数据库：
1. 如果用户需要修改或者扩展现有物理数据模型而进行的DDL操作，必然会影响线上业务的整体可用性，也可能会影响到标准数据模型，从而影响到线上功能使用。
2. 如果用户可自定义对物理模型进行扩展和定制，当平台进行模型升级的时候，极容易产生物理模型的冲突，导致新旧功能异常。
3. 由于用户在各自数据库存在各自定义的扩展和定制，则平台数据模型和功能升级，需要针对不同的租户进行分别验证，存在着极大的升级验证工作量和风险。

## APIs

``` json
{
    "status": 0,
    "msg": "",
    "data": {}
}
```

Example:
``` json
{
    "status": 0,
    "msg": "OK",
    "data": {
        "count": 8,
        "rows": [
            {
                "id": 7,
                "identify": "/tools/tpl_mall",
                "name": "模板商城",
                "source": "示例源码"
            },
            {
                "id": 8,
                "identify": "/tools/tpl_mall",
                "name": "模板商城",
                "source": "示例源码"
            }
        ]
    }
}
```

``` json
{
    "status": -1,
    "msg": "操作失败",
    "data": {}
}
```

## Capacity Planning

### DB

假设：
```
1000个org
   10个app
      10个model
         30个field
         1,000,000条记录
      10个page
```

索引空间：
- meta
   - App: 64
      - records: 10K
      - index:   640KB
   - Model: 64
      - records: 100K
      - index:   6.4MB
   - Field: 70
      - records: 3M
      - index:   210MB
   - Page: 70
      - records: 100K
      - index:   7MB
   - Totals
      - 250MB
- data
   - Data: 16
      - records: 1G
      - index:   16GB
   - Clob: 8
   - Indexes: 69/14/10
      - records: 5G
      - index:   150GB

## Deployment

![](http://www.plantuml.com/plantuml/svg/RPAnheKW38Ptdg8lu43TRwHlWiAcSgEGCvZpxiSw80rbQEZN_YdzmtHcV-IohSOiQoe1_bDmSC5z38VWiv_z6M6BsJk9-EAimb1XuyFs529yukoiJD7KYKyLXk6l6SaCNoRxwZjwY1eSAqHDZfeSM_ctyOhWwdsYVd-o2aDX0bcPIa_ezOGgGuMzcvBL80frxLGWygP-oUc0SiR3S82xo5-aJSXCSGzbMty0)

