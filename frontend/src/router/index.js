import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  },
  {
    path: '/login',
    name: 'loginUser',
    component: () => import('../views/Auth/Login.vue')
  },
  {
    path: '/signup',
    name: 'signupUser',
    component: () => import('../views/Auth/SignUp.vue')
  },
  {
    path: '/recovery',
    name: 'recoveryUser',
    component: () => import('../views/Auth/Recovery.vue')
  },
  {
    path: '/services',
    name: 'services',
    component: () => import('../views/Services.vue')
  },
  {
    path: '/calendar',
    name: 'calendar',
    component: () => import('../views/Calendar.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
