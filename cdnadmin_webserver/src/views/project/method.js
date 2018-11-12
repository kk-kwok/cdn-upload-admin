import { getProjectList, deleteProject, addProject, updateProject } from '@/api/project/project'
import { getPlatformList } from '@/api/platform/platform'

export default{
  // 条件搜索
  handleSearch () {
    this.pagination.current = 1
    this.getProjectData()
  },
  // 重置搜索条件并刷新数据
  resetForm (formName) {
    this.$refs[formName].resetFields()
    this.handleSearch()
  },
  // 设置分页大小
  handlePageSizeChange (pageSize) {
    this.pagination.pageSize = pageSize
    this.getProjectData()
  },
  // 设置页码
  handleCurrentChange (current) {
    this.pagination.current = current
    this.getProjectData()
  },
  handleCloseAddDialog () {
    this.dialogCreateProject = false
    this.projectForm = Object.assign({}, this.defaultProjectForm)
    this.$refs.addProjectForm.resetFields()
    this.$refs.createProject.close()
  },
  handleCloseEditDialog () {
    this.dialogEditProject = false
    this.editProjectForm = Object.assign({}, this.defaultProjectForm)
    this.$refs.editProjectForm.resetFields()
    this.$refs.editProject.close()
  },
  // 提交新建表单
  handleSave () {
    this.$refs.addProjectForm.validate((valid) => {
      if (valid) {
        const params = Object.assign({}, this.projectForm)
        // console.log(params)
        addProject(params).then((res) => {
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
          this.$refs.createProject.close()
          this.projectForm = Object.assign({}, this.defaultProjectForm)
          this.$refs.addProjectForm.resetFields()
          this.handleCloseAddDialog()
          this.getProjectData()
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
  editProject (scope) {
    this.dialogEditProject = true
    this.editProjectForm = Object.assign({}, {
      id: scope.row.id,
      name: scope.row.name,
      platform_name: scope.row.platform_name,
      cdn_id: scope.row.cdn_id,
      path: scope.row.path,
      domain: scope.row.domain,
      status: scope.row.status
    })
  },
  updateProjectById () {
    let params = Object.assign({}, this.editProjectForm)

    // console.log(params)
    updateProject(params).then((res) => {
      // console.log(res)
      if (res.data.code === 200) {
        this.$notify({
          title: 'success',
          type: 'success',
          message: '修改项目信息成功'
        })
      } else {
        this.$notify({
          title: 'error',
          type: 'error',
          message: res.data.msg
        })
      }
      this.dialogEditProject = false
      this.getProjectData()
    }).catch((err) => {
      console.log(err)
      this.$notify({
        title: 'error',
        type: 'error',
        message: '修改项目信息失败'
      })
    })
  },
  // 删除操作
  delProject (scope) {
    this.$confirm('此操作将删除选中项, 是否继续?', '提示', {
      confirmButtonText: '确定',
      cancelButtonText: '取消',
      type: 'error'
    }).then(() => {
      this.removeProject(scope)
    }).catch(() => {
      this.$notify({
        title: 'error',
        type: 'error',
        message: '已取消删除'
      })
    })
  },
  removeProject (scope) {
    const params = {
      id: scope.row.id,
      name: scope.row.name
    }
    // console.log(scope)
    deleteProject(params).then((res) => {
      this.getProjectData()
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
  getProjectData: function () {
    let para = {
      page_num: this.pagination.current,
      page_size: this.pagination.pageSize,
      ...this.filter
    }
    // console.log(para)
    getProjectList(para).then((res) => {
      // console.log(res)
      this.projectData = res.data.data
      this.pagination.total = res.data.total
    })
    getProjectList().then((res) => {
      // console.log(res)
      this.projectAllData = res.data.data
    })
  },
  // 获取platform
  getPlatData: function () {
    let para = {
      status: '使用'
    }
    getPlatformList(para).then((res) => {
      // console.log(res)
      this.platData = res.data.data
    })
  }
}
