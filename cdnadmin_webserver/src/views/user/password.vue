<template>
  <div style="padding: 10px;">
    <!-- -------------------------------------------User List--------------------------------------------- -->
    <div>
      <div style='margin-bottom: 20px;'>
        <el-table :data='currentUserData' border stripe
                  :default-sort="{prop: '', order: 'descending'}" max-height="800" style="width: 100%">
          <el-table-column label='序号' min-width='50' align="center">
            <template slot-scope="scope">
              <span>{{scope.$index+(pagination.current - 1) * pagination.pageSize + 1}}</span>
            </template>
          </el-table-column>
          <el-table-column label='用户名' prop='username' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='中文名' prop='name' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='状态' prop='status' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='创建时间' prop='create_time' sortable resizable min-width='100' width="200" align="center"></el-table-column>
          <el-table-column label='最后编辑时间' prop='update_time' sortable resizable min-width='100' width="200" align="center"></el-table-column>
          <el-table-column label='操作' prop='operate' min-width="100" align="center">
            <template slot-scope="scope">
              <el-button @click="editUserPwd(scope)" type="primary" icon="el-icon-refresh" size="mini">修改密码</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <!-- -------------------------------------------------修改用户密码----------------------------------------------- -->
    <el-dialog title='重置用户密码' ref="updatePwd"  :visible.sync='dialogUpdatePwd' width="600px">
      <el-form :model='updatePwdForm' ref="updatePForm" label-width="150px">
        <el-form-item
          label='当前密码'
          prop='old_password'
          required
          :rules='{ required: true, message: "密码不能为空"}'>
          <el-input type="password" v-model='updatePwdForm.old_password' :maxlength='50' style="width: 300px" ></el-input>
        </el-form-item>
        <el-form-item
          label='新密码 第一次输入'
          prop='password'
          required
          :rules='{ required: true, message: "密码不能为空"}'>
          <el-input type="password" v-model='updatePwdForm.password' :maxlength='50' style="width: 300px"
          @blur="validSamePwd()" ></el-input>
          <div>
            <span v-if="this.isSamePwd === 1" style="color:red">新旧密码不能一样</span>
            <span v-else-if="this.isSamePwd === 2" style="color:red">新密码长度不能小于6位</span>
          </div>
        </el-form-item>
        <el-form-item
          label='新密码 第二次输入'
          prop='password2'
          required
          :rules='{ required: true, message: "密码不能为空"}'>
          <el-input type="password" v-model='updatePwdForm.password2' :maxlength='50' style="width: 300px"
          @blur="validSamePwd2()"></el-input>
          <div>
            <span v-if="this.isSamePwd2 === 1" style="color:red">新密码两次输入不一致</span>
            <span v-else-if="this.isSamePwd2 === 2" style="color:red">新密码长度不能小于6位</span>
          </div>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseUpdateDialog'>取消</el-button>
        <el-button type='primary' @click='updateUserPwdById'>确定</el-button>
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
    this.getCurrentUserData()
  }
}
</script>
<style type="text/css">
</style>
