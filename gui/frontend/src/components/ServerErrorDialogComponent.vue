<script setup>
import { ref } from 'vue'
import { mdiBugOutline, mdiClose } from '@mdi/js'

defineProps({
  text: {
    type: String,
    required: true,
  },
})

const show = ref(window.settings.surfaceServerErrors)

const close = () => (show.value = false)
</script>

<template>
  <v-dialog v-model="show">
    <v-card>
      <template #title>
        <div class="float-left text-error">
          <v-icon :icon="mdiBugOutline"></v-icon>
          Server Error
        </div>
        <div class="float-right">
          <v-icon :icon="mdiClose" @click="close"></v-icon>
        </div>
      </template>
      <v-card-title> </v-card-title>
      <v-card-text>
        <pre>{{ text }}</pre>
      </v-card-text>
    </v-card>
  </v-dialog>
</template>

<style scoped>
.v-dialog {
  align-items: normal;
  /* Position dialog at top of page. */
}

.v-card {
  overflow: auto;
  /* Allow scrolling in both directions. v-card normally scrolls Y only. */
}
</style>
