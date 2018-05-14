import Vue from 'vue'
import Router from 'vue-router'
import storage from '@/utils/storage'
import { STORAGE_KEY } from '@/utils/constants'
import { types, default as store } from '@/store'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      meta: {
        requireAuth: true
      },
      redirect: '/register',
      component: resolve => require(['../components/common/Home.vue'], resolve),
      children: [
        {
          path: '/register',
          name: 'register',
          meta: {
            requireAuth: true
          },
          component: resolve =>
            require(['../components/page/Register.vue'], resolve)
        },
        {
          path: '/search',
          name: 'search',
          meta: {
            requireAuth: true
          },
          component: resolve =>
            require(['../components/page/Search.vue'], resolve)
        },
        {
          path: '/compare',
          name: 'compare',
          meta: {
            requireAuth: true
          },
          component: resolve =>
            require(['../components/page/Compare.vue'], resolve)
        },
        {
          path: '/submit',
          name: 'submit',
          meta: {
            requireAuth: true
          },
          component: resolve =>
            require(['../components/page/Submit.vue'], resolve)
        },
        {
          path: '/monitor',
          name: 'monitor',
          meta: {
            requireAuth: true
          },
          component: resolve =>
            require(['../components/page/Monitor.vue'], resolve)
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: resolve => require(['../components/page/Login.vue'], resolve)
    }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.name === 'login') {
    if (storage.get(STORAGE_KEY.AUTHED)) {
      next({
        name: 'home'
      })
      return
    }
  }
  if (to.meta.requireAuth) {
    if (!storage.get(STORAGE_KEY.AUTHED)) {
      Vue.prototype.$error('请先登录')
      store.commit(types.CHANGE_MODAL_STATUS, { mode: 'login', visible: true })
      next({
        name: 'login'
      })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
