DROP TABLE IF EXISTS `web_auth`;
CREATE TABLE `web_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
INSERT INTO `web_auth` (`id`, `username`, `password`) VALUES ('1', '1', '1');

DROP TABLE IF EXISTS `web_article`;
CREATE TABLE `web_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '论文标题',
  `journal` varchar(100) DEFAULT '' COMMENT '期刊',
  `author` varchar(100) DEFAULT '' COMMENT '第一作者',
  `authors` varchar(100) DEFAULT '' COMMENT '其他作者',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `link` varchar(100) DEFAULT '' '详情页链接',
  `papercode` varchar(100) DEFAULT '' '代码',
  `theyear` varchar(10) DEFAULT '' COMMENT '论文年份',
  `abstract` text COMMENT '摘要',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '新建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '0' COMMENT '状态 1为禁用、0为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='论文管理';



DROP TABLE IF EXISTS `web_tag`;
CREATE TABLE `web_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT '删除时间',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';