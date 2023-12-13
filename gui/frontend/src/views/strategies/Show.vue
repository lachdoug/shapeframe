<script setup>
import ShowComponent from '/src/components/crud/ShowComponent.vue'
import { mdiFileEditOutline, mdiTrashCanOutline } from '@mdi/js'

const toolbarButtons = {
  left: [{ label: 'Edit', icon: mdiFileEditOutline, path: 'edit' }],
  right: [{ label: 'Delete', icon: mdiTrashCanOutline, path: 'delete' }],
}
</script>

<template>
  <div key="strategies-edit">
    <ShowComponent :toolbar-buttons="toolbarButtons"></ShowComponent>
  </div>
</template>
