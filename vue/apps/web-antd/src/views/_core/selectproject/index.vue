<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';

import { Button, Result, Card } from 'ant-design-vue';

const router = useRouter();

function goToHome() {
  // 跳转到真正的主页
  router.push('/analytics');
}

const projectList = ref<any>([]);
</script>

<template>
  <div class="flex h-full flex-col gap-4 bg-background p-8 overflow-scroll">
    <div>
      <div class="flex items-end justify-between pb-4">
        <p class="text-2xl">我创建的项目</p>
        <div>
          <Button
            type="primary"
            @click="
              projectList.push({
                projectName: '项目1',
                projectId: '1',
                projectDescription: '项目描述1',
                projectStatus: '正常',
                projectDevice: 10,
                projectMember: 10,
              })
            "
          >
            创建项目
          </Button>
          <Button type="primary" @click="goToHome">测试</Button>
        </div>
      </div>
      <div
        v-if="projectList.length > 0"
        class="grid grid-cols-1 gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 2xl:grid-cols-6"
      >
        <Card hoverable v-for="(item, index) in projectList" :key="index">
          <p class="text-xl font-bold">
            {{ item.projectName }}
          </p>
          <p>
            {{ item.projectDescription }}
          </p>
          <template #actions>
            <div class="flex justify-between px-4">
              <div class="flex gap-2">
                <p>{{ item.projectDevice }}设备</p>
                <p>{{ item.projectMember }}成员</p>
              </div>
              <p>{{ item.projectStatus }}</p>
            </div>
          </template>
        </Card>
      </div>
      <Result
        v-else
        status="403"
        title="创建你的物联网项目"
        sub-title="支持各种单片机、嵌入式设备、传感器、执行器、控制器、网关、DTU、PLC 等，支持 MQTT、HTTP、TCP 自定义协议。"
      >
        <template #extra>
          <Button type="primary">创建项目</Button>
        </template>
      </Result>
    </div>
    <div>
      <div class="pb-4 text-2xl">我加入的项目</div>
      <Result
        status="403"
        title="没有加入的项目"
        sub-title="您可以作为管理员或访客，加入别人的项目。"
      />
    </div>
  </div>
</template>
