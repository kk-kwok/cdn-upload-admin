import { getSuffixList, deleteSuffix, addSuffix, updateSuffix } from '@/api/suffix/suffix'

export default{
  // 条件搜索
  handleSearch () {
    this.pagination.current = 1
    this.getSuffixData()
  },
  // 重置搜索条件并刷新数据
  resetForm (formName) {
    this.$refs[formName].resetFields()
    this.handleSearch()
  },
  // 设置分页大小
  handlePageSizeChange (pageSize) {
    this.pagination.pageSize = pageSize
    this.getSuffixData()
  },
  // 设置页码
  handleCurrentChange (current) {
    this.pagination.current = current
    this.getSuffixData()
  },
  handleCloseAddDialog () {
    this.dialogCreateSuffix = false
    this.suffixForm = Object.assign({}, this.defaultSuffixForm)
    this.$refs.addSuffixForm.resetFields()
    this.$refs.createSuffix.close()
  },
  handleCloseEditDialog () {
    this.dialogEditSuffix = false
    this.editSuffixForm = Object.assign({}, this.defaultSuffixForm)
    this.$refs.editSuffixForm.resetFields()
    this.$refs.editSuffix.close()
  },
  // 提交新建表单
  handleSave () {
    this.$refs.addSuffixForm.validate((valid) => {
      if (valid) {
        const params = Object.assign({}, this.suffixForm)
        // console.log(params)
        addSuffix(params).then((res) => {
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
          this.$refs.createSuffix.close()
          this.suffixForm = Object.assign({}, this.defaultSuffixForm)
          this.$refs.addSuffixForm.resetFields()
          this.handleCloseAddDialog()
          this.getSuffixData()
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
  editSuffix (scope) {
    this.dialogEditSuffix = true
    this.editSuffixForm = Object.assign({}, {
      id: scope.row.id,
      file_suffix: scope.row.file_suffix
    })
  },
  updateSuffixById () {
    let params = Object.assign({}, this.editSuffixForm)

    // console.log(params)
    updateSuffix(params).then((res) => {
      // console.log(res)
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '修改文件后缀信息成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
      this.dialogEditSuffix = false
      this.getSuffixData()
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '修改文件后缀信息失败'
      })
    })
  },
  // 删除操作
  delSuffix (scope) {
    this.$confirm('此操作将删除选中项, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(() => {
      this.removeSuffix(scope)
    }).catch(() => {
      this.$notify({
        title: 'error',
        type: 'error',
        message: '已取消删除'
      })
    })
  },
  removeSuffix (scope) {
    const params = {
      id: scope.row.id,
      file_suffix: scope.row.file_suffix
    }
    // console.log(scope)
    deleteSuffix(params).then((res) => {
      this.getSuffixData()
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
  getSuffixData: function () {
    let para = {
      page_num: this.pagination.current,
      page_size: this.pagination.pageSize,
      ...this.filter
    }
    // console.log(para)
    getSuffixList(para).then((res) => {
      // console.log(res)
      this.suffixData = res.data.data
      this.pagination.total = res.data.total
    })
  }
}
