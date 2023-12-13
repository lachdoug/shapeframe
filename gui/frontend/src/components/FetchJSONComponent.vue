<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import LoadingSpinnerComponent from './LoadingSpinnerComponent.vue'
import ServerErrorDialogComponent from './ServerErrorDialogComponent.vue'
import { mdiArrowLeft } from '@mdi/js'
import markdown from '/src/lib/markdown'

const props = defineProps({
  // URL to call.
  endpoint: {
    type: String,
    required: true,
  },
  // HTTP method.
  method: {
    type: String,
    default: 'GET',
  },
  // HTTP request body object to stringify.
  body: {
    type: Object,
    default: null,
  },
  // Text for loading spinner placeholder.
  placeholder: {
    type: String,
    default: 'Loading',
  },
  // Emit event with data, otherwise render data in a slot.
  emitFetchedEvent: {
    type: Boolean,
    default: false,
  },
  emitBackEvent: {
    type: Boolean,
    default: false,
  },
})

const emit = defineEmits(['fetched', 'back'])

const router = useRouter()

let data, failMessage

const phase = ref('load')
const serverErrorMessage = ref(null)

fetch(props.endpoint, {
  method: props.method,
  ...(props.body
    ? {
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(props.body),
      }
    : {}),
})
  .then(response => {
    if (response.status == 200) {
      response.json().then(payload => {
        data = payload
        if (props.emitFetchedEvent) {
          emit('fetched', data)
        } else {
          // If no slot given, default is to render JSON in a <pre> tag.
          phase.value = 'slot'
        }
      })
    } else {
      response.text().then(message => {
        message = message || 'No response from server.'
        console.error('Server error: ' + message)
        failMessage = 'Server error.'
        if (window.settings.surfaceServerErrors)
          serverErrorMessage.value = message
        phase.value = 'fail'
      })
    }
  })
  .catch(error => {
    console.error('Error fetching data:', error)
    failMessage = 'Fetch error.'
    phase.value = 'fail'
  })

const back = () => {
  if (props.emitBackEvent) {
    emit('back', !failMessage)
  } else {
    router.go(-1)
  }
}
</script>

<template>
  <div v-if="phase == 'load'">
    <LoadingSpinnerComponent
      :placeholder="placeholder"
    ></LoadingSpinnerComponent>
  </div>
  <div v-else-if="phase == 'slot'">
    <slot :data="data">
      <pre>{{ JSON.stringify(data.Result, null, 2) }}</pre>
    </slot>
  </div>
  <div v-else-if="phase == 'fail'">
    <v-card class="my-2">
      <v-card-text>
        <v-btn class="mr-2" @click="back"
          ><v-icon :icon="mdiArrowLeft"></v-icon
        ></v-btn>
        <!-- eslint-disable vue/no-v-html -->
        <!-- The HTML is sanitized by DOMPurify in the markdown function in /src/lib/markdown -->
        <div
          class="markdown text-error d-inline-block px-0 py-2"
          v-html="markdown(failMessage)"
        ></div>
        <!-- eslint-disable vue/no-v-html -->
      </v-card-text>
    </v-card>
  </div>
  <div v-if="serverErrorMessage">
    <ServerErrorDialogComponent
      :text="serverErrorMessage"
    ></ServerErrorDialogComponent>
  </div>
</template>

<style scoped>
.loading-spinner {
  padding: 10px;
}
</style>
