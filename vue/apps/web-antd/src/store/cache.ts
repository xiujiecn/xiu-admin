import { ref } from "vue";
import { defineStore } from "pinia";
/** 缓存开关 */
const cacheFlag = true;
export const useCacheStore = defineStore("cache", () => {
  /** 缓存核心 */
  const cache = ref<Record<string, any>>({});

  /** 获取缓存数据
   * @param params 请求参数,将作为唯一key查询
   * @param label 缓存标签,查找的缓存库
   * @param fun 缓存数据获取方法,缓存数据不存在时自动调用该方法获取数据
   * @returns 目标数据
   */
  const getCache = async (params: Record<any, any>, label: string, fun: Promise<any> | (() => Promise<any>)) => {
    /** 关闭缓存时直接返回请求方法 */
    if (!cacheFlag) {
      await (typeof fun === 'function' ? fun() : fun)
    }

    /** debug使用 */
    const debugMsg: any = {
      label: true,
      data: true
    }
    // console.log("\n\n\n")


    const strParams = JSON.stringify(params);
    if (cache.value[label] == undefined || cache.value[label] == null) {
      // console.log("cache.value  [", label, "]  为空")
      cache.value[label] = {};
      debugMsg["label"] = false
    }
    if (cache.value[label][strParams] == undefined || cache.value[label][strParams] == null) {
      // console.log("cache.value  [", label, "]  [", strParams, "]  为空")
      cache.value[label][strParams] = await (typeof fun === 'function' ? fun() : fun);
      debugMsg["data"] = false
    }

    // console.log(label, "本地存储结构:", debugMsg["label"], "  本地存在数据:", debugMsg["data"], "  参数:", strParams)
    // if (debugMsg["data"] == false) { console.log("cache.value  [", label, "]  [", strParams, "]  已缓存到本地") }
    // else { console.log("cache.value  [", label, "]  [", strParams, "]  调用本地缓存") }
    return cache.value[label][strParams];
  }

  /**
   * 清空缓存,通常在修改数据后调用
   * @param label 要清空的缓存库
   */
  const clearCache = (label: string) => {
    // console.log("cache.value  [", label, "]  已清空")
    cache.value[label] = {};
  }


  /** 退出登录时清空缓存 */
  function $reset() {
    cache.value = {};
  }
  return {
    getCache,
    clearCache,
    $reset
  }
})
