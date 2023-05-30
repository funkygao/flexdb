CREATE DATABASE IF NOT EXISTS flexdata;
USE flexdata;

DROP TABLE IF EXISTS Data, DataChangelog, Clob, Task, IndexStr, IndexStrUniq, IndexInt, IndexIntUniq, IndexTime, IndexTimeUniq;

# Data of Org:1 Data_1
CREATE TABLE Data (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',

    s1 varchar(127),
    s2 varchar(127),
    s3 varchar(127),
    s4 varchar(127),
    s5 varchar(127),
    s6 varchar(127),
    s7 varchar(127),
    s8 varchar(127),
    s9 varchar(127),
    s10 varchar(127),
    s11 varchar(127),
    s12 varchar(127),
    s13 varchar(127),
    s14 varchar(127),
    s15 varchar(127),
    s16 varchar(127),
    s17 varchar(127),
    s18 varchar(127),
    s19 varchar(127),
    s20 varchar(127),
    s21 varchar(127),
    s22 varchar(127),
    s23 varchar(127),
    s24 varchar(127),
    s25 varchar(127),
    s26 varchar(127),
    s27 varchar(127),
    s28 varchar(127),
    s29 varchar(127),
    s30 varchar(127),
    s31 varchar(127),
    s32 varchar(127),
    s33 varchar(127),
    s34 varchar(127),
    s35 varchar(127),
    s36 varchar(127),
    s37 varchar(127),
    s38 varchar(127),
    s39 varchar(127),
    s40 varchar(127),
    s41 varchar(127),
    s42 varchar(127),
    s43 varchar(127),
    s44 varchar(127),
    s45 varchar(127),
    s46 varchar(127),
    s47 varchar(127),
    s48 varchar(127),
    s49 varchar(127),
    s50 varchar(127),

    r1 varchar(127),
    r2 varchar(127),
    r3 varchar(127),
    r4 varchar(127),
    r5 varchar(127),
    r6 varchar(127),
    r7 varchar(127),
    r8 varchar(127),
    r9 varchar(127),
    r10 varchar(127),
    memo varchar(512),

    # 默认字段，都直接建立在数据表
    ctime timestamp(3) NOT NULL COMMENT '创建时间',
    mtime timestamp(3) NOT NULL COMMENT '最近更新时间',
    cuser varchar(50),
    muser varchar(50),
    ts timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'DB timestamp',
    deleted tinyint NOT NULL DEFAULT 0,

    ver smallint unsigned NOT NULL COMMENT 'optimistic lock',

    PRIMARY KEY (id),
    KEY idx_model_mtime (model_id, mtime),    # (4 + 4) = 8
    KEY idx_model_cuser (model_id, cuser(10)) # (4 + 40) = 44
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='业务数据存储';

CREATE TABLE DataChangelog (
    row_id int unsigned NOT NULL COMMENT '主键',
    diff varchar(2000),
    ctime timestamp(3) NOT NULL COMMENT '创建时间',
    cuser varchar(50),
    PRIMARY KEY (row_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='记录的变更历史';

CREATE TABLE Clob (
    row_id int unsigned NOT NULL COMMENT '主键',
    s1 varchar(2000),
    s2 varchar(2000),
    s3 varchar(2000),
    deleted tinyint NOT NULL DEFAULT 0,
    PRIMARY KEY (row_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='业务数据的大字段存储';

# IndexStr of Org:1 IndexStr_1
CREATE TABLE IndexStr (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL, # 32,767 slots is enough
    val varchar(50) NOT NULL, # 不是所有string都可以index的, multi-select picklist/textarea/加密字段都不可以建索引，null时不会录入索引
    row_id int unsigned NOT NULL COMMENT 'Data.id',
    PRIMARY KEY (id),
    KEY idx_strval (model_id, slot, val(15)) # (4 + 2 + 15*4 + 3) = 69
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='非唯一索引表，保存字符串类型索引';

CREATE TABLE IndexStrUniq (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL, # 32,767 slots is enough
    val varchar(50) NOT NULL, # 不是所有string都可以index的, multi-select picklist/textarea/加密字段都不可以建索引，null时不会录入索引
    row_id int unsigned NOT NULL COMMENT 'Data.id',
    PRIMARY KEY (id),
    UNIQUE KEY idx_strval_uniq (model_id, slot, val(15)) # (4 + 2 + 15*4 + 3) = 69
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='非唯一索引表，保存字符串类型索引';

CREATE TABLE IndexInt (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL, # 32,767 slots is enough
    val bigint NOT NULL,
    row_id int unsigned NOT NULL COMMENT 'Data.id',
    PRIMARY KEY (id),
    KEY idx_intval (model_id, slot, val) # (4 + 2 + 8) = 14
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='非唯一索引表，保存整数类型索引';

CREATE TABLE IndexIntUniq (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL, # 32,767 slots is enough
    val bigint NOT NULL,
    row_id int unsigned NOT NULL COMMENT 'Data.id',
    PRIMARY KEY (id),
    UNIQUE KEY idx_intval_uniq (model_id, slot, val) # (4 + 2 + 8) = 14
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='非唯一索引表，保存整数类型索引';

CREATE TABLE IndexTime (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL, # 32,767 slots is enough
    val timestamp NOT NULL,
    row_id int unsigned NOT NULL COMMENT 'Data.id',
    PRIMARY KEY (id),
    KEY idx_intval (model_id, slot, val) # (4 + 2 + 4) = 10
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='非唯一索引表，保存时间类型索引';

CREATE TABLE IndexTimeUniq (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL, # 32,767 slots is enough
    val timestamp NOT NULL,
    row_id int unsigned NOT NULL COMMENT 'Data.id',
    PRIMARY KEY (id),
    UNIQUE KEY idx_intval (model_id, slot, val) # (4 + 2 + 4) = 10
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='非唯一索引表，保存时间类型索引';

CREATE TABLE Task (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    app_id int unsigned NOT NULL COMMENT 'foreign key: App.id',
    status tinyint NOT NULL DEFAULT 1,
    body varchar(2000),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Task';
