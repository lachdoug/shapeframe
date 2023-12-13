import AppRouterView from '/src/components/AppRouterView.vue'
const Index = () => import('./Index.vue')
const New = () => import('./New.vue')
const Show = () => import('./Show.vue')
const Edit = () => import('./Edit.vue')
const Delete = () => import('./Delete.vue')
const Icon = () => import('./Icon.vue')
import filesRoutes from './files/routes'

export default {
  path: '/shapes',
  component: AppRouterView,
  children: [
    {
      path: '',
      component: Index,
    },
    {
      path: 'new',
      component: New,
    },
    {
      path: '@:identifier',
      component: AppRouterView,
      children: [
        {
          path: '',
          component: Show,
        },
        {
          path: 'delete',
          component: Delete,
        },

        {
          path: 'edit',
          component: Edit,
        },
        {
          path: 'icon',
          component: Icon,
        },
        filesRoutes,
      ],
    },
  ],
}