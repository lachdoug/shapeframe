import { createRouter, createWebHistory } from 'vue-router'
import routes from './views/routes'

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

export default router
