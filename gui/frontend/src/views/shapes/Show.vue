<script setup>
import ShowComponent from '/src/components/crud/ShowComponent.vue'
import {
  mdiFileEditOutline,
  mdiGit,
  mdiImageOutline,
  mdiTrashCanOutline,
  mdiFileOutline,
} from '@mdi/js'

const toolbarButtons = {
  left: [
    { label: 'Edit', icon: mdiFileEditOutline, path: 'edit' },
    { label: 'Repository', icon: mdiGit, path: 'repository' },
    { label: 'Files', icon: mdiFileOutline, path: 'files' },
    { label: 'Icon', icon: mdiImageOutline, path: 'icon' },
  ],
  right: [{ label: 'Delete', icon: mdiTrashCanOutline, path: 'delete' }],
}
</script>

<template>
  <div key="shapes-show">
    <ShowComponent :toolbar-buttons="toolbarButtons"></ShowComponent>
  </div>
</template>
