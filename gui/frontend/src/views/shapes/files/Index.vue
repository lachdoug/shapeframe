<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import ToolbarComponent from '/src/components/ToolbarComponent.vue'
import { mdiPlus } from '@mdi/js'

const phase = ref('load')
let items

const router = useRouter()
const route = useRoute()

const shapeIdentifier = route.params.identifier

const toolbarButtons = {
  left: [{ label: 'New', icon: mdiPlus, path: 'new' }],
}

const loaded = response => {
  items = response.Result
  phase.value = 'table'
}

const navigateToItem = e => {
  const fileIdentifier = e.currentTarget.getAttribute('fileIdentifier')
  router.push(`${route.path}/@${fileIdentifier}`)
}
</script>

<template>
  <div key="shapes-files-index">
    <ToolbarComponent :buttons="toolbarButtons"></ToolbarComponent>
    <div v-if="phase == 'load'" class="mt-2">
      <FetchJSONComponent
        :url="`/api/shapes/@${shapeIdentifier}/files`"
        emit-fetched-event
        @fetched="loaded"
      >
      </FetchJSONComponent>
    </div>
    <div v-else-if="phase == 'table'">
      <v-card class="my-2">
        <v-card-text v-if="!items.length">
          <i> No files </i>
        </v-card-text>
        <v-table v-else>
          <tbody>
            <tr
              v-for="(item, i) in items"
              :key="i"
              v-ripple
              :fileIdentifier="item"
              @click="navigateToItem"
            >
              <slot :item="item">
                <td>{{ item }}</td>
              </slot>
            </tr>
          </tbody>
        </v-table>
      </v-card>
    </div>
  </div>
</template>

<style scoped>
table tr:hover {
  background-color: rgba(var(--v-theme-on-background), var(--v-hover-opacity));
}

table tr {
  cursor: pointer;
}
</style>
