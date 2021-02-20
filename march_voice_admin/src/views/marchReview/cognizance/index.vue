<template>
  <div class="app-container">
    <div class="head-container">
      <eHeader :permission="permission" />
      <crudOperation :permission="permission" :is-my-export="true" />
    </div>
    <!--表格渲染-->
    <el-table
      ref="table"
      v-loading="crud.loading"
      stripe
      lazy
      :data="unsigns"
      row-key="id"
      style="width: 100%"
    >
      <el-table-column prop="nickname" label="用户昵称" />
      <el-table-column prop="dept" label="用户身份" />
      <el-table-column prop="phone" label="手机号" />
      <el-table-column prop="createTime" label="申请时间" width="85" />
      <el-table-column
        prop="status"
        class="colum-name"
        fixed="right"
        label="状态"
        width="130"
      >
        <template slot-scope="scope">
          <!-- <span > -->
          <State :data="scope.row" />
          <!-- </span> -->
        </template>
      </el-table-column>
      <!--  操作   -->
      <el-table-column
        label="操作"
        align="center"
        fixed="right"
        :width="operateWidth"
      >
        <template slot-scope="scope">
          <MarchAffirmation
            :data="scope.row"
            :content="content"
            @btnSetWidthEvent="getBtnSetWidth"
          />
        </template>
        <!-- console.log(tableDataa); -->
      </el-table-column>
    </el-table>
    <!--分页组件-->
    <pagination />
  </div>
</template>

<script>
import eHeader from "./module/header";
// import { validatePoint, validateSelectOptionId } from "@/utils/validate.js";
import CRUD, { presenter } from "@crud/crud";
import crudOperation from "@/components/AffirmationSearch/AffirmationSearch.vue";
import pagination from "@crud/Pagination";
import State from "@/components/State/State.vue";
import MarchAffirmation from "@/components/MarchAffirmation/MarchAffirmation.vue";

export default {
  name: "MarchReview",
  components: { eHeader, pagination, State, crudOperation, MarchAffirmation },
  cruds() {
    return CRUD({
      title: "加入三月圈审核",
      url: "/api/apply/march",
      // 主页操作栏显示哪些按钮
      optShow: {
        add: false,
        edit: false,
        del: false,
        download: true,
        reset: true,
      },
    });
  },
  mixins: [presenter()],

  data() {
    return {
      // tableDataa: {
      //   nickname: "",
      //   dept: "",
      //   phone: "",
      //   createTime: "",
      // },
      tableData: [
        {
          nickname: "2016-05-02",
          dept: "小虎",
          phone: "16519826435",
          createTime: "上海市普陀区金沙江路 1518 弄",
          objectState: 1,
        },
        {
          nickname: "2016-05-04",
          dept: "王小虎",
          phone: "16519826435",
          createTime: "上海市普陀区金沙江路 1518 弄",
          objectState: 2,
        },
        {
          nickname: "2016-05-01",
          dept: "王小虎",
          phone: "16519826435",
          createTime: "上海市普陀区金沙江路 1518 弄",
          objectState: 3,
        },
        {
          nickname: "2016-05-03",
          dept: "小虎",
          phone: "16519826435",
          createTime: "上海市普陀区金沙江路 1518 弄",
          objectState: 3,
        },
      ],
      content: {
        toDetialPath: "details",
      },
      permission: {
        add: ["admin", "job:add"],
        edit: ["admin", "job:edit"],
        del: ["admin", "job:del"],
      },
      // 操作列宽度
      operateWidth: 150,
    };
  },
  computed: {
    unsigns() {
      return this.tableData.filter((item) => item.objectState === 1);
    },
  },

  methods: {
    // validate(scope, identity) {
    //   validatePoint(scope, identity, this);
    // },
    getBtnSetWidth(width) {
      width += 15;
      this.operateWidth = width > this.operateWidth ? width : this.operateWidth;
    },
    // validateSelectOptionId,
  },
};
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
// @import "@/assets/styles/point.scss";
::v-deep .el-input-number .el-input__inner {
  text-align: left;
}
::v-deep.el-table .cell {
  white-space: nowrap;
  text-overflow: ellipsis;
}
</style>
