/**
 * 图片URL处理工具函数
 */

/**
 * 处理图片URL，确保能正确显示
 * @param url 原始图片URL
 * @returns 处理后的图片URL
 */
export function processImageUrl(url: string | null | undefined): string {
  if (!url) {
    return '/ProductPictures.png'; // 默认图片
  }

  // 如果是完整的URL（以http开头），直接返回
  if (url.startsWith('http')) {
    return url;
  }

  // 如果图片路径以 /upload 开头，直接返回（会被代理转发）
  if (url.startsWith('/upload')) {
    return url;
  }

  // 如果是 app/upload/... 格式，替换为 /upload/...
  if (url.startsWith('app/upload/')) {
    const result = url.replace('app/upload/', '/upload/');
    return result;
  }

  // 如果是其他相对路径，添加 /upload 前缀
  const result = `/upload${url.startsWith('/') ? '' : '/'}${url}`;
  return result;
}

/**
 * 处理图片加载失败的情况
 * @param event 错误事件
 */
export function handleImageError(event: Event) {
  const img = event.target as HTMLImageElement;
  if (img.src !== '/ProductPictures.png') {
    img.src = '/ProductPictures.png';
  }
}
