<script setup>
import NewComponent from '/src/components/crud/NewComponent.vue'
import {
  requiredValidation,
  noWhitespaceValidation,
} from '/src/lib/fieldValidation'

const filepathRule = v => {
  if (v.match(/^\//)) {
    return 'Must not start with a slash.'
  } else if (!v.match(/^(?:commissioning|packing|service_tasks)/)) {
    return "Must have a first-level directory name of either 'commissioning', 'packing', or 'service_tasks'."
  } else if (!v.match(/^[\w_]+\/[-\w./]+$/)) {
    return 'Must be a valid filepath.'
  } else {
    return true
  }
}

const deconstruct = form => ({
  model: {
    identifier: form.identifier.replace(/\//g, '::'),
  },
})

const controls = {
  identifier: {
    label: 'File path',
    rules: [requiredValidation, noWhitespaceValidation, filepathRule],
  },
}
</script>

<template>
  <div key="shapes-files-new">
    <NewComponent
      :controls="controls"
      :deconstruct="deconstruct"
    ></NewComponent>
  </div>
</template>
