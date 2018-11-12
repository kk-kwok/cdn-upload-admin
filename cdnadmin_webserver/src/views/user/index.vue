<template>
  <div style="padding: 10px;">
    <!-- -------------------------------------------查询条件--------------------------------------------- -->
    <el-collapse v-model='activeCollapse'>
      <el-collapse-item title='查询&新建' name='search'>
        <el-form ref='searchCondition' :model='filter' label-width='80px' label-position=‘left’>
          <el-row>
            <el-col :span='4'>
              <el-form-item label='用户名' prop='username'>
                <el-select v-model="filter.username" filterable value="">
                  <el-option v-for="(item, index) in userData" :key="index" :label="item.username" :value="item.username"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='3' style="margin-left:20px;">
              <el-button @click='handleSearch' type="primary" icon="el-icon-search" size="small">查询</el-button>
              <el-button @click='resetForm("searchCondition")' type="primary" icon="el-icon-refresh" size="small">重置</el-button>
            </el-col>
            <el-col :span="2" style="margin-left: 20px">
              <el-button @click='dialogCreateUser=true' type='success' icon="el-icon-edit" size="small">新建用户</el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-collapse-item>
    </el-collapse>
    <!-- -------------------------------------------User List--------------------------------------------- -->
    <div>
      <div style='margin-bottom: 20px;'>
        <el-table :data='userData' border stripe
                  :default-sort="{prop: '', order: 'descending'}" max-height="800" style="width: 100%">
          <el-table-column label='序号' min-width='50' align="center">
            <template slot-scope="scope">
              <span>{{scope.$index+(pagination.current - 1) * pagination.pageSize + 1}}</span>
            </template>
          </el-table-column>
         <el-table-column label='用户名' prop='username' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='中文名' prop='name' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='管理员' prop='is_admin' sortable resizable min-width='100' width="150" align="center">
            <template slot-scope="scope">
              <el-switch v-model="scope.row.is_admin" :active-value=0 :inactive-value=1 active-color="#13ce66"
                         @change="resetAdmin(scope.row)"></el-switch>
            </template>
          </el-table-column>
          <el-table-column label='状态' prop='status' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='创建时间' prop='create_time' sortable resizable min-width='100' width="200" align="center"></el-table-column>
          <el-table-column label='最后编辑时间' prop='update_time' sortable resizable min-width='100' width="200" align="center"></el-table-column>
          <el-table-column label='操作' prop='operate' min-width="100" align="center">
            <template slot-scope="scope">
              <el-button @click="editUser(scope)" type="primary" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-button @click="delUser(scope)" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
              <el-button @click="resetUser(scope)" type="success" icon="el-icon-refresh" size="mini">重置密码</el-button>
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
    <!-- -------------------------------------------------新建User----------------------------------------------- -->
    <el-dialog title='新建用户' ref="createUser"  :visible.sync='dialogCreateUser' width="600px">
        <el-form :model='userForm' ref="addUserForm" label-width="150px">
          <el-form-item
            label='用户名'
            prop='username'
            required
            :rules='{ required: true, message: "用户名不能为空"}'>
            <el-input v-model='userForm.username' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='密码'
            prop='password'
            required
            :rules='{ required: true, message: "密码不能为空"}'>
            <el-input :type="pwdType" v-model='userForm.password' :maxlength='50' style="width: 300px"></el-input>
            <span class="show-pwd" @click="showPwd">{{showText}}</span>
          </el-form-item>
          <el-form-item
            label='中文名'
            prop='name'
            required
            :rules='{ required: true, message: "中文名不能为空"}'>
            <el-input v-model='userForm.name' :maxlength='50' style="width: 300px" ></el-input>
          </el-form-item>
          <el-form-item
            label='状态'
            prop='status'
            required
            :rules='{ required: true, message: "状态不能为空"}'>
            <el-select v-model="userForm.status" filterable placeholder="未使用" value="">
              <el-option v-for="(item, index) in userStatus" :key="index" :disabled="item.disabled"
                         :label="item.value" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseAddDialog'>取消</el-button>
        <el-button type='primary' @click='handleSave'>确定</el-button>
      </div>
    </el-dialog>

    <!-- -------------------------------------------------编辑User----------------------------------------------- -->
    <el-dialog title='编辑用户信息' ref="editUser"  :visible.sync='dialogEditUser' width="600px">
      <el-form :model='editUserForm' ref="editUserForm" label-width="150px">
        <el-form-item
          label='中文名'
          prop='name'
          required
          :rules='{ required: true, message: "中文名不能为空"}'>
          <el-input v-model='editUserForm.name' :maxlength='50' style="width: 300px" ></el-input>
        </el-form-item>
        <el-form-item
          label='状态'
          prop='status'
          required
          :rules='{ required: true, message: "状态不能为空"}'>
          <el-select v-model="editUserForm.status" filterable placeholder="未使用" value="">
            <el-option v-for="(item, index) in userStatus" :key="index" :disabled="item.disabled"
                       :label="item.value" :value="item.value"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseEditDialog'>取消</el-button>
        <el-button type='primary' @click='updateUserById'>确定</el-button>
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
    this.getUserData()
  }
}
</script>
<style type="text/css">
</style>
