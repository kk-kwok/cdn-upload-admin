<template>
  <div style="padding: 10px;">
    <!-- -------------------------------------------查询条件--------------------------------------------- -->
    <el-collapse v-model='activeCollapse'>
      <el-collapse-item title='查询&新建' name='search'>
        <el-form ref='searchCondition' :model='filter' label-width='100px' label-position=‘left’>
          <el-row>
            <el-col :span='5'>
              <el-form-item label='所属项目' prop='project_name'>
                <el-select v-model="filter.project_name" filterable value="">
                  <el-option v-for="(item, index) in pjNameData" :key="index" :label="item.name" :value="item.name"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='5'>
              <el-form-item label='文件类型' prop='file_suffix'>
                <el-select v-model="filter.file_suffix" filterable value="">
                  <el-option v-for="(item, index) in cSuffixData" :key="index" :label="item.file_suffix" :value="item.file_suffix"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span='3' style="marginLeft: 10px;">
              <el-button @click='handleSearch' type="primary" icon="el-icon-search" size="small">查询</el-button>
              <el-button @click='resetForm("searchCondition")' type="primary" icon="el-icon-refresh" size="small">重置</el-button>
            </el-col>
            <el-col :span="3" style="margin-left: 10px">
              <el-button @click='dialogCreateCDNFile=true' type='success' icon="el-icon-upload" size="small">上传CDN文件</el-button>
            </el-col>
          </el-row>
        </el-form>
      </el-collapse-item>
    </el-collapse>
    <!-- -------------------------------------------CDNFile List--------------------------------------------- -->
    <div>
      <div style='margin-bottom: 10px'>
        <el-table :data='cdnFileData' border stripe
                  :default-sort="{prop: 'update_time', order: 'descending'}" max-height="640" style="width: auto">
          <el-table-column type="expand">
            <template slot-scope="props">
              <el-form label-position="left" inline class="cdnFile-expand">
                <el-form-item label="id">
                  <span>{{props.row.id}}</span>
                </el-form-item>
               <el-form-item label="所属项目">
                 <span>{{props.row.project_name}}</span>
               </el-form-item>
                <el-form-item label="CDN URL">
                  <span>{{props.row.domain}}/{{props.row.path}}/{{props.row.file_name}}</span>
                </el-form-item>
                <el-form-item label="文件大小">
                  <span>{{props.row.file_size}} Byte</span>
                </el-form-item>
                <el-form-item label="文件MD5值">
                  <span>{{props.row.file_md5}}</span>
                </el-form-item>
                <el-form-item label="备注">
                  <span>{{props.row.comment}}</span>
                </el-form-item>
                <el-form-item label="上传时间">
                  <span>{{props.row.create_time}}</span>
                </el-form-item>
                <el-form-item label="最后更新时间">
                  <span>{{props.row.update_time}}</span>
                </el-form-item>
              </el-form>
            </template>
          </el-table-column>
          <el-table-column label='序号' min-width='30' align="center">
            <template slot-scope="scope">
              <span>{{scope.$index+(pagination.current - 1) * pagination.pageSize + 1}}</span>
            </template>
          </el-table-column>
          <el-table-column label='CDN URL' prop='file_name' sortable resizable min-width='100' width="480" align="center">
            <template slot-scope="scope">
              <span>{{scope.row.domain}}/{{scope.row.path}}/{{scope.row.file_name}}</span>
            </template>
          </el-table-column>
          <el-table-column label='文件大小' prop='file_size' sortable resizable min-width='100' width="120" align="center">
            <template slot-scope="scope">
              <span v-if="calcFileSize(scope) === 1">
                {{Math.round(scope.row.file_size * 100 / 1024 / 1024 / 100)}} MB</span>
              <span v-else-if="calcFileSize(scope) === 2">
                {{Math.round(scope.row.file_size * 100 / 1024 / 100)}} KB</span>
              <span v-else>{{scope.row.file_size}} Byte</span>
            </template>
          </el-table-column>
          <el-table-column label='文件MD5值' prop='file_md5' sortable resizable min-width='100' width="270" align="center"></el-table-column>
          <el-table-column label='最后更新时间' prop='update_time' sortable min-width='100' width="160" align="center"></el-table-column>
          <el-table-column label='备注' prop='comment' sortable resizable min-width='100' width="100" align="center"></el-table-column>
          <el-table-column label='操作' prop='operate' min-width="100" align="center">
            <template slot-scope="scope">
              <el-button @click="editCDNFile(scope)" type="primary" icon="el-icon-edit" size="mini">编辑</el-button>
              <el-button @click="delCDNFile(scope)" type="danger" icon="el-icon-delete" size="mini">删除</el-button>
              <el-button @click="refreshCDNFile(scope)" type="success" icon="el-icon-refresh" size="mini">刷新CDN</el-button>
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
    <!-- -------------------------------------------------新建CDNFile----------------------------------------------- -->
    <el-dialog title='上传CDN文件' ref="createCDNFile"  :visible.sync='dialogCreateCDNFile' width="600px">
      <el-form :model='cdnFileForm' ref="addCDNFileForm" label-width="150px">
        <el-form-item
          label='所属项目'
          prop='project_name'
          required
          :rules='{required: true, message: "项目不能为空"}'>
          <el-select v-model="cdnFileForm.project_name" filterable value="" v-on:change="getURL($event)">
            <el-option v-for="(item, index) in pjNameData" :key="index" :label="item.name" :value="item.name"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item
          label='上传文件'>
          <el-upload
            ref="upload"
            :accept="this.file_suffix"
            :action="uploadUrl"
            :on-change="handleFile"
            :on-exceed="handleExceed"
            :file-list="fileList"
            :multiple="false"
            :limit="1"
            :auto-upload="false">
            <el-button slot="trigger" size="small" type="primary">选取文件</el-button>
            <div slot="tip" class="el-upload__tip">只能上传 {{file_suffix}} 文件,且文件大小不超过1GB</div>
          </el-upload>
        </el-form-item>
        <el-form-item
          label='新文件名'
          prop='file_name'
          required
          :rules='[{required: true, message: "文件名不能为空"}]'>
          <el-input v-model='cdnFileForm.file_name' :maxlength='50' style="width: 300px"
                    :placeholder="this.file_suffix" @blur="validFileSuffix(cdnFileForm.file_name)"></el-input>
          <div>
            <span v-if="this.file_suffix_math === 1" style="color:red">文件格式错误，格式必须为 {{file_suffix}}</span>
          </div>
        </el-form-item>
        <el-form-item
          label='URL'>
          <span>{{file_url}}/{{cdnFileForm.file_name}}</span>
        </el-form-item>
        <el-form-item
          label='备注'
          prop='comment'>
          <el-input v-model='cdnFileForm.comment' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
        <div class="el-pro" v-show="showProgressFlag">
          <el-progress  :text-inside="true" :stroke-width="20" :percentage="uploadPercent"></el-progress>
        </div>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseAddDialog'>取消</el-button>
        <el-button type='primary' @click='handleSave'>确定</el-button>
      </div>
    </el-dialog>
    <!-- -------------------------------------------------编辑CDNFile----------------------------------------------- -->
    <el-dialog title='编辑CDN文件' ref="editCDNFile"  :visible.sync='dialogEditCDNFile' width="600px">
      <el-form :model='editCDNFileForm' ref="editCDNFileForm" label-width="150px">
        <el-form-item
          label='新文件名'
          prop='file_name'>
          <el-input v-model='editCDNFileForm.file_name' :maxlength='50' style="width: 300px"
                    :placeholder="this.file_suffix" @blur="validFileSuffix(editCDNFileForm.file_name)"></el-input>
          <div>
            <span v-if="this.file_suffix_math === 1" style="color:red">文件格式错误，格式必须为 {{file_suffix}}</span>
          </div>
        </el-form-item>
        <el-form-item
          label='备注'
          prop='comment'>
          <el-input v-model='editCDNFileForm.comment' :maxlength='50' style="width: 300px"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click='handleCloseEditDialog'>取消</el-button>
        <el-button type='primary' @click='updateCDNFileById'>确定</el-button>
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
    this.getCDNFileData()
    this.getPjNameData()
    this.getCSuffixData()
  }
}
</script>
<style type="text/css">
  .cdnFile-expand {
    font-size: 0;
  }
  .cdnFile-expand label {
    width: 100px;
    color: #99a9bf;
  }
  .cdnFile-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
