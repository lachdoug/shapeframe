import AppRouterView from '/src/components/AppRouterView.vue'
const Index = () => import('./Index.vue')
const New = () => import('./New.vue')
const Show = () => import('./Show.vue')
const Edit = () => import('./Edit.vue')
const Clone = () => import('./Clone.vue')
const Pull = () => import('./Pull.vue')
const Push = () => import('./Push.vue')
const Remove = () => import('./Remove.vue')
const Delete = () => import('./Delete.vue')

export default {
  path: '/repositories',
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
          path: 'clone',
          component: Clone,
        },
        {
          path: 'pull',
          component: Pull,
        },
        {
          path: 'push',
          component: Push,
        },
        {
          path: 'remove',
          component: Remove,
        },
        {
          path: 'delete',
          component: Delete,
        },
      ],
    },
  ],
}
