import type { SysDictData } from '#/api/system/dict';
import { reactive } from 'vue';

import { defineStore } from 'pinia';

export interface DictOption extends SysDictData {
  disabled?: boolean;
  label: string;
  value: number | string;
}

export function dictToOptions(
  data: SysDictData[],
  formatNumber = false,
): DictOption[] {
  return data.map((item) => ({
    ...item,
    label: item.dictLabel,
    value: formatNumber ? Number(item.dictValue) : item.dictValue,
  }));
}

export const useDictStore = defineStore('app-dict', () => {
  const dictOptionsMap = reactive(new Map<string, DictOption[]>());
  const dictRequestCache = reactive(
    new Map<string, Promise<SysDictData[] | void>>(),
  );
  function getDictOptions(dictName: string): DictOption[] {
    if (!dictName) return [];
    // 没有key 添加一个空数组
    if (!dictOptionsMap.has(dictName)) {
      dictOptionsMap.set(dictName, []);
    }
    // 这里拿到的就不可能为空了
    return dictOptionsMap.get(dictName)!;
  }
  function resetCache() {
    dictOptionsMap.clear();
  }
  function setDictInfo(
    dictName: string,
    dictValue: SysDictData[],
    formatNumber = false,
  ) {
    if (
      dictOptionsMap.has(dictName) &&
      dictOptionsMap.get(dictName)?.length === 0
    ) {
      dictOptionsMap
        .get(dictName)
        ?.push(...dictToOptions(dictValue, formatNumber));
    } else {
      dictOptionsMap.set(dictName, dictToOptions(dictValue, formatNumber));
    }
  }
  function $reset() {
    // nothing
  }
  return {
    $reset,
    dictOptionsMap,
    dictRequestCache,
    getDictOptions,
    resetCache,
    setDictInfo,
  };
});
