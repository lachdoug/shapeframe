<script setup>
import { ref } from 'vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import EditComponent from '/src/components/crud/EditComponent.vue'
import { controls, editLayout } from './controls.js'

const phase = ref('load')

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
  <div key="keys-edit">
    <FetchJSONComponent
      v-if="phase == 'load'"
      :endpoint="`/api/keys`"
      emit-fetched-event
      @fetched="loaded"
    ></FetchJSONComponent>
    <EditComponent
      v-else-if="phase == 'form'"
      :controls="controls"
      :layout="editLayout"
    ></EditComponent>
  </div>
</template>
