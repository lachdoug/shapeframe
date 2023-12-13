import shapeValidation from './shapeValidation'
import frameValidation from './frameValidation'

const requiredValidation = v => !!v || 'Required field.'

const noWhitespaceValidation = v =>
  !v.match(/\s/g) || 'Must not include whitespace.'

const domainValidation = v =>
  !!v.match(/^[a-zA-Z0-9][a-zA-Z0-9-]{0,61}[a-zA-Z0-9](?:\.[a-zA-Z]{2,})+$/) ||
  'Must be a valid domain.'

const urlValidation = v => {
  try {
    new URL(v)
  } catch (e) {
    return 'Must be a valid URL'
  }
  return true
}

const pathValidation = v => !!v.match(/^[\w/\-. ]*$/) || 'Must be a valid path.'

export {
  shapeValidation,
  frameValidation,
  requiredValidation,
  noWhitespaceValidation,
  domainValidation,
  urlValidation,
  pathValidation,
}
