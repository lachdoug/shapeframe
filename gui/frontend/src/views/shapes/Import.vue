<script setup>
import { ref } from 'vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import NewComponent from '/src/components/crud/NewComponent.vue'
import { controls, repositoryLayout } from './controls'

const phase = ref('load')

const endpoint = `/api/shapes/import`

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
    <NewComponent
      v-else-if="phase == 'form'"
      :endpoint="endpoint"
      :controls="controls"
      :layout="repositoryLayout"
    ></NewComponent>
  </div>
</template>
