<script setup>
import ActionComponent from '/src/components/crud/ActionComponent.vue'
// import { requiredValidation } from '/src/lib/fieldValidation'
// const controls = {
//   message: {
//     as: 'textarea',
//     placeholder: 'Commit message',
//     rules: [requiredValidation],
//   },
// }
</script>

<template>
  <div key="repositories-push">
    <ActionComponent action="Push"></ActionComponent>
  </div>
</template>
