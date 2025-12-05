import { iconData } from '#/assets/icons';
import { h } from 'vue';

export interface Icon {
  /** 图标唯一键 (用于在配置中存储和查询) */
  key: string;
  /** 图标的 SVG 内容 (完整的 <svg>...</svg> 字符串) */
  svg: string;
  /** 图标显示名称 (用于在选择器中展示) */
  name?: string;
}

class IconManager {
  private static instance: IconManager;
  private icons: Icon[] = iconData;

  private constructor() {}

  public static getInstance(): IconManager {
    if (!IconManager.instance) {
      IconManager.instance = new IconManager();
    }
    return IconManager.instance;
  }
  /**
   * 根据 Key 获取单个图标
   * @param key 图标唯一键
   * @returns Icon | undefined
   */
  public getIconByKey(key: string): Icon | undefined {
    return this.icons.find((icon) => icon.key === key);
  }

  /**
   * 分页获取图标
   * @param page 页码 (从1开始)
   * @param pageSize 每页数量
   * @returns { icons: Icon[], total: number, page: number, pageSize: number }
   */
  public getIconsByPage(
    page: number = 1,
    pageSize: number = 20,
    filter?: string,
  ): {
    icons: Icon[];
    total: number;
    page: number;
    pageSize: number;
  } {
    const filteredIcons = filter
      ? this.icons.filter((icon) => (icon.key.toLowerCase()).includes(filter.toLowerCase()))
      : this.icons;
    // 计算分页
    const total = filteredIcons.length;
    const startIndex = (page - 1) * pageSize;
    const endIndex = startIndex + pageSize;
    const paginatedIcons = filteredIcons.slice(startIndex, endIndex);
    return {
      icons: paginatedIcons,
      total,
      page,
      pageSize,
    };
  }
}

export const iconManager = IconManager.getInstance();
