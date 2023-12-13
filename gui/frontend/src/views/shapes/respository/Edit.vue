<script setup>
import { useRoute } from 'vue-router';
import EditComponent from '/src/components/crud/EditComponent.vue'
import { controls, layout } from './controls.js'

const route = useRoute()
const identifier = route.params.identifier
const endpoint = `/api/shapes/@${identifier}`
</script>

<template>
  <div key="keys-edit">
    <EditComponent :endpoint="endpoint" :controls="controls" :layout="layout"></EditComponent>
  </div>
</template>
