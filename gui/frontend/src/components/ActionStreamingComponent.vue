<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import FormButtonsComponent from '/src/components/FormButtonsComponent.vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import TerminalStreamingComponent from '/src/components/TerminalStreamingComponent.vue'
import { mdiCheck } from '@mdi/js'

const phase = ref('form')
let timestamp, saveData, buttonColor

const route = useRoute()
const identifier = route.params.identifier

const submitted = () => {
  saveData = {
    background: 'on',
  }
  phase.value = 'save'
}

const saved = saveResponse => {
  timestamp = saveResponse.result.timestamp
  phase.value = 'stream'
}

const streamed = success => {
  if (success) {
    buttonColor = 'success'
  } else {
    buttonColor = 'error'
  }
  phase.value = 'complete'
}
</script>

<template>
  <div key="frames-apply">
    <div v-if="phase == 'form'">
      <p>Apply @{{ identifier }} frame?</p>
      <v-form @submit.prevent="submitted">
        <FormButtonsComponent></FormButtonsComponent>
      </v-form>
    </div>
    <div v-else-if="phase == 'save'">
      <FetchJSONComponent
        :url="`/api/frames/@${identifier}/apply`"
        method="POST"
        :body="saveData"
        emit-fetched-event
        @fetched="saved"
      ></FetchJSONComponent>
    </div>
    <div v-else-if="phase == 'stream' || phase == 'complete'">
      <TerminalStreamingComponent
        :url="`/api/streaming/frames/@${identifier}/apply?timestamp=${timestamp}`"
        @stream-closed="streamed"
      ></TerminalStreamingComponent>
      <div v-if="phase == 'complete'" class="mt-2">
        <v-btn :to="`/frames/@${identifier}`" :color="buttonColor">
          <v-icon :icon="mdiCheck" />
          Done
        </v-btn>
      </div>
    </div>
  </div>
</template>
