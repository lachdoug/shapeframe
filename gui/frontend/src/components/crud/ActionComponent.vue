<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FormComponent from '/src/components/FormComponent.vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import { mdiArrowLeft } from '@mdi/js'
import markdown from '/src/lib/markdown'

const props = defineProps({
  // Action name.
  action: {
    type: String,
    required: true,
  },
  // HTTP method for action.
  method: {
    type: String,
    default: 'POST',
  },
  // Controls for the form.
  controls: {
    type: Object,
    default: () => ({}),
  },
  // Layout of the form.
  layout: {
    type: Array,
    default: null,
  },
  // Mutator for loaded data, to marshall form values.
  construct: {
    type: Function,
    default: loaded => loaded.Result,
  },
  // Mutator for submitted data, to marshall save body.
  deconstruct: {
    type: Function,
    default: submitted => ({ model: submitted }),
  },
})

let label, report, submission, values
const phase = ref('load')

const route = useRoute()
const router = useRouter()

const table = route.path.split('/')[1]
const identifier = route.params.identifier

const loadEndpoint = `/api/${table}/@${identifier}`
const saveEndpoint = `${loadEndpoint}/${props.action.toLowerCase()}`

const loaded = response => {
  label = response.Result.Label
  values = props.construct(response)
  phase.value = 'form'
}

const submitted = form => {
  submission = props.deconstruct(form)
  values = form // Keep the form values in case there is an error and the form needs to be shown again.
  phase.value = 'save'
}

const saved = response => {
  report = response.Result
  phase.value = 'report'
}

const back = success => {
  if (success) {
    router.push(`/${table}/@${identifier}`)
  } else {
    phase.value = 'form'
  }
}
</script>

<template>
  <div v-if="phase == 'load'">
    <FetchJSONComponent
      :endpoint="loadEndpoint"
      emit-fetched-event
      @fetched="loaded"
    ></FetchJSONComponent>
  </div>
  <v-card v-if="phase == 'form'" class="my-0">
    <v-card-text>
      <h4>{{ action }} {{ label }}?</h4>
      <div>
        <FormComponent
          :controls="controls"
          :values="values"
          :layout="layout"
          @submitted="submitted"
        ></FormComponent>
      </div>
    </v-card-text>
  </v-card>
  <div v-else-if="phase == 'save'">
    <FetchJSONComponent
      :endpoint="saveEndpoint"
      :method="method"
      :body="submission"
      emit-back-event
      emit-fetched-event
      @back="back"
      @fetched="saved"
    ></FetchJSONComponent>
  </div>
  <v-card v-else-if="phase == 'report'" class="my-0">
    <v-card-text>
      <v-btn class="mr-2" @click="back"
        ><v-icon :icon="mdiArrowLeft"></v-icon
      ></v-btn>
      <div class="mt-4">
        <h4 :class="report.Status == 'fail' ? 'text-error' : 'text-success'">
          {{ report.Title }}
        </h4>
        <!-- eslint-disable vue/no-v-html -->
        <!-- The HTML is sanitized by DOMPurify in the markdown function in /src/lib/markdown -->
        <div
          class="markdown d-inline-block px-0 py-2"
          v-html="markdown(report.Text)"
        ></div>
        <!-- eslint-disable vue/no-v-html -->
      </div>
    </v-card-text>
  </v-card>
</template>
