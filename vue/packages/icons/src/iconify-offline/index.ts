import { createIconifyOfflineIcon } from '@vben-core/icons';
import userOutlined from '@iconify/icons-ant-design/user-outlined';
import comandLine from '@iconify/icons-flat-color-icons/command-line';
import redisIcon from '@iconify/icons-logos/redis';
import memoryIcon from '@iconify/icons-la/memory';
import taobaoIconFill from '@iconify/icons-ri/taobao-fill';
import alipayIcon from '@iconify/icons-fa-brands/alipay';
import dingdingFill from '@iconify/icons-ri/dingding-fill';
import giteeIcon from '@iconify/icons-simple-icons/gitee';
import githubOAuthIcon from '@iconify/icons-uiw/github';
// 顶部菜单
export const UserOutlined = createIconifyOfflineIcon(
  'ant-design:user-outlined',
  userOutlined,
);
// 缓存监控使用
export const RedisIcon = createIconifyOfflineIcon('logos:redis', redisIcon);
export const CommandLineIcon = createIconifyOfflineIcon(
  'flat-color-icons:command-line',
  comandLine,
);
export const MemoryIcon = createIconifyOfflineIcon('la:memory', memoryIcon);

// 第三方登录相关图标
export const TaobaoIcon = createIconifyOfflineIcon(
  'ri:taobao-fill',
  taobaoIconFill,
);
export const AlipayIcon = createIconifyOfflineIcon(
  'fa-brands:alipay',
  alipayIcon,
);
export const DingdingIcon = createIconifyOfflineIcon(
  'ri:dingding-fill',
  dingdingFill,
);
export const GiteeIcon = createIconifyOfflineIcon(
  'simple-icons:gitee',
  giteeIcon,
);
export const GithubOAuthIcon = createIconifyOfflineIcon(
  'uiw:github',
  githubOAuthIcon,
);