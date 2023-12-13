<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import FetchJSONComponent from '/src/components/FetchJSONComponent.vue'
import ToolbarComponent from '/src/components/ToolbarComponent.vue'

const props = defineProps({
  toolbarButtons: {
    type: Object,
    required: true,
  },
  construct: {
    type: Function,
    default: loaded => loaded.Result,
  },
})

const phase = ref('load')
let items

const router = useRouter()
const route = useRoute()

const table = route.path.split('/')[1]

const loaded = response => {
  items = props.construct(response)
  phase.value = 'table'
}

const navigateToItem = e => {
  const path = route.path
  const ID = e.currentTarget.getAttribute('ID')
  router.push(`${path}/@${ID}`)
}
</script>

<template>
  <ToolbarComponent :buttons="toolbarButtons"></ToolbarComponent>
  <div v-if="phase == 'load'" class="mt-2">
    <FetchJSONComponent
      :endpoint="`/api/${table}`"
      emit-fetched-event
      @fetched="loaded"
    >
    </FetchJSONComponent>
  </div>
  <div v-else-if="phase == 'table'">
    <v-card class="my-2">
      <v-card-text v-if="!items.length">
        <i> No {{ table.replace(/_/g, ' ') }} </i>
      </v-card-text>
      <v-table v-else>
        <tbody>
          <tr
            v-for="(item, i) in items"
            :key="i"
            v-ripple
            :ID="item.ID"
            @click="navigateToItem"
          >
            <slot :item="item">
              <td>
                <div>
                  {{ item.Label }}
                </div>
                <div>
                  <small>
                    {{ item.Description }}
                  </small>
                </div>
              </td>
            </slot>
          </tr>
        </tbody>
      </v-table>
    </v-card>
  </div>
</template>

<style scoped>
table tr:hover {
  background-color: rgba(var(--v-theme-on-background), var(--v-hover-opacity));
}

table tr {
  cursor: pointer;
}
</style>
