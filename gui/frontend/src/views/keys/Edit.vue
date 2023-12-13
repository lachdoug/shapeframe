<script setup>
import EditComponent from '/src/components/crud/EditComponent.vue'
// import { controls, layout } from './controls.js'
</script>

<template>
  <div key="shapes-edit">
    <EditComponent :controls="controls" :layout="layout"></EditComponent>
  </div>
</template>
