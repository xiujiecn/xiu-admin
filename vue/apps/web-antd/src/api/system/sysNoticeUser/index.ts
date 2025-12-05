import { requestClient } from '#/api/request';

// 获取用户通知公告表列表
export function List(params:any) {
  return requestClient.get<any>('system/sysNoticeUser/list', { params });
}

// 删除/批量删除用户通知公告表
export function Delete(params:any) {
  return requestClient.post<any>('system/sysNoticeUser/delete', { ...params });
}

// 添加/编辑用户通知公告表
export function Edit(params:any) {
  return requestClient.post<any>('system/sysNoticeUser/edit', { ...params });
}

// 修改用户通知公告表状态
export function Status(params:any) {
  return requestClient.post<any>('system/sysNoticeUser/status', { ...params });
}

// 获取用户通知公告表指定详情
export function View(params:any) {
  return requestClient.get<any>('system/sysNoticeUser/view', { params });
}

// 导出用户通知公告表
export function Export(params:any) {
  return requestClient.post<Blob>('/system/sysNoticeUser/export',  { ...params } , {
    headers: { 'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8' },
    responseType: 'blob',
  });
}
/** 已读 */
export function Read(params:any) {
  return requestClient.post<any>('system/sysNoticeUser/read', { ...params });
}