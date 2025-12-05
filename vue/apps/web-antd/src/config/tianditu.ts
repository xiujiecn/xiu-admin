/**
 * 天地图配置文件
 * 统一管理天地图API Key和相关配置
 */

// 天地图API Key - 统一配置
export const TIANDITU_CONFIG = {
  // 天地图API Key
  API_KEY: 'dff404b950f323fcfd1b87285092cef4',
  
  // 天地图API基础URL
  API_BASE_URL: 'https://api.tianditu.gov.cn',
  
  // 天地图JS API URL
  JS_API_URL: 'https://api.tianditu.gov.cn/api',
  
  // 默认地图配置
  DEFAULT_CONFIG: {
    center: [116.921280, 36.650390] as [number, number], // 默认中心点（济南）
    zoom: 15,
    mapType: 'vec' as 'vec' | 'img' | 'ter',
  },
  
  // 搜索API配置
  SEARCH_CONFIG: {
    maxResults: 10,
    level: 18,
  }
};

/**
 * 获取天地图API Key
 */
export function getTiandituApiKey(): string {
  return TIANDITU_CONFIG.API_KEY;
}

/**
 * 获取天地图JS API URL
 */
export function getTiandituJsApiUrl(): string {
  return `${TIANDITU_CONFIG.JS_API_URL}?v=4.0&tk=${TIANDITU_CONFIG.API_KEY}`;
}

/**
 * 获取天地图搜索API URL
 */
export function getTiandituSearchUrl(keyword: string): string {
  const { API_BASE_URL, API_KEY, SEARCH_CONFIG } = TIANDITU_CONFIG;
  const encodedKeyword = encodeURIComponent(keyword);

  const postStr = JSON.stringify({
    keyWord: keyword,
    level: SEARCH_CONFIG.level.toString(),
    mapBound: "0,0,0,0",
    queryType: "7",
    start: "0",
    count: SEARCH_CONFIG.maxResults.toString()
  });

  return `${API_BASE_URL}/v2/search?postStr=${encodeURIComponent(postStr)}&type=query&tk=${API_KEY}`;
}

/**
 * 获取天地图逆地理编码API URL
 * 根据经纬度获取地址信息
 */
export function getTiandituReverseGeocodeUrl(lng: number, lat: number): string {
  const { API_KEY } = TIANDITU_CONFIG;

  const postStr = JSON.stringify({
    lon: lng,
    lat: lat,
    ver: 1
  });

  return `https://api.tianditu.gov.cn/geocoder?postStr=${encodeURIComponent(postStr)}&type=geocode&tk=${API_KEY}`;
}

/**
 * 获取默认地图配置
 */
export function getDefaultMapConfig() {
  return {
    ...TIANDITU_CONFIG.DEFAULT_CONFIG,
    apiKey: TIANDITU_CONFIG.API_KEY,
  };
}
