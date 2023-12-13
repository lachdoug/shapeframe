<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FormButtonsComponent from '/src/components/FormButtonsComponent.vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'

let label
const phase = ref('load')

const route = useRoute()
const router = useRouter()

const table = route.path.split('/')[1]
const identifier = route.params.identifier

const endpoint = `/api/${table}/@${identifier}`

const loaded = response => {
  label = response.Result.Label
  phase.value = 'form'
}

const submitted = () => {
  phase.value = 'save'
}

const saved = () => {
  router.push(`/${table}`)
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
  <v-card v-if="phase == 'form'" class="my-2">
    <v-card-text>
      <h4>Delete {{ label }} from {{ table.replace(/_/g, ' ') }}?</h4>
      <div>
        <v-form @submit.prevent="submitted">
          <FormButtonsComponent></FormButtonsComponent>
        </v-form>
      </div>
    </v-card-text>
  </v-card>
  <div v-else-if="phase == 'save'">
    <FetchJSONComponent
      :endpoint="endpoint"
      method="DELETE"
      emit-fetched-event
      @fetched="saved"
    ></FetchJSONComponent>
  </div>
</template>
