import bg from '@/assets/bg.jpg'

export default {
  init: function () {
    return {
      img: bg,
      logining: false,
      checked: true,
      showText: '显示密码',
      pwdType: 'password',
      userInfo: {
        user_id: '',
        username: '',
        password: '',
        is_admin: ''
      },
      show: false
    }
  }
}
