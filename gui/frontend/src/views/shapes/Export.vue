<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import FormButtonsComponent from '/src/components/FormButtonsComponent.vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import TerminalStreamingComponent from '/src/components/TerminalStreamingComponent.vue'
import { requiredValidation } from '/src/lib/fieldValidation'
import { mdiCheck } from '@mdi/js'

const phase = ref('form')
let timestamp, formValid, saveData, buttonColor

const messageField = ref('')
const messageFieldRules = [requiredValidation]

const route = useRoute()
const identifier = route.params.identifier

const submitted = () => {
  if (formValid) {
    saveData = {
      background: 'on',
      message: messageField.value,
    }
    phase.value = 'save'
  }
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
  <div key="shapes-export">
    <div v-if="phase == 'form'">
      <v-form v-model="formValid" @submit.prevent="submitted">
        <v-row>
          <v-col>
            <v-text-field
              v-model="messageField"
              label="Commit message"
              name="messageField"
              :rules="messageFieldRules"
            />
          </v-col>
        </v-row>
        <FormButtonsComponent></FormButtonsComponent>
      </v-form>
    </div>
    <div v-else-if="phase == 'save'">
      <FetchJSONComponent
        :url="`/api/publications/@${identifier}/export`"
        method="POST"
        :body="saveData"
        emit-fetched-event
        @fetched="saved"
      ></FetchJSONComponent>
    </div>
    <div v-else-if="phase == 'stream' || phase == 'complete'">
      <TerminalStreamingComponent
        :url="`/api/streaming/publications/@${identifier}/exporting?timestamp=${timestamp}`"
        @stream-closed="streamed"
      ></TerminalStreamingComponent>
      <div v-if="phase == 'complete'" class="mt-2">
        <v-btn :to="`/shapes/@${identifier}`" :color="buttonColor">
          <v-icon :icon="mdiCheck" />
          Done
        </v-btn>
      </div>
    </div>
  </div>
</template>
