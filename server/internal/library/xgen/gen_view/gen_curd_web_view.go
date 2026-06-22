// package genview
// @Link  https://github.com/xiujiecn/xiu-admin
// @Copyright  Copyright (c) 2025 LiXiujie
// @Author  Lxj <li@xiujie.cn>
// @License  https://github.com/xiujiecn/xiu-admin/blob/master/LICENSE
package genview

import (
	"bytes"
	"context"
	"fmt"
	genmodel "xiuadmin/internal/library/xgen/gen_model"

	"github.com/gogf/gf/v2/frame/g"
)

func (l *gCurd) webViewTplData(ctx context.Context, in *genmodel.CurdPreviewParam) (data g.Map, err error) {
	data = make(g.Map)
	data["item"] = l.generateWebViewItem(ctx, in)
	return
}

func (l *gCurd) generateWebViewItem(ctx context.Context, in *genmodel.CurdPreviewParam) string {
	buffer := bytes.NewBuffer(nil)
	for _, field := range in.MasterFields {
		if !field.IsEdit {
			continue
		}

		var (
			defaultComponent = fmt.Sprintf(`<DescriptionsItem label="%s">{{ formValue.%s }}</DescriptionsItem>`, field.Dc, field.TsName)
			component        string
		)

		switch field.FormMode {
		case FMInputTextarea, FMInputEditor:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <span v-html="formValue.%s"></span>
        </DescriptionsItem>`, field.Dc, field.TsName)

		case FMInputDynamic:
			component = defaultComponent

		case FMDate:
			component = defaultComponent

		case FMTime:
			component = defaultComponent

		case FMRadio, FMSelect:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <component :is="renderDict(String(formValue.%s), '%s')" />
        </DescriptionsItem>`, field.Dc, field.TsName, in.Options.DictMap[field.TsName])

		case FMCheckbox, FMSelectMultiple:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <component :is="renderDictTags(formValue.%s, '%s')" />
        </DescriptionsItem>`, field.Dc, field.TsName, in.Options.DictMap[field.TsName])

		case FMUploadImage:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <Image style="margin-left: 10px; height: 100px; width: 100px" :src="formValue.%s" />
        </DescriptionsItem>`, field.Dc, field.TsName)

		case FMUploadImages:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <Image.PreviewGroup>
            <Space>
              <Image
                v-for="(item, key) in formValue.%s"
                :key="key"
                style="margin-left: 10px; height: 100px; width: 100px"
                :src="item"
              />
            </Space>
          </Image.PreviewGroup>
        </DescriptionsItem>`, field.Dc, field.TsName)

		case FMUploadFile:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <div class="upload-card" v-show="formValue.%s !== ''" @click="download(formValue.%s)">
            <div class="upload-card-item" style="height: 100px; width: 100px">
              <div class="upload-card-item-info">
                <div class="img-box">
                  <Avatar :style="fileAvatarCSS">{{ getFileExt(formValue.%s) }}</Avatar>
                </div>
              </div>
            </div>
          </div>
        </DescriptionsItem>`, field.Dc, field.TsName, field.TsName, field.TsName)

		case FMUploadFiles:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <div class="upload-card">
            <Space :size="0">
              <div
                class="upload-card-item"
                style="height: 100px; width: 100px"
                v-for="(item, key) in formValue.%s"
                :key="key"
              >
                <div class="upload-card-item-info">
                  <div class="img-box">
                    <Avatar :style="fileAvatarCSS" @click="download(item)">{{
                      getFileExt(item)
                    }}</Avatar>
                  </div>
                </div>
              </div>
            </Space>
          </div>
        </DescriptionsItem>`, field.Dc, field.TsName)

		case FMSwitch:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <Switch :checked="formValue.%s === 1" disabled />
        </DescriptionsItem>`, field.Dc, field.TsName)

		case FMRate:
			component = fmt.Sprintf(`<DescriptionsItem label="%s">
          <Rate :value="formValue.%s" disabled allow-half />
        </DescriptionsItem>`, field.Dc, field.TsName)

		default:
			component = defaultComponent
		}

		buffer.WriteString("        " + component + "\n\n")
	}
	return buffer.String()
}
