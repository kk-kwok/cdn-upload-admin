<template>
  <div style="padding: 10px;">
    <!-- -------------------------------------------查询条件--------------------------------------------- -->
    <el-collapse v-model='activeCollapse'>
      <el-collapse-item title='查询&新建' name='search'>
        <el-form ref='searchCondition' :model='filter' label-width='100px' label-position=‘left’>
          <el-row>
            <el-col :span='5'>
              <el-form-item label='CDN平台名称' prop='name'>
                <el-select v-model="filter.name" filterable value="">
                  <el-option v-for="(item, index) in platformData" :key="index" :label="item.name" :value="item.name"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='3'>
              <el-form-item label='使用状态' prop='status'>
                <el-select v-model="filter.status" filterable value="">
                  <el-option v-for="(item, index) in platformStatus" :key="index" :label="item.value" :value="item.value"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='3' style="margin-left:20px;">
              <el-button @click='handleSearch' type="primary" icon="el-icon-search" size="small">查询</el-button>
              <el-button @click='resetForm("searchCondition")' type="primary" icon="el-icon-refresh" size="small">重置</el-button>
            </el-col>
            <el-col :span="3" style="margin-left: 20px">
              <el-button @click='dialogCreatePlatform=true' type='success' icon="el-icon-edit" size="small">新建CDN平台</el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-collapse-item>
    </el-collapse>
    <!-- -------------------------------------------Platform List--------------------------------------------- -->
    <div>
      <div style='margin-bottom: 10px'>
        <el-table :data='platformData' border stripe
                  :default-sort="{prop: '', order: 'descending'}" max-height="640" style="width: auto">
          <el-table-column label='序号' min-width='50' align="center">
            <template slot-scope="scope">
              <span>{{scope.$index+(pagination.current - 1) * pagination.pageSize + 1}}</span>
            </template>
          </el-table-column>
          <el-table-column label='CDN平台名称' prop='name' sortable resizable min-width='100' width="150" align="center"></el-table-column>
          <el-table-column label='CDN加密ID' prop='secret_id' sortable resizable min-width='100' width="350" align="center"></el-table-column>
          <el-table-column label='CDN加密KEY' prop='secret_key' sortable resizable min-width='100' width="300" align="center"></el-table-column>
          <el-table-column label='API URL' prop='api_url' sortable resizable min-width='100' width="300" align="center"></el-table-column>
          <el-table-column label='API调用方法' prop='action' sortable resizable min-width='100' width="160" align="center"></el-table-column>
          <el-table-column label='状态' prop='status' sortable min-width='100' width="100" align="center"></el-table-column>
          <el-table-column label='操作' prop='operate' min-width="200" align="center">
            <template slot-scope="scope">
              <el-button @click="editPlatform(scope)" type="primary" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-button @click="delPlatform(scope)" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
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
    <!-- -------------------------------------------------新建Platform----------------------------------------------- -->
    <el-dialog title='新建CDN平台' ref="createPlatform"  :visible.sync='dialogCreatePlatform' width="600px">
        <el-form :model='platformForm' ref="addPlatformForm" label-width="150px">
          <el-form-item
            label='CDN名称'
            prop='name'
            required
            :rules='{ required: true, message: "名称不能为空！"}'>
            <el-input v-model='platformForm.name' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='加密ID'
            prop='secret_id'>
            <el-input v-model='platformForm.secret_id' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='加密KEY'
            prop='secret_id'>
            <el-input v-model='platformForm.secret_key' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='API URL'
            prop='api_url'>
            <el-input v-model='platformForm.api_url' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='API调用方法'
            prop='action'>
            <el-input v-model='platformForm.action' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
          <el-form-item
            label='状态'
            prop='status'>
            <el-select v-model="platformForm.status" filterable placeholder="未使用" value="">
              <el-option v-for="(item, index) in platformStatus" :key="index" :disabled="item.disabled"
                         :label="item.value" :value="item.value"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseAddDialog'>取消</el-button>
        <el-button type='primary' @click='handleSave'>确定</el-button>
      </div>
    </el-dialog>

    <!-- -------------------------------------------------编辑Platform----------------------------------------------- -->
    <el-dialog title='编辑CDN' ref="editPlatform"  :visible.sync='dialogEditPlatform' width="600px">
      <el-form :model='editPlatformForm' ref="editPlatformForm" label-width="150px">
        <el-form-item
          label='CDN名称'
          prop='name'
          required
          :rules='{ required: true, message: "名称不能为空！"}'>
          <el-input v-model='editPlatformForm.name' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='加密ID'
          prop='secret_id'>
          <el-input v-model='editPlatformForm.secret_id' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='加密KEY'
          prop='secret_id'>
          <el-input v-model='editPlatformForm.secret_key' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='API URL'
          prop='api_url'>
          <el-input v-model='editPlatformForm.api_url' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='API调用方法'
          prop='action'>
          <el-input v-model='editPlatformForm.action' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <el-form-item
          label='状态'
          prop='status'>
          <el-select v-model="editPlatformForm.status" filterable  value="">
            <el-option v-for="(item, index) in platformStatus" :key="index" :disabled="item.disabled"
                       :label="item.value" :value="item.value"></el-option>
          </el-select>
          <!--<el-input v-model='editForm.status' :maxlength='50' style="width: 200px"></el-input>-->
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseEditDialog'>取消</el-button>
        <el-button type='primary' @click='updatePlatformById'>确定</el-button>
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
    this.getPlatformData()
  }
}
</script>
<style type="text/css">
</style>
