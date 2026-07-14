import { computed } from 'vue';

import { preferences, updatePreferences } from '@vben/preferences';
import { useAccessStore, useUserStore } from '@vben/stores';

/**
 * 检查权限码是否匹配，支持通配符 *
 * @param pattern 用户拥有的权限码（可能包含通配符 *）
 * @param code 需要校验的权限码
 * @returns 是否匹配
 */
function matchAccessCode(pattern: string, code: string): boolean {
  if (pattern === code) {
    return true;
  }
  if (pattern.endsWith('*')) {
    const prefix = pattern.slice(0, -1);
    return code.startsWith(prefix);
  }
  return false;
}

function useAccess() {
  const accessStore = useAccessStore();
  const userStore = useUserStore();
  const accessMode = computed(() => {
    return preferences.app.accessMode;
  });

  /**
   * 基于角色判断是否有权限
   * @description: Determine whether there is permission，The role is judged by the user's role
   * @param roles
   */
  function hasAccessByRoles(roles: string[]) {
    const userRoleSet = new Set(userStore.userRoles);
    const intersection = roles.filter((item) => userRoleSet.has(item));
    return intersection.length > 0;
  }

  /**
   * 基于权限码判断是否有权限
   * @description: Determine whether there is permission，The permission code is judged by the user's permission code
   * 支持通配符匹配，如用户拥有 "cpm:system:user:*" 可匹配 "cpm:system:user:query"
   * @param codes
   */
  function hasAccessByCodes(codes: string[]) {
    const userCodes = accessStore.accessCodes;

    const intersection = codes.filter((item) =>
      userCodes.some((userCode) => matchAccessCode(userCode, item)),
    );
    return intersection.length > 0;
  }

  async function toggleAccessMode() {
    updatePreferences({
      app: {
        accessMode:
          preferences.app.accessMode === 'frontend' ? 'backend' : 'frontend',
      },
    });
  }

  return {
    accessMode,
    hasAccessByCodes,
    hasAccessByRoles,
    toggleAccessMode,
  };
}

export { useAccess };
