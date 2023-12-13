<script setup>
import { ref } from 'vue'
import { mdiEye, mdiEyeOff } from '@mdi/js'
import 'text-security/text-security.css'

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
})

const emit = defineEmits(['update'])

const model = ref(props.value)
const show = ref(false)

const update = value => emit('update', value)

const toggle = e => {
  const classes = e.currentTarget
    .closest('.v-input')
    .querySelector('input').classList
  if (classes.contains('show-text')) {
    show.value = false
    classes.remove('show-text')
  } else {
    show.value = true
    classes.add('show-text')
  }
}
</script>

<template>
  <div class="secret-field-wrapper">
    <v-text-field
      v-model="model"
      :label="label"
      :placeholder="placeholder"
      :name="name"
      :rules="rules"
      :append-icon="show ? mdiEyeOff : mdiEye"
      autocomplete="off"
      @click:append="toggle"
      @update:model-value="update"
    ></v-text-field>
  </div>
</template>

<style scoped>
.secret-field-wrapper :deep(input) {
  font-family: text-security-disc;
}

.secret-field-wrapper :deep(input.show-text) {
  font-family: monospace;
  font-size: unset;
}
</style>
