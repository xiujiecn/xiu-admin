import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Xiu Admin",
  description: "XiuAdmin基于全新GoFrame2+Vue3+VueVbenAdmin开发的全栖框架，适合企业级完整应用开发。",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: '主页', link: '/' },
      { text: '文档', link: '/guide' }
    ],

    sidebar: [

          { text: '简介',
            collapsed: false,
            items: [
              { text: '平台简介', link: 'guide' },
              { text: '演示图例', link: 'guide#演示图例' },
              { text: '跟新日志', link: 'changlog' },
            ]
          },
          {
            text: '快速开始',
            collapsed: false,
            items: [
              { text: '环境部署', link: 'start/environment' },
              { text: '项目启动', link: 'start/install' },
              { text: '生产部署', link: 'start/deploy' },
              { text: '常见问题', link: 'start/issue' },
            ]
          },
          {
            text: '系统开发',
            collapsed: false,
            items: [
              { text: '目录结构', link: 'sys/catalog' }
            ]
          }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/xiujiecn/xiu-admin' }
    ],
    footer: {
      message: '基于 Apache2.0 许可发布',
      copyright: `版权所有 © 2019-${new Date().getFullYear()} 李秀杰`
    },
  }
})
