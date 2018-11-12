import { getUserList, deleteUser, addUser, updateUser, updateUserPwd, resetUserPwd } from '@/api/user/user'
import store from '@/store/store.js'
import crypto from 'crypto'
import { setUserAdmin } from '../../api/user/user'

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
  // 条件搜索
  handleSearch () {
    this.pagination.current = 1
    this.getUserData()
  },
  // 重置搜索条件并刷新数据
  resetForm (formName) {
    this.$refs[formName].resetFields()
    this.handleSearch()
  },
  // 设置分页大小
  handlePageSizeChange (pageSize) {
    this.pagination.pageSize = pageSize
    this.getUserData()
  },
  // 设置页码
  handleCurrentChange (current) {
    this.pagination.current = current
    this.getUserData()
  },
  handleCloseAddDialog () {
    this.dialogCreateUser = false
    this.userForm = Object.assign({}, this.defaultUserForm)
    this.$refs.addUserForm.resetFields()
    this.$refs.createUser.close()
  },
  handleCloseEditDialog () {
    this.dialogEditUser = false
    this.editUserForm = Object.assign({}, this.defaultUserForm)
    this.$refs.editUserForm.resetFields()
    this.$refs.editUser.close()
  },
  handleCloseUpdateDialog () {
    this.isSamePwd = 0
    this.isSamePwd2 = 0
    this.dialogUpdatePwd = false
    this.updatePwdForm = Object.assign({}, this.defaultUpdatePwdForm)
    this.$refs.updatePForm.resetFields()
    this.$refs.updatePwd.close()
  },
  // 提交新建表单
  handleSave () {
    this.$refs.addUserForm.validate((valid) => {
      if (valid) {
        const params = Object.assign({}, this.userForm)
        // console.log(params)
        addUser(params).then((res) => {
          if (res.data.code === 200) {
            this.$notify({
              title: 'success',
              type: 'success',
              message: '新建成功'
            })
          } else {
            this.$notify({
              title: 'error',
              type: 'error',
              message: res.data.msg
            })
          }
          this.$refs.createUser.close()
          this.userForm = Object.assign({}, this.defaultUserForm)
          this.$refs.addUserForm.resetFields()
          this.handleCloseAddDialog()
          this.getUserData()
        }).catch((err) => {
          this.$notify({
            title: 'error',
            type: 'error',
            message: '新建失败'
          })
          console.log(err)
        })
      } else {
        console.log('error submit!!')
        return false
      }
    })
  },
  // 编辑操作
  editUser (scope) {
    this.dialogEditUser = true
    this.editUserForm = Object.assign({}, {
      id: scope.row.id,
      name: scope.row.name,
      status: scope.row.status
    })
  },
  updateUserById () {
    let params = Object.assign({}, this.editUserForm)

    // console.log(params)
    updateUser(params).then((res) => {
      // console.log(res)
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '修改用户信息成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
      this.dialogEditUser = false
      this.getUserData()
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '修改用户信息失败'
      })
    })
  },
  // 修改密码
  editUserPwd (scope) {
    this.dialogUpdatePwd = true
    this.updatePwdForm = Object.assign({}, {
      id: scope.row.id
    })
  },
  updateUserPwdById (algorithm, data) {
    let md5 = crypto.createHash('md5')
    md5.update(this.updatePwdForm.old_password)
    this.cryptoOldPwd = md5.digest('hex', data)

    md5 = crypto.createHash('md5')
    md5.update(this.updatePwdForm.password2)
    this.cryptoNewPwd = md5.digest('hex', data)

    let params = {
      id: this.updatePwdForm.id,
      old_password: this.cryptoOldPwd,
      password: this.cryptoNewPwd
    }
    if (this.isSamePwd2 !== 0) {
      this.$message.warning('输入错误请重新填写')
    } else {
      // console.log(params)
      updateUserPwd(params).then((res) => {
        // console.log(res)
        if (res.data.code === 200) {
          this.$notify({
            title: 'success',
            type: 'success',
            message: '修改用户密码成功'
          })
        } else {
          this.$notify({
            title: 'error',
            type: 'error',
            message: res.data.msg
          })
        }
        this.dialogUpdatePwd = false
      }).catch((err) => {
        console.log(err)
        this.$notify({
          title: 'error',
          type: 'error',
          message: '修改用户密码失败'
        })
      })
    }
  },
  // 重置用户默认密码
  resetUser (scope) {
    const params = {
      id: scope.row.id
    }
    resetUserPwd(params).then((res) => {
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '重置用户密码成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '重置用户密码失败'
      })
    })
  },
  // 重置用户默认密码
  resetAdmin (row) {
    const params = {
      id: row.id,
      is_admin: row.is_admin
    }
    setUserAdmin(params).then((res) => {
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '修改管理员状态成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '修改管理员状态失败'
      })
    })
  },
  // 删除操作
  delUser (scope) {
    this.$confirm('此操作将删除选中项, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(() => {
      this.removeUser(scope)
    }).catch(() => {
      this.$notify({
        title: 'error',
        type: 'error',
        message: '已取消删除'
      })
    })
  },
  removeUser (scope) {
    const params = {
      id: scope.row.id
    }
    // console.log(scope)
    deleteUser(params).then((res) => {
      this.getUserData()
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '删除成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
    }).catch((err) => {
      console.log(err)
    })
  },
  // 获取列表数据
  getUserData: function () {
    let para = {
      page_num: this.pagination.current,
      page_size: this.pagination.pageSize,
      ...this.filter
    }
    // console.log(para)
    getUserList(para).then((res) => {
      // console.log(res)
      this.userData = res.data.data
      this.pagination.total = res.data.total
    })
  },
  // 获取当前用户信息
  getCurrentUserData: function () {
    // noinspection JSAnnotator
    let para = {
      username: store.state.userInfo.username ? store.state.userInfo.username : window.sessionStorage.getItem('username')
    }
    // console.log(para)
    getUserList(para).then((res) => {
    // console.log(res)
      this.currentUserData = res.data.data
    })
  },
  // 验证新旧密码是否一样
  validSamePwd: function () {
    if (this.updatePwdForm.old_password === this.updatePwdForm.password) {
      this.isSamePwd = 1
    } else if (this.updatePwdForm.password.length < 6) {
      this.isSamePwd = 2
    } else {
      this.isSamePwd = 0
    }
    return this.isSamePwd
  },
  // 验证两次新密码是否一样
  validSamePwd2: function () {
    if (this.updatePwdForm.password !== this.updatePwdForm.password2) {
      this.isSamePwd2 = 1
    } else if (this.updatePwdForm.password2.length < 6) {
      this.isSamePwd2 = 2
    } else {
      this.isSamePwd2 = 0
    }
    return this.isSamePwd2
  }
}
