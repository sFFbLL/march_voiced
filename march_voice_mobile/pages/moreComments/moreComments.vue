<template>
	<view>
		<view class="comment-list">
			<comment :type="type" @getMore="getMore" :kidsCommentCount="kidsCommentCount" :commentList="commentList" @childFn="comment"></comment>
		</view>
	</view>
</template>

<script>
	import comment from '../../marchVoiceComponents/comment/index.vue'
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			
		}
	}
</script>

<style>

</style>
