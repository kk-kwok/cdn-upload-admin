export default{
  init: function () {
    let filter = {
    }
    return {
      suffixData: [],
      filter: filter, // 查询条件
      activeCollapse: 'search', // 开关查询折叠面板
      pagination: {
        total: 0,
        current: 1,
        pageSize: 10
      }, // 分页数据
      editVisible: false,
      dialogCreateSuffix: false,
      dialogEditSuffix: false,
      suffixForm: {
        file_suffix: ''
      },
      editSuffixForm: {
        id: '',
        file_suffix: ''
      },
      defaultSuffixForm: {
        file_suffix: ''
      }
    }
  }
}
