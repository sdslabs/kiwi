const { description } = require('../../package')

module.exports = {
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#title
   */
  title: 'Kiwi',
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#description
   */
  description: "A minimalistic in-memory key value store",

  /**
   * Extra tags to be injected to the page HTML `<head>`
   *
   * ref：https://v1.vuepress.vuejs.org/config/#head
   */
  head: [
    ['meta', { name: 'theme-color', content: '#A1CE48' }],
    ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }]
  ],

  /**
   * Theme configuration, here is the default theme configuration for VuePress.
   *
   * ref：https://v1.vuepress.vuejs.org/theme/default-theme-config.html
   */
  themeConfig: {
    repo: 'sdslabs/kiwi',
    editLinks: true,
    docsDir: 'docs',
    lastUpdated: 'Last Updated',
    smoothScroll: true,
    nav: [
      {
        text: 'Documentation',
        link: '/docs/',
      },
      {
        text: 'Go Reference',
        link: 'https://pkg.go.dev/github.com/sdslabs/kiwi'
      },
    ],
    sidebar: {
      '/docs/': [
        {
          title: 'Welcome',
          collapsable: false,
          children: [
            '',
            'get-started',
          ]
        }
      ],
    }
  },

  /**
   * Apply plugins，ref：https://v1.vuepress.vuejs.org/zh/plugin/
   */
  plugins: [
    '@vuepress/plugin-back-to-top',
    '@vuepress/plugin-medium-zoom',
  ]
}
