<script setup>
import { ref } from 'vue'
import { basicSetup, EditorView } from 'codemirror'
import { Codemirror } from 'vue-codemirror'
import { vim } from '@replit/codemirror-vim'
import { mdiFullscreen, mdiFullscreenExit } from '@mdi/js'

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

const model = ref(props.value, null, 2)

const useVim = window.settings.editorVimKeys

const codeMirrorExtensions = [
  basicSetup,
  useVim ? vim() : [],
  EditorView.lineWrapping,
]

const ready = codemirror => {
  setTimeout(() => codemirror.view.scrollDOM.scroll({ top }), 0)
}

const update = value => emit('update', value)

const enterFullscreen = e => {
  e.currentTarget
    .closest('.json-field-wrapper')
    .querySelector('.codemirror-wrapper')
    .classList.add('fullscreen')
}

const exitFullscreen = e => {
  e.currentTarget
    .closest('.json-field-wrapper')
    .querySelector('.codemirror-wrapper')
    .classList.remove('fullscreen')
}
</script>

<template>
  <div class="json-field-wrapper">
    <VFieldLabel>{{ label }}</VFieldLabel>
    <div class="text-right">
      <v-btn class="mb-2" @click="enterFullscreen">
        <v-icon :icon="mdiFullscreen" />
        Fullscreen
      </v-btn>
    </div>
    <div class="codemirror-wrapper">
      <v-btn class="exit-fullscreen-button" @click="exitFullscreen">
        <v-icon :icon="mdiFullscreenExit" />
      </v-btn>
      <Codemirror
        v-model="model"
        placeholder="Enter file content"
        :style="{ height: '100%' }"
        :autofocus="true"
        :indent-with-tab="true"
        :tab-size="2"
        :extensions="codeMirrorExtensions"
        @ready="ready"
        @change="update"
      />
    </div>
    <v-textarea
      v-model="model"
      :rules="rules"
      class="error-display-field mb-2"
    ></v-textarea>
  </div>
</template>

<style scoped>
.json-field-wrapper :deep(.error-display-field .v-field__input) {
  display: none;
}

.json-field-wrapper :deep(.error-display-field .v-field__outline::before) {
  border: none;
}

.json-field-wrapper .codemirror-wrapper {
  height: 400px;
  border: 1px solid rgba(var(--v-border-color), var(--v-border-opacity));
}

.json-field-wrapper .codemirror-wrapper .exit-fullscreen-button {
  display: none;
}

.json-field-wrapper .codemirror-wrapper.fullscreen .exit-fullscreen-button {
  display: block;
  background-color: transparent;
  position: fixed;
  top: 5px;
  right: 5px;
  z-index: 2001;
}

.json-field-wrapper
  .codemirror-wrapper.fullscreen
  .exit-fullscreen-button:hover {
  display: block;
  background-color: inherit;
}

.json-field-wrapper .codemirror-wrapper.fullscreen {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  height: 100vh;
  background-color: rgb(var(--v-theme-background));
  z-index: 2000;
}
</style>
