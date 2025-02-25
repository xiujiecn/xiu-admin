/**
 * Returns the parent node of the given element or the document body if the element is not provided.it
 */
export function getPopupContainer(node?: HTMLElement): HTMLElement {
  return (node?.parentNode as HTMLElement) ?? document.body;
}

export function getVxePopupContainer(
  _node?: HTMLElement,
  id?: string,
): HTMLElement {
  let selector = 'div.vxe-table--body-wrapper.body--wrapper';
  if (id) {
    selector = `div#${id} ${selector}`;
  }
  // 挂载到vxe-table的滚动区域
  const vxeTableContainerNode = document.querySelector(selector);
  if (!vxeTableContainerNode) {
    console.warn('无法找到vxe-table元素, 将会挂载到body.');
    return document.body;
  }
  return vxeTableContainerNode as HTMLElement;
}
