import { defineConfig } from 'vitepress'
import type { DefaultTheme } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Xiu Admin",
  description: "XiuAdmin 基于 GoFrame2、Vue3、Vue Vben Admin 技术栈开发的全栈框架，适用于中后台应用开发。",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: '主页', link: '/' },
      { text: '文档', link: '/guide' },
      { text: '捐赠', link: '/donate' },
      { text: '关于我们', link: '/about' }
    ],

    sidebar: {
      '/guide/': sidebarGuide(),
      '/start/': sidebarGuide(),
      '/changlog/': sidebarGuide(),
      '/sys/': sidebarGuide(),
      '/code/': sidebarGuide(),
      '/web/': sidebarGuide()
    },
    socialLinks: [
      { icon: 'github', link: 'https://github.com/xiujiecn/xiu-admin' }
    ],
    footer: {
      message: '基于 Apache License 2.0 开源许可发布',
      copyright: `版权所有 © 2024-${new Date().getFullYear()} 李秀杰`
    },
  }
})
function sidebarGuide(): DefaultTheme.SidebarItem[] {
  return [
    { 
    text: '简介',
    collapsed: false,
    items: [
        { text: '平台简介', link: '/guide' },
        { text: '演示图例', link: '/guide#演示图例' },
        { text: '更新日志', link: '/changlog' },
      ]
    },
    {
      text: '快速开始',
      collapsed: false,
      items: [
        { text: '环境部署', link: '/start/environment' },
        { text: '系统安装', link: '/start/install' },
        { text: '生产部署', link: '/start/deploy' },
        { text: '常见问题', link: '/start/issue' },
      ]
    },
    {
      text: '系统开发',
      collapsed: false,
      items: [
        { text: '目录结构', link: '/sys/catalog' },
        { text: '开发规范', link: '/sys/exploit' },
        { text: '中间件', link: '/sys/middleware' },
        { text: '权限控制', link: '/sys/auth' },
        { text: '定时任务', link: '/sys/cron' },
        { text: '队列', link: '/sys/queue' },
        { text: '事件', link: '/sys/event' },
        { text: 'WebSocket', link: '/sys/websocket' },
      ]
    },
    {
      text: '代码生成',
      collapsed: false,
      items: [
        { text: '使用前提', link: '/code/start' },
        { text: '数据库', link: '/sys/db' },
        { text: '生成配置', link: '/code/config' },
        { text: '生成CURD', link: '/code/curd' },
        { text: '生成关联表', link: '/code/curd-join' },
        { text: '生成模板', link: '/code/template' },
        { text: '生成常见问题', link: '/code/help' },
      ]
    },
    {
      text: '前端开发',
      collapsed: false,
      items: [
        { text: '前端开发', link: '/web' },
      ]
    }
  ];
}