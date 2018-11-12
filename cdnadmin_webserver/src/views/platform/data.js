export default{
  init: function () {
    let filter = {
    }
    return {
      platformData: [],
      filter: filter, // 查询条件
      activeCollapse: 'search', // 开关查询折叠面板
      pagination: {
        total: 0,
        current: 1,
        pageSize: 10
      }, // 分页数据
      editVisible: false,
      dialogCreatePlatform: false,
      dialogEditPlatform: false,
      platformForm: {
        name: '',
        secret_id: '',
        secret_key: '',
        api_url: '',
        action: '',
        status: ''
      },
      editPlatformForm: {
        id: '',
        name: '',
        secret_id: '',
        secret_key: '',
        api_url: '',
        action: '',
        status: ''
      },
      defaultPlatformForm: {
        name: '',
        secret_id: '',
        secret_key: '',
        api_url: '',
        action: '',
        status: ''
      },
      platformStatus: [{
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
