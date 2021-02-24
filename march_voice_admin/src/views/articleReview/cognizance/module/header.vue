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
      v-model="query.tag"
      clearable
      size="small"
      style="width: 200px"
      class="filter-item"
      placeholder="文章类型"
    >
      <el-option
        v-for="item in arr"
        :key="item.id"
        :label="item.tag"
        :value="item.tag"
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
export default {
  components: { DateRangePicker },
  mixins: [header(), crud()],
  dicts: ['dept_status'],
  props: {
    permission: {
      type: Object,
      required: true
    }
  },
  data() {
    return {
      invitationNameArray: [
        {
          id: '1', // 这里的id一定要加引号
          light: '审核通过'
        },
        {
          id: '0',
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
    this.arrnew = this.crud.data.map((obj, index) => {
      return obj.tag
    }).join(',').split(',')
    for (var i = 0; i < this.arrnew.length; i++) { // 循环遍历当前数组
      if (this.arr.indexOf(this.arrnew[i]) == -1) {
        this.arr.push(this.arrnew[i])
      }
    }
    console.log(222222222)
    console.log(this.arr)
    // let arrr = [];
    // for (let index = 0; index < this.arr.length; index++) {
    //   arrr.push({ tag: "" });
    //   console.log(arrr)
    //   arrr[index].tag = this.arr[index];
    //   console.log(arrr)
    // }

    // this.arr.forEach(item => {
    //   arrr.push({ tag: item });
    // })
    // console.log(arrr)
    // this.arr = arrr;
    // console.log(2222222222)
    // console.log(this.arr)
  },
  methods: {
    // [CRUD.HOOK.beforeRefresh] (crud) {//将参数转为对象
    //   this.arrnew = this.crud.data.map((obj, index) => {
    //     return obj.tag;
    //   }).join(",").split(',')
    //   for (var i = 0; i < this.arrnew.length; i++) {    //循环遍历当前数组
    //     if (this.arr.indexOf(this.arrnew[i]) == -1) {
    //       this.arr.push(this.arrnew[i]);
    //     }
    //   }
    //   console.log(222222222)
    //   console.log(this.arr)
    //   this.arr.forEach(item => {
    //     arrr.push({ tag: item });
    //   })
    //   console.log(arrr)
    //   this.arr = arrr;
    //   console.log(2222222222)
    //   console.log(this.arr)
    // },

  }
}
</script>
