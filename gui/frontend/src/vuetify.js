import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import { aliases, mdi } from 'vuetify/iconsets/mdi-svg'

export default createVuetify({
  components,
  directives,
  defaults: {
    VToolbar: {
      density: 'compact',
    },
  },
  theme: {
    themes: {
      // defaultTheme: 'dark',
      light: {
        colors: {
          primary: '#0000FF',
          // 'on-surface': '#AAAAFF',
        },
      },
    },
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi,
    },
  },
})
