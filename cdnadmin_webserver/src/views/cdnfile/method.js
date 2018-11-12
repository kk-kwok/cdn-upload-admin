import { getCDNFileList, deleteCDNFile, addCDNFile, updateCDNFile, pushCDNFile } from '@/api/cdnfile/cdnfile'
import { getProjectList } from '@/api/project/project'
import { getSuffixList } from '@/api/suffix/suffix'

export default{
  // 条件搜索
  handleSearch () {
    this.pagination.current = 1
    this.getCDNFileData()
  },
  // 重置搜索条件并刷新数据
  resetForm (formName) {
    this.$refs[formName].resetFields()
    this.handleSearch()
  },
  // 设置分页大小
  handlePageSizeChange (pageSize) {
    this.pagination.pageSize = pageSize
    this.getCDNFileData()
  },
  // 设置页码
  handleCurrentChange (current) {
    this.pagination.current = current
    this.getCDNFileData()
  },
  handleCloseAddDialog () {
    this.file_suffix_math = ''
    this.file_url = ''
    this.cfParam = new FormData()
    this.showProgressFlag = false
    this.uploadPercent = 0
    this.dialogCreateCDNFile = false
    this.cdnFileForm = Object.assign({}, this.defaultCDNFileForm)
    this.$refs.upload.clearFiles()
    this.$refs.addCDNFileForm.resetFields()
    this.$refs.createCDNFile.close()
  },
  handleCloseEditDialog () {
    this.file_suffix_math = ''
    this.file_url = ''
    this.dialogEditCDNFile = false
    this.editCDNFileForm = Object.assign({}, this.defaultCDNFileForm)
    this.$refs.editCDNFileForm.resetFields()
    this.$refs.editCDNFile.close()
  },
  // 提交新建表单
  handleFile (file, fileList) {
    this.fileList = fileList
  },
  handleExceed (files, fileList) {
    this.$message.warning(`最多上传 ${files.length} 个文件`)
  },
  handleSave () {
    this.$refs.addCDNFileForm.validate((valid) => {
      if (valid) {
        // const params = Object.assign({}, this.cdnFileForm)
        this.cfParam.append('file', this.fileList[0].raw)
        this.cfParam.append('project_name', this.cdnFileForm.project_name)
        this.cfParam.append('file_name', this.cdnFileForm.file_name)
        this.cfParam.append('comment', this.cdnFileForm.comment)
        const token = this.$store.state.userInfo.token ? this.$store.state.userInfo.token : window.sessionStorage.getItem('token')
        let config = {
          headers: {
            // 'Content-Type': 'multipart/form-data',
            'authorization': token
          },
          onUploadProgress: progressEvent => {
            this.uploadPercent = Math.floor(progressEvent.loaded / progressEvent.total * 100)
          }
        }
        if (this.validFileSuffix(this.fileList[0].name) === 1) {
          this.$message.warning('上传文件格式错误, 格式只能为: ' + this.file_suffix)
          this.file_suffix_math = 0
        } else if (this.validFileSuffix(this.cdnFileForm.file_name) === 1) {
          this.$message.warning('新文件名格式错误, 格式只能为: ' + this.file_suffix)
          this.file_suffix_math = 0
        } else if (this.fileList[0].size > 1024 * 1024 * 1024) {
          this.$message.warning('上传文件大小不能超过1GB')
        } else {
          this.showProgressFlag = true
          addCDNFile(this.cfParam, config).then((res) => {
            if (res.data.code === 200) {
              this.$notify({
                title: 'success',
                type: 'success',
                message: '上传成功, 除越南和泰国资源服以外大文件刷新CDN需要20-30分钟'
              })
            } else {
              this.$notify({
                title: 'error',
                type: 'error',
                message: res.data.msg
              })
            }
            this.$refs.createCDNFile.close()
            this.cdnFileForm = Object.assign({}, this.defaultCDNFileForm)
            this.$refs.addCDNFileForm.resetFields()
            this.handleCloseAddDialog()
            this.getCDNFileData()
          }).catch((err) => {
            this.$notify({
              title: 'error',
              type: 'error',
              message: '上传失败'
            })
            console.log(err)
          })
        }
      } else {
        console.log('error submit!!')
        return false
      }
    })
  },
  // 编辑操作
  editCDNFile (scope) {
    this.dialogEditCDNFile = true
    this.editCDNFileForm = Object.assign({}, {
      id: scope.row.id,
      file_name: scope.row.file_name,
      comment: scope.row.comment
    })
  },
  updateCDNFileById () {
    let params = Object.assign({}, this.editCDNFileForm)

    console.log(params)
    updateCDNFile(params).then((res) => {
      // console.log(res)
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '修改CDN文件信息成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
      this.dialogEditCDNFile = false
      this.getCDNFileData()
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '修改CDN文件信息失败'
      })
    })
  },
  // 刷新CDN
  refreshCDNFile (scope) {
    this.$confirm('此操作将进行刷新CDN文件，视文件大小大概需要30分钟生效，是否继续？(越南和泰国资源服不用进行此操作)', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(() => {
      this.refreshCDNFileById(scope)
    }).catch(() => {
      this.$notify({
        title: 'error',
        type: 'error',
        message: '已取消刷新CDN文件'
      })
    })
  },
  refreshCDNFileById (scope) {
    const params = {
      id: scope.row.id
    }
    pushCDNFile(params).then((res) => {
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: res.data.msg
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
  // 删除操作
  delCDNFile (scope) {
    this.$confirm('此操作将删除选中项, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(() => {
      this.removeCDNFile(scope)
    }).catch(() => {
      this.$notify({
        title: 'error',
        type: 'error',
        message: '已取消删除'
      })
    })
  },
  removeCDNFile (scope) {
    const params = {
      id: scope.row.id,
      file_name: scope.row.file_name
    }
    // console.log(scope)
    deleteCDNFile(params).then((res) => {
      this.getCDNFileData()
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
  getCDNFileData: function () {
    let para = {
      page_num: this.pagination.current,
      page_size: this.pagination.pageSize,
      ...this.filter
    }
    // console.log(para)
    getCDNFileList(para).then((res) => {
      // console.log(res)
      this.cdnFileData = res.data.data
      this.pagination.total = res.data.total
    })
  },
  // 获取project_name
  getPjNameData: function () {
    getProjectList().then((res) => {
      this.pjNameData = res.data.data
    })
  },
  // 获取file suffix
  getCSuffixData: function () {
    getSuffixList().then((res) => {
      this.cSuffixData = res.data.data
      for (let val of this.cSuffixData) {
        if (this.file_suffix === '') {
          this.file_suffix = '.' + val.file_suffix
        } else {
          this.file_suffix = this.file_suffix + ',.' + val.file_suffix
        }
      }
    })
  },
  // get url
  getURL: function (name) {
    for (let val of this.pjNameData) {
      if (name === val.name) {
        this.file_url = 'http://' + val.domain + '/' + val.path
      }
    }
  },
  // 判断文件后缀名
  validFileSuffix: function (name) {
    let n = 0
    for (let val of this.cSuffixData) {
      if (name.endsWith(val.file_suffix)) {
        n = n + 1
      }
    }
    if (n === 0) {
      this.file_suffix_math = 1
      return 1
    } else {
      this.file_suffix_math = 0
      return 0
    }
  },
  // 判断文件大小单位
  calcFileSize: function (scope) {
    if (scope.row.file_size > 1048576) {
      return 1
    } else if (scope.row.file_size > 1024) {
      return 2
    } else {
      return 3
    }
  }
}
