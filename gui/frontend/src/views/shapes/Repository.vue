<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import EditComponent from '/src/components/crud/EditComponent.vue'
import { controls, repositoryLayout } from './controls'

const phase = ref('load')

const route = useRoute()
const identifier = route.params.identifier
const endpoint = `/api/shapes/@${identifier}`

const loaded = response => {
  for (let key of response.Result) {
    controls.KeyID.options[key.ID] = key.Description
      ? `${key.Label} (${key.Description})`
      : key.Label
  }
  phase.value = 'form'
}
</script>

<template>
  <div key="shapes-new">
    <FetchJSONComponent
      v-if="phase == 'load'"
      :endpoint="`/api/keys`"
      emit-fetched-event
      @fetched="loaded"
    ></FetchJSONComponent>
    <EditComponent
      v-else-if="phase == 'form'"
      :endpoint="endpoint"
      :controls="controls"
      :layout="repositoryLayout"
    ></EditComponent>
  </div>
</template>
