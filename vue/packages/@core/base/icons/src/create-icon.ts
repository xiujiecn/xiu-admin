import { defineComponent, h } from 'vue';

import { addIcon, Icon, type IconifyIcon } from '@iconify/vue';

function createIconifyIcon(icon: string) {
  return defineComponent({
    name: `Icon-${icon}`,
    setup(props, { attrs }) {
      return () => h(Icon, { icon, ...props, ...attrs });
    },
  });
}

function createIconifyOfflineIcon(icon: string, iconComponent: IconifyIcon) {
  return defineComponent({
    name: `Icon-${icon}`,
    setup(props, { attrs }) {
      addIcon(icon, iconComponent);
      return () => h(Icon, { icon, ...props, ...attrs });
    },
  });
}

export { createIconifyIcon, createIconifyOfflineIcon };
