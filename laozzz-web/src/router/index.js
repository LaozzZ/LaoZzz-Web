import { createRouter, createWebHistory } from 'vue-router'

const routes  = [
  {
    path: '/',
    component: () => import('@/layouts/DefaultLayout.vue'),
    children: [
      {path: '', name: 'home', component: () => import('@/views/Index/Index.vue')},
      {path: 'ApexMap', name: 'ApexMap', component: () => import('@/views/ApexMap/ApexMap.vue')},
    ]
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

export default router
