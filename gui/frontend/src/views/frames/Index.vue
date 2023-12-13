<script setup>
import IndexComponent from '/src/components/crud/IndexComponent.vue'
import { mdiPlus } from '@mdi/js'

const toolbarButtons = {
  left: [{ label: 'New', icon: mdiPlus, path: 'new' }],
}
</script>

<template>
  <div key="frames-index">
    <IndexComponent :toolbar-buttons="toolbarButtons"></IndexComponent>
  </div>
</template>
