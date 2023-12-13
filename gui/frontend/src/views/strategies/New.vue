<script setup>
import NewComponent from '/src/components/crud/NewComponent.vue'
import { controls, layout } from './controls.js'
</script>

<template>
  <div key="strategies-edit">
    <NewComponent :controls="controls" :layout="layout"></NewComponent>
  </div>
</template>
