export default{
  init: function () {
    let filter = {
      project_name: ''
    }
    return {
      cfParam: new FormData(),
      cdnFileData: [],
      pjNameData: [],
      cSuffixData: [],
      fileList: [],
      file_suffix: '',
      file_suffix_math: '',
      file_url: '',
      filter: filter, // 查询条件
      uploadUrl: '',
      uploadPercent: 0,
      showProgressFlag: false,
      activeCollapse: 'search', // 开关查询折叠面板
      pagination: {
        total: 0,
        current: 1,
        pageSize: 10
      }, // 分页数据
      editVisible: false,
      dialogCreateCDNFile: false,
      dialogEditCDNFile: false,
      cdnFileForm: {
        user_id: '',
        project_name: '',
        file_name: '',
        comment: ''
      },
      editCDNFileForm: {
        id: '',
        project_name: '',
        file_name: '',
        comment: ''
      },
      defaultCDNFileForm: {
        project_name: '',
        file_name: '',
        comment: ''
      }
    }
  }
}
