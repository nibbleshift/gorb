// Composables
import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/default/Default.vue'),
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import(/* webpackChunkName: "dashboard" */ '@/views/Dashboard.vue'),
      },
      {
        path: 'tests',
        name: 'Tests',
        component: () => import(/* webpackChunkName: "tests" */ '@/views/Tests.vue'),
      },
      {
        path: 'benchmarks',
        name: 'Benchmarks',
        component: () => import(/* webpackChunkName: "benchmarks" */ '@/views/Benchmarks.vue'),
      },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
