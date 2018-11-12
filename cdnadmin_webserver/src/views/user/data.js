export default{
  init: function () {
    let filter = {
    }
    return {
      currentUserData: [],
      userData: [],
      isSamePwd: 0,
      isSamePwd2: 0,
      showText: '显示密码',
      pwdType: 'password',
      filter: filter, // 查询条件
      activeCollapse: 'search', // 开关查询折叠面板
      pagination: {
        total: 0,
        current: 1,
        pageSize: 10
      }, // 分页数据
      editVisible: false,
      dialogCreateUser: false,
      dialogEditUser: false,
      dialogUpdatePwd: false,
      userForm: {
        username: '',
        name: '',
        status: ''
      },
      editUserForm: {
        id: '',
        name: '',
        status: ''
      },
      defaultUserForm: {
        id: '',
        name: '',
        status: ''
      },
      updatePwdForm: {
        id: '',
        old_password: '',
        password: '',
        password2: ''
      },
      defaultUpdatePwdForm: {
        id: '',
        old_password: '',
        password: '',
        password2: ''
      },
      userStatus: [{
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
