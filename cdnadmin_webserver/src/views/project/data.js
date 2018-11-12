export default{
  init: function () {
    let filter = {
    }
    return {
      projectData: [],
      projectAllData: [],
      platData: [],
      filter: filter, // 查询条件
      activeCollapse: 'search', // 开关查询折叠面板
      pagination: {
        total: 0,
        current: 1,
        pageSize: 10
      }, // 分页数据
      editVisible: false,
      dialogCreateProject: false,
      dialogEditProject: false,
      projectForm: {
        id: '',
        name: '',
        platform_name: '',
        cdn_id: '',
        path: '',
        domain: '',
        status: ''
      },
      editProjectForm: {
        id: '',
        name: '',
        platform_name: '',
        cdn_id: '',
        path: '',
        domain: '',
        status: ''
      },
      defaultProjectForm: {
        id: '',
        name: '',
        platform_name: '',
        cdn_id: '',
        path: '',
        domain: '',
        status: ''
      },
      projectStatus: [{
        id: 1,
        value: '使用中'
      }, {
        id: 2,
        value: '未使用'
      }, {
        id: 3,
        value: '已删除',
        disabled: true
      }]
    }
  }
}
