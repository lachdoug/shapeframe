<script setup>
import EditComponent from '/src/components/crud/EditComponent.vue'
import { shapeValidation } from '/src/lib/fieldValidation'

const controls = {
  specification: {
    // label: 'Specification',
    as: 'json',
    rules: [shapeValidation],
  },
}

const construct = object => ({
  specification: JSON.stringify(object.result, null, 2),
})
const deconstruct = object => ({ model: JSON.parse(object.specification) })
</script>

<template>
  <div key="shapes-edit">
    <EditComponent
      :controls="controls"
      :construct="construct"
      :deconstruct="deconstruct"
    ></EditComponent>
  </div>
</template>
