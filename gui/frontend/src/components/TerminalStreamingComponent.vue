<script setup>
import 'xterm/css/xterm.css'
import { ref, onMounted, onUnmounted } from 'vue'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'
import ServerErrorDialogComponent from './ServerErrorDialogComponent.vue'
import { mdiFullscreen, mdiFullscreenExit } from '@mdi/js'

const props = defineProps({
  // URL endpoint for stream
  endpoint: {
    type: String,
    required: true,
  },
})

const emit = defineEmits(['stream-closed'])

const xtermContainer = ref(null)
const error = ref(null)
let terminal, eventsource, errorReported, fitAddon, resizeTimer

onMounted(() => {
  window.addEventListener('resize', resizeHandler)
  initTerminal()
  initEventsource()
})

onUnmounted(() => {
  window.removeEventListener('resize', resizeHandler)
})

const resizeHandler = () => {
  clearTimeout(resizeTimer)
  resizeTimer = setTimeout(fitTerminalToContainer, 100)
}

const fitTerminalToContainer = () => fitAddon.fit()

const initTerminal = () => {
  terminal = new Terminal({
    disableStdin: true,
    cursorStyle: 'bar',
    convertEol: true,
    fontSize: 16,
  })

  fitAddon = new FitAddon()
  terminal.loadAddon(fitAddon)
  terminal.open(xtermContainer.value)
  fitTerminalToContainer()
}

const initEventsource = () => {
  eventsource = new EventSource(props.endpoint)

  eventsource.addEventListener('output', e => {
    const data = JSON.parse(e.data)
    if (data.output) write(data.output)
    if (data.error) {
      errorReported = true
      write(`${data.error}`, { color: 'yellow' })
    }
    if (data.error) {
      write(`Failed!`, { color: 'red' })
      errorReported = true
      const message = data.error
      if (window.settings.surfaceServerErrors) error.value = message
      console.error(message)
    }
  })

  eventsource.addEventListener('error', e => {
    write(`Failed!`, { color: 'red' })
    errorReported = true
    let message = e.data
    if (window.settings.surfaceServerErrors) error.value = message
    console.error(e.data)
  })
  eventsource.addEventListener('eot', close)
  eventsource.timeout = () => {
    errorReported = true
    console.error('Terminal Fetch Stream Timeout.')
    close()
  }
  eventsource.onerror = e => {
    errorReported = true
    console.error(e)
    close()
  }
}

const write = (text, opts = {}) => {
  const color = {
    black: '30',
    red: '31',
    green: '32',
    yellow: '33',
    blue: '34',
    magenta: '35',
    cyan: '36',
    white: '37',
  }[opts.color || 'white']
  const boldness = opts.bold ? '1' : '0'

  terminal.write(`\u001b[${boldness};${color}m${text}\u001b[0m`)
}

const enterFullscreen = e => {
  e.currentTarget
    .closest('.terminal-streaming-wrapper')
    .querySelector('.xterm-wrapper')
    .classList.add('fullscreen')
  fitTerminalToContainer()
}

const exitFullscreen = e => {
  e.currentTarget
    .closest('.terminal-streaming-wrapper')
    .querySelector('.xterm-wrapper')
    .classList.remove('fullscreen')
  fitTerminalToContainer()
}

const close = () => {
  if (eventsource) {
    eventsource.close()
    eventsource = null
  }
  emit('stream-closed', !errorReported)
}
</script>

<template>
  <div class="terminal-streaming-wrapper">
    <div class="text-right">
      <v-btn class="mb-2" @click="enterFullscreen">
        <v-icon :icon="mdiFullscreen" />
        Fullscreen
      </v-btn>
    </div>
    <div class="xterm-wrapper">
      <v-btn class="exit-fullscreen-button" @click="exitFullscreen">
        <v-icon :icon="mdiFullscreenExit" />
      </v-btn>
      <div ref="xtermContainer" class="xterm-container"></div>
      <div v-if="error">
        <ServerErrorDialogComponent :text="error"></ServerErrorDialogComponent>
      </div>
    </div>
  </div>
</template>

<style scoped>
.terminal-streaming-wrapper .xterm-wrapper {
  width: 100%;
  height: 400px;
}

.terminal-streaming-wrapper .xterm-wrapper .xterm-container {
  height: 100%;
}

.terminal-streaming-wrapper .exit-fullscreen-button {
  display: none;
}

.terminal-streaming-wrapper .xterm-wrapper.fullscreen .exit-fullscreen-button {
  display: block;
  color: #fff9;
  border: 1px solid #fff9;
  background-color: transparent;
  position: fixed;
  top: 5px;
  right: 5px;
  z-index: 2001;
}

.terminal-streaming-wrapper
  .xterm-wrapper.fullscreen
  .exit-fullscreen-button:hover {
  color: unset;
  border: unset;
  display: block;
  background-color: inherit;
}

.terminal-streaming-wrapper .xterm-wrapper.fullscreen {
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
