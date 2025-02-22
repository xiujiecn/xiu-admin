import { createIconifyOfflineIcon } from '@vben-core/icons';
import comandLine from '@iconify/icons-flat-color-icons/command-line';
import redisIcon from '@iconify/icons-logos/redis';
import memoryIcon from '@iconify/icons-la/memory';

// 缓存监控使用
export const RedisIcon = createIconifyOfflineIcon('logos:redis', redisIcon);
export const CommandLineIcon = createIconifyOfflineIcon(
  'flat-color-icons:command-line',
  comandLine,
);
export const MemoryIcon = createIconifyOfflineIcon('la:memory', memoryIcon);
