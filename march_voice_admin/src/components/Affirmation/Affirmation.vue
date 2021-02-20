<template>
  <div ref="btnSet" class="wrap">
    <el-button size="mini" class="filter-item" type="primary">详情</el-button>
    <el-button
      v-if="data.objectState == 3 || data.objectState == 1"
      :disabled="disabled"
      size="mini"
      type="success"
      :class="{ active: disabled }"
      class="filter-item"
      @click="doAdopt()"
      >通过</el-button
    >
    <el-button
      v-if="data.objectState == 2"
      size="mini"
      type="warning"
      :class="{ active: disabled }"
      class="filter-item"
      >推荐</el-button
    >
    <el-button
      v-if="data.objectState == 3 || data.objectState == 1"
      :disabled="disabled"
      class="filter-item"
      type="danger"
      size="mini"
      :class="{ active: disabled }"
      @click="judgeText()"
      >驳回</el-button
    >
  </div>
</template>

<script>
// import crudMenu from "@/api/system/menu";
import CRUD, { crud } from "@crud/crud";
import { articleAdopt } from "@/api/review/articleReview.js";
// import Cookies from "js-cookie";

export default {
  name: "ProjectAffirmation",
  mixins: [crud()],
  cruds() {
    return CRUD({ title: "组件", url: "", crudMethod: { ...crudMenu } });
  },
  props: {
    data: {
      type: Object,
      required: true,
    },
    content: {
      type: Object,
      required: true,
    },
  },
  data() {
    return {
      disabled: false,
      superadmin: false,
      roleStatus: null,
    };
  },
  updated() {
    if (this.data.objectState == 3) {
      this.disabled = true;
    } else {
      this.disabled = false;
    }
  },
  created() {
    if (this.data.objectState == 3) {
      this.disabled = true;
    } else {
      this.disabled = false;
    }
  },
  mounted() {
    this.$emit("btnSetWidthEvent", this.$refs.btnSet.clientWidth);
  },
  methods: {
    // // 路由跳转详情方法
    // toDetails(reject) {
    //   const workPoint = this.data.workPoint;
    //   const bonusPoint = this.data.bonusPoint;
    //   const objectState = parseInt(this.data.objectState);
    //   const identity = this.content.identity;
    //   const dataId = this.data.id;
    //   const showReject = true;
    //   const values = {
    //     workPoint,
    //     bonusPoint,
    //     identity,
    //     dataId,
    //     objectState,
    //     reject,
    //     showReject,
    //   };
    //   if (identity === "project") {
    //     values.secondPlanCloseDate = this.data.secondPlanCloseDate;
    //     values.actualCloseDate = this.data.actualCloseDate;
    //   }
    //   Cookies.set(identity, values);
    //   const path = this.content.toDetialPath;
    //   this.$router.push({
    //     path: `${path}`,
    //     query: {
    //       values,
    //     },
    //   });
    // },

    judgeText() {
      this.$confirm("请再次确定是否驳回", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        const dataId = parseInt(this.data.id);
        const objectState = parseInt(this.data.objectState);
        const isShow = objectState === 1;
        const status = 1;
        const doAdoptData = {
          status,
          dataId,
        };
        console.log(doAdoptData);
        if (isShow) {
          articleAdopt(doAdoptData)
            .then((rep) => {})
            .then(() => {
              this.$message({
                title: "成功",
                message: "成功驳回",
                type: "success",
              });
              objectState = 3;
            })
            .catch(() => {
              this.$message({
                type: "info",
                message: "已取消驳回",
              });
              objectState = 1;
            });
        }
      });
    },

    // 通过方法
    doAdopt() {
      this.$confirm("请再次确定是否通过", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        const dataId = parseInt(this.data.id);
        const objectState = parseInt(this.data.objectState);
        const isShow = objectState === 1;
        const status = 1;
        const doAdoptData = {
          status,
          dataId,
        };
        console.log(doAdoptData);
        if (isShow) {
          articleAdopt(doAdoptData)
            .then((rep) => {})
            .then(() => {
              this.$message({
                type: "success",
                message: "通过成功！",
              });
              objectState = 2;
              // this.disabled = false;
            })
            .catch(() => {
              this.$message({
                type: "info",
                message: "通过失败",
              });
            });
        }
      });
    },
  },
};
</script>

<style scoped>
.wrap {
  float: left;
  display: flex;
  flex-flow: row nowrap;
  justify-content: center;
  align-items: center;
}
.filter-item {
  margin: auto 3px;
  padding: 5px 5px;
}
.active {
  background-color: #909399 !important;
  border-color: #909399 !important;
}
</style>
