/** 文件类型
 *  @desc 1.文件需要加路径,图片不需要
 */
export const fileType = {
  // IOT平台
  /** IOT平台 - 用户文件 */
  UserFIle: 1,
  /** IOT平台 - 用户图片 */
  UserImg: 2,
  /** IOT平台 - 设备文件 */
  DeviceFile: 3,
  /** IOT平台 - Iot配置图片 */
  IotConfigImg: 4,
  /** IOT平台 - Iot配置文件 */
  IotConfigFile: 5,
  /** IOT平台 - 项目图片 */
  ProjectImg: 6,
  /** IOT平台 - 项目文件 */
  ProjectFile: 7,

  // 瀚臻平台
  /** 瀚臻平台 - 图库路径 */
  HzImg: 101,
  /** 瀚臻平台 - 固件路径 */
  HzFirmware: 102,
  /** 瀚臻平台 - 分组配置路径 */
  HzGroupConfig: 103,
  /** 瀚臻平台 - 设备配置路径 */
  HzDeviceConfig: 104,
  /** 瀚臻平台 - 设备文件 */
  HzDeviceFile: 105,
};

const FILEPATHTYPE = {
  /** ====================IOT平台==================== */
  // vue\apps\web-antd\src\views\_core\profile\profile-panel.vue
  /** 个人中心/个人信息面板 */
  profilPanel: fileType.UserImg,
  // vue\apps\web-antd\src\views\device\iotDevice\edit.vue
  /** 设备管理/设备列表 */
  iotDevice: fileType.IotConfigImg,
  // vue\apps\web-antd\src\views\device\iotProduct\edit.vue
  /** 设备管理/产品管理 */
  iotProduct: fileType.IotConfigImg,
  // vue\apps\web-antd\src\views\device\iotProductPreview\index.vue
  /** 设备管理/产品管理/产品详情 */
  iotProductPreview: fileType.IotConfigImg,
  // vue\apps\web-antd\src\views\device\iotProductPreview\Management\edit.vue
  /** 设备管理/产品管理/产品详情/固件管理 */
  iotProductPreviewManagement: fileType.IotConfigFile,
  // vue\apps\web-antd\src\views\manage\iotFirmware\index.vue
  /** 运维管理/产品固件 */
  iotFirmware: fileType.IotConfigFile,

  /** ====================瀚臻平台==================== */
  // vue\apps\web-antd\src\views\hzproductioncenter\hzPcProductInfo\edit.vue
  /** 瀚臻/产品图片 */
  hzPcProductInfo: fileType.HzImg,
  // vue\apps\web-antd\src\views\hzdevelopercenter\hzDcFirmware\edit.vue
  /** 瀚臻/升级固件 */
  hzDcFirmware: fileType.HzFirmware,
  // vue\apps\web-antd\src\views\hzproductioncenter\hzPcProductInfo\firmware\update\edit.vue
  /** 瀚臻/产品/升级固件 */
  hzPcProductInfoUpdate: fileType.HzFirmware,
  // vue\apps\web-antd\src\views\hzproductioncenter\hzPcProductInfo\firmware\flash\edit.vue
  /** 瀚臻/产品/固件下载 */
  hzPcProductInfoFlash: fileType.HzFirmware,
};
export default FILEPATHTYPE;

/** 模板 */
// <template #fileUrl="slotProps">
//   <div style="display: flex; flex-direction: column; gap: 8px">
//     <Upload :file-list="fileList" :before-upload="beforeUpload" :custom-request="customRequest"
//       accept=".bin,.hex,.fw,.zip" :multiple="false" list-type="text" @remove="handleRemove">
//       <Button :disabled="slotProps.disabled" type="primary">
//         <template #icon>
//           <svg viewBox="0 0 1024 1024" width="1em" height="1em" fill="currentColor">
//             <path
//               d="M400 317.7h73.9V656c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V317.7H624c6.7 0 10.4-7.7 6.3-12.9L518.3 163a8 8 0 0 0-12.6 0l-112 141.8c-4.1 5.2-.4 12.9 6.3 12.9z" />
//             <path
//               d="M878 626h-60c-4.4 0-8 3.6-8 8v154H214V634c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v198c0 17.7 14.3 32 32 32h684c17.7 0 32-14.3 32-32V634c0-4.4-3.6-8-8-8z" />
//           </svg>
//         </template>
//         上传固件文件
//       </Button>
//     </Upload>
//     <div style="color: #666; font-size: 12px; margin-top: 4px">
//       请上传不超过
//       <span style="color: #1890ff; font-weight: bold">100MB</span> 的
//       <span style="color: #1890ff; font-weight: bold">bin/hex/fw/zip</span>
//       格式文件
//     </div>
//     <div v-if="slotProps.value" style="color: #666; font-size: 12px">
//       当前文件: {{ slotProps.value }}
//     </div>
//   </div>
// </template>

// const productCode = (await formApi.getValues()).productCode;
// const response = await uploadApi(
//   file,
//   progressEvent,
//   undefined,
//   FILEPATHTYPE.hzPcProductInfoUpdate,
//   productCode,
// );

// dependencies: {
//   disabled(values) {
//     return !!values.field3Switch;
//   },
//   triggerFields: ['field3Switch'],
// },
