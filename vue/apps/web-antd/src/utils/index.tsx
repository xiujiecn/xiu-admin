import { h, unref } from 'vue';
import { Tag, Popover, Button, Table, Avatar, } from 'ant-design-vue';
export * from './render';
import { renderIcon } from './render';

export interface MemberSumma {
  userId: number;
  userName: string;
  nickName: string;
  avatar: string;
  tenantId: string;
  deptId: number;
}

// render 操作人摘要
export const renderPopoverMemberSumma = (member: MemberSumma | null | undefined) => {
    if (!member) {
      return '';
    }
    // return '';
    const columns = [
      {title: '用户ID', dataIndex: 'userId', key: 'userId'},
      {title: '头像', dataIndex: 'avatar', key: 'avatar',
        customRender({text,record}: any) {
          return h(Avatar, {
              size: 'small',
              src: record.avatar,
          });
        },
      },
      {title: '昵称', dataIndex: 'nickName', key: 'nickName'},
      {title: '用户名', dataIndex: 'userName', key: 'userName'},
    ];
    const dataSource = [
      {
        userId: member.userId,
        avatar: member.avatar,
        nickName: member.nickName,
        userName: member.userName,
      },
    ];
    return h(
      Popover,
      { trigger: 'hover' },
      {
        default: () =>
          h(
            Button,
            {
              strong: true,
              size: 'small',
              text: true,
              iconPlacement: 'right',
            },
            { default: () => member.nickName,}
          ),
        content: () =>
          h(
            Table,
            {
                bordered: false,
                'single-line': false,
                pagination: false,
                size: 'small',
                columns,
                dataSource,
            }
          ),
      }
    );
  };