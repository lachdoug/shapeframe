<script setup>
import { useRoute } from 'vue-router'
import NewComponent from '/src/components/crud/NewComponent.vue'
import { tokenControls, tokenLayout } from './controls.js'

const route = useRoute()
const itemPath = `/keys/@${route.params.identifier}`
const saveUrl = `/api${itemPath}`
</script>

<template>
  <div key="keys-token">
    <NewComponent
      :controls="tokenControls"
      :layout="tokenLayout"
      :save-url="saveUrl"
      save-method="PUT"
      :return-path="itemPath"
    ></NewComponent>
  </div>
</template>
