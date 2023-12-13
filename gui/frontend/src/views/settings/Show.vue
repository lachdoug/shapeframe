<script setup>
import { useRouter } from 'vue-router'
import FormComponent from '/src/components/FormComponent.vue'

import { mdiArrowLeft } from '@mdi/js'

const values = {
  graphScale: window.settings.graphScale,
  graphFontSize: window.settings.graphFontSize,
  graphFit: window.settings.graphFit,
  graphSizeX: window.settings.graphSizeX,
  graphSizeY: window.settings.graphSizeY,
  graphInterconnect: window.settings.graphInterconnect,
  editorVimKeys: window.settings.editorVimKeys,
  shapesSchemaValidationSkip: window.settings.shapesSchemaValidationSkip,
  framesSchemaValidationSkip: window.settings.framesSchemaValidationSkip,
  surfaceServerErrors: window.settings.surfaceServerErrors,
}

const controls = {
  graphCaption: {
    as: 'caption',
    text: 'Graph',
  },
  graphScale: {
    label: 'Scale',
    type: 'number',
    min: 1,
    step: 0.1,
    rules: [v => (v < 1 ? 'Must be at least 1' : true)],
  },
  graphFontSize: {
    label: 'Font size',
    type: 'number',
    min: '0',
    step: '1',
    rules: [v => (v <= 0 ? 'Must not be negative' : true)],
  },
  graphFit: {
    label: 'Fit to view',
    as: 'checkbox',
  },
  graphSizeX: {
    label: 'Horizontal',
    type: 'number',
    min: '100',
    step: '1',
    rules: [v => (v < 100 ? 'Must be at least 100' : true)],
    hide: form => form.graphFit,
  },
  graphSizeY: {
    label: 'Vertical',
    type: 'number',
    min: '100',
    step: '1',
    rules: [v => (v < 100 ? 'Must be at least 100' : true)],
    hide: form => form.graphFit,
  },
  graphInterconnect: {
    label: 'Interconnect',
    as: 'checkbox',
  },
  editorKeysCaption: {
    as: 'caption',
    text: 'Editor Keys',
  },
  editorVimKeys: {
    label: 'VIM',
    as: 'checkbox',
  },
  skipSchemaValidationCaption: {
    as: 'caption',
    text: 'Skip schema validation',
  },
  shapesSchemaValidationSkip: {
    label: 'Shapes',
    as: 'checkbox',
  },
  framesSchemaValidationSkip: {
    label: 'Frames',
    as: 'checkbox',
  },
  errorsCaption: {
    as: 'caption',
    text: 'Errors',
  },
  surfaceServerErrors: {
    label: 'Surface server errors',
    as: 'checkbox',
  },
}

const layout = [
  ['graphCaption'],
  ['graphScale', 'graphFontSize'],
  ['graphFit', 'graphSizeX', 'graphSizeY'],
  ['graphInterconnect'],
  ['editorKeysCaption'],
  ['editorVimKeys'],
  ['skipSchemaValidationCaption'],
  ['shapesSchemaValidationSkip', 'framesSchemaValidationSkip'],
  ['errorsCaption'],
  ['surfaceServerErrors'],
]

const changed = form => {
  Object.keys(form).forEach(key => (window.settings[key] = form[key]))
}
const router = useRouter()
const back = () => router.push('/')
</script>

<template>
  <div key="settings-show">
    <v-btn @click="back"><v-icon :icon="mdiArrowLeft"></v-icon></v-btn>
    <div class="mt-2">
      <FormComponent
        :controls="controls"
        :values="values"
        :layout="layout"
        :buttonless="true"
        @changed="changed"
      >
      </FormComponent>
    </div>
  </div>
</template>
