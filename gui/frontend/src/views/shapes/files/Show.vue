<script setup>
import { useRoute } from 'vue-router'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import ToolbarComponent from '/src/components/ToolbarComponent.vue'
import { mdiFileEditOutline, mdiTrashCanOutline } from '@mdi/js'

const route = useRoute()
const identifier = route.params.identifier
const fileIdentifier = route.params.fileIdentifier

const toolbarButtons = {
  left: [{ label: 'Edit', icon: mdiFileEditOutline, path: 'edit' }],
  right: [{ label: 'Delete', icon: mdiTrashCanOutline, path: 'delete' }],
}
</script>

<template>
  <div key="shapes-files-show">
    <ToolbarComponent :buttons="toolbarButtons"></ToolbarComponent>
    <FetchJSONComponent
      :url="`/api/shapes/@${identifier}/files/@${fileIdentifier}`"
    >
      <template #default="{ data }">
        <v-card class="my-2">
          <v-card-text>
            <i v-if="!data.Result.content"> No content</i>
            <pre v-else>{{ data.Result.content }}</pre>
          </v-card-text>
        </v-card>
      </template>
    </FetchJSONComponent>
  </div>
</template>
