import { doLogin } from '@/api/login/index'
import crypto from 'crypto'

export default{
  showPwd () {
    if (this.pwdType === 'password') {
      this.pwdType = ''
      this.showText = '隐藏密码'
    } else {
      this.pwdType = 'password'
      this.showText = '显示密码'
    }
  },
  handleLogin (algorithm, data) {
    if (this.userInfo.username === '' || this.userInfo.username === 'undefined') {
      this.$notify({message: '用户名不能为空', type: 'warning'})
      return false
    }
    if (this.userInfo.password === '' || this.userInfo.password === 'undefined') {
      this.$notify({message: '密码不能为空', type: 'warning'})
      return false
    }
    if (this.userInfo.password.length < 6) {
      this.$notify({message: '密码长度不能小于6位', type: 'warning'})
      return false
    }
    let md5 = crypto.createHash('md5')
    md5.update(this.userInfo.password)
    this.cryptoPwd = md5.digest('hex', data)
    let para = {
      username: this.userInfo.username,
      password: this.cryptoPwd
    }
    // console.log(para)
    doLogin(para).then((res) => {
      // console.log(res)
      if (res.data.code === 200) {
        this.$store.commit('setToken', res.data.data)
        this.$store.commit('setUserName', res.data.data)
        this.$store.commit('setRoles', res.data.data)
        this.$notify({
          title: 'Success',
          message: 'Welcome ' + this.userInfo.username,
          type: 'success'
        })
        this.$router.push({path: '/cdn'})
      } else {
        this.$notify({
          title: 'Error',
          message: res.data.msg,
          type: 'error'
        })
      }
    }).catch(err => {
      console.log(err)
      this.$notify({
        title: 'Error',
        message: err,
        type: 'error'
      })
      return false
    })
  }
}
