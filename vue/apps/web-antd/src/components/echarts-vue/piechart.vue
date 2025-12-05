<template>
    <div ref="chartRef" :style="{ width, height }"></div>
</template>
  
  <script setup>
  import { ref, onMounted, onBeforeUnmount, watch } from 'vue';
  import { EchartsUI, useEcharts, echarts } from '@vben/plugins/echarts';
  
  const props = defineProps({
    options: {
      type: Object,
      required: true,
    },
    width: {
      type: String,
      default: '100%',
    },
    height: {
      type: String,
      default: '100%',
    },
  });
  
  const chartRef = ref(null);
  let chartInstance = null;
  
  const initChart = () => {
    if (!chartRef.value) return;
    chartInstance = echarts.init(chartRef.value);
    chartInstance.setOption(props.options);
  };
  
  const handleResize = () => {
    chartInstance?.resize();
  };
  
  onMounted(() => {
    initChart();
    window.addEventListener('resize', handleResize);
  });
  
  onBeforeUnmount(() => {
    window.removeEventListener('resize', handleResize);
    chartInstance?.dispose();
  });
  
  watch(
    () => props.options,
    (newOptions) => {
      chartInstance?.setOption(newOptions);
    },
    { deep: true }
  );
  </script>