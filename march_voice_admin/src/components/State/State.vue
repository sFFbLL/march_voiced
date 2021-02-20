<template>
  <div class="status">
    <p v-if="data.objectState == 1" class="status">待审核</p>
    <p
      v-else-if="data.objectState == 2"
      class="status"
      style="color: rgba(67, 207, 124, 1)"
    >
      审核通过
    </p>
    <p
      v-else-if="data.objectState == 3"
      class="status"
      style="color: rgba(255, 87, 51, 1)"
    >
      驳回
    </p>
  </div>
</template>

<script>
export default {
  name: "State",
  props: {
    data: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style>
.status {
  font-size: 13px;
  line-height: 0;
  color: rgba(42, 130, 228, 1);
}
</style>
