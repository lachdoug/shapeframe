<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const files = ref(null)
const route = useRoute()
const router = useRouter()

const upload = () => {
  const formData = new FormData()
  formData.append('icon', files.value[0])

  fetch(`/api/shapes/@${route.params.identifier}/icon`, {
    method: 'PUT',
    body: formData,
  })
    .then(() => router.push(`/shapes/@${route.params.identifier}`))
    .catch(error => {
      console.error('Error fetching data:', error)
    })
}
</script>

<template>
  <div key="shapes-icon">
    <v-file-input
      v-model="files"
      label="Icon"
      accept="image/*"
      @change="upload"
    ></v-file-input>
  </div>
</template>
