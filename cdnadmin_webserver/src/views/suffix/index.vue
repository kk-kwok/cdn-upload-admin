<template>
  <div style="padding: 10px;">
    <div style="margin: 10px;">
    <el-row>
    <el-col :span="3" style="margin-left: 20px">
      <el-button @click='dialogCreateSuffix=true' type='primary' icon="el-icon-edit" size="small">新建文件后缀</el-button>
    </el-col>
    </el-row>
    </div>
    <!-- -------------------------------------------Suffix List--------------------------------------------- -->
    <div>
      <div style='margin-bottom: 20px;'>
        <el-table :data='suffixData' border stripe
                  :default-sort="{prop: '', order: 'descending'}" max-height="800" style="width: 800px">
          <el-table-column label='序号' min-width='50' align="center">
            <template slot-scope="scope">
              <span>{{scope.$index+(pagination.current - 1) * pagination.pageSize + 1}}</span>
            </template>
          </el-table-column>
         <el-table-column label='文件后缀名' prop='file_suffix' sortable resizable min-width='100' width="300" align="center"></el-table-column>
          <el-table-column label='操作' prop='operate' min-width="100" align="center">
            <template slot-scope="scope">
              <el-button @click="editSuffix(scope)" type="primary" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-button @click="delSuffix(scope)" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
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
    <!-- -------------------------------------------------新建Suffix----------------------------------------------- -->
    <el-dialog title='新建文件后缀' ref="createSuffix"  :visible.sync='dialogCreateSuffix' width="600px">
        <el-form :model='suffixForm' ref="addSuffixForm" label-width="150px">
          <el-form-item
            label='文件后缀名'
            prop='file_suffix'
            required
            :rules='{ required: true, message: "文件后缀名不能为空"}'>
            <el-input v-model='suffixForm.file_suffix' :maxlength='50' style="width: 300px"></el-input>
          </el-form-item>
        </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseAddDialog'>取消</el-button>
        <el-button type='primary' @click='handleSave'>确定</el-button>
      </div>
    </el-dialog>

    <!-- -------------------------------------------------编辑Suffix----------------------------------------------- -->
    <el-dialog title='编辑文件后缀' ref="editSuffix"  :visible.sync='dialogEditSuffix' width="600px">
      <el-form :model='editSuffixForm' ref="editSuffixForm" label-width="150px">
        <el-form-item
          label='文件后缀名'
          prop='file_suffix'
          required
          :rules='{ required: true, message: "文件后缀名不能为空"}'>
          <el-input v-model='editSuffixForm.file_suffix' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseEditDialog'>取消</el-button>
        <el-button type='primary' @click='updateSuffixById'>确定</el-button>
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
    this.getSuffixData()
  }
}
</script>
<style type="text/css">
</style>
