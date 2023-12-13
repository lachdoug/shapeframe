<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FormButtonsComponent from '/src/components/FormButtonsComponent.vue'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'

const phase = ref('form')

const route = useRoute()
const router = useRouter()

const identifier = route.params.identifier
const fileIdentifier = route.params.fileIdentifier

const submitted = () => {
  phase.value = 'save'
}

const saved = () => {
  router.push(`/shapes/@${identifier}/files`)
}
</script>

<template>
  <div key="shapes-files-delete">
    <v-card class="my-2">
      <v-card-text>
        <p>Delete @{{ fileIdentifier }} from shape @{{ identifier }}?</p>
        <div v-if="phase == 'form'">
          <v-form @submit.prevent="submitted">
            <FormButtonsComponent></FormButtonsComponent>
          </v-form>
        </div>
        <div v-else-if="phase == 'save'">
          <FetchJSONComponent
            :url="`/api/shapes/@${identifier}/files/@${fileIdentifier}`"
            method="DELETE"
            emit-fetched-event
            @fetched="saved"
          ></FetchJSONComponent>
        </div>
      </v-card-text>
    </v-card>
  </div>
</template>
