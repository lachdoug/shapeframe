<script setup>
import { useRoute } from 'vue-router'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import ToolbarComponent from '/src/components/ToolbarComponent.vue'

const props = defineProps({
  toolbarButtons: {
    type: Object,
    required: true,
  },
  endpoint: {
    type: String,
    default: '',
  },
})

const route = useRoute()
const table = route.path.split('/')[1]
const fetchEndpoint =
  props.endpoint || `/api/${table}/@${route.params.identifier}`
</script>

<template>
  <ToolbarComponent :buttons="toolbarButtons"></ToolbarComponent>
  <FetchJSONComponent :endpoint="fetchEndpoint">
    <template #default="{ data }">
      <v-card class="my-2">
        <v-card-text>
          <pre>{{ JSON.stringify(data.Result, null, 2) }}</pre>
        </v-card-text>
      </v-card>
    </template>
  </FetchJSONComponent>
</template>
