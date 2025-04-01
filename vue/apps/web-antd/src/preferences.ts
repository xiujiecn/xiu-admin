import { defineOverridesPreferences } from '@vben/preferences';

/**
 * @description 项目配置文件
 * 只需要覆盖项目中的一部分配置，不需要的配置不用覆盖，会自动使用默认配置
 * !!! 更改配置后请清空缓存，否则可能不生效
 */
export const overridesPreferences = defineOverridesPreferences({
  // overrides
  app: {
    name: import.meta.env.VITE_APP_TITLE,
    accessMode: 'backend',
  },
  theme: {
    mode: "light"
  },
  logo: {
    enable: true,
    source: "https://iot.xiujie.cn/wp-content/uploads/2025/02/96x96.webp",
  },
  copyright: {
    enable: true,
    companyName: "李秀杰",
    companySiteLink: "https://www.xiujiezhilian.cn",
    date: "2024-2025",
    icp: "鲁ICP备2024117944号-1",
    icpLink: "https://beian.miit.gov.cn/",
  },
});
