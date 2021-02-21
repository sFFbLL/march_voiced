<template>
  <!-- eladmin自带的布局方式 -->
  <div v-loading="loading" class="app-container">
    <!-- 头部盒子 -->
    <!-- <el-row> -->
    <!-- <el-col :span="10"> -->
    <div class="style-a4">
      <div class="position-box">
        <div class="body-box">
          <!-- 状态显示 -->
          <div v-if="data.objectState == 2" class="status-box success">
            <span>已通过</span>
          </div>
          <div v-if="this.data.objectState == 1" class="status-box wait">
            <span>待审核</span>
          </div>
          <div v-if="this.data.objectState == 3" class="status-box failed">
            <span>已驳回</span>
          </div>
          <div class="button-box">
            <Reject
              v-if="this.showReject"
              :content="this.$route.query.values"
              :data="data"
            />
          </div>
          <div class="head-box">
            <!-- 获奖成果标题 -->
            <div class="tittle-box">
              <span class="ql-editor" v-html="showHtml(contents)" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Reject from '@/components/MarchAffirmation/MarchAffirmation.vue'
// import DetailState from '@/components/State/DetailState.vue'
// import { getInfoById } from "@/api/data";
// import Cookies from "js-cookie";
// import { validateSelectOptionId, getAllName } from "@/utils/validate.js";

export default {
  name: 'Detail',
  components: { Reject },

  data() {
    return {
      data: { objectState: 1 },
      loading: false,
      showReject: true,
      contents:
        '<p>爱搜IDno为阿斯兰的框架内</p><hr><p><br></p><hr><h2>我是test</h2><p>我是<strong>加粗</strong></p><p><em>我是斜体</em><span class="ql-cursor">﻿</span></p>'
    }
  },
  created() {
    this.data.objectState = 3
  },

  methods: {
    showHtml(contents) {
      return contents
        .replace(contents ? /&(?!#?\w+;)/g : /&/g, '&amp;')
        .replace(/&lt;/g, '<')
        .replace(/&gt;/g, '>')
        .replace(/&quot;/g, '"')
        .replace(/&#39;/g, "'")
    }
    // getInfoById,
    // getAllName,
    // validateSelectOptionId
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@import "@/assets/styles/detail.scss";
</style>
