/**
 * @description 设备管理类型定义
 * @Author  HuanTian
 * @date 2025-01-08
 */

/** 设备查询表单 */
export interface FormType {
  /** 设备组 */
  group: boolean;
  /** 设备查询 */
  query: boolean;
  /** 设备类型 */
  type: string;
  /** 在线状态 */
  online: string;
  /** 活跃状态 */
  active: string;
  /** 告警状态 */
  alert: string;
  /** 接入类型 */
  access: string;
  /** 搜索设备 */
  search: string;
  /** 排序方式 */
  sort: string;
  /** 展示模式: 1-网格 2-列表 3-地图 */
  mode: 1 | 2 | 3;
}
