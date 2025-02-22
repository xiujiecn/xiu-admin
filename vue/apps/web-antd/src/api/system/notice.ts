import { requestClient } from '#/api/request';

export interface SysNoticeListReq {
  noticeTitle: string;
  noticeType: string;
  createdBy: string;
}

export interface SysNotice {
    noticeId: number;
    tenantId:string;
    noticeTitle:string;
    noticeType:string;
    noticeContent:string;
    status:string;
    createdDept:number;
    createdBy:number;
    createdAt:string;
    remark:string;
}

export interface SysNoticeListRes {
  items: SysNotice[];
  total: number;
}

export async function getSysNoticeListApi(params: SysNoticeListReq) {
  return requestClient.get<SysNoticeListRes>('/system/notice/list', { params });
}
