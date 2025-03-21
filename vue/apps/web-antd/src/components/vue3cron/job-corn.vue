<script setup lang="ts">
import { ref, computed, defineProps, defineEmits, reactive } from 'vue';
import Language from '#/components/vue3cron/language';
import { Card, Radio, RadioGroup, InputNumber, Select, SelectOption, Tag, Button, Row, Col} from  'ant-design-vue';
// 定义 Props
const props = defineProps({
  type: { type: String, default: '' },
  i18n: { type: String, default: 'cn' },
  maxHeight: { type: String, default: '300px' },
});

// 定义 Emits
const emit = defineEmits<{
(e: 'cancel'): void;
(e: 'confirm', value: { cron: string}): void;
}>();

// 定义状态
// 定义 Language 的键类型
type LanguageKey = keyof typeof Language;
const language = ref<LanguageKey>(props.i18n as LanguageKey);
const second = ref({ cronEvery: "1", incrementStart: 3, incrementIncrement: 5, rangeStart: 0, rangeEnd: 0, specificSpecific: [] });
const minute = ref({ cronEvery: '1', incrementStart: 3, incrementIncrement: 5, rangeStart: 0, rangeEnd: 0, specificSpecific: [] });
const hour = ref({ cronEvery: '1', incrementStart: 3, incrementIncrement: 5, rangeStart: 0, rangeEnd: 0, specificSpecific: [] });
const day = ref({ cronEvery: '1', incrementStart: 1, incrementIncrement: 1, rangeStart: 0, rangeEnd: 0, specificSpecific: [], cronLastSpecificDomDay: 1, cronDaysBeforeEomMinus: 0, cronDaysNearestWeekday: 0 });
const week = ref({ cronEvery: '1', incrementStart: 1, incrementIncrement: 1, specificSpecific: [], cronNthDayDay: 1, cronNthDayNth: 1 });
const month = ref({ cronEvery: '1', incrementStart: 3, incrementIncrement: 5, rangeStart: 0, rangeEnd: 0, specificSpecific: [] });
const year = ref({ cronEvery: '1', incrementStart: 2017, incrementIncrement: 1, rangeStart: 0, rangeEnd: 0, specificSpecific: [] });

// 计算属性
const text = computed(() => Language[language.value || 'cn']);
const secondsText = computed(() => {
let seconds = '';
console.log(second.value.cronEvery);
let val = String(second.value.cronEvery);
switch (val) {
  case '1': 
    seconds = '*'; console.log("1 seconds  ", seconds); break;
  case '2': 
    seconds = `${second.value.incrementStart}/${second.value.incrementIncrement}`;console.log("2 seconds  ", seconds); break;
  case '3': seconds = second.value.specificSpecific.join(','); console.log("3 seconds  ", seconds);break;
  case '4': seconds = `${second.value.rangeStart}-${second.value.rangeEnd}`; console.log("4 seconds  ", seconds);break;
  default: seconds = ''; console.log("default seconds  ", seconds);;
}
return seconds;
});
const minutesText = computed(() => {
let minutes = '';
switch (minute.value.cronEvery) {
  case '1': minutes = '*'; break;
  case '2': minutes = `${minute.value.incrementStart}/${minute.value.incrementIncrement}`; break;
  case '3': minutes = minute.value.specificSpecific.join(','); break;
  case '4': minutes = `${minute.value.rangeStart}-${minute.value.rangeEnd}`; break;
}
return minutes;
});
const hoursText = computed(() => {
let hours = '';
switch (hour.value.cronEvery) {
  case '1': hours = '*'; break;
  case '2': hours = `${hour.value.incrementStart}/${hour.value.incrementIncrement}`; break;
  case '3': hours = hour.value.specificSpecific.join(','); break;
  case '4': hours = `${hour.value.rangeStart}-${hour.value.rangeEnd}`; break;
}
return hours;
});
const daysText = computed(() => {
let days = '';
switch (day.value.cronEvery) {
  case '1': break;
  case '2':
  case '4':
  case '11': days = '?'; break;
  case '3': days = `${day.value.incrementStart}/${day.value.incrementIncrement}`; break;
  case '5': days = day.value.specificSpecific.join(','); break;
  case '6': days = 'L'; break;
  case '7': days = 'LW'; break;
  case '8': days = `${day.value.cronLastSpecificDomDay}L`; break;
  case '9': days = `L-${day.value.cronDaysBeforeEomMinus}`; break;
  case '10': days = `${day.value.cronDaysNearestWeekday}W`; break;
}
return days;
});
const weeksText = computed(() => {
let weeks = '';
switch (day.value.cronEvery) {
  case '1':
  case '3':
  case '5': weeks = '?'; break;
  case '2': weeks = `${week.value.incrementStart}/${week.value.incrementIncrement}`; break;
  case '4': weeks = week.value.specificSpecific.join(','); break;
  case '6':
  case '7':
  case '8':
  case '9':
  case '10': weeks = '?'; break;
  case '11': weeks = `${week.value.cronNthDayDay}#${week.value.cronNthDayNth}`; break;
}
return weeks;
});
const monthsText = computed(() => {
let months = '';
switch (month.value.cronEvery) {
  case '1': months = '*'; break;
  case '2': months = `${month.value.incrementStart}/${month.value.incrementIncrement}`; break;
  case '3': months = month.value.specificSpecific.join(','); break;
  case '4': months = `${month.value.rangeStart}-${month.value.rangeEnd}`; break;
}
return months;
});
const yearsText = computed(() => {
let years = '';
switch (year.value.cronEvery) {
  case '1': years = '*'; break;
  case '2': years = `${year.value.incrementStart}/${year.value.incrementIncrement}`; break;
  case '3': years = year.value.specificSpecific.join(','); break;
  case '4': years = `${year.value.rangeStart}-${year.value.rangeEnd}`; break;
}
return years;
});
const cron = computed(() => {
  return `${secondsText.value || '*'} ${minutesText.value || '*'} ${hoursText.value || '*'} ${daysText.value || '*'} ${monthsText.value || '*'} ${weeksText.value || '?'} ${yearsText.value || '*'}`;
});

// 方法
const close = () => {
  emit('cancel');
}
const handleChange = () => {
    emit('confirm', { cron: cron.value});
    close();
};

const radioStyle = reactive({
  display: 'flex',
  height: '30px',
  lineHeight: '30px',
});

const selectStyle = reactive({
  width: 'fit-content', 
  minWidth: '80px',
});

const tabListNoTitle = [
  {
    key: '1',
    tab: text.value.Seconds.name,
  },
  {
    key: '2',
    tab: text.value.Minutes.name,
  },
  {
    key: '3',
    tab: text.value.Hours.name,
  },
  {
    key: '4',
    tab: text.value.Day.name,
  },
  {
    key: '5',
    tab: text.value.Month.name,
  },
  {
    key: '6',
    tab: text.value.Year.name,
  },
];
const noTitleKey = ref('1');
const onTabChange = (value: string) => {
  noTitleKey.value = value;
};

defineExpose({
    cron,
});
</script>
<template>
  <Card
    :tab-list="tabListNoTitle"
    :active-tab-key="noTitleKey"
    @tabChange="key => onTabChange(key)"
  >
  <div v-if="noTitleKey === '1'">
        <RadioGroup v-model:value="second.cronEvery">
          <Radio value='1' :style="radioStyle">{{ text.Seconds.every }}</Radio>
          <Radio value='2' :style="radioStyle">{{ text.Seconds.interval[0] }}
            <InputNumber size="small" v-model:value="second.incrementIncrement" :min="1" :max="60"></InputNumber>
            {{ text.Seconds.interval[1] || '' }}
            <InputNumber size="small" v-model:value="second.incrementStart" :min="0" :max="59"></InputNumber>
            {{ text.Seconds.interval[2] || '' }}</Radio>

          <Radio value='3' :style="radioStyle">{{ text.Seconds.specific }}
            <Select v-model:value="second.specificSpecific" mode="multiple" :max-tag-count="10" placeholder="请选择"  :token-separators="[',']" size="small" :style="selectStyle" >
              <SelectOption v-for="(val, index) in 60" :key="index" :label="index" :value="val-1" ></SelectOption>
            </Select>
          </Radio>

          <Radio value='4' :style="radioStyle">{{ text.Seconds.cycle[0] }}
            <InputNumber size="small" v-model:value="second.rangeStart" :min="1" :max="60"></InputNumber>
            {{ text.Seconds.cycle[1] || '' }}
            <InputNumber size="small" v-model:value="second.rangeEnd" :min="0" :max="59"></InputNumber>
            {{ text.Seconds.cycle[2] || '' }}
          </Radio>
      </RadioGroup>
    </div>
    <div v-if="noTitleKey === '2'">
        <RadioGroup v-model:value="minute.cronEvery">
          <Radio :style="radioStyle" value="1" >{{ text.Minutes.every }}</Radio>
          <Radio :style="radioStyle" value="2">{{ text.Minutes.interval[0] }}
            <InputNumber size="small" v-model:value="minute.incrementIncrement" :min="1" :max="60"></InputNumber>
            {{ text.Minutes.interval[1] }}
            <InputNumber size="small" v-model:value="minute.incrementStart" :min="0" :max="59"></InputNumber>
            {{ text.Minutes.interval[2] || '' }}
          </Radio>
          <Radio :style="radioStyle" value="3">{{ text.Minutes.specific }}
            <Select  mode="multiple" v-model:value="minute.specificSpecific" :max-tag-count="10" placeholder="请选择"  :token-separators="[',']" size="small" :style="selectStyle" >
              <SelectOption v-for="(val, index) in 60" :key="index" :value="val - 1">{{
                val - 1
              }}</SelectOption>
            </Select>
          </Radio>
          <Radio :style="radioStyle" value="4">{{ text.Minutes.cycle[0] }}
            <InputNumber size="small" v-model:value="minute.rangeStart" :min="1" :max="60"></InputNumber>
            {{ text.Minutes.cycle[1] }}
            <InputNumber size="small" v-model:value="minute.rangeEnd" :min="0" :max="59"></InputNumber>
            {{ text.Minutes.cycle[2] }}
          </Radio>
        </RadioGroup>
    </div>
    <div v-if="noTitleKey === '3'">
        <RadioGroup v-model:value="hour.cronEvery">
          <Radio :style="radioStyle" value="1">{{ text.Hours.every }}</Radio>
          <Radio :style="radioStyle" value="2">{{ text.Hours.interval[0] }}
            <InputNumber size="small" v-model:value="hour.incrementIncrement" :min="0" :max="23"></InputNumber>
            {{ text.Hours.interval[1] }}
            <InputNumber size="small" v-model:value="hour.incrementStart" :min="0" :max="23"></InputNumber>
            {{ text.Hours.interval[2] }}
          </Radio>
          <Radio :style="radioStyle" value="3">{{ text.Hours.specific }}
            <Select v-model:value="hour.specificSpecific" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 24" :key="index" :value="val - 1">{{
                val - 1
              }}</SelectOption>
            </Select>
          </Radio>
          <Radio :style="radioStyle" value="4">{{ text.Hours.cycle[0] }}
            <InputNumber size="small" v-model:value="hour.rangeStart" :min="0" :max="23"></InputNumber>
            {{ text.Hours.cycle[1] }}
            <InputNumber size="small" v-model:value="hour.rangeEnd" :min="0" :max="23"></InputNumber>
            {{ text.Hours.cycle[2] }}
          </Radio>
        </RadioGroup>
    </div>
    <div v-if="noTitleKey === '4'">
        <RadioGroup v-model:value="day.cronEvery">
          <Radio :style="radioStyle" value="1">{{ text.Day.every }}</Radio>
          <Radio :style="radioStyle" value="2">{{ text.Day.intervalWeek[0] }}
            <InputNumber size="small" v-model:value="week.incrementIncrement" :min="1" :max="7"></InputNumber>
            {{ text.Day.intervalWeek[1] }}
            <Select v-model:value="week.incrementStart" size="small"  placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 7" :key="index" :value="val">{{text.Week[val - 1]}}</SelectOption>
            </Select>
            {{ text.Day.intervalWeek[2] }}
          </Radio>
          <Radio :style="radioStyle" value="3">{{ text.Day.intervalDay[0] }}
            <InputNumber size="small" v-model:value="day.incrementIncrement" :min="1" :max="31"></InputNumber>
            {{ text.Day.intervalDay[1] }}
            <InputNumber size="small" v-model:value="day.incrementStart" :min="1" :max="31"></InputNumber>
           {{ text.Day.intervalDay[2] }}
          </Radio>
          <Radio :style="radioStyle" value="4">{{ text.Day.specificWeek }}
            <Select v-model:value="week.specificSpecific" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 7" :key="index" :value="val">{{text.Week[val - 1]}}</SelectOption>
            </Select>
          </Radio>
          
          <Radio :style="radioStyle" value="5">{{ text.Day.specificDay }}
            <Select v-model:value="day.specificSpecific" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 31" :key="index" :value="val">{{ val }}</SelectOption>
            </Select>
          </Radio>
          <Radio :style="radioStyle" value="6">{{ text.Day.lastDay }}</Radio>
          <Radio :style="radioStyle" value="7">{{ text.Day.lastWeekday }}</Radio>
          <Radio :style="radioStyle" value="8">{{ text.Day.lastWeek[0] }}
            <Select v-model:value="day.cronLastSpecificDomDay" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 7" :key="index" :value="val">{{ text.Week[val - 1] }}</SelectOption>
            </Select>
            {{ text.Day.lastWeek[1] || '' }}
          </Radio>
          <Radio :style="radioStyle" value="9">
            <InputNumber size="small" v-model:value="day.cronDaysBeforeEomMinus" :min="1" :max="31"></InputNumber>
            {{ text.Day.beforeEndMonth[0] }}
          </Radio>
          <Radio :style="radioStyle" value="10">{{ text.Day.nearestWeekday[0] }}
            <InputNumber size="small" v-model:value="day.cronDaysNearestWeekday" :min="1" :max="31"></InputNumber>
            {{ text.Day.nearestWeekday[1] }}
          </Radio>
          <Radio :style="radioStyle" value="11">{{ text.Day.someWeekday[0] }}
            <InputNumber size="small" v-model:value="week.cronNthDayNth" :min="1" :max="5"></InputNumber>{{ text.Day.someWeekday[1] }}
            <Select v-model:value="week.cronNthDayDay" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 7" :key="index" :value="val">{{ text.Week[val - 1] }}</SelectOption>
            </Select>
          </Radio>  
        </RadioGroup>
    </div>
    <div v-if="noTitleKey === '5'">
        <RadioGroup v-model:value="month.cronEvery">
          <Radio :style="radioStyle" value="1">{{ text.Month.every }}</Radio>
          <Radio :style="radioStyle" value="2">{{ text.Month.interval[0] }}
            <InputNumber size="small" v-model:value="month.incrementIncrement" :min="0" :max="12"></InputNumber>
            {{ text.Month.interval[1] }}
            <InputNumber size="small" v-model:value="month.incrementStart" :min="0" :max="12"></InputNumber>
          </Radio>
          <Radio :style="radioStyle" value="3">{{ text.Month.specific }}
            <Select v-model:value="month.specificSpecific" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 12" :key="index" :value="val">{{ val }}</SelectOption>
            </Select>
          </Radio>
          <Radio :style="radioStyle" value="4">{{ text.Month.cycle[0] }}
            <InputNumber size="small" v-model:value="month.rangeStart" :min="1" :max="12"></InputNumber>
            {{ text.Month.cycle[1] }}
            <InputNumber size="small" v-model:value="month.rangeEnd" :min="1" :max="12"></InputNumber>
          </Radio>
        </RadioGroup>
    </div>
    <div v-if="noTitleKey === '6'">
        <RadioGroup v-model:value="year.cronEvery">
          <Radio :style="radioStyle" value="1">{{ text.Year.every }}</Radio>
          <Radio :style="radioStyle" value="2">{{ text.Year.interval[0] }}
            <InputNumber size="small" v-model:value="year.incrementIncrement" :min="1" :max="99"></InputNumber>
            {{ text.Year.interval[1] }}
            <InputNumber size="small" v-model:value="year.incrementStart" :min="2018" :max="2118"></InputNumber>
          </Radio>
          <Radio :style="radioStyle" value="3">{{ text.Year.specific }}
            <Select v-model:value="year.specificSpecific" size="small"  mode="multiple" :token-separators="[',']"  :max-tag-count="10" placeholder="请选择" :style="selectStyle">
              <SelectOption v-for="(val, index) in 100" :key="index" :value="2017 + val">{{ 2017 + val }}</SelectOption>
            </Select>
          </Radio>
          <Radio :style="radioStyle" value="4">{{ text.Year.cycle[0] }}
            <InputNumber size="small" v-model:value="year.rangeStart" :min="2018" :max="2118"></InputNumber>
            {{ text.Year.cycle[1] }}
            <InputNumber size="small" v-model:value="year.rangeEnd" :min="2018" :max="2118"></InputNumber>
          </Radio>
        </RadioGroup>
    </div>
  </Card>
  <Row >
    <Col class="value" :span="20">
      <span> cron预览: </span>
      <Tag color="blue">
        {{ cron }}
      </Tag>
      <span>{秒数} {分钟} {小时} {日期} {月份} {?} {年份}</span>
    </Col>
    <Col class="button" :span="4">
      <Button @click="handleChange">{{ text.Save }}</Button>
      <Button type="primary" @click="close">{{ text.Close }}</Button>
    </Col>
  </Row>
</template>
