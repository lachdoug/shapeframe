<script setup>
import { ref } from 'vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import NewComponent from '/src/components/crud/NewComponent.vue'
import { controls, layout } from './controls.js'

const phase = ref('load')

const loaded = response => {
  for (let repository of response.Result) {
    controls.RepositoryID.options[repository.ID] = repository.Description
      ? `${repository.Label} (${repository.Description})`
      : repository.Label
  }
  phase.value = 'form'
}
</script>

<template>
  <div key="shapes-new">
    <FetchJSONComponent
      v-if="phase == 'load'"
      :endpoint="`/api/repositories`"
      emit-fetched-event
      @fetched="loaded"
    ></FetchJSONComponent>
    <NewComponent
      v-else-if="phase == 'form'"
      :controls="controls"
      :layout="layout"
    ></NewComponent>
  </div>
</template>
