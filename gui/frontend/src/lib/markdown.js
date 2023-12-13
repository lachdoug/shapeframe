import { marked } from 'marked'
import DOMPurify from 'dompurify'

export default text => DOMPurify.sanitize(marked.parse(text))
