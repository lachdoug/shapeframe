<script setup>
import { useRoute } from 'vue-router'

const route = useRoute()
const path = route.path
</script>

<template>
  <v-alert type="warning" title="Page does not exist" :text="path"></v-alert>
</template>
