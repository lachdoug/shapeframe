<script setup>
import NewComponent from '/src/components/crud/NewComponent.vue'
import { newControls, newLayout } from './controls.js'
</script>

<template>
  <div key="keys-edit">
    <NewComponent :controls="newControls" :layout="newLayout"></NewComponent>
  </div>
</template>
