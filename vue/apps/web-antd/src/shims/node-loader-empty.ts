const noop = () => undefined;

const proxy = new Proxy(noop, {
  get() {
    return proxy;
  },
  apply() {
    return proxy;
  },
});

export const createJiti = () => proxy;
export default proxy;
