import AppRouterView from '/src/components/AppRouterView.vue'
const Index = () => import('./Index.vue')
const New = () => import('./New.vue')
const Show = () => import('./Show.vue')
const Edit = () => import('./Edit.vue')
const Delete = () => import('./Delete.vue')

export default {
  path: '/strategies',
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
          path: 'edit',
          component: Edit,
        },
        {
          path: 'delete',
          component: Delete,
        },
      ],
    },
  ],
}
