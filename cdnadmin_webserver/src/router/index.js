import Vue from 'vue'
import Router from 'vue-router'
import layout from '@/views/layout/index'

// 懒加载方式，当路由被访问的时候才加载对应组件
const Login = resolve => require(['@/views/login/index'], resolve)

Vue.use(Router)

export const constantRouterMap = [
  {
    path: '/',
    redirect: '/cdn',
    meta: { roles: '' },
    hidden: true
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: { roles: '' },
    hidden: true
  },
  {
    path: '/cdn',
    component: layout,
    redirect: '/cdn',
    iconCls: 'el-icon-star-on',
    leaf: true,
    meta: { roles: '' },
    children: [
      {
        path: '/cdn',
        name: 'CDN上传',
        component: () => import('@/views/cdnfile/index'),
        meta: { requireAuth: true }
      }
    ]
  },
  {
    path: '/updatePwd',
    component: layout,
    redirect: '/updatePwd',
    iconCls: 'el-icon-setting',
    leaf: true,
    meta: { roles: '' },
    children: [
      {
        name: '修改密码',
        path: '/updatePwd',
        component: () => import('@/views/user/password'),
        meta: { requireAuth: true }
      }
    ]
  },
  {
    path: '/admin',
    name: '系统管理',
    component: layout,
    iconCls: 'el-icon-setting',
    meta: { roles: 'admin' },
    children: [
      {
        path: '/project',
        name: '业务管理',
        component: () => import('@/views/project/index'),
        meta: { requireAuth: true, roles: 'admin' }
      },
      {
        path: '/platform',
        name: 'CDN平台',
        component: () => import('@/views/platform/index'),
        meta: { requireAuth: true, roles: 'admin' }
      },
      {
        path: '/suffix',
        name: '文件后缀',
        component: () => import('@/views/suffix/index'),
        meta: { requireAuth: true, roles: 'admin' }
      },
      {
        path: '/user',
        name: '用户管理',
        component: () => import('@/views/user/index'),
        meta: { requireAuth: true, roles: 'admin' }
      }
    ]
  },
  { path: '*', redirect: '/', hidden: true, meta: { roles: '' } }
]

export default new Router({
  mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
