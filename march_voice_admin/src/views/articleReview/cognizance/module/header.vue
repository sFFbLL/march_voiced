<template>
  <div class="header">
    <!-- 搜索 -->
    <el-input
      v-model="articleSearch"
      clearable
      size="small"
      placeholder="文章搜索框"
      style="width: 200px"
      class="filter-item"
    />
    <el-input
      v-model="authorNickname"
      clearable
      size="small"
      placeholder="作者昵称"
      style="width: 200px"
      class="filter-item"
    />
    <el-select
      v-model="articleType"
      clearable
      size="small"
      style="width: 200px"
      class="filter-item"
      placeholder="文章类型"
    >
      <el-option
        v-for="(item, key) in options"
        :key="key"
        :label="item['label']"
        :value="item['value']"
      />
    </el-select>
    <el-select
      v-model="objectState"
      clearable
      size="small"
      style="width: 200px"
      class="filter-item"
      placeholder="全部状态"
    >
      <el-option
        v-for="(item, key) in options"
        :key="key"
        :label="item['label']"
        :value="item['value']"
      />
    </el-select>
    <date-range-picker v-model="data" class="date-item" @change="go()" />
  </div>
</template>

<script>
import DateRangePicker from '@/components/DateRangePicker'
export default {
  components: { DateRangePicker },
  // mixins: [header()],
  dicts: ['dept_status'],
  props: {
    permission: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      articleSearch: '',
      authorNickname: '',
      articleType: '',
      objectState: '',
      data: '',
      options: [
        {
          value: '选项1',
          label: '黄金糕'
        },
        {
          value: '选项2',
          label: '双皮奶'
        },
        {
          value: '选项3',
          label: '蚵仔煎'
        },
        {
          value: '选项4',
          label: '龙须面'
        },
        {
          value: '选项5',
          label: '北京烤鸭'
        }
      ]
    }
  },
  methods: {
    go() {
      console.log(this.data)
    }
  }
}
</script>
