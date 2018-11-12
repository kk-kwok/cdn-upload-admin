<template>
  <el-row class="container">
    <el-col :span="24" class="header">
      <el-col :span="20" class="logo">
        <span style="margin-left: 30px">CDN UPLOAD ADMIN</span>
      </el-col>
      <el-col :span="4" class="userinfo">
        <el-dropdown trigger="click">
          <span class="el-dropdown-link userinfo-inner"><img src="@/assets/user.jpg" />Welcome {{sysUserName}}</span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item divided @click.native="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </el-col>
    </el-col>
    <el-col :span="24" class="main">
      <aside>
        <el-menu :default-active="$router.path" class="el-menu-vertical-demo"
                 @open="handleOpen" @close="handleClose" @select="handleSelect" unique-opened router>
          <template v-for="(item,index) in $router.options.routes"
                    v-if="!item.hidden && handlePermission(item.meta.roles)">
            <el-submenu :index="index+''" :key="item.leaf" v-if="!item.leaf">
              <template slot="title"><i :class="item.iconCls"></i>{{item.name}}</template>
              <el-menu-item v-for="child in item.children" :key="child.path"
                            :index="child.path" v-if="!child.hidden">{{child.name}}</el-menu-item>
            </el-submenu>
            <el-menu-item v-if="item.leaf&&item.children.length>0"
                          :key="item.children[0].path" :index="item.children[0].path">
              <i :class="item.iconCls"></i>{{item.children[0].name}}</el-menu-item>
          </template>
          <template v-else-if="!item.hidden && isAdmin()">
            <el-submenu :index="index+''" :key="item.leaf" v-if="!item.leaf">
              <template slot="title"><i :class="item.iconCls"></i>{{item.name}}</template>
              <el-menu-item v-for="child in item.children" :key="child.path"
                            :index="child.path" v-if="!child.hidden">{{child.name}}</el-menu-item>
            </el-submenu>
            <el-menu-item v-if="item.leaf&&item.children.length>0"
                          :key="item.children[0].path" :index="item.children[0].path">
              <i :class="item.iconCls"></i>{{item.children[0].name}}</el-menu-item>
          </template>
        </el-menu>
      </aside>
      <section class="content-container">
        <div class="grid-content bg-purple-light">
          <el-col :span="24" class="breadcrumb-container">
            <strong class="title">{{$route.name}}</strong>
            <el-breadcrumb separator="/" class="breadcrumb-inner">
              <el-breadcrumb-item v-for="item in $route.matched" :key="item.name">
                {{ item.name }}
              </el-breadcrumb-item>
            </el-breadcrumb>
          </el-col>
          <el-col :span="24" class="content-wrapper">
            <transition>
              <router-view></router-view>
            </transition>
          </el-col>
        </div>
      </section>
    </el-col>
  </el-row>
</template>

<script>
import store from '@/store/store.js'

export default{
  data () {
    return {
      isCollapse: false,
      sysUserName: store.state.userInfo.username ? store.state.userInfo.username : window.sessionStorage.getItem('username'),
      sysUserAvatar: '',
      roles: store.state.userInfo.roles ? store.state.userInfo.roles : window.sessionStorage.getItem('roles')
    }
  },
  methods: {
    isAdmin () {
      return this.roles === 'admin'
    },
    handlePermission (data) {
      return this.roles === data
    },
    onSubmit () {
      console.log('submit!')
    },
    handleOpen (key, keyPath) {
      console.log(key, keyPath)
    },
    handleClose (key, keyPath) {
      console.log(key, keyPath)
    },
    handleSelect: function (a, b) {
    },
    // 退出登录
    logout: function () {
      const _this = this
      this.$confirm('确认退出吗?', '提示', {
        // type: 'warning'
      }).then(() => {
        this.$store.commit('logout')
        _this.$router.push('/login')
      }).catch(() => {

      })
    }
  }
}
</script>
<style scoped lang="scss">
  .container {
    position: absolute;
    top: 0;
    bottom: 0;
    width: 100%;
    .header {
      height: 60px;
      line-height: 60px;
      background: #1F2D3D;
      color: #ffffff;
      .userinfo {
        text-align: right;
        padding-right: 35px;
        .userinfo-inner {
          color: #c0ccda;
          cursor: pointer;
          img {
            width: 40px;
            height: 40px;
            border-radius: 20px;
            margin: 10px 0 10px 10px;
            float: right;
          }
        }
      }
      .logo {
        font-size: 22px;
        img {
          width: 40px;
          float: left;
          margin: 10px 10px 10px 18px;
        }
        .txt {
          color: #20a0ff
        }
      }
    }
    .main {
      background: #324057;
      position: absolute;
      top: 60px;
      bottom: 0;
      overflow: hidden;
      aside {
        width: 230px;
      }
      .content-container {
        background: #f1f2f7;
        position: absolute;
        right: 0;
        top: 0;
        bottom: 0;
        left: 230px;
        overflow-y: scroll;
        padding: 20px;
        .breadcrumb-container {
          margin-bottom: 15px;
          .title {
            width: 200px;
            float: left;
            color: #475669;
          }
          .breadcrumb-inner {
            float: right;
          }
        }
        .content-wrapper {
          background-color: #fff;
          box-sizing: border-box;
        }
      }
    }

  }
</style>
