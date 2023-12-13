<script setup>
import EditComponent from '/src/components/crud/EditComponent.vue'

const controls = {
  content: {
    label: 'Content',
    as: 'file',
  },
}
</script>

<template>
  <div key="shapes-files-edit">
    <EditComponent :controls="controls"></EditComponent>
  </div>
</template>
