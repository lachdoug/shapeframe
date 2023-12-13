<script setup>
import ShowComponent from '/src/components/crud/ShowComponent.vue'
import {
  mdiFileEditOutline,
  mdiFolderMultipleOutline,
  mdiFolderSyncOutline,
  mdiFolderRemoveOutline,
  mdiTrashCanOutline,
  mdiFolderCheckOutline,
} from '@mdi/js'

const toolbarButtons = {
  left: [
    { label: 'Edit', icon: mdiFileEditOutline, path: 'edit' },
    { label: 'Clone', icon: mdiFolderMultipleOutline, path: 'clone' },
    { label: 'Pull', icon: mdiFolderSyncOutline, path: 'pull' },
    { label: 'Push', icon: mdiFolderCheckOutline, path: 'push' },
    { label: 'Remove', icon: mdiFolderRemoveOutline, path: 'remove' },
  ],
  right: [{ label: 'Delete', icon: mdiTrashCanOutline, path: 'delete' }],
}
</script>

<template>
  <div key="repositories-edit">
    <ShowComponent :toolbar-buttons="toolbarButtons"></ShowComponent>
  </div>
</template>
