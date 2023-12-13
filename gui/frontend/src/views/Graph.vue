<script setup>
import { ref } from 'vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import TopologyGraphComponent from '/src/components/TopologyGraphComponent.vue'

const phase = ref('load')
let data

const loaded = response => {
  data = response.Result
  phase.value = 'chart'
}
</script>
<template>
  <div key="graph">
    <FetchJSONComponent v-if="phase == 'load'" endpoint="/api/topology">
      <!-- emit-fetched-event
      @fetched="loaded" -->
    </FetchJSONComponent>
    <TopologyGraphComponent
      v-else-if="phase == 'chart'"
      :data="data"
    ></TopologyGraphComponent>
  </div>
</template>
