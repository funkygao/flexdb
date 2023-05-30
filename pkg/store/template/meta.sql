CREATE DATABASE IF NOT EXISTS flexmeta;
USE flexmeta;

# meta
DROP TABLE IF EXISTS Org, App, Page, Model, Field, Relationship, AppPage, Share;
# meta extension
DROP TABLE IF EXISTS ModelTrigger, ModelWebhook, Picklist;

CREATE TABLE Org (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    name varchar(20) NOT NULL,
    description varchar(500),
    secret char(20) COMMENT '对称加密的密钥',
    api_ver varchar(10) NOT NULL DEFAULT '0.1',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='租户';

CREATE TABLE App (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键', # 4,294,967,296 is enough
    org_id int unsigned NOT NULL,
    name varchar(20) NOT NULL,
    description varchar(500),
    uuid char(22) BINARY CHARACTER SET latin1 NOT NULL,
    logo varchar(100),
    status tinyint NOT NULL DEFAULT 1,
    visibility tinyint NOT NULL DEFAULT 0,

    ctime timestamp(3) NULL COMMENT '创建时间',
    mtime timestamp(3) NULL COMMENT '最近更新时间',
    cuser varchar(50),
    muser varchar(50),
    ts timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'DB timestamp',

    PRIMARY KEY (id),
    UNIQUE KEY idx_org_name (org_id, name), # (4 + 20*3) = 64
    UNIQUE KEY idx_org_uuid (org_id, uuid)  # (4 + 22) = 26
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用';

CREATE TABLE Relationship (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    from_model_id int unsigned NOT NULL,
    from_slot smallint unsigned NOT NULL,
    ref_kind smallint unsigned NOT NULL,
    to_model_id int unsigned NOT NULL,
    to_slot smallint unsigned NOT NULL,
    PRIMARY KEY (id),
    KEY idx_master_model (to_model_id, to_slot)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='模型间关系';

CREATE TABLE Share (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键', # 4,294,967,296 is enough
    app_id int unsigned NOT NULL COMMENT 'foreign key: App.id',
    subject varchar(20) NOT NULL,
    ctime timestamp(3) NULL COMMENT '创建时间',
    cuser varchar(50),
    PRIMARY KEY (id),
    UNIQUE KEY idx_app_subject (app_id, subject) # (4 + 20*3) = 64
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用共享关系';

CREATE TABLE Model (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键', # 4,294,967,296 is enough
    app_id int unsigned NOT NULL COMMENT 'foreign key: App.id',
    name varchar(20) NOT NULL,
    remark varchar(500),
    ident_name varchar(20), # Data.ident的显示名称
    kind smallint unsigned NOT NULL,
    store varchar(20) NOT NULL DEFAULT '',
    feature char(100) NOT NULL DEFAULT '',

    ctime timestamp(3) NULL COMMENT '创建时间',
    mtime timestamp(3) NULL COMMENT '最近更新时间',
    cuser varchar(50),
    muser varchar(50),
    ts timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'DB timestamp',
    deleted tinyint NOT NULL DEFAULT 0,
    ver smallint NOT NULL DEFAULT 0,

    PRIMARY KEY (id),
    UNIQUE KEY idx_app_name (app_id, name) # (4 + 20*3) = 64
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='业务模型定义';

CREATE TABLE Picklist (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    slot smallint unsigned NOT NULL,
    val varchar(127),
    parent_id int unsigned,

    ctime timestamp(3) NULL COMMENT '创建时间',
    mtime timestamp(3) NULL COMMENT '最近更新时间',
    cuser varchar(50),
    muser varchar(50),

    ts timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'DB timestamp',
    PRIMARY KEY (id),
    KEY idx_model_slot_parent (model_id, slot, parent_id) # (4 + 2 + 4) = 10
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='业务枚举项仓库';

CREATE TABLE ModelTrigger (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键', # 4,294,967,296 is enough
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    trigger_id int unsigned NOT NULL,
    payload varchar(2000),
    before_insert varchar(2000),
    after_insert varchar(2000),
    before_update varchar(2000),
    after_update varchar(2000),
    before_delete varchar(2000),
    after_delete varchar(2000),
    PRIMARY KEY (id),
    UNIQUE KEY idx_model (model_id, trigger_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='业务模型触发器';

CREATE TABLE ModelWebhook (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键', # 4,294,967,296 is enough
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    payload_url varchar(500),
    secret varchar(80),
    content_type smallint,
    event_mask smallint, # 哪些事件会推送
    active tinyint default 1,
    PRIMARY KEY (id),
    KEY idx_model (model_id, active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='业务模型webhooks';

CREATE TABLE Field (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    model_id int unsigned NOT NULL COMMENT 'foreign key: Model.id',
    name varchar(20) NOT NULL,
    label varchar(50),
    remark varchar(500),
    kind smallint unsigned NOT NULL,
    ordinal smallint unsigned NOT NULL default 0,
    indexed tinyint DEFAULT 0,
    slot smallint unsigned NOT NULL,
    required tinyint DEFAULT 0,
    deprecated tinyint DEFAULT 0,
    ro tinyint DEFAULT 0 COMMENT 'read only',
    uniq tinyint DEFAULT 0 COMMENT 'unique',
    sortable  tinyint DEFAULT 0,
    rule varchar(500),
    dftval varchar(127) COMMENT 'default value',

    ref_model_id int unsigned COMMENT 'reference model if kind is many2one,lookup',
    ref_slot smallint unsigned COMMENT 'reference model if kind is many2one,lookup',
    choices varchar(2000) COMMENT 'comma separated choices for kind=ColumnChoice',

    ctime timestamp(3) NULL COMMENT '创建时间',
    mtime timestamp(3) NULL COMMENT '最近更新时间',
    cuser varchar(50),
    muser varchar(50),
    ts timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'DB timestamp',

    PRIMARY KEY (id),
    UNIQUE KEY idx_model_name (model_id, name) # (4 + 80) = 84
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='业务字段定义';

CREATE TABLE AppPage (
    id int unsigned NOT NULL AUTO_INCREMENT COMMENT '物理主键',
    app_id int unsigned NOT NULL COMMENT 'foreign key: App.id',
    name varchar(30) NOT NULL,
    ordinal smallint unsigned NOT NULL default 0,
    render_schema text NOT NULL,

    ctime timestamp(3) NULL COMMENT '创建时间',
    mtime timestamp(3) NULL COMMENT '最近更新时间',
    cuser varchar(50),
    muser varchar(50),
    ts timestamp(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT 'DB timestamp',

    PRIMARY KEY (id),
    UNIQUE KEY idx_app_name (app_id, name),        # (4 + 30*3) = 94
    UNIQUE KEY idx_app_ordinal (app_id, ordinal)   # (4 + 2)    = 6
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用的页面';
