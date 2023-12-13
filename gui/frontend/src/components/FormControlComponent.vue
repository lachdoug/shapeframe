<script setup>
import { defineAsyncComponent } from 'vue'
import { useRoute } from 'vue-router'
import InputFieldComponent from '/src/components/fields/InputFieldComponent.vue'
import SelectFieldComponent from '/src/components/fields/SelectFieldComponent.vue'
import CheckboxFieldComponent from '/src/components/fields/CheckboxFieldComponent.vue'
import SecretFieldComponent from '/src/components/fields/SecretFieldComponent.vue'
import TextareaFieldComponent from '/src/components/fields/TextareaFieldComponent.vue'
import TextCaptionComponent from '/src/components/TextCaptionComponent.vue'
import LoadingSpinnerComponent from '/src/components/LoadingSpinnerComponent.vue'
// Dynamically load JSONFieldComponent and FileFieldComponent since they are big and want to keep bundle sizes under 500kB.
const JSONFieldComponent = defineAsyncComponent({
  loader: () => import('/src/components/fields/JSONFieldComponent.vue'),
  delay: 0,
  loadingComponent: LoadingSpinnerComponent,
})
const FileFieldComponent = defineAsyncComponent({
  loader: () => import('/src/components/fields/FileFieldComponent.vue'),
  delay: 0,
  loadingComponent: LoadingSpinnerComponent,
})

const props = defineProps({
  control: {
    type: Object,
    required: true,
  },
})

const emit = defineEmits(['update'])

const route = useRoute()

const name =
  props.control.name ||
  `${route.path.replace(/\//, '').replace(/\//g, '-')}-${props.control.key}`

const value = props.control.value == undefined ? '' : props.control.value

const update = value => emit('update', value)
</script>

<template>
  <TextCaptionComponent
    v-if="control.as == 'caption'"
    :text="control.text"
  ></TextCaptionComponent>
  <SelectFieldComponent
    v-else-if="control.as == 'select'"
    :value="value"
    :label="control.label"
    :placeholder="control.placeholder"
    :name="name"
    :rules="control.rules"
    :options="control.options"
    @update="update"
  >
  </SelectFieldComponent>
  <CheckboxFieldComponent
    v-else-if="control.as == 'checkbox'"
    :value="value"
    :label="control.label"
    :name="name"
    :rules="control.rules"
    @update="update"
  ></CheckboxFieldComponent>
  <JSONFieldComponent
    v-else-if="control.as == 'json'"
    :value="value"
    :label="control.label"
    :placeholder="control.placeholder"
    :name="name"
    :rules="control.rules"
    @update="update"
  ></JSONFieldComponent>
  <FileFieldComponent
    v-else-if="control.as == 'file'"
    :value="value"
    :label="control.label"
    :name="name"
    :rules="control.rules"
    @update="update"
  ></FileFieldComponent>
  <SecretFieldComponent
    v-else-if="control.as == 'secret'"
    :value="value"
    :label="control.label"
    :name="name"
    :rules="control.rules"
    @update="update"
  ></SecretFieldComponent>
  <TextareaFieldComponent
    v-else-if="control.as == 'textarea'"
    :value="value"
    :label="control.label"
    :placeholder="control.placeholder"
    :name="name"
    :rules="control.rules"
    @update="update"
  ></TextareaFieldComponent>
  <InputFieldComponent
    v-else
    :value="value"
    :label="control.label"
    :placeholder="control.placeholder"
    :name="name"
    :rules="control.rules"
    :autocomplete="control.autocomplete"
    :type="control.type"
    @update="update"
  >
  </InputFieldComponent>
</template>
