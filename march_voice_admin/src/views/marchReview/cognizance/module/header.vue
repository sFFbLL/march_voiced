<template>
  <div class="header">
    <el-input
      v-model="query.authorNickname"
      clearable
      size="small"
      placeholder="作者昵称"
      style="width: 200px"
      class="filter-item"
      @keyup.enter.native="crud.toQuery"
    />
    <date-range-picker
      v-model="query.createTime"
      class="date-item"
      @change="go()"
    />
  </div>
</template>

<script>
import header from '@crud/crud'
import DateRangePicker from '@/components/DateRangePicker'
export default {
  components: { DateRangePicker },
  mixins: [header()],
  dicts: ['dept_status'],
  props: {
    permission: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      authorNickname: '',
      createTime: ''
    }
  },
  methods: {
    go() {
      console.log(this.crud)
    }
  }
}
</script>
