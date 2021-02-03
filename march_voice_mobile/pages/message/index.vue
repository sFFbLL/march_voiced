<template>
	<view>
		<view class="header">
			<view class="head-nav-items">
				<tabs :tabs='tablist' v-on:tabActive='tabActive' />
			</view>
			
		</view>
		
	</view>
</template>

<script>
	import tabs from '../../marchVoiceComponents/tabCard.vue'
	export default {
		components: {
			tabs
		},
		data() {
			return {
				tablist: [{
						index: 0,
						value: '互动',
						isActive: true
					},
					{
						index: 1,
						value: '关注',
						isActive: false
					},
					{
						index: 2,
						value: '其他',
						isActive: false
					}
				],
				contentIndex: '',
			}
		},
		methods: {
			tabActive(tabIndex) {
				this.tablist.map((value, index) => {
					value.isActive = tabIndex == index ? true : false;
				})
				if (!tabIndex) {
					this.recommend();
				} else if (tabIndex) {
					this.attention();
				}
				this.contentIndex = tabIndex;
			},
			recommend() {
				console.log(1);
			},
			attention() {
				console.log(2);
			}
		},
	}
	
	
</script>

<style>
</style>
