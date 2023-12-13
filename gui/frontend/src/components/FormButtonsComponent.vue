<script setup>
import { useRoute } from 'vue-router'
import { mdiCancel, mdiCheck } from '@mdi/js'

const route = useRoute()

const segments = route.path.split('/')
segments.pop()
const cancelPath = segments.join('/')
</script>

<template>
  <v-btn :to="cancelPath" color="error">
    <v-icon :icon="mdiCancel" />
    Cancel
  </v-btn>
  <v-btn type="submit" color="success" class="ml-2">
    <v-icon :icon="mdiCheck" />
    Submit
  </v-btn>
</template>
