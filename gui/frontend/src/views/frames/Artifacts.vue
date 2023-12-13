<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'

const tab = ref('one')
const route = useRoute()
const identifier = route.params.identifier
</script>

<template>
  <div key="frames-artifacts">
    <FetchJSONComponent :url="`/api/frames/@${identifier}/artifacts`">
      <template #default="{ data }">
        <v-card class="my-2">
          <v-tabs v-model="tab" bg-color="primary">
            <v-tab value="one">Orchestrate</v-tab>
            <v-tab value="two">Build</v-tab>
          </v-tabs>
          <v-card-text>
            <v-window v-model="tab">
              <v-window-item value="one">
                <i v-if="!data.Result.orchestrate.outputs.length">None</i>
                <div
                  v-else
                  v-for="(output, i) in data.Result.orchestrate.outputs"
                  :key="i"
                >
                  <pre>{{ output }}</pre>
                  <br />
                </div>
              </v-window-item>
              <v-window-item value="two">
                <i v-if="!data.Result.build.length">None</i>
                <template
                  v-else
                  v-for="(build, i) in data.Result.build"
                  :key="i"
                >
                  <div class="text-h6">{{ build.pack.split('::')[1] }}</div>
                  <div v-for="(output, j) in build.outputs" :key="j">
                    <pre>{{ output }}</pre>
                    <br />
                  </div>
                </template>
              </v-window-item>
            </v-window>
          </v-card-text>
        </v-card>
      </template>
    </FetchJSONComponent>
  </div>
</template>
