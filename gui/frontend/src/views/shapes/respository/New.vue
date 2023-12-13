<script setup>
import { ref } from 'vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import NewComponent from '/src/components/crud/NewComponent.vue'
import { controls, layout } from './controls'

let keys = {}
const phase = ref('load')

const loaded = response => {
  for (let key of response.Result) {
    keys[key.ID] = key.Description
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
      :url="`/api/keys`"
      emit-fetched-event
      @fetched="loaded"
    ></FetchJSONComponent>
    <NewComponent
      v-else-if="phase == 'form'"
      :controls="controls(keys)"
      :layout="layout"
    ></NewComponent>
  </div>
</template>
