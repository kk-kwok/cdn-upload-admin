<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
  body {
    margin: 0;
    padding: 0;
    background: #1F2D3D;
    font-family: Helvetica Neue, Helvetica, PingFang SC, Hiragino Sans GB, Microsoft YaHei, SimSun, sans-serif;
    font-size: 14px;
    -webkit-font-smoothing: antialiased;
  }
  #app {
    position: absolute;
    top: 0;
    bottom: 0;
    width: 100%;
    min-width: 1200px;
  }
  .el-submenu__title {
    background-color: #324057;
    color: white;
    width: 230px;
  }
  .el-submenu [class^=el-icon] {
    vertical-align: baseline;
    margin-right: 10px;
  }
  .el-menu-item {
    background-color: #324057;
    color: white;
    width: 230px;
  }
  .el-menu-item [class^=el-icon] {
    // background-color: #324057;
    vertical-align: baseline;
    margin-right: 10px;
  }
  .toolbar {
    background: #2a2a2a;
    padding: 10px;
    .el-form-item {
      margin-bottom: 10px;
    }
  }
</style>
