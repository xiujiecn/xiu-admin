import { requestClient } from '#/api/request';


export interface CommandStats {
  name: string;
  value: string;
}

export interface RedisInfo {
  [key: string]: string;
}

export interface CacheInfo {
  commandStats: CommandStats[];
  dbSize: number;
  info: RedisInfo;
}

export async function redisCacheInfo() {
  return requestClient.get<CacheInfo>('/monitor/cache');
}
  