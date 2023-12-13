<script setup>
import { useRoute } from 'vue-router'

defineProps({
  buttons: {
    type: [Object, Array, Object],
    required: true,
  },
})

const route = useRoute()

const pathTo = path => (path[0] == '/' ? path : `${route.path}/${path}`)
</script>

<template>
  <v-toolbar>
    <v-btn
      v-for="(button, i) in buttons.left"
      :key="i"
      :to="pathTo(button.path)"
      :prepend-icon="button.icon"
    >
      {{ button.label }}
    </v-btn>
    <v-spacer></v-spacer>
    <v-btn
      v-for="(button, i) in buttons.right"
      :key="i"
      :to="pathTo(button.path)"
      :prepend-icon="button.icon"
    >
      {{ button.label }}
    </v-btn>
  </v-toolbar>
</template>
