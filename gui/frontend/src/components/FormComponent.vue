<script setup>
import { reactive } from 'vue'
import FormControlComponent from '/src/components/FormControlComponent.vue'
import ButtonsComponent from '/src/components/FormButtonsComponent.vue'

const props = defineProps({
  controls: {
    type: Object,
    default: () => ({}),
  },
  values: {
    type: Object,
    default: () => ({}),
  },
  layout: {
    type: Array,
    default: null,
  },
  buttonless: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['submitted', 'changed'])
const hideControl = reactive({})

let valid = true

const result = props.values

const submitted = () => {
  if (valid) {
    emit('submitted', result)
  }
}

const changed = () => {
  if (valid) {
    emit('changed', result)
  }
}

const update = updated => {
  if (valid) {
    const deconstruct = props.controls[updated.key].deconstruct
    result[updated.key] = deconstruct
      ? deconstruct(updated.value)
      : updated.value
    updateHideControl()
  }
}

const layout = props.layout || Object.keys(props.controls).map(key => [key])

const controlFor = key => {
  let control = props.controls[key]
  control.key = key
  let raw = props.values[key]
  const construct = props.controls[key].construct
  control.value = construct ? construct(raw) : raw
  return control
}

const updateHideControl = () =>
  Object.keys(props.controls).forEach(
    key =>
      (hideControl[key] = (props.controls[key].hide || (() => false))(result))
  )
updateHideControl()
</script>

<template>
  <div class="form-wrapper">
    <v-form v-model="valid" @submit.prevent="submitted" @change="changed">
      <v-row v-for="(row, i) in layout" :key="i">
        <v-col v-for="(key, j) in row" :key="j">
          <FormControlComponent
            v-if="!hideControl[key]"
            :control="controlFor(key)"
            @update="value => update({ key: key, value: value })"
          ></FormControlComponent>
        </v-col>
      </v-row>
      <ButtonsComponent v-if="!buttonless"></ButtonsComponent>
    </v-form>
  </div>
</template>

<style scoped>
/* For form layout needs to be no vertical gutters, but keep horizinal gutters. */
.form-wrapper :deep(.v-form) {
  margin-top: 1rem;
}

.form-wrapper :deep(.v-row) {
  margin-top: 0rem;
}

.form-wrapper :deep(.v-row [class*='v-col']) {
  padding-top: 0rem;
  padding-bottom: 1.25rem;
}

.form-wrapper :deep(.v-row [class*='v-col']:empty) {
  padding-bottom: 12px;
}
</style>
