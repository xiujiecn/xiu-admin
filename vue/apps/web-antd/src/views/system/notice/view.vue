<script setup lang="ts">
import { useVbenDrawer } from '@vben/common-ui';
import { cloneDeep } from '@vben/utils';
import { Description, useDescription } from '#/components/description';
import { viewSchema } from './model';
import { getSysNoticeApi } from '#/api/system/notice';
import { getSysUserListApi } from '#/api/system/user';
import type { AnyFunction } from '@vben/types';
const [BasicDrawer, drawerApi] = useVbenDrawer({
  onOpenChange: handleOpenChange,
});
const [registerDescription, { setDescProps }] = useDescription({
  column: 1,
  schema: viewSchema,
});
async function handleOpenChange(open: boolean) {
  if (!open) {
    return null;
  }
  const { id, userList, deptList } = drawerApi.getData();
  const record2 = await getSysNoticeApi({ noticeId: id });
  let currentSchema = cloneDeep(viewSchema);

  // if (userList) {
  //   record2.userIdList.forEach(item => {
  //     const user = userList.find((item2: any) => item2.userId === item)
  //     if (user)
  //       ul.push(
  //         user.nickName
  //       )
  //   })
  //   record2.userIdList = ul
  // }
  if (record2.userIdList) {
    const ul: any[] = []
    record2.userIdList = record2.userIdList.map((item: any) => {
      const user = userList.find((item2: any) => item2.userId === item)
      if (user)
        ul.push(user.nickName)
    })
    record2.userIdList = ul
  }

  if (record2.deptIdList) {
    const dl: any = []
    record2.deptIdList.forEach((item: any) => {
      dl.push(findFullName(deptList, item))
    })
    record2.deptIdList = dl
  }
  setDescProps({ data: record2, schema: currentSchema }, true);
}
function findFullName(tree: any, deptId: any): any {
  if (!tree) return null;
  // 如果是数组，遍历每个子树
  if (Array.isArray(tree)) {
    for (const node of tree) {
      const result = findFullName(node, deptId);
      if (result) return result;
    }
    return null;
  }
  // 检查当前节点
  if (tree.deptId === deptId) {
    return tree.fullName;
  }
  // 递归子节点
  if (tree.children && tree.children.length > 0) {
    for (const child of tree.children) {
      const result = findFullName(child, deptId);
      if (result) return result;
    }
  }
  return null;
}
</script>
<template>
  <BasicDrawer :footer="false" class="w-[600px]" title="查看">
    <Description @register="registerDescription">

    </Description>
  </BasicDrawer>
</template>
