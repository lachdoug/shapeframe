import { requiredValidation, pathValidation } from '/src/lib/fieldValidation'

const controls = {
  Label: {
    label: 'Label',
    rules: [requiredValidation],
  },
  Description: {
    label: 'Description',
    placeholder: 'Optional',
  },
  RepositoryID: {
    label: 'Repository',
    as: 'select',
    cast: Number,
    options: {},
  },
  Directory: {
    label: 'Directory',
    placeholder: 'Director in the Git repository',
    rules: [pathValidation],
  },
}

const layout = [['Label'], ['Description'], ['RepositoryID', 'Directory']]
export { controls, layout }
