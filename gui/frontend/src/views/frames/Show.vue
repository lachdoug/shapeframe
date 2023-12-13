<script setup>
import ShowComponent from '/src/components/crud/ShowComponent.vue'
import { mdiFileEditOutline, mdiPlay, mdiCodeJson, mdiTrashCanOutline } from '@mdi/js'

const toolbarButtons = {
  left: [
    { label: 'Edit', icon: mdiFileEditOutline, path: 'edit' },
    { label: 'Artifacts', icon: mdiCodeJson, path: 'artifacts' },
    { label: 'Apply', icon: mdiPlay, path: 'apply' },
  ],
  right: [{ label: 'Delete', icon: mdiTrashCanOutline, path: 'delete' }],
}
</script>

<template>
  <div key="frames-show">
    <ShowComponent :toolbar-buttons="toolbarButtons"></ShowComponent>
  </div>
</template>
