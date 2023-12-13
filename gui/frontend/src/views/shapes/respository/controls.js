import {
  requiredValidation,
  urlValidation,
  noWhitespaceValidation,
} from '/src/lib/fieldValidation'

const controls = keys => ({
  Repository: {
    label: 'Repository',
    placeholder: 'URL for Git repository',
    type: 'url',
    rules: [requiredValidation, urlValidation],
  },
  Branch: {
    label: "Branch",
    placeholder: "Optional, defaults to 'main'",
    rules: [noWhitespaceValidation],
  },
  KeyID: {
    label: 'Key',
    as: 'select',
    placeholder: 'Optional',
    cast: Number,
    options: keys,
  },
})

const layout = [['Repository'], ['Branch', 'KeyID']]

export {controls, layout}
