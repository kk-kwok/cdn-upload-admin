import { getPlatformList, deletePlatform, addPlatform, updatePlatform } from '@/api/platform/platform'

export default{
  // 条件搜索
  handleSearch () {
    this.pagination.current = 1
    this.getPlatformData()
  },
  // 重置搜索条件并刷新数据
  resetForm (formName) {
    this.$refs[formName].resetFields()
    this.handleSearch()
  },
  // 设置分页大小
  handlePageSizeChange (pageSize) {
    this.pagination.pageSize = pageSize
    this.getPlatformData()
  },
  // 设置页码
  handleCurrentChange (current) {
    this.pagination.current = current
    this.getPlatformData()
  },
  handleCloseAddDialog () {
    this.dialogCreatePlatform = false
    this.platformForm = Object.assign({}, this.defaultPlatformForm)
    this.$refs.addPlatformForm.resetFields()
    this.$refs.createPlatform.close()
  },
  handleCloseEditDialog () {
    this.dialogEditPlatform = false
    this.editPlatformForm = Object.assign({}, this.defaultPlatformForm)
    this.$refs.editPlatformForm.resetFields()
    this.$refs.editPlatform.close()
  },
  // 提交新建表单
  handleSave () {
    this.$refs.addPlatformForm.validate((valid) => {
      if (valid) {
        const params = Object.assign({}, this.platformForm)
        // console.log(params)
        addPlatform(params).then((res) => {
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
          this.$refs.createPlatform.close()
          this.platformForm = Object.assign({}, this.defaultPlatformForm)
          this.$refs.addPlatformForm.resetFields()
          this.handleCloseAddDialog()
          this.getPlatformData()
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
  editPlatform (scope) {
    this.dialogEditPlatform = true
    this.editPlatformForm = Object.assign({}, {
      id: scope.row.id,
      name: scope.row.name,
      secret_id: scope.row.secret_id,
      secret_key: scope.row.secret_key,
      api_url: scope.row.api_url,
      action: scope.row.action,
      status: scope.row.status
    })
  },
  updatePlatformById () {
    let params = Object.assign({}, this.editPlatformForm)

    // console.log(params)
    updatePlatform(params).then((res) => {
      // console.log(res)
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '修改CDN平台信息成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
      this.dialogEditPlatform = false
      this.getPlatformData()
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '修改CDN平台信息失败'
      })
    })
  },
  // 删除操作
  delPlatform (scope) {
    this.$confirm('此操作将删除选中项, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(() => {
      this.removePlatform(scope)
    }).catch(() => {
      this.$notify({
        title: 'error',
        type: 'error',
        message: '已取消删除'
      })
    })
  },
  removePlatform (scope) {
    const params = {
      id: scope.row.id,
      name: scope.row.name
    }
    // console.log(scope)
    deletePlatform(params).then((res) => {
      this.getPlatformData()
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
  getPlatformData: function () {
    let para = {
      page_num: this.pagination.current,
      page_size: this.pagination.pageSize,
      ...this.filter
    }
    // console.log(para)
    getPlatformList(para).then((res) => {
      // console.log(res)
      this.platformData = res.data.data
      this.pagination.total = res.data.total
    })
  }
}
