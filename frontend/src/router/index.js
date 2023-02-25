import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'calendar',
    component: () => import('../views/Calendar.vue')
  },
  {
    path: '/admin',
    name: 'admin',
    component: () => import('../views/AdminCalendar.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
