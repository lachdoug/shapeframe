<script setup>
import { ref } from 'vue'

const props = defineProps({
  value: {
    type: String,
    default: '',
  },
  label: {
    type: String,
    default: '',
  },
  name: {
    type: String,
    default: '',
  },
  rules: {
    type: Array,
    default: () => [],
  },
})

const emit = defineEmits(['update'])

const model = ref(Boolean(props.value))

const update = value => emit('update', !!value)
</script>

<template>
  <v-checkbox
    v-model="model"
    :label="label"
    :name="name"
    :rules="rules"
    @update:model-value="update"
  ></v-checkbox>
</template>
