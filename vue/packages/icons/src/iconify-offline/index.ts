import { createIconifyOfflineIcon } from '@vben-core/icons';
import userOutlined from '@iconify/icons-ant-design/user-outlined';
import comandLine from '@iconify/icons-flat-color-icons/command-line';
import redisIcon from '@iconify/icons-logos/redis';
import memoryIcon from '@iconify/icons-la/memory';
import taobaoIconFill from '@iconify/icons-ri/taobao-fill';
import alipayIcon from '@iconify/icons-fa-brands/alipay';
import dingdingFill from '@iconify/icons-ri/dingding-fill';
import giteeIcon from '@iconify/icons-simple-icons/gitee';
import macosIcon from '@iconify/icons-simple-icons/macos';
import githubOAuthIcon from '@iconify/icons-uiw/github';
import appleIcon from '@iconify/icons-uiw/apple';
import linuxIcon from '@iconify/icons-uiw/linux';
import wechatIcon from '@iconify/icons-uiw/linux';
import firefoxbrowserIcon from '@iconify/icons-simple-icons/firefoxbrowser';
import operaIcon from '@iconify/icons-uiw/opera';
import qqIcon from '@iconify/icons-uiw/qq';
import safariIcon from '@iconify/icons-uiw/safari';
import ucIcon from '@iconify/icons-arcticons/uc-browser';
import windowsIcon from '@iconify/icons-uiw/windows';

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
// 系统相关
export const AndroidIcon = createIconifyOfflineIcon(
  'ant-design:android-outlined',
  userOutlined,
);
export const BaiduIcon = createIconifyOfflineIcon(
  'ant-design:baidu-outlined',
  userOutlined,
);
export const ChromeIcon = createIconifyOfflineIcon(
  'ant-design:chrome-outlined',
  userOutlined,
);
export const DefaultBrowserIcon = createIconifyOfflineIcon(
  'ant-design:ie-outlined',
    userOutlined,
  );
export const DefaultOsIcon = createIconifyOfflineIcon(
  'ant-design:windows-outlined',
  userOutlined,
);
export const DingtalkIcon = createIconifyOfflineIcon(
  'ant-design:dingtalk-outlined',
  userOutlined,
);
export const EdgeIcon = createIconifyOfflineIcon(
  'ant-design:ie-outlined',
  userOutlined,
);
export const FirefoxIcon = createIconifyOfflineIcon(
  'simple-icons:firefoxbrowser',
  firefoxbrowserIcon,
);
export const IPhoneIcon = createIconifyOfflineIcon(
  'uiw:apple',
  appleIcon,
);
export const LinuxIcon = createIconifyOfflineIcon(
  'uiw:linux',
  linuxIcon,
);
export const MicromessengerIcon = createIconifyOfflineIcon(
    'uiw:weixin',
    wechatIcon,
  );
export const OperaIcon = createIconifyOfflineIcon(
  'uiw:opera',
  operaIcon,
);
export const OSXIcon = createIconifyOfflineIcon(
  'simple-icons:macos',
  macosIcon,
);
export const QQIcon = createIconifyOfflineIcon(
  'uiw:qq',
  qqIcon,
);

export const SafariIcon = createIconifyOfflineIcon(
  'uiw:safari',
  safariIcon,
);
export const UcIcon = createIconifyOfflineIcon(
  'arcticons:uc-browser',
  ucIcon,
);
export const WindowsIcon = createIconifyOfflineIcon(
  'uiw:windows',
  windowsIcon,
);
