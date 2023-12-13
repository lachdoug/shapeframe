import AppRouterView from '/src/components/AppRouterView.vue'
const Index = () => import('/src/views/shapes/files/Index.vue')
const New = () => import('/src/views/shapes/files/New.vue')
const Show = () => import('/src/views/shapes/files/Show.vue')
const Edit = () => import('/src/views/shapes/files/Edit.vue')
const Delete = () => import('/src/views/shapes/files/Delete.vue')

export default {
  path: 'files',
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
      path: '@:fileIdentifier',
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
      ],
    },
  ],
}
