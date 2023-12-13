<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FormComponent from '/src/components/FormComponent.vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import isType from '/src/lib/isType'

const props = defineProps({
  controls: {
    type: Object,
    default: () => ({}),
  },
  layout: {
    type: Array,
    default: null,
  },
  endpoint: {
    type: String,
    default: null,
  },
  saveMethod: {
    type: String,
    default: 'POST',
  },
  deconstruct: {
    type: Function,
    default: form => ({ model: form }),
  },
  returnPath: {
    type: [Function, String],
    default: () => (result, parentPath) => `${parentPath}/@${result}`,
  },
})

const phase = ref('form')
let result, values

const router = useRouter()
const route = useRoute()
const path = route.path
const parentPath = path.replace(/\/new$/, '')
const defaultEndpoint = `/api${parentPath}`

const endpoint = props.endpoint || defaultEndpoint

const submitted = form => {
  result = props.deconstruct(form)
  values = form // Keep the form values in case there is an error and the form needs to be shown again.
  phase.value = 'save'
}

const saved = response => {
  const path = isType('function', props.returnPath)
    ? props.returnPath(response.Result.ID, parentPath)
    : props.returnPath
  router.push(path)
}
</script>

<template>
  <div v-if="phase == 'form'">
    <FormComponent
      :controls="controls"
      :layout="layout"
      :values="values"
      @submitted="submitted"
    ></FormComponent>
  </div>
  <div v-else-if="phase == 'save'">
    <FetchJSONComponent
      :endpoint="endpoint"
      :method="saveMethod"
      :body="result"
      emit-back-event
      emit-fetched-event
      @back="phase = 'form'"
      @fetched="saved"
    >
    </FetchJSONComponent>
  </div>
</template>
