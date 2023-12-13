import { requiredValidation, urlValidation } from '/src/lib/fieldValidation'

const controls = {
  Label: {
    label: 'Label',
    rules: [requiredValidation],
  },
  Description: {
    label: 'Description',
    placeholder: 'Optional',
  },
  KeyID: {
    label: 'Key',
    as: 'select',
    placeholder: 'Optional',
    construct: v => (v == 0 ? '' : v),
    deconstruct: Number,
    options: {},
  },
  URL: {
    label: 'URL',
    placeholder: 'URL for Git repository',
    type: 'url',
    rules: [requiredValidation, urlValidation],
  },
}

const newLayout = [['Label'], ['Description'], ['URL', 'KeyID']]
const editLayout = [['Label'], ['Description'], ['KeyID']]
export { controls, newLayout, editLayout }
