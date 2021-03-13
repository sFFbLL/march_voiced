<template>
  <div v-if="crud.props.searchToggle" style="max-width: 1500px" class="header">
    <!-- 搜索 -->
    <el-input
      v-model="query.title"
      clearable
      size="small"
      placeholder="文章搜索框"
      style="width: 200px"
      class="filter-item"
      @keyup.enter.native="crud.toQuery"
    />
    <el-input
      v-model="query.nickname"
      clearable
      size="small"
      placeholder="作者昵称"
      style="width: 200px"
      class="filter-item"
      @keyup.enter.native="crud.toQuery"
    />
    <el-select
      v-if="once"
      v-model="query.tag"
      clearable
      size="small"
      style="width: 200px"
      class="filter-item"
      placeholder="文章类型"
    >
      <el-option
        v-for="item in tag"
        :key="item.id"
        :label="item.tag"
        :value="item.id"
      />
    </el-select>
    <el-select
      v-model="query.status"
      clearable
      size="small"
      style="width: 200px"
      class="filter-item"
      placeholder="全部状态"
    >
      <el-option
        v-for="item in invitationNameArray"
        :key="item.id"
        :label="item.light"
        :value="item.id"
      />
    </el-select>
    <date-range-picker
      v-model="query.createTime"
      class="date-item"
      @change="go()"
    />
  </div>
</template>

<script>
import { header, crud } from '@crud/crud'
import DateRangePicker from '@/components/DateRangePicker'
import { getInfoBy } from '@/api/review/articleReview.js'
export default {
  components: { DateRangePicker },
  mixins: [header(), crud()],
  dicts: ['dept_status'],
  props: {
    data: {
      type: Object,
      required: true
    },
    permission: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      aa: 1,
      once: true,
      invitationNameArray: [
        {
          id: '1', // 这里的id一定要加引号
          light: '审核通过'
        },
        {
          id: '2',
          light: '待审核'
        }
      ],
      selectOptionArrays: this.$store.getters.selectOptionArrays,
      title: '',
      nickname: '',
      tag: '',
      status: '',
      createTime: '',
      arrnew: [],
      arr: []
      // res: []
    }
  },
  created() {
    getInfoBy('/api/article/tag').then(
      data => {
        console.log(data.data)
        this.tag = data.data
        // console.log(this.tag)
        // this.contents = data.data.content
      }
    )
  },
  methods: {
    // getInfoBy,
  }
}
</script>
