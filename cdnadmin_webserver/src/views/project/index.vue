<template>
  <div style="padding: 10px;">
    <!-- -------------------------------------------查询条件--------------------------------------------- -->
    <el-collapse v-model='activeCollapse'>
      <el-collapse-item title='查询&新建' name='search'>
        <el-form ref='searchCondition' :model='filter' label-width='80px' label-position=‘left’>
          <el-row>
            <el-col :span='4'>
              <el-form-item label='项目名称' prop='name'>
                <el-select v-model="filter.name" filterable value="">
                  <el-option v-for="(item, index) in projectAllData" :key="index" :label="item.name" :value="item.name"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='4'>
              <el-form-item label='CDN平台' prop='platform_name'>
                <el-select v-model="filter.platform_name" filterable value="">
                  <el-option v-for="(item, index) in platData" :key="index" :label="item.name" :value="item.name"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='3'>
              <el-form-item label='使用状态' prop='status'>
                <el-select v-model="filter.status" filterable value="">
                  <el-option v-for="(item, index) in projectStatus" :key="index" :label="item.value" :value="item.value"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='3' style="margin-left:20px;">
              <el-button @click='handleSearch' type="primary" icon="el-icon-search" size="small">查询</el-button>
              <el-button @click='resetForm("searchCondition")' type="primary" icon="el-icon-refresh" size="small">重置</el-button>
            </el-col>
            <el-col :span="2" style="margin-left: 20px">
              <el-button @click='dialogCreateProject=true' type='success' icon="el-icon-edit" size="small">新建项目</el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-collapse-item>
    </el-collapse>
    <!-- -------------------------------------------Project List--------------------------------------------- -->
    <div>
      <div style='margin-bottom: 10px'>
        <el-table :data='projectData' border stripe
                  :default-sort="{prop: '', order: 'descending'}" max-height="640" style="width: auto">
          <el-table-column label='序号' min-width='50' align="center">
            <template slot-scope="scope">
              <span>{{scope.$index+(pagination.current - 1) * pagination.pageSize + 1}}</span>
            </template>
          </el-table-column>
          <el-table-column label='ID' prop='id' sortable min-width='30' width="100" align="center"></el-table-column>
          <el-table-column label='名称' prop='name' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='CDN平台名称' prop='platform_name' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='CDN域名' prop='domain' sortable resizable min-width='100' width="300" align="center"></el-table-column>
          <el-table-column label='CDN域名ID' prop='cdn_id' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='目录' prop='path' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='状态' prop='status' sortable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='操作' prop='operate' min-width="200" align="center">
            <template slot-scope="scope">
              <el-button @click="editProject(scope)" type="primary" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-button @click="delProject(scope)" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <el-row type='flex' justify='end'>
        <el-pagination
          background
          @size-change='handlePageSizeChange'
          @current-change='handleCurrentChange'
          :current-page='pagination.current'
          :page-sizes='[10, 50, 100]'
          :page-size='pagination.pageSize'
          :total='pagination.total'
          layout='prev,pager,next,jumper,total,sizes'
        ></el-pagination>
      </el-row>
    </div>
    <!-- -------------------------------------------------新建Project----------------------------------------------- -->
    <el-dialog title='新建项目' ref="createProject"  :visible.sync='dialogCreateProject' width="600px">
        <el-form :model='projectForm' ref="addProjectForm" label-width="150px">
          <el-form-item
            label='ID'
            prop='id'
            required
            :rules='{ required: true, message: "ID不能为空！"}'>
            <el-input v-model='projectForm.id' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='项目名称'
            prop='name'
            required
            :rules='{ required: true, message: "项目名称不能为空！"}'>
            <el-input v-model='projectForm.name' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='CDN平台名称'
            prop='platform_name'
            required
            :rules='{ required: true, message: "CDN平台名称不能为空！"}'>
            <el-select v-model="projectForm.platform_name" filterable  value="" style="width: 300px">
              <el-option v-for="(item, index) in platData" :key="index" :label="item.name" :value="item.name" ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item
            label='CDN域名'
            prop='domain'
            required
            :rules='{ required: true, message: "CDN域名不能为空！"}'>
            <el-input v-model='projectForm.domain' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='CDN域名ID'
            prop='cdn_id' >
            <el-input v-model='projectForm.cdn_id' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='目录'
            prop='path'
            required
            :rules='{ required: true, message: "目录不能为空！"}'>
            <el-input v-model='projectForm.path' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='状态'
            prop='status'>
            <el-select v-model="projectForm.status" filterable placeholder="未使用" value="">
              <el-option v-for="(item, index) in projectStatus" :key="index" :disabled="item.disabled"
                         :label="item.value" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseAddDialog'>取消</el-button>
        <el-button type='primary' @click='handleSave'>确定</el-button>
      </div>
    </el-dialog>

    <!-- -------------------------------------------------编辑Project----------------------------------------------- -->
    <el-dialog title='编辑项目' ref="editProject"  :visible.sync='dialogEditProject' width="600px">
      <el-form :model='editProjectForm' ref="editProjectForm" label-width="150px">
        <el-form-item
          label='名称'
          prop='name'
          required
          :rules='{ required: true, message: "名称不能为空！"}'>
          <el-input v-model='editProjectForm.name' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='CDN平台名称'
          prop='platform_name'
          required
          :rules='{ required: true, message: "CDN平台名称不能为空！"}'>
          <el-select v-model="editProjectForm.platform_name" filterable  value="" style="width: 300px">
           <el-option v-for="(item, index) in platData" :key="index" :label="item.name" :value="item.name" ></el-option>
          </el-select>
          <!--<el-input v-model='editProjectForm.platform_name' :maxlength='50' style="width: 300px"></el-input>-->
        </el-form-item>
        <el-form-item
          label='CDN域名'
          prop='domain'
          required
          :rules='{ required: true, message: "CDN域名不能为空！"}'>
          <el-input v-model='editProjectForm.domain' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='CDN域名ID'
          prop='cdn_id' >
          <el-input v-model='editProjectForm.cdn_id' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='目录'
          prop='path'
          required
          :rules='{ required: true, message: "目录不能为空！"}'>
          <el-input v-model='editProjectForm.path' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='状态'
          prop='status'>
          <el-select v-model="editProjectForm.status" filterable  value="">
            <el-option v-for="(item, index) in projectStatus" :key="index" :disabled="item.disabled"
                       :label="item.value" :value="item.value"></el-option>
          </el-select>
          <!--<el-input v-model='editForm.status' :maxlength='50' style="width: 200px"></el-input>-->
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseEditDialog'>取消</el-button>
        <el-button type='primary' @click='updateProjectById'>确定</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script type="text/javascript">

import methods from './method'
import data from './data'

export default{
  data () {
    return data.init()
  },
  methods: methods,
  mounted () {
    this.getProjectData()
    this.getPlatData()
  }
}
</script>
<style type="text/css">
</style>
