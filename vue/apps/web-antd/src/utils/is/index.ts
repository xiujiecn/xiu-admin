// 判断值是否未某个类型
export function is(val: unknown, type: string) {
  return toString.call(val) === `[object ${type}]`;
}
// 判断值是否为函数
export function isFunction<T = Function>(val: unknown): val is T {
  return is(val, 'Function') || is(val, 'AsyncFunction');
}
// 判断值是否为JSON字符串
export function isJsonString(value: any) {
  try {
    const toObj = JSON.parse(value);
    if (toObj && typeof toObj === 'object') {
      return true;
    }
  } catch {}
  return false;
}
// 判断值是否为null
export function isNull(val: unknown): val is null {
  return val === null;
}
// 判断值是否为undefined
export const isDef = <T = unknown>(val?: T): val is T => {
  return typeof val !== 'undefined';
};
// 判断值是否为非undefined
export const isUnDef = <T = unknown>(val?: T): val is T => {
  return !isDef(val);
};
// 判断值是否为对象
export const isObject = (val: any): val is Record<any, any> => {
  return val !== null && is(val, 'Object');
};
// 判断值是否为null或undefined
export function isNullOrUnDef(val: unknown): val is null | undefined {
  return isUnDef(val) || isNull(val);
}
// 判断值是否为空对象
export function isNullObject(value: any) {
  return (
    isNullOrUnDef(value) ||
    JSON.stringify(value) === '{}' ||
    JSON.stringify(value) === '[]' ||
    value === ''
  );
}
