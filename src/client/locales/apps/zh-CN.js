export default {
  app: {
    empty: '还没有任何应用哦~',
    plugin: '插件',
    create: {
      name: '请输入新应用的名称',
      minlength: '请输入3位以上的名称',
      character: '应用名称只能为数字字母下划线和减号',
      btn: '创建新应用'
    },
    tab: {
      overview: '概览',
      resources: '资源',
      deploy: '部署',
      activity: '事件',
      settings: '设置'
    },
    resources: {
      tip: '当前只对官方插件进行简单的管理操作，其它插件只提供查询、删除和启用/禁用功能',
      searchTip: '搜索插件',
      link: '关联',
      unlink: '解除关联',
      linkSuccess: '已关联{name}插件',
      linkFailed: '关联{name}插件失败',
      empty: '暂时没有安装插件哦',
      toggleFailed: '切换状态失败',
      plugin: {
        name: '名称',
        version: '版本',
        author: '作者',
        status: '状态',
        description: '描述'
      }
    },
    deploy: {
      repo: '代码来源',
      tipDesc: '',
      dokkuGit: 'Git仓库地址 ',
      fromEmpty: '创建一个新的项目',
      fromExisted: '从已有的项目开始',
      uploadText: '拖拽或<em>点击上传</em>文件，仅支持.tar和.tar.gz文件',
      label: {
        git: '使用Dokku Git作为仓库地址',
        tar: '上传tar压缩文件进行部署'
      },
      origin: {
        git: 'Dokku git仓库',
        tar: '压缩包'
      }
    },
    activity: {
      deployed: '推送了'
    },
    settings: {
      name: '应用名称',
      environment: '环境变量',
      info: '应用信息',
      buildpacks: 'Buildpacks',
      domainAndSSL: '域名和证书',
      deleteApp: '删除应用',
      removeAppConfirm: '确认要删除该应用么？'
    }
  }
}
