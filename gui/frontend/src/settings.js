const defaults = {
  graphScale: 1,
  graphFontSize: 16,
  graphFit: true,
  graphSizeX: 1000,
  graphSizeY: 1000,
  graphInterconnect: false,
  editorVimKeys: false,
  shapesSchemaValidationSkip: false,
  framesSchemaValidationSkip: false,
  surfaceServerErrors: false,
}

const types = {
  graphScale: Number,
  graphFontSize: Number,
  graphFit: Boolean,
  graphSizeX: Number,
  graphSizeY: Number,
  graphInterconnect: Boolean,
  editorVimKeys: Boolean,
  shapesSchemaValidationSkip: Boolean,
  framesSchemaValidationSkip: Boolean,
  surfaceServerErrors: Boolean,
}

const stored = window.localStorage.settings

const target = stored ? JSON.parse(stored) : defaults

const getValue = property => {
  const value = target[property]
  return types[property](value == undefined ? defaults[property] : value)
}

const setValue = (property, value) => {
  target[property] = value
  save()
  return true
}

const handler = {
  get(_, property) {
    return getValue(property)
  },
  set(_, property, value) {
    return setValue(property, value)
  },
}

const proxy = new Proxy(target, handler)

const save = () => (window.localStorage.settings = JSON.stringify(target))

const apply = () => (window.settings = proxy)

export default apply
