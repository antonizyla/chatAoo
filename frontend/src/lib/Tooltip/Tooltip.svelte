<script lang="ts">
	export let outer: string;
</script>

<div class="tooltip">
	{outer}
	<span class="tooltiptext"><slot /></span>
</div>

<style lang="postcss">
	.tooltip {
		@apply relative inline-block;
	}

	.tooltip .tooltiptext {
		@apply rounded-md p-2 left-0 bottom-full z-[1] absolute text-center invisible bg-gray-900 text-white;
	}

	.tooltip:hover .tooltiptext {
		@apply visible;
	}
</style>
