import {
  requiredValidation,
} from '/src/lib/fieldValidation'

const controls = {
  Label: {
    label: 'Label',
    rules: [requiredValidation],
  },
  Description: {
    label: 'Description',
    placeholder: 'Optional',
  },
  Token: {
    label: 'Token',
    as: 'secret',
    rules: [requiredValidation],
    autocomplete: 'off',
  },
}

const newControls = controls
const newLayout = [['Label'], ['Description'], ['Token']]
const editControls = {
  Label: controls.Label,
  Description: controls.Description,
}
const editLayout = [['Label'], ['Description']]
const tokenControls = {
  Token: controls.Token,
}
const tokenLayout = [['Token']]

export {
  newControls,
  newLayout,
  editControls,
  editLayout,
  tokenControls,
  tokenLayout,
}
