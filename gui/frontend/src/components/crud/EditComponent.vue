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
    default: '',
  },
  construct: {
    type: Function,
    default: loaded => loaded.Result,
  },
  deconstruct: {
    type: Function,
    default: submitted => ({ model: submitted }),
  },
  returnPath: {
    type: [Function, String],
    default: () => (identifier, parentPath) => `${parentPath}/@${identifier}`,
  },
})

const phase = ref('load')
let submission, values

const router = useRouter()
const route = useRoute()
const itemPath = route.path.replace(/\/\w+$/, '')
const defaultUrl = `/api${itemPath}`

const endpoint = props.endpoint || defaultUrl

const loaded = response => {
  values = props.construct(response)
  phase.value = 'form'
}

const submitted = form => {
  submission = props.deconstruct(form)
  values = form // Keep the form values in case there is an error and the form needs to be shown again.
  phase.value = 'save'
}

const saved = response => {
  const parentPath = itemPath.replace(/\/[@\d]+$/, '')
  const path = isType('function', props.returnPath)
    ? props.returnPath(response.Result.ID, parentPath)
    : props.returnPath
  router.push(path)
}
</script>

<template>
  <div v-if="phase == 'load'">
    <FetchJSONComponent
      :endpoint="endpoint"
      emit-fetched-event
      @fetched="loaded"
    ></FetchJSONComponent>
  </div>

  <div v-else-if="phase == 'form'">
    <FormComponent
      :controls="controls"
      :values="values"
      :layout="layout"
      @submitted="submitted"
    ></FormComponent>
  </div>

  <div v-else-if="phase == 'save'">
    <FetchJSONComponent
      :endpoint="endpoint"
      method="PUT"
      :body="submission"
      emit-back-event
      emit-fetched-event
      @back="phase = 'form'"
      @fetched="saved"
    >
    </FetchJSONComponent>
  </div>
</template>
