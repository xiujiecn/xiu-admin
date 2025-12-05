<template>
  <VModal title="选择图片" class="h-[80%] w-[60%]" centered>
    <div class="w-full h-full flex flex-col gap-3">
      <div class="w-full h-full flex-1 flex flex-col gap-3">
        <!-- 操作栏 -->
        <div class="flex justify-between items-center pb-3 border-b border-solid">
          <Space>
            <Button type="primary" @click="() => { createFolderFlag = true; createFolderForm.name = '' }">新建文件夹</Button>
            <Button type="primary" @click="() => { uploadFlag = true; fileList = [] }">上传图片</Button>
            <Button @click="refreshData">刷新</Button>
          </Space>
          <Space size="large">
            <!-- <div>
              已用容量:
              <span class="text-red-500">{{ "1000MB" }}</span>
            </div>
            <div>
              剩余容量:
              <span class="text-green-500">{{ "1000MB" }}</span>
            </div> -->
            <InputSearch v-model:value="searchValue" placeholder="搜索图片" style="width: 200px"></InputSearch>
            <Select v-model:value="viewMode" style="width: 120px">
              <Select.Option value="grid">网格视图</Select.Option>
              <Select.Option value="list">列表视图</Select.Option>
            </Select>
            <Popconfirm title="你确定要删除吗？" :disabled="selectedRowKeys.length <= 0" @confirm="deleteItem(selectedRowKeys)">
              <Button type="primary" danger :disabled="selectedRowKeys.length <= 0">
                批量删除
              </Button>
            </Popconfirm>
          </Space>
        </div>
        <!-- 面包屑 -->
        <Breadcrumb>
          <BreadcrumbItem>
            <a @click="navigateToFolder({ displayName: '根目录', galleryId: 0 }, 'backward')">根目录</a>
          </BreadcrumbItem>
          <BreadcrumbItem v-for="folder in breadcrumbPath" :key="folder.galleryId">
            <a @click="navigateToFolder(folder, 'backward')">{{ folder.displayName }}</a>
          </BreadcrumbItem>
        </Breadcrumb>
        <!-- 文件夹 -->
        <div class="min-h-[400px]">
          <!-- 网格视图 -->
          <div v-if="viewMode == 'grid'" class="grid gap-4 grid-cols-[repeat(auto-fill_minmax(200px, 1fr))]"
            :style="{ gridTemplateColumns: 'repeat(auto-fill,minmax(200px, 1fr))' }">
            <!-- 文件夹 -->
            <div v-for="folder in currentFolders" :key="folder.galleryId" @click="navigateToFolder(folder, 'forward')"
              @contextmenu="showContextMenu($event, 'folder', folder)" :style="{ transition: 'all 0.3s' }"
              class="grid-item h-[200px] border-[1px] border-solid border-slate-500 rounded-md p-4 cursor-pointer relative">
              <div class="flex flex-col justify-center  items-center gap-2 h-full">
                <AppIcon :size="66" class="text-blue-400" icon="FolderOutlined"></AppIcon>
                <div class="font-bold break-all overflow-hidden text-ellipsis">
                  <Tooltip autoAdjustOverflow :mouseEnterDelay="1">
                    <template #title>
                      <div>{{ folder.displayName }}</div>
                    </template>
                    {{ folder.displayName }}
                  </Tooltip>
                </div>
                <div class="text-sm text-slate-500">{{ folder.imageCount }} 张图片</div>
              </div>
            </div>
            <!-- 图片 -->
            <div v-for="image in currentImages" :key="image.galleryId" :style="{ transition: 'all 0.3s' }"
              class="grid-item  h-[200px] border-[1px] border-solid border-slate-500 rounded-md p-4 cursor-pointer relative"
              :class="{ 'selected': selectedImage == image.galleryId }" @click="imageSelection(image)"
              @contextmenu="showContextMenu($event, 'image', image)">
              <div class="flex flex-col justify-center items-center gap-2 h-full">
                <img :src="image.filePath" :alt="image.name" class="flex-1 min-h-0 rounded-sm object-cover"></img>
                <div class="font-bold break-all overflow-hidden text-ellipsis">{{ image.displayName }}</div>
                <div class="text-sm text-slate-500">{{ formatFileSize(image.fileSize) }}</div>
                <div v-if="selectedImage == image.galleryId"
                  class="absolute top-2 right-2 bg-blue-500 text-blue-300 rounded-full w-6 h-6 flex items-center justify-center">
                  <AppIcon :size="16" icon="CheckCircle"></AppIcon>
                </div>
              </div>
            </div>
          </div>
          <!-- 列表视图 -->
          <Table v-else :columns="columns" :data-source="allItems" :pagination="false" rowKey="galleryId"
            :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: (sk) => { selectedRowKeys = sk } }">
            <template #bodyCell="{ column, record }">
              <template v-if="column.key === 'preview'">
                <img v-if="record.type === 'image'" :src="record.filePath" class="w-10 h-10 object-cover rounded-sm" />
                <AppIcon v-else :size="40" class="text-blue-400" icon="FolderOutlined"></AppIcon>
              </template>
              <template v-if="column.key === 'displayName'">
                <a @click="record.type === 'folder' ? navigateToFolder(record, 'forward') : imageSelection(record)">
                  {{ record.displayName }}
                </a>
              </template>
              <template v-if="column.key === 'fileSize'">
                {{ record.type === 'image' ? formatFileSize(record.fileSize) : '-' }}
              </template>
              <template v-if="column.key === 'actions'">
                <Space>
                  <Button size="small" @click="openEdit(record)">重命名</Button>
                  <Popconfirm title="你确定要删除吗？" @confirm="deleteItem([record.galleryId])">
                    <Button size="small" danger>删除</Button>
                  </Popconfirm>
                </Space>
              </template>
            </template>
          </Table>
        </div>
      </div>
    </div>
  </VModal>

  <!-- 新建文件夹模态框 -->
  <Modal v-model:open="createFolderFlag" title="新建文件夹" @cancel="createFolderFlag = false" @ok="createFolder">
    <Form :model="createFolderForm" layout="vertical">
      <FormItem label="文件夹名称" required>
        <Input v-model:value="createFolderForm.name" placeholder="请输入文件夹名称" />
      </FormItem>
    </Form>
  </Modal>
  <!-- 重命名模态框 -->
  <Modal v-model:open="editFlag" title="重命名" @cancel="editFlag = false" @ok="editItem">
    <Form :model="editForm" layout="vertical">
      <FormItem label="文件夹名称" required>
        <Input v-model:value="editForm.name" placeholder="请输入文件夹名称" />
      </FormItem>
    </Form>
  </Modal>
  <!-- 上传图片模态框 -->
  <Modal v-model:open="uploadFlag" title="上传图片" @cancel="uploadFlag = false" @ok="uploadImages">
    <UploadDragger v-model:file-list="fileList" :before-upload="beforeUpload" :multiple="true" accept="image/*"
      :show-upload-list="true">
      <p class="ant-upload-drag-icon flex item-center justify-center">
        <AppIcon class="text-neutral-500 dark:text-neutral-600" :size="66" icon="FolderOpenOutlined"></AppIcon>
      </p>
      <p class="ant-upload-text">点击或拖拽图片到此区域上传</p>
      <p class="ant-upload-hint">支持单个或批量上传，支持常见图片格式</p>
    </UploadDragger>
  </Modal>

  <!-- 右键菜单 -->
  <Dropdown>

  </Dropdown>
</template>

<script setup lang='ts' name=''>
import AppIcon from '../icon/AppIcon.vue';
import {
  Modal, Form, FormItem, Input, UploadDragger, Tooltip, Dropdown, Table,
  Button, Space, InputSearch, Select, Breadcrumb, BreadcrumbItem, message, Popconfirm
} from 'ant-design-vue';
import { useVbenModal } from '@vben/common-ui'
import { computed, ref } from 'vue';
import { List, Create, Rename, Upload, Delete } from '#/api/utils';
const emit = defineEmits(['confirm'])
const { fileType, filePath } = defineProps({
  fileType: {
    type: Number,
    default: 1
  },
  filePath: {
    type: String,
    default: ''
  }
})

/** 视图模式 */
const viewMode = ref<'list' | 'grid'>('list');
/** 搜索条件 */
const searchValue = ref('');

/** 当前文件ID */
const currentFolderId = ref<number>(0)
/** 选中的对象 */
const selectItems = ref<any>()
/** 选中的图片 */
const selectedImage = ref<any>()
/** 列表模式选中内容 */
const selectedRowKeys = ref<any>([])

/** 文件夹列表 */
const foldersList = ref<any>([])
/** 图片列表 */
const imageList = ref<any>([])
/** 面包屑列表 */
const breadcrumbPath = ref<any>([])

/** 当前路径文件夹列表 */
const currentFolders = computed(() => {
  return foldersList.value.filter((folder: any) => folder.parentId === currentFolderId.value && folder.displayName.toLowerCase().includes(searchValue.value.toLowerCase()))
})
/** 当前路径图片列表 */
const currentImages = computed(() => {
  return imageList.value.filter((image: any) => image.parentId === currentFolderId.value && image.displayName.toLowerCase().includes(searchValue.value.toLowerCase()))
})
/** 所有文件列表 */
const allItems = computed(() => {
  const items = [
    ...currentFolders.value.map((f: any) => ({ ...f, type: 'folder' })),
    ...currentImages.value.map((i: any) => ({ ...i, type: 'image' }))
  ]
  return items
})

/** 表格列配置 */
const columns = [
  { title: '预览', key: 'preview', width: 80 },
  { title: '名称', key: 'displayName', dataIndex: 'displayName' },
  { title: '大小', key: 'fileSize', dataIndex: 'fileSize' },
  { title: '上传时间', key: 'createdAt', dataIndex: 'createdAt' },
  { title: '操作', key: 'actions', width: 120 }
]

/** 核心模态框实例 */
const [VModal, modalApi] = useVbenModal({
  confirmText: '确定选择',
  onConfirm: () => {
    // 列表模式
    if (viewMode.value === 'list') {
      // 单选
      if (selectedRowKeys.value.length === 1) {
        // 过滤图片
        const selected = imageList.value.find((item: any) => item.galleryId === selectedRowKeys.value[0])
        if (selected == undefined) {
          message.error('请选择一张图片')
          return
        } else {
          emit('confirm', selected.filePath)
          modalApi.close()
          return
        }
      } else {
        message.error('请选择一张图片')
        return
      }
    } else {
      if (selectedImage.value != null) {
        const selected = imageList.value.find((item: any) => item.galleryId === selectedImage.value)
        emit('confirm', selected.filePath)
        modalApi.close()
        return
      } else {
        message.error('请选择一张图片')
        return
      }
    }
  }
});


// ========== API ==========
/** 获取文件夹列表 */
async function getFolders() {
  const res = await List({
    parentId: currentFolderId.value == null ? 0 : currentFolderId.value,
    type: 'folder'
  })
  foldersList.value = res.list
}
/** 获取图片列表 */
async function getImages() {
  const res = await List({
    parentId: currentFolderId.value == null ? 0 : currentFolderId.value,
    type: 'image'
  })
  imageList.value = res.list
}
/** 刷新数据 */
async function refreshData() {
  await getFolders()
  await getImages()
}
refreshData()


// ========== 方法 ==========
/** 切换文件夹 */
function navigateToFolder(folder: Record<string, any>, direction: 'forward' | 'backward') {
  // 判断是否是返回
  if (direction === 'backward') {
    // 找到返回的目录
    const find = breadcrumbPath.value.find((path: any) => path.galleryId === folder.galleryId)
    if (find) {
      // 存在,截取路径
      breadcrumbPath.value = breadcrumbPath.value.slice(0, breadcrumbPath.value.indexOf(find) + 1)
    } else {
      // 不存在,返回上一级
      breadcrumbPath.value = breadcrumbPath.value.slice(0, breadcrumbPath.value.length - 1)
    }
    // 如果是根目录,则清空路径
    if (folder.galleryId === 0) {
      breadcrumbPath.value = []
    }
  } else {
    // 将最新目录加入路径
    breadcrumbPath.value.push({
      galleryId: folder.galleryId,
      displayName: folder.displayName
    })
  }
  // 设置激活的文件夹
  currentFolderId.value = folder.galleryId
  refreshData()
}

/** 图片选择 */
function imageSelection(image: any) {
  if (selectedImage.value == image.galleryId) {
    selectedImage.value = null
  } else {
    selectedImage.value = image.galleryId
  }
}

/** 格式化文件大小 */
function formatFileSize(bytes: number) {
  if (!bytes) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}


// ========== 新建文件夹 ==========
/** 模态框flag */
const createFolderFlag = ref(false)
/** 新建文件夹表单数据 */
const createFolderForm = ref({
  name: '',
})
/** 新建文件夹提交函数 */
async function createFolder() {
  // 验证
  if (!createFolderForm.value.name.trim()) {
    message.error('请输入文件夹名称')
    return
  }
  // 请求
  const res = await Create({
    parentId: currentFolderId.value,
    displayName: createFolderForm.value.name.trim(),
  })
  if (res) {
    createFolderFlag.value = false
    message.success('文件夹创建成功')
    refreshData()
  }
}


// ========== 编辑/重命名 ==========
const editFlag = ref(false)
const editForm = ref({
  name: '',
})
function openEdit(item: any) {
  editFlag.value = true

  selectItems.value = item
  editForm.value = {
    name: item.displayName,
  }
}
async function editItem() {
  // 验证
  if (!editForm.value.name.trim()) {
    message.error('请输入名称')
    return
  }
  // 请求
  const res = await Rename({
    galleryId: selectItems.value.galleryId,
    displayName: editForm.value.name.trim(),
  })
  if (JSON.stringify(res) === '{}') {
    message.success('重命名成功')
    editFlag.value = false
    selectItems.value = null
    editForm.value = { name: '' }
    refreshData()
  }
}


// ========== 删除 ==========
async function deleteItem(items: any) {
  await Delete({
    galleryIds: [...items],
  })
  selectedRowKeys.value = []
  message.success('删除成功')
  refreshData()
}


// ========== 上传图片 ==========
/** 模态框flag */
const uploadFlag = ref(false)
/** 上传图片列表 */
const fileList = ref<any>([])

/** 上传检查函数 */
function beforeUpload(file: any) {
  const isImage = file.type.startsWith('image/')
  if (!isImage) {
    message.error('请上传图片')
    return false
  }
  const isLt10M = file.size / 1024 / 1024 < 10
  if (!isLt10M) {
    message.error('图片大小不能超过10M')
    return false
  }
  // 阻止自动上传
  return false
}

/** 上传函数 */
async function uploadImages() {
  if (fileList.value == 0) {
    message.error('请选择图片')
    return
  }
  // 处理图片数据
  fileList.value.forEach(async (fileItem: any) => {
    const file = fileItem.originFileObj || fileItem
    if (!file || !(file instanceof File)) {
      message.error('无效的文件对象')
      return
    }
    // 上传
    await Upload({
      parentId: currentFolderId.value,
      file,
      newFileType: fileType,
      ...(filePath ? { subDirName: filePath } : {})
    })
    refreshData()
  });
  uploadFlag.value = false
  message.success(`成功上传 ${fileList.value.length} 张图片`)
}

// ========== 右键菜单 ==========
function showContextMenu(event: any, type: any, item: any) {
  event.preventDefault()
}


</script>

<style scoped lang="scss">
/** 悬浮 */
.grid-item:hover {
  border-color: #1890ff;
  box-shadow: 0 4px 12px rgba(24, 144, 255, 0.15);
  transform: translateY(-2px);
}

/** 选中 */
.grid-item.selected {
  border-color: #1890ff;
  background: #b9b9b955;
}
</style>
