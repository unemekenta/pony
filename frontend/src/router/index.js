import Vue from 'vue'
import Router from 'vue-router'
import TopList from '@/components/TopList'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import MyPage from '@/components/MyPage'
import Edit from '@/components/Edit'
import firebase from 'firebase'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '/',
      name: 'TopList',
      component: TopList,
      meta: { requiresAuth: true }
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    },
    {
      path: '/mypage',
      name: 'MyPage',
      component: MyPage,
      meta: { requiresAuth: true }
    },
    {
      path: '/edit',
      name: 'Edit',
      component: Edit,
      meta: { requiresAuth: true }
    }
  ]
})
router.beforeEach((to, from, next) => {
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
  if (requiresAuth) {
    // サインインの有無判断
    firebase.auth().onAuthStateChanged(function (user) {
      if (user) {
        next()
      } else {
        next({
          path: '/signin'
        })
      }
    })
  } else {
    next()
  }
})
export default router
