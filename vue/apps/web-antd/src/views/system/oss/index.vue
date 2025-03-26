<script lang="ts" setup>
import { h, ref } from 'vue';
import type { VbenFormProps } from '#/adapter/form';
import type { VxeTableGridOptions, VxeGridListeners } from '#/adapter/vxe-table';
import type { SysOssViewModel } from '#/api/system/oss';
import { Page, useVbenModal } from '@vben/common-ui';

import { Button, message, Switch, Tag, Modal,Popconfirm, Tooltip ,Spin, Image } from 'ant-design-vue';
import dayjs from 'dayjs';
import { $t } from '@vben/locales';
import type { DeepPartial } from '@vben/types';
import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { getSysOssListApi, ossDownload, deleteSysOssApi } from '#/api/system/oss';
import { getVxePopupContainer } from '@vben/utils';
import { calculateFileSize } from '#/utils/file';
import { downloadByData } from '#/utils/file/download';
import {
  MdiPlus,
  MdiEdit,
  MdiDelete,
} from '@vben/icons';
import { useRouter } from 'vue-router';
import fileUploadModal from './file-upload-modal.vue';
import imageUploadModal from './image-upload-modal.vue';
import { fallbackImageBase64 } from './model';
interface RowType {
  category: string;
  color: string;
  id: string;
  price: string;
  productName: string;
  releaseDate: string;
}

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  fieldMappingTime: [['date', ['start', 'end']]],
  schema: [
    {
      component: 'Input',
      fieldName: 'fileName',
      label: '文件名称',
    },
    {
      component: 'Input',
      fieldName: 'originalName',
      label: '原名',
    },
    {
      component: 'Input',
      fieldName: 'fileSuffix',
      label: '文件后缀',
    },
    {
      component: 'Input',
      fieldName: 'service',
      label: '服务商',
    },
    {
      component: 'RangePicker',
      componentProps: {
        format: "YYYY-MM-DD",
        valueFormat: "YYYY-MM-DD",
      },
      // defaultValue: [dayjs().subtract(7, 'days'), dayjs()],
      fieldName: 'createdAt',
      label: '创建时间',
    },
  ],
  // 控制表单是否显示折叠按钮
  showCollapseButton: true,
  // 是否在字段值改变时提交表单
  submitOnChange: true,
  // 按下回车时是否提交表单
  submitOnEnter: false,
};

const gridOptions: VxeTableGridOptions<RowType> = {
  checkboxConfig: {
    highlight: true,
    labelField: 'infoId',
  },
  columns: [
    { align: 'left', title: 'ID', type: 'checkbox', width: 80 },
    { field: 'fileName', title: '文件名称' },
    { field: 'originalName', title: '原名', showOverflow: true,},
    { field: 'fileSuffix', title: '文件后缀' },
    { field: 'url', title: 'URL地址', slots: { default: 'url' },showOverflow: true,},
    { field: 'service', title: '服务商' },
    { field: 'createdAt', title: '创建时间' },
    { field: 'createdBy', title: '创建者' },
    { title: '操作', width: 80, slots: { default: 'action' } }
  ],
  exportConfig: {},
  height: 'auto',
  keepSource: true,
  pagerConfig: {},
  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        // message.success(`Query params: ${JSON.stringify(formValues)}`);
        return await getSysOssListApi({
          page: page.currentPage,
          pageSize: page.pageSize,
          ...formValues,
        });
      },
    },
  },
  rowConfig: {
    keyField: 'ossId',
    height: 65,
  },
  toolbarConfig: {
    custom: true,
    export: true,
    refresh: true,
    resizable: true,
    search: true,
    zoom: true,
  },
};


const gridEvents: DeepPartial<VxeGridListeners> = {
  checkboxChange: handleCheckboxChange,
  checkboxAll: handleCheckboxChange,
};

const CheckboxChecked = ref(false);
function handleCheckboxChange() {
  CheckboxChecked.value = tableApi.grid.getCheckboxRecords().length > 0;
}

const [Grid, tableApi] = useVbenVxeGrid({
  formOptions,
  gridOptions,
  gridEvents,
});


async function handleDownload(row: SysOssViewModel) {
  const downloadSize = ref($t('pages.common.downloadLoading'));
  const hideLoading = message.loading({
    content: () => downloadSize.value,
    duration: 0,
  });
  try {
    const data = await ossDownload(row.ossId, (e) => {
      // 计算下载进度
      const percent = Math.floor((e.loaded / e.total!) * 100);
      // 已经下载
      const current = calculateFileSize(e.loaded);
      // 总大小
      const total = calculateFileSize(e.total!);
      downloadSize.value = `已下载: ${current}/${total} (${percent}%)`;
    });
    downloadByData(data, row.originalName);
    message.success('下载完成');
  } finally {
    hideLoading();
  }
}

async function handleDelete(row: SysOssViewModel) {
  await deleteSysOssApi({ ossIds: [row.ossId] });
  await tableApi.query();
}

function handleMultiDelete() {
  const rows = tableApi.grid.getCheckboxRecords();
  const ids = rows.map((row: SysOssViewModel) => row.ossId);
  Modal.confirm({
    title: '提示',
    okType: 'danger',
    content: `确认删除选中的${ids.length}条记录吗？`,
    onOk: async () => {
      await deleteSysOssApi({ ossIds: ids });
      await tableApi.query();
    },
  });
}

const router = useRouter();
const handleClickOssConfig = () => {
  router.push(`/system/oss-config`);
}

const [ImageUploadModal, imageUploadApi] = useVbenModal({
  connectedComponent: imageUploadModal,
});

const [FileUploadModal, fileUploadApi] = useVbenModal({
  connectedComponent: fileUploadModal,
});
const preview = ref(false);
function isImageFile(ext: string) {
  const supportList = ['jpg', 'jpeg', 'png', 'gif', 'webp'];
  return supportList.some((item) => ext.toLocaleLowerCase().includes(item));
}

</script>

<template>
  <Page auto-content-height>
    <Grid table-title="文件列表">
      <template #toolbar-tools>
        <Tooltip title="预览图片" class="mr-2 flex items-center ">
            <Switch v-model:checked="preview" />
        </Tooltip>
        <Button class="mr-2 flex items-center " @click="handleClickOssConfig" v-access:code="'cpm:system:ossConfig:list'">OSS配置</Button>
        <Button class="mr-2 flex items-center " type="primary" @click="fileUploadApi.open" v-access:code="'cpm:system:oss:upload'">文件上传</Button>
        <Button class="mr-2 flex items-center " type="primary" @click="imageUploadApi.open" v-access:code="'cpm:system:oss:upload'">图片上传</Button>
        <Button class="mr-2 flex items-center " type="primary" @click="handleMultiDelete" v-access:code="'cpm:system:oss:remove'">删除</Button>
      </template>
      <template #url="{ row }">
        <!-- placeholder为图片未加载时显示的占位图 -->
        <!-- fallback为图片加载失败时显示 -->
        <!-- 需要设置key属性 否则切换翻页会有延迟 -->
        <Image
          :key="row.ossId"
          v-if="preview && isImageFile(row.url)"
          :src="row.url"
          height="50px"
          :fallback="fallbackImageBase64"
        >
          <template #placeholder>
            <div class="flex size-full items-center justify-center">
              <Spin />
            </div>
          </template>
        </Image>
        <span v-else>{{ row.url }}</span>
      </template>
      <template #status="{ row }">
        <Tag :color="row.status == '0' ? 'green' : 'red'">{{ row.status == '0' ? '正常' : '关闭' }}</Tag>
      </template>
      <template #action="{ row }">
        <div class="flex items-center">
          <Button class="mr-2 border-none p-0" :block="false" type="link" @click="handleDownload(row)" v-access:code="'cpm:system:oss:download'">下载</Button>
          <Popconfirm :get-popup-container="getVxePopupContainer" placement="left" title="确认删除？"
            @confirm="handleDelete(row)" v-access:code="'cpm:system:oss:remove'">
            <Button class="mr-2 border-none p-0" :block="false" type="link" danger v-access:code="'cpm:system:oss:remove'">删除</Button>
          </Popconfirm>
        </div>
      </template>
    </Grid>
    <ImageUploadModal @reload="tableApi.query" />
    <FileUploadModal @reload="tableApi.query" />
  </Page>
</template>
