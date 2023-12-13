<script setup>
import { ref } from 'vue'

const props = defineProps({
  value: {
    type: [String, Number],
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
  options: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['update'])

const model = ref(String(props.value) || null)

const items = Object.keys(props.options).map(k => ({
  value: String(k),
  title: props.options[k],
}))

const update = value => emit('update', value)
</script>

<template>
  <v-select
    v-model="model"
    :label="label"
    :placeholder="placeholder"
    :name="name"
    :rules="rules"
    :items="items"
    no-data-text="None"
    @update:model-value="update"
  >
  </v-select>
</template>
