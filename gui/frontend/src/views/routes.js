// Use dynamic import in routes to facilitate build bundle chunking.
const Graph = () => import('/src/views/Graph.vue')
const PageNotFound = () => import('/src/components/PageNotFoundComponent.vue')
import shapesRoutes from './shapes/routes'
import framesRoutes from './frames/routes'
import keysRoutes from './keys/routes'
import repositoriesRoutes from './repositories/routes'
import strategiesRoutes from './strategies/routes'
import settingsRoutes from './settings/routes'

export default [
  {
    path: '/',
    component: Graph,
  },
  shapesRoutes,
  framesRoutes,
  keysRoutes,
  repositoriesRoutes,
  strategiesRoutes,
  settingsRoutes,
  {
    path: '/:noPage(.*)*',
    component: PageNotFound,
  },
]
