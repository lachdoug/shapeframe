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
  placeholder: {
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
  type: {
    type: String,
    default: 'text',
  },
  autocomplete: {
    type: String,
    default: '',
  },
})

const emit = defineEmits(['update'])

const model = ref(props.value)

const update = value => emit('update', value)
</script>

<template>
  <v-textarea
    v-model="model"
    :label="label"
    :placeholder="placeholder"
    :name="name"
    :rules="rules"
    :autocomplete="autocomplete"
    :type="type"
    @update:model-value="update"
  ></v-textarea>
</template>
